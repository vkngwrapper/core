package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestVulkanLoader1_0_CreateDevice_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockPhysicalDevice := mocks.EasyMockPhysicalDevice(ctrl, mockDriver)
	deviceHandle := mocks.NewFakeDeviceHandle()

	mockDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(mockDriver, nil)
	mockDriver.EXPECT().VkCreateDevice(mockPhysicalDevice.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(physicalDevice core.VkPhysicalDevice, pCreateInfo *core.VkDeviceCreateInfo, pAllocator *core.VkAllocationCallbacks, pDevice *core.VkDevice) (core.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(3), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), v.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), v.FieldByName("queueCreateInfoCount").Uint())
			require.Equal(t, uint64(3), v.FieldByName("enabledExtensionCount").Uint())
			require.Equal(t, uint64(2), v.FieldByName("enabledLayerCount").Uint())

			featuresV := v.FieldByName("pEnabledFeatures").Elem()

			require.Equal(t, uint64(1), featuresV.FieldByName("occlusionQueryPrecise").Uint())
			require.Equal(t, uint64(1), featuresV.FieldByName("tessellationShader").Uint())
			require.Equal(t, uint64(0), featuresV.FieldByName("samplerAnisotropy").Uint())

			extensionNamePtr := (**core.Char)(unsafe.Pointer(v.FieldByName("ppEnabledExtensionNames").Elem().UnsafeAddr()))
			extensionNameSlice := ([]*core.Char)(unsafe.Slice(extensionNamePtr, 3))

			var extensionNames []string
			for _, extensionNameBytes := range extensionNameSlice {
				var extensionNameRunes []rune
				extensionNameByteSlice := ([]core.Char)(unsafe.Slice(extensionNameBytes, 1<<30))
				for _, nameByte := range extensionNameByteSlice {
					if nameByte == 0 {
						break
					}

					extensionNameRunes = append(extensionNameRunes, rune(nameByte))
				}

				extensionNames = append(extensionNames, string(extensionNameRunes))
			}

			require.ElementsMatch(t, []string{"A", "B", "C"}, extensionNames)

			layerNamePtr := (**core.Char)(unsafe.Pointer(v.FieldByName("ppEnabledLayerNames").Elem().UnsafeAddr()))
			layerNameSlice := ([]*core.Char)(unsafe.Slice(layerNamePtr, 2))

			var layerNames []string
			for _, layerNameBytes := range layerNameSlice {
				var layerNameRunes []rune
				layerNameByteSlice := ([]core.Char)(unsafe.Slice(layerNameBytes, 1<<30))
				for _, nameByte := range layerNameByteSlice {
					if nameByte == 0 {
						break
					}

					layerNameRunes = append(layerNameRunes, rune(nameByte))
				}

				layerNames = append(layerNames, string(layerNameRunes))
			}

			require.ElementsMatch(t, []string{"D", "E"}, layerNames)

			queueCreateInfoPtr := (*core.VkDeviceQueueCreateInfo)(unsafe.Pointer(v.FieldByName("pQueueCreateInfos").Elem().UnsafeAddr()))
			queueCreateInfoSlice := ([]core.VkDeviceQueueCreateInfo)(unsafe.Slice(queueCreateInfoPtr, 2))

			queueInfoV := reflect.ValueOf(queueCreateInfoSlice[0])
			require.Equal(t, uint64(2), queueInfoV.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
			require.True(t, queueInfoV.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), queueInfoV.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), queueInfoV.FieldByName("queueFamilyIndex").Uint())
			require.Equal(t, uint64(3), queueInfoV.FieldByName("queueCount").Uint())

			priorityPtr := (*float32)(unsafe.Pointer(queueInfoV.FieldByName("pQueuePriorities").Elem().UnsafeAddr()))
			prioritySlice := ([]float32)(unsafe.Slice(priorityPtr, 3))
			require.Equal(t, float32(1.0), prioritySlice[0])
			require.Equal(t, float32(0.0), prioritySlice[1])
			require.Equal(t, float32(0.5), prioritySlice[2])

			queueInfoV = reflect.ValueOf(queueCreateInfoSlice[1])
			require.Equal(t, uint64(2), queueInfoV.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
			require.True(t, queueInfoV.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), queueInfoV.FieldByName("flags").Uint())
			require.Equal(t, uint64(3), queueInfoV.FieldByName("queueFamilyIndex").Uint())
			require.Equal(t, uint64(1), queueInfoV.FieldByName("queueCount").Uint())

			priorityPtr = (*float32)(unsafe.Pointer(queueInfoV.FieldByName("pQueuePriorities").Elem().UnsafeAddr()))
			prioritySlice = ([]float32)(unsafe.Slice(priorityPtr, 1))
			require.Equal(t, float32(0.5), prioritySlice[0])

			*pDevice = deviceHandle
			return core.VKSuccess, nil
		})

	device, _, err := loader.CreateDevice(mockPhysicalDevice, &core.DeviceOptions{
		QueueFamilies: []*core.QueueFamilyOptions{
			{
				QueueFamilyIndex: 1,
				QueuePriorities:  []float32{1, 0, 0.5},
			},
			{
				QueueFamilyIndex: 3,
				QueuePriorities:  []float32{0.5},
			},
		},
		ExtensionNames: []string{"A", "B", "C"},
		LayerNames:     []string{"D", "E"},
		EnabledFeatures: &common.PhysicalDeviceFeatures{
			OcclusionQueryPrecise: true,
			TessellationShader:    true,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, deviceHandle, device.Handle())
}

func TestVulkanLoader1_0_CreateDevice_FailNoQueueFamilies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	mockPhysicalDevice := mocks.EasyMockPhysicalDevice(ctrl, mockDriver)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	_, _, err = loader.CreateDevice(mockPhysicalDevice, &core.DeviceOptions{
		QueueFamilies:  []*core.QueueFamilyOptions{},
		ExtensionNames: []string{"A", "B", "C"},
		LayerNames:     []string{"D", "E"},
		EnabledFeatures: &common.PhysicalDeviceFeatures{
			OcclusionQueryPrecise: true,
			TessellationShader:    true,
		},
	})
	require.EqualError(t, err, "alloc DeviceOptions: no queue families added")
}

