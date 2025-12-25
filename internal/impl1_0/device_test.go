package impl1_0_test

import (
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateDevice_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	expectedDevice := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	mockLoader.EXPECT().VkCreateDevice(physicalDevice.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice, pCreateInfo *loader.VkDeviceCreateInfo, pAllocator *loader.VkAllocationCallbacks, pDevice *loader.VkDevice) (common.VkResult, error) {
			v := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(3), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
			require.True(t, v.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), v.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), v.FieldByName("queueCreateInfoCount").Uint())
			require.Equal(t, uint64(3), v.FieldByName("enabledExtensionCount").Uint())
			require.Equal(t, uint64(0), v.FieldByName("enabledLayerCount").Uint())

			featuresV := v.FieldByName("pEnabledFeatures").Elem()

			require.Equal(t, uint64(1), featuresV.FieldByName("occlusionQueryPrecise").Uint())
			require.Equal(t, uint64(1), featuresV.FieldByName("tessellationShader").Uint())
			require.Equal(t, uint64(0), featuresV.FieldByName("samplerAnisotropy").Uint())

			extensionNamePtr := (**loader.Char)(unsafe.Pointer(v.FieldByName("ppEnabledExtensionNames").Elem().UnsafeAddr()))
			extensionNameSlice := ([]*loader.Char)(unsafe.Slice(extensionNamePtr, 3))

			var extensionNames []string
			for _, extensionNameBytes := range extensionNameSlice {
				var extensionNameRunes []rune
				extensionNameByteSlice := ([]loader.Char)(unsafe.Slice(extensionNameBytes, 1<<30))
				for _, nameByte := range extensionNameByteSlice {
					if nameByte == 0 {
						break
					}

					extensionNameRunes = append(extensionNameRunes, rune(nameByte))
				}

				extensionNames = append(extensionNames, string(extensionNameRunes))
			}

			require.ElementsMatch(t, []string{"A", "B", "C"}, extensionNames)

			require.True(t, v.FieldByName("ppEnabledLayerNames").IsNil())

			queueCreateInfoPtr := (*loader.VkDeviceQueueCreateInfo)(unsafe.Pointer(v.FieldByName("pQueueCreateInfos").Elem().UnsafeAddr()))
			queueCreateInfoSlice := ([]loader.VkDeviceQueueCreateInfo)(unsafe.Slice(queueCreateInfoPtr, 2))

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

			*pDevice = expectedDevice.Handle()
			return core1_0.VKSuccess, nil
		})

	device, _, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueueFamilyIndex: 1,
				QueuePriorities:  []float32{1, 0, 0.5},
			},
			{
				QueueFamilyIndex: 3,
				QueuePriorities:  []float32{0.5},
			},
		},
		EnabledExtensionNames: []string{"A", "B", "C"},
		EnabledFeatures: &core1_0.PhysicalDeviceFeatures{
			OcclusionQueryPrecise: true,
			TessellationShader:    true,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Handle())
}

func TestVulkanLoader1_0_CreateDevice_FailNoQueueFamilies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	_, _, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos:      []core1_0.DeviceQueueCreateInfo{},
		EnabledExtensionNames: []string{"A", "B", "C"},
		EnabledFeatures: &core1_0.PhysicalDeviceFeatures{
			OcclusionQueryPrecise: true,
			TessellationShader:    true,
		},
	})
	require.EqualError(t, err, "alloc DeviceCreateInfo: no queue families added")
}

func TestVulkanLoader1_0_CreateDevice_FailFamilyWithoutPriorities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewInstanceDriver(mockLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	_, _, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueueFamilyIndex: 1,
				QueuePriorities:  []float32{1, 0, 0.5},
			},
			{
				QueueFamilyIndex: 3,
				QueuePriorities:  []float32{},
			},
		},
		EnabledExtensionNames: []string{"A", "B", "C"},
		EnabledFeatures: &core1_0.PhysicalDeviceFeatures{
			OcclusionQueryPrecise: true,
			TessellationShader:    true,
		},
	})
	require.EqualError(t, err, "alloc DeviceCreateInfo: queue family 3 had no queue priorities")
}

