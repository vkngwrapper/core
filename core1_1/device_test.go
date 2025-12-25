package core1_1_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestBindBufferMemoryDeviceGroupOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	buffer := mocks.NewDummyBuffer(device)
	memory := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader.EXPECT().VkBindBufferMemory2(
		device.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		infoCount loader.Uint32,
		pInfo *loader.VkBindBufferMemoryInfo,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000157000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
		require.Equal(t, buffer.Handle(), (loader.VkBuffer)(val.FieldByName("buffer").UnsafePointer()))
		require.Equal(t, memory.Handle(), (loader.VkDeviceMemory)(val.FieldByName("memory").UnsafePointer()))
		require.Equal(t, uint64(1), val.FieldByName("memoryOffset").Uint())

		next := (*loader.VkBindBufferMemoryDeviceGroupInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060013), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceIndexCount").Uint())

		indices := (*loader.Uint32)(val.FieldByName("pDeviceIndices").UnsafePointer())
		indexSlice := ([]loader.Uint32)(unsafe.Slice(indices, 3))
		val = reflect.ValueOf(indexSlice)

		require.Equal(t, uint64(1), val.Index(0).Uint())
		require.Equal(t, uint64(2), val.Index(1).Uint())
		require.Equal(t, uint64(7), val.Index(2).Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := driver.BindBufferMemory2(
		core1_1.BindBufferMemoryInfo{
			Buffer:       buffer,
			Memory:       memory,
			MemoryOffset: 1,

			NextOptions: common.NextOptions{
				core1_1.BindBufferMemoryDeviceGroupInfo{
					DeviceIndices: []int{1, 2, 7},
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestBindImageMemoryDeviceGroupOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	image := mocks.NewDummyImage(device)
	memory := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader.EXPECT().VkBindImageMemory2(
		device.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		infoCount loader.Uint32,
		pInfo *loader.VkBindImageMemoryInfo,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000157001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
		require.Equal(t, image.Handle(), (loader.VkImage)(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, memory.Handle(), (loader.VkDeviceMemory)(val.FieldByName("memory").UnsafePointer()))
		require.Equal(t, uint64(1), val.FieldByName("memoryOffset").Uint())

		next := (*loader.VkBindImageMemoryDeviceGroupInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060014), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceIndexCount").Uint())

		indices := (*loader.Uint32)(val.FieldByName("pDeviceIndices").UnsafePointer())
		indexSlice := ([]loader.Uint32)(unsafe.Slice(indices, 3))
		indexVal := reflect.ValueOf(indexSlice)

		require.Equal(t, uint64(1), indexVal.Index(0).Uint())
		require.Equal(t, uint64(2), indexVal.Index(1).Uint())
		require.Equal(t, uint64(7), indexVal.Index(2).Uint())

		require.Equal(t, uint64(2), val.FieldByName("splitInstanceBindRegionCount").Uint())

		regions := (*loader.VkRect2D)(val.FieldByName("pSplitInstanceBindRegions").UnsafePointer())
		regionSlice := ([]loader.VkRect2D)(unsafe.Slice(regions, 2))
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

	_, err := driver.BindImageMemory2(
		core1_1.BindImageMemoryInfo{
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
	)
	require.NoError(t, err)
}

func TestBindImagePlaneMemoryOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	image := mocks.NewDummyImage(device)
	memory := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader.EXPECT().VkBindImageMemory2(
		device.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		bindInfoCount loader.Uint32,
		pBindInfos *loader.VkBindImageMemoryInfo,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pBindInfos).Elem()

		require.Equal(t, uint64(1000157001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
		require.Equal(t, image.Handle(), loader.VkImage(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, memory.Handle(), loader.VkDeviceMemory(val.FieldByName("memory").UnsafePointer()))

		next := (*loader.VkBindImagePlaneMemoryInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000156002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x40), val.FieldByName("planeAspect").Uint()) // VK_IMAGE_ASPECT_PLANE_2_BIT

		return core1_0.VKSuccess, nil
	})

	_, err := driver.BindImageMemory2(
		core1_1.BindImageMemoryInfo{
			Image:  image,
			Memory: memory,

			NextOptions: common.NextOptions{
				core1_1.BindImagePlaneMemoryInfo{
					PlaneAspect: core1_1.ImageAspectPlane2,
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestDeviceGroupBindSparseOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	fence := mocks.NewDummyFence(device)

	semaphore1 := mocks.NewDummySemaphore(device)
	semaphore2 := mocks.NewDummySemaphore(device)
	semaphore3 := mocks.NewDummySemaphore(device)

	queue := mocks.NewDummyQueue(device)

	coreLoader.EXPECT().VkQueueBindSparse(
		queue.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(
		queue loader.VkQueue,
		optionCount loader.Uint32,
		pSparseOptions *loader.VkBindSparseInfo,
		fence loader.VkFence,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pSparseOptions).Elem()

		require.Equal(t, uint64(7), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_SPARSE_INFO
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, semaphore1.Handle(), loader.VkSemaphore(val.FieldByName("pWaitSemaphores").Elem().UnsafePointer()))
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		semaphores := (*loader.VkSemaphore)(val.FieldByName("pSignalSemaphores").UnsafePointer())
		semaphoreSlice := ([]loader.VkSemaphore)(unsafe.Slice(semaphores, 2))
		require.Equal(t, semaphore2.Handle(), semaphoreSlice[0])
		require.Equal(t, semaphore3.Handle(), semaphoreSlice[1])

		next := (*loader.VkDeviceGroupBindSparseInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("resourceDeviceIndex").Uint())
		require.Equal(t, uint64(3), val.FieldByName("memoryDeviceIndex").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := driver.QueueBindSparse(queue, &fence,
		core1_0.BindSparseInfo{
			WaitSemaphores:   []core.Semaphore{semaphore1},
			SignalSemaphores: []core.Semaphore{semaphore2, semaphore3},
			NextOptions: common.NextOptions{
				core1_1.DeviceGroupBindSparseInfo{
					ResourceDeviceIndex: 1,
					MemoryDeviceIndex:   3,
				},
			},
		},
	)
	require.NoError(t, err)
}