func TestVulkanLoader1_0_CreateDevice_FailFamilyWithoutPriorities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	mockPhysicalDevice := mocks.EasyMockPhysicalDevice(ctrl, mockDriver)

	_, _, err = loader.CreateDevice(mockPhysicalDevice, &core.DeviceOptions{
		QueueFamilies: []*core.QueueFamilyOptions{
			{
				QueueFamilyIndex: 1,
				QueuePriorities:  []float32{1, 0, 0.5},
			},
			{
				QueueFamilyIndex: 3,
				QueuePriorities:  []float32{},
			},
		},
		ExtensionNames: []string{"A", "B", "C"},
		LayerNames:     []string{"D", "E"},
		EnabledFeatures: &common.PhysicalDeviceFeatures{
			OcclusionQueryPrecise: true,
			TessellationShader:    true,
		},
	})
	require.EqualError(t, err, "alloc DeviceOptions: queue family 3 had no queue priorities")
}

func TestDevice_GetQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	queueHandle := mocks.NewFakeQueue()

	mockDriver.EXPECT().VkGetDeviceQueue(device.Handle(), core.Uint32(1), core.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, queueFamilyIndex, queueIndex core.Uint32, pQueue *core.VkQueue) {
			*pQueue = queueHandle
		})

	queue := device.GetQueue(1, 2)
	require.NotNil(t, queue)
	require.Equal(t, queueHandle, queue.Handle())
}