func TestDevice_GetQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	expectedQueue := mocks.NewDummyQueue(device)

	mockLoader.EXPECT().VkGetDeviceQueue(device.Handle(), loader.Uint32(1), loader.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, queueFamilyIndex, queueIndex loader.Uint32, pQueue *loader.VkQueue) {
			*pQueue = expectedQueue.Handle()
		})

	queue := driver.GetQueue(device, 1, 2)
	require.NotNil(t, queue)
	require.Equal(t, expectedQueue.Handle(), queue.Handle())
}

func TestDevice_WaitForFences_Timeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	mockLoader.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	driver := impl1_0.NewDeviceDriver(mockLoader)
	fence1 := mocks.NewDummyFence(device)
	fence2 := mocks.NewDummyFence(device)

	mockLoader.EXPECT().VkWaitForFences(device.Handle(), loader.Uint32(2), gomock.Not(nil), loader.VkBool32(1), loader.Uint64(1)).DoAndReturn(
		func(device loader.VkDevice, fenceCount loader.Uint32, pFences *loader.VkFence, waitAll loader.VkBool32, timeout loader.Uint64) (common.VkResult, error) {
			fenceSlice := ([]loader.VkFence)(unsafe.Slice(pFences, 2))
			require.Equal(t, fence1.Handle(), fenceSlice[0])
			require.Equal(t, fence2.Handle(), fenceSlice[1])

			return core1_0.VKSuccess, nil
		})

	_, err := driver.WaitForFences(true, time.Nanosecond, fence1, fence2)
	require.NoError(t, err)
}

func TestDevice_WaitForFences_NoTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	mockLoader.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	driver := impl1_0.NewDeviceDriver(mockLoader)
	fence1 := mocks.NewDummyFence(device)

	mockLoader.EXPECT().VkWaitForFences(device.Handle(), loader.Uint32(1), gomock.Not(nil), loader.VkBool32(0), loader.Uint64(0xffffffffffffffff)).DoAndReturn(
		func(device loader.VkDevice, fenceCount loader.Uint32, pFences *loader.VkFence, waitAll loader.VkBool32, timeout loader.Uint64) (common.VkResult, error) {
			fenceSlice := ([]loader.VkFence)(unsafe.Slice(pFences, 1))
			require.Equal(t, fence1.Handle(), fenceSlice[0])

			return core1_0.VKSuccess, nil
		})

	_, err := driver.WaitForFences(false, common.NoTimeout, fence1)
	require.NoError(t, err)
}

func TestDevice_WaitForIdle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	mockLoader.EXPECT().VkDeviceWaitIdle(device.Handle()).Return(core1_0.VKSuccess, nil)
	_, err := driver.DeviceWaitIdle(device)
	require.NoError(t, err)
}

func TestVulkanDevice_ResetFences(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	mockLoader.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	driver := impl1_0.NewDeviceDriver(mockLoader)
	fence1 := mocks.NewDummyFence(device)
	fence2 := mocks.NewDummyFence(device)

	mockLoader.EXPECT().VkResetFences(device.Handle(), loader.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, fenceCount loader.Uint32, pFence *loader.VkFence) (common.VkResult, error) {
			fences := ([]loader.VkFence)(unsafe.Slice(pFence, 2))

			require.Equal(t, fence1.Handle(), fences[0])
			require.Equal(t, fence2.Handle(), fences[1])
			return core1_0.VKSuccess, nil
		})

	_, err := driver.ResetFences(fence1, fence2)
	require.NoError(t, err)
}

