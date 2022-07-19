package core1_1_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
	mock_driver "github.com/vkngwrapper/core/driver/mocks"
	"github.com/vkngwrapper/core/internal/dummies"
	"github.com/vkngwrapper/core/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanDevice_DescriptorSetLayoutSupport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))

	coreDriver.EXPECT().VkGetDescriptorSetLayoutSupport(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pSupport *driver.VkDescriptorSetLayoutSupport) {
			optionVal := reflect.ValueOf(pCreateInfo).Elem()

			require.Equal(t, uint64(32), optionVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
			require.True(t, optionVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), optionVal.FieldByName("bindingCount").Uint())

			bindingPtr := (*driver.VkDescriptorSetLayoutBinding)(optionVal.FieldByName("pBindings").UnsafePointer())
			binding := reflect.ValueOf(bindingPtr).Elem()
			require.Equal(t, uint64(1), binding.FieldByName("binding").Uint())
			require.Equal(t, uint64(3), binding.FieldByName("descriptorCount").Uint())

			outDataVal := reflect.ValueOf(pSupport).Elem()

			require.Equal(t, uint64(1000168001), outDataVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
			require.True(t, outDataVal.FieldByName("pNext").IsNil())

			*(*driver.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("supported").UnsafeAddr())) = driver.VkBool32(1)
		})

	outData := &core1_1.DescriptorSetLayoutSupport{}
	err := device.DescriptorSetLayoutSupport(core1_0.DescriptorSetLayoutCreateInfo{
		Bindings: []core1_0.DescriptorSetLayoutBinding{
			{
				Binding:         1,
				DescriptorCount: 3,
			},
		},
	}, outData)
	require.NoError(t, err)
	require.True(t, outData.Supported)
}