func TestDevice_WaitForFences_Timeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	fence1 := mocks.EasyDummyFence(t, loader, device)
	fence2 := mocks.EasyDummyFence(t, loader, device)

	mockDriver.EXPECT().VkWaitForFences(device.Handle(), core.Uint32(2), gomock.Not(nil), core.VkBool32(1), core.Uint64(1)).DoAndReturn(
		func(device core.VkDevice, fenceCount core.Uint32, pFences *core.VkFence, waitAll core.VkBool32, timeout core.Uint64) (core.VkResult, error) {
			fenceSlice := ([]core.VkFence)(unsafe.Slice(pFences, 2))
			require.ElementsMatch(t, []core.VkFence{fence1.Handle(), fence2.Handle()}, fenceSlice)

			return core.VKSuccess, nil
		})

	_, err = device.WaitForFences(true, time.Nanosecond, []core.Fence{
		fence1, fence2,
	})
	require.NoError(t, err)
}

func TestDevice_WaitForFences_NoTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	fence1 := mocks.EasyDummyFence(t, loader, device)

	mockDriver.EXPECT().VkWaitForFences(device.Handle(), core.Uint32(1), gomock.Not(nil), core.VkBool32(0), core.Uint64(0xffffffffffffffff)).DoAndReturn(
		func(device core.VkDevice, fenceCount core.Uint32, pFences *core.VkFence, waitAll core.VkBool32, timeout core.Uint64) (core.VkResult, error) {
			fenceSlice := ([]core.VkFence)(unsafe.Slice(pFences, 1))
			require.ElementsMatch(t, []core.VkFence{fence1.Handle()}, fenceSlice)

			return core.VKSuccess, nil
		})

	_, err = device.WaitForFences(false, common.NoTimeout, []core.Fence{
		fence1,
	})
	require.NoError(t, err)
}

func TestDevice_WaitForIdle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)

	mockDriver.EXPECT().VkDeviceWaitIdle(device.Handle()).Return(core.VKSuccess, nil)
	_, err = device.WaitForIdle()
	require.NoError(t, err)
}

func TestVulkanDevice_ResetFences(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	fence1 := mocks.EasyDummyFence(t, loader, device)
	fence2 := mocks.EasyDummyFence(t, loader, device)

	mockDriver.EXPECT().VkResetFences(device.Handle(), core.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, fenceCount core.Uint32, pFence *core.VkFence) (core.VkResult, error) {
			fences := ([]core.VkFence)(unsafe.Slice(pFence, 2))

			require.ElementsMatch(t, []core.VkFence{fence1.Handle(), fence2.Handle()}, fences)
			return core.VKSuccess, nil
		})

	_, err = device.ResetFences([]core.Fence{fence1, fence2})
	require.NoError(t, err)
}

func TestVulkanDevice_AllocateAndFreeMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	memoryHandle := mocks.NewFakeDeviceMemoryHandle()

	mockDriver.EXPECT().VkAllocateMemory(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkMemoryAllocateInfo, pAllocator *core.VkAllocationCallbacks, pMemory *core.VkDeviceMemory) (core.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(7), val.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), val.FieldByName("memoryTypeIndex").Uint())

			*pMemory = memoryHandle
			return core.VKSuccess, nil
		})
	mockDriver.EXPECT().VkFreeMemory(device.Handle(), memoryHandle, nil)

	memory, _, err := device.AllocateMemory(&core.DeviceMemoryOptions{
		AllocationSize:  7,
		MemoryTypeIndex: 3,
	})
	require.NoError(t, err)
	require.NotNil(t, memory)
	require.Equal(t, memoryHandle, memory.Handle())

	device.FreeMemory(memory)
}