func TestVulkanDevice_UpdateDescriptorSets_WriteImageInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	destDescriptor := mocks.NewDummyDescriptorSet(pool, device)
	sampler1 := mocks.NewDummySampler(device)
	sampler2 := mocks.NewDummySampler(device)
	imageView1 := mocks.NewDummyImageView(device)
	imageView2 := mocks.NewDummyImageView(device)

	mockLoader.EXPECT().VkUpdateDescriptorSets(device.Handle(), loader.Uint32(1), gomock.Not(nil), loader.Uint32(0), nil).DoAndReturn(
		func(device loader.VkDevice, descriptorWriteCount loader.Uint32, pDescriptorWrites *loader.VkWriteDescriptorSet, descriptorCopyCount loader.Uint32, pDescriptorCopies *loader.VkCopyDescriptorSet) error {
			writeSlice := ([]loader.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), loader.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pBufferInfo").IsNil())
			require.True(t, writeVal.FieldByName("pTexelBufferView").IsNil())

			imageInfoPtr := (*loader.VkDescriptorImageInfo)(unsafe.Pointer(writeVal.FieldByName("pImageInfo").Elem().UnsafeAddr()))
			imageInfoSlice := ([]loader.VkDescriptorImageInfo)(unsafe.Slice(imageInfoPtr, 2))

			require.Len(t, imageInfoSlice, 2)

			imageInfoVal := reflect.ValueOf(imageInfoSlice[0])
			require.Equal(t, sampler1.Handle(), (loader.VkSampler)(unsafe.Pointer(imageInfoVal.FieldByName("sampler").Elem().UnsafeAddr())))
			require.Equal(t, imageView1.Handle(), (loader.VkImageView)(unsafe.Pointer(imageInfoVal.FieldByName("imageView").Elem().UnsafeAddr())))
			require.Equal(t, uint64(3), imageInfoVal.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL

			imageInfoVal = reflect.ValueOf(imageInfoSlice[1])
			require.Equal(t, sampler2.Handle(), (loader.VkSampler)(unsafe.Pointer(imageInfoVal.FieldByName("sampler").Elem().UnsafeAddr())))
			require.Equal(t, imageView2.Handle(), (loader.VkImageView)(unsafe.Pointer(imageInfoVal.FieldByName("imageView").Elem().UnsafeAddr())))
			require.Equal(t, uint64(8), imageInfoVal.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_PREINITIALIZED

			return nil
		})

	err := driver.UpdateDescriptorSets(device, []core1_0.WriteDescriptorSet{
		{
			DstSet:          destDescriptor,
			DstBinding:      1,
			DstArrayElement: 2,
			DescriptorType:  core1_0.DescriptorTypeUniformBuffer,
			ImageInfo: []core1_0.DescriptorImageInfo{
				{
					Sampler:     sampler1,
					ImageView:   imageView1,
					ImageLayout: core1_0.ImageLayoutDepthStencilAttachmentOptimal,
				},
				{
					Sampler:     sampler2,
					ImageView:   imageView2,
					ImageLayout: core1_0.ImageLayoutPreInitialized,
				},
			},
		},
	}, nil)

	require.NoError(t, err)
}
func TestVulkanDevice_UpdateDescriptorSets_WriteBufferInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	destDescriptor := mocks.NewDummyDescriptorSet(pool, device)
	buffer1 := mocks.NewDummyBuffer(device)
	buffer2 := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkUpdateDescriptorSets(device.Handle(), loader.Uint32(1), gomock.Not(nil), loader.Uint32(0), nil).DoAndReturn(
		func(device loader.VkDevice, descriptorWriteCount loader.Uint32, pDescriptorWrites *loader.VkWriteDescriptorSet, descriptorCopyCount loader.Uint32, pDescriptorCopies *loader.VkCopyDescriptorSet) error {
			writeSlice := ([]loader.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), loader.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(3), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pImageInfo").IsNil())
			require.True(t, writeVal.FieldByName("pTexelBufferView").IsNil())

			bufferInfoPtr := (*loader.VkDescriptorBufferInfo)(unsafe.Pointer(writeVal.FieldByName("pBufferInfo").Elem().UnsafeAddr()))
			bufferInfoSlice := ([]loader.VkDescriptorBufferInfo)(unsafe.Slice(bufferInfoPtr, 2))

			require.Len(t, bufferInfoSlice, 2)

			bufferInfoVal := reflect.ValueOf(bufferInfoSlice[0])
			require.Equal(t, buffer1.Handle(), (loader.VkBuffer)(unsafe.Pointer(bufferInfoVal.FieldByName("buffer").Elem().UnsafeAddr())))
			require.Equal(t, uint64(7), bufferInfoVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(11), bufferInfoVal.FieldByName("_range").Uint())

			bufferInfoVal = reflect.ValueOf(bufferInfoSlice[1])
			require.Equal(t, buffer2.Handle(), (loader.VkBuffer)(unsafe.Pointer(bufferInfoVal.FieldByName("buffer").Elem().UnsafeAddr())))
			require.Equal(t, uint64(13), bufferInfoVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(17), bufferInfoVal.FieldByName("_range").Uint())

			return nil
		})

	err := driver.UpdateDescriptorSets(device, []core1_0.WriteDescriptorSet{
		{
			DstSet:          destDescriptor,
			DstBinding:      1,
			DstArrayElement: 3,
			DescriptorType:  core1_0.DescriptorTypeUniformBuffer,
			BufferInfo: []core1_0.DescriptorBufferInfo{
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

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	destDescriptor := mocks.NewDummyDescriptorSet(pool, device)
	bufferView1 := mocks.NewDummyBufferView(device)
	bufferView2 := mocks.NewDummyBufferView(device)

	mockLoader.EXPECT().VkUpdateDescriptorSets(device.Handle(), loader.Uint32(1), gomock.Not(nil), loader.Uint32(0), nil).DoAndReturn(
		func(device loader.VkDevice, descriptorWriteCount loader.Uint32, pDescriptorWrites *loader.VkWriteDescriptorSet, descriptorCopyCount loader.Uint32, pDescriptorCopies *loader.VkCopyDescriptorSet) error {
			writeSlice := ([]loader.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), loader.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(3), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pImageInfo").IsNil())
			require.True(t, writeVal.FieldByName("pBufferInfo").IsNil())

			bufferInfoPtr := (*loader.VkBufferView)(unsafe.Pointer(writeVal.FieldByName("pTexelBufferView").Elem().UnsafeAddr()))
			bufferInfoSlice := ([]loader.VkBufferView)(unsafe.Slice(bufferInfoPtr, 2))

			require.Len(t, bufferInfoSlice, 2)

			require.Equal(t, bufferView1.Handle(), (loader.VkBufferView)(unsafe.Pointer(bufferInfoSlice[0])))
			require.Equal(t, bufferView2.Handle(), (loader.VkBufferView)(unsafe.Pointer(bufferInfoSlice[1])))

			return nil
		})

	err := driver.UpdateDescriptorSets(device, []core1_0.WriteDescriptorSet{
		{
			DstSet:          destDescriptor,
			DstBinding:      1,
			DstArrayElement: 3,
			DescriptorType:  core1_0.DescriptorTypeUniformBuffer,
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

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	srcDescriptor := mocks.NewDummyDescriptorSet(pool, device)
	destDescriptor := mocks.NewDummyDescriptorSet(pool, device)

	mockLoader.EXPECT().VkUpdateDescriptorSets(device.Handle(), loader.Uint32(0), nil, loader.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, descriptorWriteCount loader.Uint32, pDescriptorWrites *loader.VkWriteDescriptorSet, descriptorCopyCount loader.Uint32, pDescriptorCopies *loader.VkCopyDescriptorSet) error {
			copySlice := ([]loader.VkCopyDescriptorSet)(unsafe.Slice(pDescriptorCopies, 1))
			copyVal := reflect.ValueOf(copySlice[0])

			require.Equal(t, uint64(36), copyVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_DESCRIPTOR_SET
			require.True(t, copyVal.FieldByName("pNext").IsNil())
			require.Equal(t, srcDescriptor.Handle(), (loader.VkDescriptorSet)(unsafe.Pointer(copyVal.FieldByName("srcSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(3), copyVal.FieldByName("srcBinding").Uint())
			require.Equal(t, uint64(5), copyVal.FieldByName("srcArrayElement").Uint())
			require.Equal(t, destDescriptor.Handle(), (loader.VkDescriptorSet)(unsafe.Pointer(copyVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(7), copyVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(11), copyVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(13), copyVal.FieldByName("descriptorCount").Uint())

			return nil
		})

	err := driver.UpdateDescriptorSets(device, nil, []core1_0.CopyDescriptorSet{
		{
			SrcSet:          srcDescriptor,
			SrcBinding:      3,
			SrcArrayElement: 5,

			DstSet:          destDescriptor,
			DstBinding:      7,
			DstArrayElement: 11,

			DescriptorCount: 13,
		},
	})

	require.NoError(t, err)
}

func TestVulkanDevice_UpdateDescriptorSets_FailureImageInfoAndBufferInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	destDescriptor := mocks.NewDummyDescriptorSet(pool, device)
	buffer1 := mocks.NewDummyBuffer(device)
	buffer2 := mocks.NewDummyBuffer(device)
	sampler1 := mocks.NewDummySampler(device)
	sampler2 := mocks.NewDummySampler(device)
	imageView1 := mocks.NewDummyImageView(device)
	imageView2 := mocks.NewDummyImageView(device)

	err := driver.UpdateDescriptorSets(device, []core1_0.WriteDescriptorSet{
		{
			DstSet:          destDescriptor,
			DstBinding:      1,
			DstArrayElement: 3,
			DescriptorType:  core1_0.DescriptorTypeUniformBuffer,
			BufferInfo: []core1_0.DescriptorBufferInfo{
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
			ImageInfo: []core1_0.DescriptorImageInfo{
				{
					Sampler:     sampler1,
					ImageView:   imageView1,
					ImageLayout: core1_0.ImageLayoutDepthStencilAttachmentOptimal,
				},
				{
					Sampler:     sampler2,
					ImageView:   imageView2,
					ImageLayout: core1_0.ImageLayoutPreInitialized,
				},
			},
		},
	}, nil)

	require.EqualError(t, err, "a WriteDescriptorSet may have one or more ImageInfo sources OR one or more BufferInfo sources, but not both")
}

func TestVulkanDevice_UpdateDescriptorSets_FailureImageInfoAndBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	destDescriptor := mocks.NewDummyDescriptorSet(pool, device)
	bufferView1 := mocks.NewDummyBufferView(device)
	bufferView2 := mocks.NewDummyBufferView(device)
	sampler1 := mocks.NewDummySampler(device)
	sampler2 := mocks.NewDummySampler(device)
	imageView1 := mocks.NewDummyImageView(device)
	imageView2 := mocks.NewDummyImageView(device)

	err := driver.UpdateDescriptorSets(device, []core1_0.WriteDescriptorSet{
		{
			DstSet:          destDescriptor,
			DstBinding:      1,
			DstArrayElement: 3,
			DescriptorType:  core1_0.DescriptorTypeUniformBuffer,
			ImageInfo: []core1_0.DescriptorImageInfo{
				{
					Sampler:     sampler1,
					ImageView:   imageView1,
					ImageLayout: core1_0.ImageLayoutDepthStencilAttachmentOptimal,
				},
				{
					Sampler:     sampler2,
					ImageView:   imageView2,
					ImageLayout: core1_0.ImageLayoutPreInitialized,
				},
			},
			TexelBufferView: []core.BufferView{
				bufferView1, bufferView2,
			},
		},
	}, nil)

	require.EqualError(t, err, "a WriteDescriptorSet may have one or more ImageInfo sources OR one or more TexelBufferView sources, but not both")
}

func TestVulkanDevice_UpdateDescriptorSets_FailureBufferInfoAndBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)

	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	destDescriptor := mocks.NewDummyDescriptorSet(pool, device)
	buffer1 := mocks.NewDummyBuffer(device)
	buffer2 := mocks.NewDummyBuffer(device)
	bufferView1 := mocks.NewDummyBufferView(device)
	bufferView2 := mocks.NewDummyBufferView(device)

	err := driver.UpdateDescriptorSets(device, []core1_0.WriteDescriptorSet{
		{
			DstSet:          destDescriptor,
			DstBinding:      1,
			DstArrayElement: 3,
			DescriptorType:  core1_0.DescriptorTypeUniformBuffer,
			BufferInfo: []core1_0.DescriptorBufferInfo{
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

	require.EqualError(t, err, "a WriteDescriptorSet may have one or more BufferInfo sources OR one or more TexelBufferView sources, but not both")
}

func TestVulkanDevice_UpdateDescriptorSets_FailureNoSource(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := impl1_0.NewDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyDescriptorPool(device)
	destDescriptor := mocks.NewDummyDescriptorSet(pool, device)

	err := driver.UpdateDescriptorSets(device, []core1_0.WriteDescriptorSet{
		{
			DstSet:          destDescriptor,
			DstBinding:      1,
			DstArrayElement: 3,
			DescriptorType:  core1_0.DescriptorTypeUniformBuffer,
		},
	}, nil)

	require.EqualError(t, err, "a WriteDescriptorSet must have a source to write the descriptor from: ImageInfo, BufferInfo, TexelBufferView, or an extension source")
}

func TestVulkanDevice_FlushMappedMemoryRanges(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	mockLoader.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	driver := impl1_0.NewDeviceDriver(mockLoader)
	mem1 := mocks.NewDummyDeviceMemory(device, 1)
	mem2 := mocks.NewDummyDeviceMemory(device, 1)

	mockLoader.EXPECT().VkFlushMappedMemoryRanges(device.Handle(), loader.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, rangeCount loader.Uint32, pRanges *loader.VkMappedMemoryRange) (common.VkResult, error) {
			val := reflect.ValueOf([]loader.VkMappedMemoryRange(unsafe.Slice(pRanges, 2)))

			r := val.Index(0)
			require.Equal(t, uint64(6), r.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, r.FieldByName("pNext").IsNil())
			require.Equal(t, mem1.Handle(), loader.VkDeviceMemory(unsafe.Pointer(r.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), r.FieldByName("offset").Uint())
			require.Equal(t, uint64(3), r.FieldByName("size").Uint())

			r = val.Index(1)
			require.Equal(t, uint64(6), r.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, r.FieldByName("pNext").IsNil())
			require.Equal(t, mem2.Handle(), loader.VkDeviceMemory(unsafe.Pointer(r.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(5), r.FieldByName("offset").Uint())
			require.Equal(t, uint64(7), r.FieldByName("size").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := driver.FlushMappedMemoryRanges(
		core1_0.MappedMemoryRange{
			Memory: mem1,
			Offset: 1,
			Size:   3,
		},
		core1_0.MappedMemoryRange{
			Memory: mem2,
			Offset: 5,
			Size:   7,
		},
	)
	require.NoError(t, err)
}

func TestVulkanDevice_InvalidateMappedMemoryRanges(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	mockLoader.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	driver := impl1_0.NewDeviceDriver(mockLoader)
	mem1 := mocks.NewDummyDeviceMemory(device, 1)
	mem2 := mocks.NewDummyDeviceMemory(device, 1)

	mockLoader.EXPECT().VkInvalidateMappedMemoryRanges(device.Handle(), loader.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, rangeCount loader.Uint32, pRanges *loader.VkMappedMemoryRange) (common.VkResult, error) {
			val := reflect.ValueOf([]loader.VkMappedMemoryRange(unsafe.Slice(pRanges, 2)))

			r := val.Index(0)
			require.Equal(t, uint64(6), r.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, r.FieldByName("pNext").IsNil())
			require.Equal(t, mem1.Handle(), loader.VkDeviceMemory(unsafe.Pointer(r.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), r.FieldByName("offset").Uint())
			require.Equal(t, uint64(3), r.FieldByName("size").Uint())

			r = val.Index(1)
			require.Equal(t, uint64(6), r.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, r.FieldByName("pNext").IsNil())
			require.Equal(t, mem2.Handle(), loader.VkDeviceMemory(unsafe.Pointer(r.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(5), r.FieldByName("offset").Uint())
			require.Equal(t, uint64(7), r.FieldByName("size").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := driver.InvalidateMappedMemoryRanges(
		core1_0.MappedMemoryRange{
			Memory: mem1,
			Offset: 1,
			Size:   3,
		},
		core1_0.MappedMemoryRange{
			Memory: mem2,
			Offset: 5,
			Size:   7,
		},
	)
	require.NoError(t, err)
}