func TestVulkanDevice_BindBufferMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))

	buffer1 := mocks.EasyMockBuffer(ctrl)
	buffer2 := mocks.EasyMockBuffer(ctrl)

	memory1 := mocks.EasyMockDeviceMemory(ctrl)
	memory2 := mocks.EasyMockDeviceMemory(ctrl)

	coreDriver.EXPECT().VkBindBufferMemory2(device.Handle(), driver.Uint32(2), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, bindInfoCount driver.Uint32, pBindInfos *driver.VkBindBufferMemoryInfo) (common.VkResult, error) {
			bindInfoSlice := unsafe.Slice(pBindInfos, 2)
			val := reflect.ValueOf(bindInfoSlice)

			bind := val.Index(0)
			require.Equal(t, uint64(1000157000), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, buffer1.Handle(), (driver.VkBuffer)(bind.FieldByName("buffer").UnsafePointer()))
			require.Equal(t, memory1.Handle(), (driver.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(1), bind.FieldByName("memoryOffset").Uint())

			bind = val.Index(1)
			require.Equal(t, uint64(1000157000), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, buffer2.Handle(), (driver.VkBuffer)(bind.FieldByName("buffer").UnsafePointer()))
			require.Equal(t, memory2.Handle(), (driver.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(3), bind.FieldByName("memoryOffset").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := device.BindBufferMemory2([]core1_1.BindBufferMemoryInfo{
		{
			Buffer:       buffer1,
			Memory:       memory1,
			MemoryOffset: 1,
		},
		{
			Buffer:       buffer2,
			Memory:       memory2,
			MemoryOffset: 3,
		},
	})
	require.NoError(t, err)
}

func TestVulkanDevice_BindImageMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))

	image1 := mocks.EasyMockImage(ctrl)
	image2 := mocks.EasyMockImage(ctrl)

	memory1 := mocks.EasyMockDeviceMemory(ctrl)
	memory2 := mocks.EasyMockDeviceMemory(ctrl)

	coreDriver.EXPECT().VkBindImageMemory2(device.Handle(), driver.Uint32(2), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, bindInfoCount driver.Uint32, pBindInfos *driver.VkBindImageMemoryInfo) (common.VkResult, error) {
			bindInfoSlice := unsafe.Slice(pBindInfos, 2)
			val := reflect.ValueOf(bindInfoSlice)

			bind := val.Index(0)
			require.Equal(t, uint64(1000157001), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, image1.Handle(), (driver.VkImage)(bind.FieldByName("image").UnsafePointer()))
			require.Equal(t, memory1.Handle(), (driver.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(1), bind.FieldByName("memoryOffset").Uint())

			bind = val.Index(1)
			require.Equal(t, uint64(1000157001), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, image2.Handle(), (driver.VkImage)(bind.FieldByName("image").UnsafePointer()))
			require.Equal(t, memory2.Handle(), (driver.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(3), bind.FieldByName("memoryOffset").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := device.BindImageMemory2([]core1_1.BindImageMemoryInfo{
		{
			Image:        image1,
			Memory:       memory1,
			MemoryOffset: 1,
		},
		{
			Image:        image2,
			Memory:       memory2,
			MemoryOffset: 3,
		},
	})
	require.NoError(t, err)
}

func TestBindBufferMemoryDeviceGroupOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))
	buffer := mocks.EasyMockBuffer(ctrl)
	memory := mocks.EasyMockDeviceMemory(ctrl)

	coreDriver.EXPECT().VkBindBufferMemory2(
		device.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		infoCount driver.Uint32,
		pInfo *driver.VkBindBufferMemoryInfo,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000157000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
		require.Equal(t, buffer.Handle(), (driver.VkBuffer)(val.FieldByName("buffer").UnsafePointer()))
		require.Equal(t, memory.Handle(), (driver.VkDeviceMemory)(val.FieldByName("memory").UnsafePointer()))
		require.Equal(t, uint64(1), val.FieldByName("memoryOffset").Uint())

		next := (*driver.VkBindBufferMemoryDeviceGroupInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060013), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceIndexCount").Uint())

		indices := (*driver.Uint32)(val.FieldByName("pDeviceIndices").UnsafePointer())
		indexSlice := ([]driver.Uint32)(unsafe.Slice(indices, 3))
		val = reflect.ValueOf(indexSlice)

		require.Equal(t, uint64(1), val.Index(0).Uint())
		require.Equal(t, uint64(2), val.Index(1).Uint())
		require.Equal(t, uint64(7), val.Index(2).Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := device.BindBufferMemory2([]core1_1.BindBufferMemoryInfo{
		{
			Buffer:       buffer,
			Memory:       memory,
			MemoryOffset: 1,

			NextOptions: common.NextOptions{
				core1_1.BindBufferMemoryDeviceGroupInfo{
					DeviceIndices: []int{1, 2, 7},
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestBindImageMemoryDeviceGroupOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))
	image := mocks.EasyMockImage(ctrl)
	memory := mocks.EasyMockDeviceMemory(ctrl)

	coreDriver.EXPECT().VkBindImageMemory2(
		device.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		infoCount driver.Uint32,
		pInfo *driver.VkBindImageMemoryInfo,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000157001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
		require.Equal(t, image.Handle(), (driver.VkImage)(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, memory.Handle(), (driver.VkDeviceMemory)(val.FieldByName("memory").UnsafePointer()))
		require.Equal(t, uint64(1), val.FieldByName("memoryOffset").Uint())

		next := (*driver.VkBindImageMemoryDeviceGroupInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060014), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceIndexCount").Uint())

		indices := (*driver.Uint32)(val.FieldByName("pDeviceIndices").UnsafePointer())
		indexSlice := ([]driver.Uint32)(unsafe.Slice(indices, 3))
		indexVal := reflect.ValueOf(indexSlice)

		require.Equal(t, uint64(1), indexVal.Index(0).Uint())
		require.Equal(t, uint64(2), indexVal.Index(1).Uint())
		require.Equal(t, uint64(7), indexVal.Index(2).Uint())

		require.Equal(t, uint64(2), val.FieldByName("splitInstanceBindRegionCount").Uint())

		regions := (*driver.VkRect2D)(val.FieldByName("pSplitInstanceBindRegions").UnsafePointer())
		regionSlice := ([]driver.VkRect2D)(unsafe.Slice(regions, 2))
		regionVal := reflect.ValueOf(regionSlice)

		oneRegion := regionVal.Index(0)
		require.Equal(t, int64(3), oneRegion.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(5), oneRegion.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(7), oneRegion.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(11), oneRegion.FieldByName("extent").FieldByName("height").Uint())

		oneRegion = regionVal.Index(1)
		require.Equal(t, int64(13), oneRegion.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(17), oneRegion.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(19), oneRegion.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(23), oneRegion.FieldByName("extent").FieldByName("height").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := device.BindImageMemory2([]core1_1.BindImageMemoryInfo{
		{
			Image:        image,
			Memory:       memory,
			MemoryOffset: 1,

			NextOptions: common.NextOptions{
				core1_1.BindImageMemoryDeviceGroupInfo{
					DeviceIndices: []int{1, 2, 7},
					SplitInstanceBindRegions: []core1_0.Rect2D{
						{
							Offset: core1_0.Offset2D{X: 3, Y: 5},
							Extent: core1_0.Extent2D{Width: 7, Height: 11},
						},
						{
							Offset: core1_0.Offset2D{X: 13, Y: 17},
							Extent: core1_0.Extent2D{Width: 19, Height: 23},
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestBindImagePlaneMemoryOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))

	image := mocks.EasyMockImage(ctrl)
	memory := mocks.EasyMockDeviceMemory(ctrl)

	coreDriver.EXPECT().VkBindImageMemory2(
		device.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		bindInfoCount driver.Uint32,
		pBindInfos *driver.VkBindImageMemoryInfo,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pBindInfos).Elem()

		require.Equal(t, uint64(1000157001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
		require.Equal(t, image.Handle(), driver.VkImage(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, memory.Handle(), driver.VkDeviceMemory(val.FieldByName("memory").UnsafePointer()))

		next := (*driver.VkBindImagePlaneMemoryInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000156002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x40), val.FieldByName("planeAspect").Uint()) // VK_IMAGE_ASPECT_PLANE_2_BIT

		return core1_0.VKSuccess, nil
	})

	_, err := device.BindImageMemory2([]core1_1.BindImageMemoryInfo{
		{
			Image:  image,
			Memory: memory,

			NextOptions: common.NextOptions{
				core1_1.BindImagePlaneMemoryInfo{
					PlaneAspect: core1_1.ImageAspectPlane2,
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestDeviceGroupBindSparseOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := dummies.EasyDummyDevice(coreDriver)
	fence := mocks.EasyMockFence(ctrl)

	semaphore1 := mocks.EasyMockSemaphore(ctrl)
	semaphore2 := mocks.EasyMockSemaphore(ctrl)
	semaphore3 := mocks.EasyMockSemaphore(ctrl)

	queue := dummies.EasyDummyQueue(coreDriver, device)

	coreDriver.EXPECT().VkQueueBindSparse(
		queue.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(
		queue driver.VkQueue,
		optionCount driver.Uint32,
		pSparseOptions *driver.VkBindSparseInfo,
		fence driver.VkFence,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pSparseOptions).Elem()

		require.Equal(t, uint64(7), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_SPARSE_INFO
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, semaphore1.Handle(), driver.VkSemaphore(val.FieldByName("pWaitSemaphores").Elem().UnsafePointer()))
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		semaphores := (*driver.VkSemaphore)(val.FieldByName("pSignalSemaphores").UnsafePointer())
		semaphoreSlice := ([]driver.VkSemaphore)(unsafe.Slice(semaphores, 2))
		require.Equal(t, semaphore2.Handle(), semaphoreSlice[0])
		require.Equal(t, semaphore3.Handle(), semaphoreSlice[1])

		next := (*driver.VkDeviceGroupBindSparseInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("resourceDeviceIndex").Uint())
		require.Equal(t, uint64(3), val.FieldByName("memoryDeviceIndex").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := queue.BindSparse(fence, []core1_0.BindSparseInfo{
		{
			WaitSemaphores:   []core1_0.Semaphore{semaphore1},
			SignalSemaphores: []core1_0.Semaphore{semaphore2, semaphore3},
			NextOptions: common.NextOptions{
				core1_1.DeviceGroupBindSparseInfo{
					ResourceDeviceIndex: 1,
					MemoryDeviceIndex:   3,
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestVulkanDevice_BufferMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)

	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))
	buffer := mocks.EasyMockBuffer(ctrl)

	coreDriver.EXPECT().VkGetBufferMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pInfo *driver.VkBufferMemoryRequirementsInfo2,
		pMemoryRequirements *driver.VkMemoryRequirements2,
	) {
		val := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000146000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, buffer.Handle(), driver.VkBuffer(val.FieldByName("buffer").UnsafePointer()))

		val = reflect.ValueOf(pMemoryRequirements).Elem()
		require.Equal(t, uint64(1000146003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
		require.True(t, val.FieldByName("pNext").IsNil())

		val = val.FieldByName("memoryRequirements")
		*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("size").UnsafeAddr())) = driver.VkDeviceSize(1)
		*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("alignment").UnsafeAddr())) = driver.VkDeviceSize(3)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("memoryTypeBits").UnsafeAddr())) = driver.Uint32(5)
	})

	var outData core1_1.MemoryRequirements2
	err := device.BufferMemoryRequirements2(core1_1.BufferMemoryRequirementsInfo2{
		Buffer: buffer,
	}, &outData)
	require.NoError(t, err)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryTypeBits)
}

func TestVulkanDevice_ImageMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)

	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))
	image := mocks.EasyMockImage(ctrl)

	coreDriver.EXPECT().VkGetImageMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pInfo *driver.VkImageMemoryRequirementsInfo2,
		pMemoryRequirements *driver.VkMemoryRequirements2,
	) {
		val := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000146001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, image.Handle(), driver.VkImage(val.FieldByName("image").UnsafePointer()))

		val = reflect.ValueOf(pMemoryRequirements).Elem()
		require.Equal(t, uint64(1000146003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
		require.True(t, val.FieldByName("pNext").IsNil())

		val = val.FieldByName("memoryRequirements")
		*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("size").UnsafeAddr())) = driver.VkDeviceSize(1)
		*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("alignment").UnsafeAddr())) = driver.VkDeviceSize(3)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("memoryTypeBits").UnsafeAddr())) = driver.Uint32(5)
	})

	var outData core1_1.MemoryRequirements2
	err := device.ImageMemoryRequirements2(core1_1.ImageMemoryRequirementsInfo2{
		Image: image,
	}, &outData)
	require.NoError(t, err)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryTypeBits)
}

func TestVulkanExtension_SparseImageMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))
	image := mocks.EasyMockImage(ctrl)

	coreDriver.EXPECT().VkGetImageSparseMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkDevice,
			pInfo *driver.VkImageSparseMemoryRequirementsInfo2,
			pSparseMemoryRequirementCount *driver.Uint32,
			pSparseMemoryRequirements *driver.VkSparseImageMemoryRequirements2) {

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146002), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, image.Handle(), (driver.VkImage)(options.FieldByName("image").UnsafePointer()))

			*pSparseMemoryRequirementCount = driver.Uint32(2)
		})

	coreDriver.EXPECT().VkGetImageSparseMemoryRequirements2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pInfo *driver.VkImageSparseMemoryRequirementsInfo2,
			pSparseMemoryRequirementCount *driver.Uint32,
			pSparseMemoryRequirements *driver.VkSparseImageMemoryRequirements2) {

			require.Equal(t, driver.Uint32(2), *pSparseMemoryRequirementCount)

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146002), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, image.Handle(), (driver.VkImage)(options.FieldByName("image").UnsafePointer()))

			requirementSlice := ([]driver.VkSparseImageMemoryRequirements2)(unsafe.Slice(pSparseMemoryRequirements, 2))
			outData := reflect.ValueOf(requirementSlice)
			element := outData.Index(0)
			require.Equal(t, uint64(1000146004), element.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
			require.True(t, element.FieldByName("pNext").IsNil())

			memReqs := element.FieldByName("memoryRequirements")
			imageAspectFlags := (*driver.VkImageAspectFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr()))
			*imageAspectFlags = driver.VkImageAspectFlags(0x00000008) // VK_IMAGE_ASPECT_METADATA_BIT
			width := (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr()))
			*width = driver.Uint32(1)
			height := (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr()))
			*height = driver.Uint32(3)
			depth := (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr()))
			*depth = driver.Uint32(5)
			flags := (*driver.VkSparseImageFormatFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr()))
			*flags = driver.VkSparseImageFormatFlags(0x00000004) // VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT
			*(*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = driver.Uint32(7)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailSize").UnsafeAddr())) = driver.VkDeviceSize(17)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailOffset").UnsafeAddr())) = driver.VkDeviceSize(11)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailStride").UnsafeAddr())) = driver.VkDeviceSize(13)

			element = outData.Index(1)
			require.Equal(t, uint64(1000146004), element.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
			require.True(t, element.FieldByName("pNext").IsNil())

			memReqs = element.FieldByName("memoryRequirements")
			imageAspectFlags = (*driver.VkImageAspectFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr()))
			*imageAspectFlags = driver.VkImageAspectFlags(0x00000004) // VK_IMAGE_ASPECT_STENCIL_BIT
			width = (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr()))
			*width = driver.Uint32(19)
			height = (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr()))
			*height = driver.Uint32(23)
			depth = (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr()))
			*depth = driver.Uint32(29)
			flags = (*driver.VkSparseImageFormatFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr()))
			*flags = driver.VkSparseImageFormatFlags(0)
			*(*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = driver.Uint32(43)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailSize").UnsafeAddr())) = driver.VkDeviceSize(31)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailOffset").UnsafeAddr())) = driver.VkDeviceSize(41)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailStride").UnsafeAddr())) = driver.VkDeviceSize(37)
		})

	outData, err := device.ImageSparseMemoryRequirements2(core1_1.ImageSparseMemoryRequirementsInfo2{
		Image: image,
	}, nil)
	require.NoError(t, err)
	require.Equal(t, []*core1_1.SparseImageMemoryRequirements2{
		{
			MemoryRequirements: core1_0.SparseImageMemoryRequirements{
				FormatProperties: core1_0.SparseImageFormatProperties{
					AspectMask: core1_0.ImageAspectMetadata,
					ImageGranularity: core1_0.Extent3D{
						Width:  1,
						Height: 3,
						Depth:  5,
					},
					Flags: core1_0.SparseImageFormatNonstandardBlockSize,
				},
				ImageMipTailFirstLod: 7,
				ImageMipTailOffset:   11,
				ImageMipTailStride:   13,
				ImageMipTailSize:     17,
			},
		},
		{
			MemoryRequirements: core1_0.SparseImageMemoryRequirements{
				FormatProperties: core1_0.SparseImageFormatProperties{
					AspectMask: core1_0.ImageAspectStencil,
					ImageGranularity: core1_0.Extent3D{
						Width:  19,
						Height: 23,
						Depth:  29,
					},
					Flags: 0,
				},
				ImageMipTailSize:     31,
				ImageMipTailStride:   37,
				ImageMipTailOffset:   41,
				ImageMipTailFirstLod: 43,
			},
		},
	}, outData)
}