func TestVulkanDevice_UpdateDescriptorSets_WriteImageInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	sampler1 := mocks.EasyMockSampler(ctrl)
	sampler2 := mocks.EasyMockSampler(ctrl)
	imageView1 := mocks.EasyMockImageView(ctrl)
	imageView2 := mocks.EasyMockImageView(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), core.Uint32(1), gomock.Not(nil), core.Uint32(0), nil).DoAndReturn(
		func(device core.VkDevice, descriptorWriteCount core.Uint32, pDescriptorWrites *core.VkWriteDescriptorSet, descriptorCopyCount core.Uint32, pDescriptorCopies *core.VkCopyDescriptorSet) error {
			writeSlice := ([]core.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), core.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pBufferInfo").IsNil())
			require.True(t, writeVal.FieldByName("pTexelBufferView").IsNil())

			imageInfoPtr := (*core.VkDescriptorImageInfo)(unsafe.Pointer(writeVal.FieldByName("pImageInfo").Elem().UnsafeAddr()))
			imageInfoSlice := ([]core.VkDescriptorImageInfo)(unsafe.Slice(imageInfoPtr, 2))

			require.Len(t, imageInfoSlice, 2)

			imageInfoVal := reflect.ValueOf(imageInfoSlice[0])
			require.Equal(t, sampler1.Handle(), (core.VkSampler)(unsafe.Pointer(imageInfoVal.FieldByName("sampler").Elem().UnsafeAddr())))
			require.Equal(t, imageView1.Handle(), (core.VkImageView)(unsafe.Pointer(imageInfoVal.FieldByName("imageView").Elem().UnsafeAddr())))
			require.Equal(t, uint64(3), imageInfoVal.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL

			imageInfoVal = reflect.ValueOf(imageInfoSlice[1])
			require.Equal(t, sampler2.Handle(), (core.VkSampler)(unsafe.Pointer(imageInfoVal.FieldByName("sampler").Elem().UnsafeAddr())))
			require.Equal(t, imageView2.Handle(), (core.VkImageView)(unsafe.Pointer(imageInfoVal.FieldByName("imageView").Elem().UnsafeAddr())))
			require.Equal(t, uint64(8), imageInfoVal.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_PREINITIALIZED

			return nil
		})

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 2,
			DescriptorType:          common.DescriptorUniformBuffer,
			ImageInfo: []core.DescriptorImageInfo{
				{
					Sampler:     sampler1,
					ImageView:   imageView1,
					ImageLayout: common.LayoutDepthStencilAttachmentOptimal,
				},
				{
					Sampler:     sampler2,
					ImageView:   imageView2,
					ImageLayout: common.LayoutPreInitialized,
				},
			},
		},
	}, nil)

	require.NoError(t, err)
}
func TestVulkanDevice_UpdateDescriptorSets_WriteBufferInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	buffer1 := mocks.EasyMockBuffer(ctrl)
	buffer2 := mocks.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), core.Uint32(1), gomock.Not(nil), core.Uint32(0), nil).DoAndReturn(
		func(device core.VkDevice, descriptorWriteCount core.Uint32, pDescriptorWrites *core.VkWriteDescriptorSet, descriptorCopyCount core.Uint32, pDescriptorCopies *core.VkCopyDescriptorSet) error {
			writeSlice := ([]core.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), core.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(3), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pImageInfo").IsNil())
			require.True(t, writeVal.FieldByName("pTexelBufferView").IsNil())

			bufferInfoPtr := (*core.VkDescriptorBufferInfo)(unsafe.Pointer(writeVal.FieldByName("pBufferInfo").Elem().UnsafeAddr()))
			bufferInfoSlice := ([]core.VkDescriptorBufferInfo)(unsafe.Slice(bufferInfoPtr, 2))

			require.Len(t, bufferInfoSlice, 2)

			bufferInfoVal := reflect.ValueOf(bufferInfoSlice[0])
			require.Equal(t, buffer1.Handle(), (core.VkBuffer)(unsafe.Pointer(bufferInfoVal.FieldByName("buffer").Elem().UnsafeAddr())))
			require.Equal(t, uint64(7), bufferInfoVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(11), bufferInfoVal.FieldByName("_range").Uint())

			bufferInfoVal = reflect.ValueOf(bufferInfoSlice[1])
			require.Equal(t, buffer2.Handle(), (core.VkBuffer)(unsafe.Pointer(bufferInfoVal.FieldByName("buffer").Elem().UnsafeAddr())))
			require.Equal(t, uint64(13), bufferInfoVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(17), bufferInfoVal.FieldByName("_range").Uint())

			return nil
		})

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
			BufferInfo: []core.DescriptorBufferInfo{
				{
					Buffer: buffer1,
					Offset: 7,
					Range:  11,
				},
				{
					Buffer: buffer2,
					Offset: 13,
					Range:  17,
				},
			},
		},
	}, nil)

	require.NoError(t, err)
}

