package impl1_0_test

import (
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateDevice_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := mocks1_0.EasyMockInstance(ctrl, mockDriver)

	builder := &impl1_0.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(mockDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_0, common.Vulkan1_0)

	deviceHandle := mocks.NewFakeDeviceHandle()

	mockDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(mockDriver, nil)
	mockDriver.EXPECT().VkCreateDevice(physicalDevice.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice, pCreateInfo *driver.VkDeviceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDevice *driver.VkDevice) (common.VkResult, error) {
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

			extensionNamePtr := (**driver.Char)(unsafe.Pointer(v.FieldByName("ppEnabledExtensionNames").Elem().UnsafeAddr()))
			extensionNameSlice := ([]*driver.Char)(unsafe.Slice(extensionNamePtr, 3))

			var extensionNames []string
			for _, extensionNameBytes := range extensionNameSlice {
				var extensionNameRunes []rune
				extensionNameByteSlice := ([]driver.Char)(unsafe.Slice(extensionNameBytes, 1<<30))
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

			queueCreateInfoPtr := (*driver.VkDeviceQueueCreateInfo)(unsafe.Pointer(v.FieldByName("pQueueCreateInfos").Elem().UnsafeAddr()))
			queueCreateInfoSlice := ([]driver.VkDeviceQueueCreateInfo)(unsafe.Slice(queueCreateInfoPtr, 2))

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
			return core1_0.VKSuccess, nil
		})

	device, _, err := physicalDevice.CreateDevice(nil, core1_0.DeviceCreateInfo{
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
	require.Equal(t, deviceHandle, device.Handle())
}

func TestVulkanLoader1_0_CreateDevice_FailNoQueueFamilies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := mocks1_0.EasyMockInstance(ctrl, mockDriver)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(mockDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_0, common.Vulkan1_0)

	_, _, err := physicalDevice.CreateDevice(nil, core1_0.DeviceCreateInfo{
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := mocks1_0.EasyMockInstance(ctrl, mockDriver)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	physicalDevice := builder.CreatePhysicalDeviceObject(mockDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_0, common.Vulkan1_0)

	_, _, err := physicalDevice.CreateDevice(nil, core1_0.DeviceCreateInfo{
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	queueHandle := mocks.NewFakeQueue()

	mockDriver.EXPECT().VkGetDeviceQueue(device.Handle(), driver.Uint32(1), driver.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, queueFamilyIndex, queueIndex driver.Uint32, pQueue *driver.VkQueue) {
			*pQueue = queueHandle
		})

	queue := device.GetQueue(1, 2)
	require.NotNil(t, queue)
	require.Equal(t, queueHandle, queue.Handle())
}

func TestDevice_WaitForFences_Timeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	fence1 := mocks1_0.EasyMockFence(ctrl)
	fence2 := mocks1_0.EasyMockFence(ctrl)

	mockDriver.EXPECT().VkWaitForFences(device.Handle(), driver.Uint32(2), gomock.Not(nil), driver.VkBool32(1), driver.Uint64(1)).DoAndReturn(
		func(device driver.VkDevice, fenceCount driver.Uint32, pFences *driver.VkFence, waitAll driver.VkBool32, timeout driver.Uint64) (common.VkResult, error) {
			fenceSlice := ([]driver.VkFence)(unsafe.Slice(pFences, 2))
			require.Equal(t, fence1.Handle(), fenceSlice[0])
			require.Equal(t, fence2.Handle(), fenceSlice[1])

			return core1_0.VKSuccess, nil
		})

	_, err := device.WaitForFences(true, time.Nanosecond, []core1_0.Fence{
		fence1, fence2,
	})
	require.NoError(t, err)
}

func TestDevice_WaitForFences_NoTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	fence1 := mocks1_0.EasyMockFence(ctrl)

	mockDriver.EXPECT().VkWaitForFences(device.Handle(), driver.Uint32(1), gomock.Not(nil), driver.VkBool32(0), driver.Uint64(0xffffffffffffffff)).DoAndReturn(
		func(device driver.VkDevice, fenceCount driver.Uint32, pFences *driver.VkFence, waitAll driver.VkBool32, timeout driver.Uint64) (common.VkResult, error) {
			fenceSlice := ([]driver.VkFence)(unsafe.Slice(pFences, 1))
			require.Equal(t, fence1.Handle(), fenceSlice[0])

			return core1_0.VKSuccess, nil
		})

	_, err := device.WaitForFences(false, common.NoTimeout, []core1_0.Fence{
		fence1,
	})
	require.NoError(t, err)
}

func TestDevice_WaitForIdle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})

	mockDriver.EXPECT().VkDeviceWaitIdle(device.Handle()).Return(core1_0.VKSuccess, nil)
	_, err := device.WaitIdle()
	require.NoError(t, err)
}