func TestVulkanDevice_GetDeviceGroupPeerMemoryFeatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))

	coreDriver.EXPECT().VkGetDeviceGroupPeerMemoryFeatures(
		device.Handle(),
		driver.Uint32(1),
		driver.Uint32(3),
		driver.Uint32(5),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		heapIndex, localDeviceIndex, remoteDeviceIndex driver.Uint32,
		pPeerMemoryFeatures *driver.VkPeerMemoryFeatureFlags,
	) {
		*pPeerMemoryFeatures = driver.VkPeerMemoryFeatureFlags(1) // VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT
	})

	features := device.DeviceGroupPeerMemoryFeatures(
		1, 3, 5,
	)
	require.Equal(t, core1_1.PeerMemoryFeatureCopySrc, features)
}

func TestVulkanLoader_CreateSamplerYcbcrConversion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))
	mockYcbcr := mocks.EasyMockSamplerYcbcrConversion(ctrl)

	coreDriver.EXPECT().VkCreateSamplerYcbcrConversion(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pCreateInfo *driver.VkSamplerYcbcrConversionCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pYcbcrConversion *driver.VkSamplerYcbcrConversion,
		) (common.VkResult, error) {
			*pYcbcrConversion = mockYcbcr.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(1000156000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1000156021), val.FieldByName("format").Uint())             // VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16_KHR
			require.Equal(t, uint64(2), val.FieldByName("ycbcrModel").Uint())                  // VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR
			require.Equal(t, uint64(1), val.FieldByName("ycbcrRange").Uint())                  // VK_SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR
			require.Equal(t, uint64(4), val.FieldByName("components").FieldByName("r").Uint()) // VK_COMPONENT_SWIZZLE_G
			require.Equal(t, uint64(6), val.FieldByName("components").FieldByName("g").Uint()) // VK_COMPONENT_SWIZZLE_A
			require.Equal(t, uint64(0), val.FieldByName("components").FieldByName("b").Uint()) // VK_COMPONENT_SWIZZLE_IDENTITY
			require.Equal(t, uint64(2), val.FieldByName("components").FieldByName("a").Uint()) // VK_COMPONENT_SWIZZLE_ONE
			require.Equal(t, uint64(0), val.FieldByName("yChromaOffset").Uint())               // VK_CHROMA_LOCATION_COSITED_EVEN_KHR
			require.Equal(t, uint64(1), val.FieldByName("xChromaOffset").Uint())               // VK_CHROMA_LOCATION_MIDPOINT_KHR
			require.Equal(t, uint64(1), val.FieldByName("forceExplicitReconstruction").Uint())

			return core1_0.VKSuccess, nil
		})

	ycbcr, _, err := device.CreateSamplerYcbcrConversion(
		core1_1.SamplerYcbcrConversionCreateInfo{
			Format:     core1_1.FormatB12X4G12X4R12X4G12X4HorizontalChromaComponentPacked,
			YcbcrModel: core1_1.SamplerYcbcrModelConversionYcbcr709,
			YcbcrRange: core1_1.SamplerYcbcrRangeITUNarrow,
			Components: core1_0.ComponentMapping{
				R: core1_0.ComponentSwizzleGreen,
				G: core1_0.ComponentSwizzleAlpha,
				B: core1_0.ComponentSwizzleIdentity,
				A: core1_0.ComponentSwizzleOne,
			},
			YChromaOffset:               core1_1.ChromaLocationCositedEven,
			XChromaOffset:               core1_1.ChromaLocationMidpoint,
			ChromaFilter:                core1_0.FilterLinear,
			ForceExplicitReconstruction: true,
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, mockYcbcr.Handle(), ycbcr.Handle())

	coreDriver.EXPECT().VkDestroySamplerYcbcrConversion(
		device.Handle(),
		ycbcr.Handle(),
		gomock.Nil(),
	)

	ycbcr.Destroy(nil)
}

func TestVulkanLoader1_1_CreateDescriptorUpdateTemplate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := core1_1.PromoteDevice(dummies.EasyDummyDevice(coreDriver))
	mockQueue := mocks.EasyMockQueue(ctrl)

	coreDriver.EXPECT().VkGetDeviceQueue2(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pInfo *driver.VkDeviceQueueInfo2,
		pQueue *driver.VkQueue,
	) {
		*pQueue = mockQueue.Handle()

		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000145003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT
		require.Equal(t, uint64(3), val.FieldByName("queueFamilyIndex").Uint())
		require.Equal(t, uint64(5), val.FieldByName("queueIndex").Uint())
	})

	queue, err := device.GetQueue2(
		core1_1.DeviceQueueInfo2{
			Flags:            core1_1.DeviceQueueCreateProtected,
			QueueFamilyIndex: 3,
			QueueIndex:       5,
		})

	require.NoError(t, err)
	require.Equal(t, mockQueue.Handle(), queue.Handle())
}