func TestVulkanDevice_UpdateDescriptorSets_TexelBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	bufferView1 := mocks.EasyMockBufferView(ctrl)
	bufferView2 := mocks.EasyMockBufferView(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), core.Uint32(1), gomock.Not(nil), core.Uint32(0), nil).DoAndReturn(
		func(device core.VkDevice, descriptorWriteCount core.Uint32, pDescriptorWrites *core.VkWriteDescriptorSet, descriptorCopyCount core.Uint32, pDescriptorCopies *core.VkCopyDescriptorSet) error {
			writeSlice := ([]core.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), core.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(3), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pImageInfo").IsNil())
			require.True(t, writeVal.FieldByName("pBufferInfo").IsNil())

			bufferInfoPtr := (*core.VkBufferView)(unsafe.Pointer(writeVal.FieldByName("pTexelBufferView").Elem().UnsafeAddr()))
			bufferInfoSlice := ([]core.VkBufferView)(unsafe.Slice(bufferInfoPtr, 2))

			require.Len(t, bufferInfoSlice, 2)

			require.Equal(t, bufferView1.Handle(), (core.VkBufferView)(unsafe.Pointer(bufferInfoSlice[0])))
			require.Equal(t, bufferView2.Handle(), (core.VkBufferView)(unsafe.Pointer(bufferInfoSlice[1])))

			return nil
		})

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
			TexelBufferView: []core.BufferView{
				bufferView1, bufferView2,
			},
		},
	}, nil)

	require.NoError(t, err)
}
func TestVulkanDevice_UpdateDescriptorSets_Copy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	srcDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), core.Uint32(0), nil, core.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, descriptorWriteCount core.Uint32, pDescriptorWrites *core.VkWriteDescriptorSet, descriptorCopyCount core.Uint32, pDescriptorCopies *core.VkCopyDescriptorSet) error {
			copySlice := ([]core.VkCopyDescriptorSet)(unsafe.Slice(pDescriptorCopies, 1))
			copyVal := reflect.ValueOf(copySlice[0])

			require.Equal(t, uint64(36), copyVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_DESCRIPTOR_SET
			require.True(t, copyVal.FieldByName("pNext").IsNil())
			require.Equal(t, srcDescriptor.Handle(), (core.VkDescriptorSet)(unsafe.Pointer(copyVal.FieldByName("srcSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(3), copyVal.FieldByName("srcBinding").Uint())
			require.Equal(t, uint64(5), copyVal.FieldByName("srcArrayElement").Uint())
			require.Equal(t, destDescriptor.Handle(), (core.VkDescriptorSet)(unsafe.Pointer(copyVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(7), copyVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(11), copyVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(13), copyVal.FieldByName("descriptorCount").Uint())

			return nil
		})

	err = device.UpdateDescriptorSets(nil, []core.CopyDescriptorSetOptions{
		{
			Source:             srcDescriptor,
			SourceBinding:      3,
			SourceArrayElement: 5,

			Destination:             destDescriptor,
			DestinationBinding:      7,
			DestinationArrayElement: 11,

			Count: 13,
		},
	})

	require.NoError(t, err)
}

func TestVulkanDevice_UpdateDescriptorSets_FailureImageInfoAndBufferInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	buffer1 := mocks.EasyMockBuffer(ctrl)
	buffer2 := mocks.EasyMockBuffer(ctrl)
	sampler1 := mocks.EasyMockSampler(ctrl)
	sampler2 := mocks.EasyMockSampler(ctrl)
	imageView1 := mocks.EasyMockImageView(ctrl)
	imageView2 := mocks.EasyMockImageView(ctrl)

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
			BufferInfo: []core.DescriptorBufferInfo{
				{
					Buffer: buffer1,
					Offset: 7,
					Range:  11,
				},
				{
					Buffer: buffer2,
					Offset: 13,
					Range:  17,
				},
			},
			ImageInfo: []core.DescriptorImageInfo{
				{
					Sampler:     sampler1,
					ImageView:   imageView1,
					ImageLayout: common.LayoutDepthStencilAttachmentOptimal,
				},
				{
					Sampler:     sampler2,
					ImageView:   imageView2,
					ImageLayout: common.LayoutPreInitialized,
				},
			},
		},
	}, nil)

	require.EqualError(t, err, "a WriteDescriptorSetOptions may have one or more ImageInfo sources OR one or more BufferInfo sources, but not both")
}

func TestVulkanDevice_UpdateDescriptorSets_FailureImageInfoAndBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	bufferView1 := mocks.EasyMockBufferView(ctrl)
	bufferView2 := mocks.EasyMockBufferView(ctrl)
	sampler1 := mocks.EasyMockSampler(ctrl)
	sampler2 := mocks.EasyMockSampler(ctrl)
	imageView1 := mocks.EasyMockImageView(ctrl)
	imageView2 := mocks.EasyMockImageView(ctrl)

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
			ImageInfo: []core.DescriptorImageInfo{
				{
					Sampler:     sampler1,
					ImageView:   imageView1,
					ImageLayout: common.LayoutDepthStencilAttachmentOptimal,
				},
				{
					Sampler:     sampler2,
					ImageView:   imageView2,
					ImageLayout: common.LayoutPreInitialized,
				},
			},
			TexelBufferView: []core.BufferView{
				bufferView1, bufferView2,
			},
		},
	}, nil)

	require.EqualError(t, err, "a WriteDescriptorSetOptions may have one or more ImageInfo sources OR one or more TexelBufferView sources, but not both")
}