func TestVulkanDevice_ResetFences(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	fence1 := mocks1_0.EasyMockFence(ctrl)
	fence2 := mocks1_0.EasyMockFence(ctrl)

	mockDriver.EXPECT().VkResetFences(device.Handle(), driver.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, fenceCount driver.Uint32, pFence *driver.VkFence) (common.VkResult, error) {
			fences := ([]driver.VkFence)(unsafe.Slice(pFence, 2))

			require.Equal(t, fence1.Handle(), fences[0])
			require.Equal(t, fence2.Handle(), fences[1])
			return core1_0.VKSuccess, nil
		})

	_, err := device.ResetFences([]core1_0.Fence{fence1, fence2})
	require.NoError(t, err)
}

func TestVulkanDevice_UpdateDescriptorSets_WriteImageInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	destDescriptor := mocks1_0.EasyMockDescriptorSet(ctrl)
	sampler1 := mocks1_0.EasyMockSampler(ctrl)
	sampler2 := mocks1_0.EasyMockSampler(ctrl)
	imageView1 := mocks1_0.EasyMockImageView(ctrl)
	imageView2 := mocks1_0.EasyMockImageView(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), driver.Uint32(1), gomock.Not(nil), driver.Uint32(0), nil).DoAndReturn(
		func(device driver.VkDevice, descriptorWriteCount driver.Uint32, pDescriptorWrites *driver.VkWriteDescriptorSet, descriptorCopyCount driver.Uint32, pDescriptorCopies *driver.VkCopyDescriptorSet) error {
			writeSlice := ([]driver.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), driver.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pBufferInfo").IsNil())
			require.True(t, writeVal.FieldByName("pTexelBufferView").IsNil())

			imageInfoPtr := (*driver.VkDescriptorImageInfo)(unsafe.Pointer(writeVal.FieldByName("pImageInfo").Elem().UnsafeAddr()))
			imageInfoSlice := ([]driver.VkDescriptorImageInfo)(unsafe.Slice(imageInfoPtr, 2))

			require.Len(t, imageInfoSlice, 2)

			imageInfoVal := reflect.ValueOf(imageInfoSlice[0])
			require.Equal(t, sampler1.Handle(), (driver.VkSampler)(unsafe.Pointer(imageInfoVal.FieldByName("sampler").Elem().UnsafeAddr())))
			require.Equal(t, imageView1.Handle(), (driver.VkImageView)(unsafe.Pointer(imageInfoVal.FieldByName("imageView").Elem().UnsafeAddr())))
			require.Equal(t, uint64(3), imageInfoVal.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL

			imageInfoVal = reflect.ValueOf(imageInfoSlice[1])
			require.Equal(t, sampler2.Handle(), (driver.VkSampler)(unsafe.Pointer(imageInfoVal.FieldByName("sampler").Elem().UnsafeAddr())))
			require.Equal(t, imageView2.Handle(), (driver.VkImageView)(unsafe.Pointer(imageInfoVal.FieldByName("imageView").Elem().UnsafeAddr())))
			require.Equal(t, uint64(8), imageInfoVal.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_PREINITIALIZED

			return nil
		})

	err := device.UpdateDescriptorSets([]core1_0.WriteDescriptorSet{
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	destDescriptor := mocks1_0.EasyMockDescriptorSet(ctrl)
	buffer1 := mocks1_0.EasyMockBuffer(ctrl)
	buffer2 := mocks1_0.EasyMockBuffer(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), driver.Uint32(1), gomock.Not(nil), driver.Uint32(0), nil).DoAndReturn(
		func(device driver.VkDevice, descriptorWriteCount driver.Uint32, pDescriptorWrites *driver.VkWriteDescriptorSet, descriptorCopyCount driver.Uint32, pDescriptorCopies *driver.VkCopyDescriptorSet) error {
			writeSlice := ([]driver.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), driver.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(3), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pImageInfo").IsNil())
			require.True(t, writeVal.FieldByName("pTexelBufferView").IsNil())

			bufferInfoPtr := (*driver.VkDescriptorBufferInfo)(unsafe.Pointer(writeVal.FieldByName("pBufferInfo").Elem().UnsafeAddr()))
			bufferInfoSlice := ([]driver.VkDescriptorBufferInfo)(unsafe.Slice(bufferInfoPtr, 2))

			require.Len(t, bufferInfoSlice, 2)

			bufferInfoVal := reflect.ValueOf(bufferInfoSlice[0])
			require.Equal(t, buffer1.Handle(), (driver.VkBuffer)(unsafe.Pointer(bufferInfoVal.FieldByName("buffer").Elem().UnsafeAddr())))
			require.Equal(t, uint64(7), bufferInfoVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(11), bufferInfoVal.FieldByName("_range").Uint())

			bufferInfoVal = reflect.ValueOf(bufferInfoSlice[1])
			require.Equal(t, buffer2.Handle(), (driver.VkBuffer)(unsafe.Pointer(bufferInfoVal.FieldByName("buffer").Elem().UnsafeAddr())))
			require.Equal(t, uint64(13), bufferInfoVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(17), bufferInfoVal.FieldByName("_range").Uint())

			return nil
		})

	err := device.UpdateDescriptorSets([]core1_0.WriteDescriptorSet{
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	destDescriptor := mocks1_0.EasyMockDescriptorSet(ctrl)
	bufferView1 := mocks1_0.EasyMockBufferView(ctrl)
	bufferView2 := mocks1_0.EasyMockBufferView(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), driver.Uint32(1), gomock.Not(nil), driver.Uint32(0), nil).DoAndReturn(
		func(device driver.VkDevice, descriptorWriteCount driver.Uint32, pDescriptorWrites *driver.VkWriteDescriptorSet, descriptorCopyCount driver.Uint32, pDescriptorCopies *driver.VkCopyDescriptorSet) error {
			writeSlice := ([]driver.VkWriteDescriptorSet)(unsafe.Slice(pDescriptorWrites, 1))
			writeVal := reflect.ValueOf(writeSlice[0])

			require.Equal(t, uint64(35), writeVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
			require.True(t, writeVal.FieldByName("pNext").IsNil())
			require.Equal(t, destDescriptor.Handle(), driver.VkDescriptorSet(unsafe.Pointer(writeVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), writeVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(3), writeVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(2), writeVal.FieldByName("descriptorCount").Uint())
			require.Equal(t, uint64(6), writeVal.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
			require.True(t, writeVal.FieldByName("pImageInfo").IsNil())
			require.True(t, writeVal.FieldByName("pBufferInfo").IsNil())

			bufferInfoPtr := (*driver.VkBufferView)(unsafe.Pointer(writeVal.FieldByName("pTexelBufferView").Elem().UnsafeAddr()))
			bufferInfoSlice := ([]driver.VkBufferView)(unsafe.Slice(bufferInfoPtr, 2))

			require.Len(t, bufferInfoSlice, 2)

			require.Equal(t, bufferView1.Handle(), (driver.VkBufferView)(unsafe.Pointer(bufferInfoSlice[0])))
			require.Equal(t, bufferView2.Handle(), (driver.VkBufferView)(unsafe.Pointer(bufferInfoSlice[1])))

			return nil
		})

	err := device.UpdateDescriptorSets([]core1_0.WriteDescriptorSet{
		{
			DstSet:          destDescriptor,
			DstBinding:      1,
			DstArrayElement: 3,
			DescriptorType:  core1_0.DescriptorTypeUniformBuffer,
			TexelBufferView: []core1_0.BufferView{
				bufferView1, bufferView2,
			},
		},
	}, nil)

	require.NoError(t, err)
}
func TestVulkanDevice_UpdateDescriptorSets_Copy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	srcDescriptor := mocks1_0.EasyMockDescriptorSet(ctrl)
	destDescriptor := mocks1_0.EasyMockDescriptorSet(ctrl)

	mockDriver.EXPECT().VkUpdateDescriptorSets(device.Handle(), driver.Uint32(0), nil, driver.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, descriptorWriteCount driver.Uint32, pDescriptorWrites *driver.VkWriteDescriptorSet, descriptorCopyCount driver.Uint32, pDescriptorCopies *driver.VkCopyDescriptorSet) error {
			copySlice := ([]driver.VkCopyDescriptorSet)(unsafe.Slice(pDescriptorCopies, 1))
			copyVal := reflect.ValueOf(copySlice[0])

			require.Equal(t, uint64(36), copyVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_DESCRIPTOR_SET
			require.True(t, copyVal.FieldByName("pNext").IsNil())
			require.Equal(t, srcDescriptor.Handle(), (driver.VkDescriptorSet)(unsafe.Pointer(copyVal.FieldByName("srcSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(3), copyVal.FieldByName("srcBinding").Uint())
			require.Equal(t, uint64(5), copyVal.FieldByName("srcArrayElement").Uint())
			require.Equal(t, destDescriptor.Handle(), (driver.VkDescriptorSet)(unsafe.Pointer(copyVal.FieldByName("dstSet").Elem().UnsafeAddr())))
			require.Equal(t, uint64(7), copyVal.FieldByName("dstBinding").Uint())
			require.Equal(t, uint64(11), copyVal.FieldByName("dstArrayElement").Uint())
			require.Equal(t, uint64(13), copyVal.FieldByName("descriptorCount").Uint())

			return nil
		})

	err := device.UpdateDescriptorSets(nil, []core1_0.CopyDescriptorSet{
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	destDescriptor := mocks1_0.EasyMockDescriptorSet(ctrl)
	buffer1 := mocks1_0.EasyMockBuffer(ctrl)
	buffer2 := mocks1_0.EasyMockBuffer(ctrl)
	sampler1 := mocks1_0.EasyMockSampler(ctrl)
	sampler2 := mocks1_0.EasyMockSampler(ctrl)
	imageView1 := mocks1_0.EasyMockImageView(ctrl)
	imageView2 := mocks1_0.EasyMockImageView(ctrl)

	err := device.UpdateDescriptorSets([]core1_0.WriteDescriptorSet{
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	destDescriptor := mocks1_0.EasyMockDescriptorSet(ctrl)
	bufferView1 := mocks1_0.EasyMockBufferView(ctrl)
	bufferView2 := mocks1_0.EasyMockBufferView(ctrl)
	sampler1 := mocks1_0.EasyMockSampler(ctrl)
	sampler2 := mocks1_0.EasyMockSampler(ctrl)
	imageView1 := mocks1_0.EasyMockImageView(ctrl)
	imageView2 := mocks1_0.EasyMockImageView(ctrl)

	err := device.UpdateDescriptorSets([]core1_0.WriteDescriptorSet{
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
			TexelBufferView: []core1_0.BufferView{
				bufferView1, bufferView2,
			},
		},
	}, nil)

	require.EqualError(t, err, "a WriteDescriptorSet may have one or more ImageInfo sources OR one or more TexelBufferView sources, but not both")
}

func TestVulkanDevice_UpdateDescriptorSets_FailureBufferInfoAndBufferView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	destDescriptor := mocks1_0.EasyMockDescriptorSet(ctrl)
	buffer1 := mocks1_0.EasyMockBuffer(ctrl)
	buffer2 := mocks1_0.EasyMockBuffer(ctrl)
	bufferView1 := mocks1_0.EasyMockBufferView(ctrl)
	bufferView2 := mocks1_0.EasyMockBufferView(ctrl)

	err := device.UpdateDescriptorSets([]core1_0.WriteDescriptorSet{
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
			TexelBufferView: []core1_0.BufferView{
				bufferView1, bufferView2,
			},
		},
	}, nil)

	require.EqualError(t, err, "a WriteDescriptorSet may have one or more BufferInfo sources OR one or more TexelBufferView sources, but not both")
}

func TestVulkanDevice_UpdateDescriptorSets_FailureNoSource(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	destDescriptor := mocks1_0.EasyMockDescriptorSet(ctrl)

	err := device.UpdateDescriptorSets([]core1_0.WriteDescriptorSet{
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

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	mem1 := mocks1_0.EasyMockDeviceMemory(ctrl)
	mem2 := mocks1_0.EasyMockDeviceMemory(ctrl)

	mockDriver.EXPECT().VkFlushMappedMemoryRanges(device.Handle(), driver.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, rangeCount driver.Uint32, pRanges *driver.VkMappedMemoryRange) (common.VkResult, error) {
			val := reflect.ValueOf([]driver.VkMappedMemoryRange(unsafe.Slice(pRanges, 2)))

			r := val.Index(0)
			require.Equal(t, uint64(6), r.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, r.FieldByName("pNext").IsNil())
			require.Equal(t, mem1.Handle(), driver.VkDeviceMemory(unsafe.Pointer(r.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), r.FieldByName("offset").Uint())
			require.Equal(t, uint64(3), r.FieldByName("size").Uint())

			r = val.Index(1)
			require.Equal(t, uint64(6), r.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, r.FieldByName("pNext").IsNil())
			require.Equal(t, mem2.Handle(), driver.VkDeviceMemory(unsafe.Pointer(r.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(5), r.FieldByName("offset").Uint())
			require.Equal(t, uint64(7), r.FieldByName("size").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := device.FlushMappedMemoryRanges([]core1_0.MappedMemoryRange{
		{
			Memory: mem1,
			Offset: 1,
			Size:   3,
		},
		{
			Memory: mem2,
			Offset: 5,
			Size:   7,
		},
	})
	require.NoError(t, err)
}

func TestVulkanDevice_InvalidateMappedMemoryRanges(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	mem1 := mocks1_0.EasyMockDeviceMemory(ctrl)
	mem2 := mocks1_0.EasyMockDeviceMemory(ctrl)

	mockDriver.EXPECT().VkInvalidateMappedMemoryRanges(device.Handle(), driver.Uint32(2), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, rangeCount driver.Uint32, pRanges *driver.VkMappedMemoryRange) (common.VkResult, error) {
			val := reflect.ValueOf([]driver.VkMappedMemoryRange(unsafe.Slice(pRanges, 2)))

			r := val.Index(0)
			require.Equal(t, uint64(6), r.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, r.FieldByName("pNext").IsNil())
			require.Equal(t, mem1.Handle(), driver.VkDeviceMemory(unsafe.Pointer(r.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(1), r.FieldByName("offset").Uint())
			require.Equal(t, uint64(3), r.FieldByName("size").Uint())

			r = val.Index(1)
			require.Equal(t, uint64(6), r.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
			require.True(t, r.FieldByName("pNext").IsNil())
			require.Equal(t, mem2.Handle(), driver.VkDeviceMemory(unsafe.Pointer(r.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(5), r.FieldByName("offset").Uint())
			require.Equal(t, uint64(7), r.FieldByName("size").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := device.InvalidateMappedMemoryRanges([]core1_0.MappedMemoryRange{
		{
			Memory: mem1,
			Offset: 1,
			Size:   3,
		},
		{
			Memory: mem2,
			Offset: 5,
			Size:   7,
		},
	})
	require.NoError(t, err)
}