func TestVulkanDevice_UpdateDescriptorSets_FailureBufferInfoAndBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)
	buffer1 := mocks.EasyMockBuffer(ctrl)
	buffer2 := mocks.EasyMockBuffer(ctrl)
	bufferView1 := mocks.EasyMockBufferView(ctrl)
	bufferView2 := mocks.EasyMockBufferView(ctrl)

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
			BufferInfo: []core.DescriptorBufferInfo{
				{
					Buffer: buffer1,
					Offset: 7,
					Range:  11,
				},
				{
					Buffer: buffer2,
					Offset: 13,
					Range:  17,
				},
			},
			TexelBufferView: []core.BufferView{
				bufferView1, bufferView2,
			},
		},
	}, nil)

	require.EqualError(t, err, "a WriteDescriptorSetOptions may have one or more BufferInfo sources OR one or more TexelBufferView sources, but not both")
}

func TestVulkanDevice_UpdateDescriptorSets_FailureNoSource(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	destDescriptor := mocks.EasyMockDescriptorSet(ctrl)

	err = device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
		{
			Destination:             destDescriptor,
			DestinationBinding:      1,
			DestinationArrayElement: 3,
			DescriptorType:          common.DescriptorUniformBuffer,
		},
	}, nil)

	require.EqualError(t, err, "a WriteDescriptorSetOptions must have a source to write the descriptor from: ImageInfo, BufferInfo, TexelBufferView, or an extension source")
}
