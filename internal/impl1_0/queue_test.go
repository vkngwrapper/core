package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanQueue_WaitForIdle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	queue := mocks.NewDummyQueue(device)

	mockLoader.EXPECT().VkQueueWaitIdle(queue.Handle()).Return(core1_0.VKSuccess, nil)

	_, err := driver.QueueWaitIdle(queue)
	require.NoError(t, err)
}

func TestVulkanQueue_BindSparse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	queue := mocks.NewDummyQueue(device)

	semaphore1 := mocks.NewDummySemaphore(device)
	semaphore2 := mocks.NewDummySemaphore(device)
	semaphore3 := mocks.NewDummySemaphore(device)

	buffer := mocks.NewDummyBuffer(device)
	image1 := mocks.NewDummyImage(device)
	image2 := mocks.NewDummyImage(device)
	memory1 := mocks.NewDummyDeviceMemory(device, 1)
	memory2 := mocks.NewDummyDeviceMemory(device, 1)
	memory3 := mocks.NewDummyDeviceMemory(device, 1)
	memory4 := mocks.NewDummyDeviceMemory(device, 1)

	mockLoader.EXPECT().VkQueueBindSparse(queue.Handle(), loader.Uint32(1), gomock.Not(nil), loader.VkFence(loader.NullHandle)).DoAndReturn(
		func(queue loader.VkQueue, bindInfoCount loader.Uint32, pBindInfo *loader.VkBindSparseInfo, fence loader.VkFence) (common.VkResult, error) {
			bindSlice := ([]loader.VkBindSparseInfo)(unsafe.Slice(pBindInfo, 1))
			val := reflect.ValueOf(bindSlice).Index(0)

			// Root
			require.Equal(t, uint64(7), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_SPARSE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
			require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("bufferBindCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("imageOpaqueBindCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("imageBindCount").Uint())

			// Wait Semaphores
			waitSemaphorePtr := (*loader.VkSemaphore)(unsafe.Pointer(val.FieldByName("pWaitSemaphores").Elem().UnsafeAddr()))
			waitSemaphoreSlice := ([]loader.VkSemaphore)(unsafe.Slice(waitSemaphorePtr, 1))

			require.Equal(t, semaphore1.Handle(), waitSemaphoreSlice[0])

			// Signal Semaphores
			signalSemaphorePtr := (*loader.VkSemaphore)(unsafe.Pointer(val.FieldByName("pSignalSemaphores").Elem().UnsafeAddr()))
			signalSemaphoreSlice := ([]loader.VkSemaphore)(unsafe.Slice(signalSemaphorePtr, 2))

			require.Equal(t, semaphore2.Handle(), signalSemaphoreSlice[0])
			require.Equal(t, semaphore3.Handle(), signalSemaphoreSlice[1])

			// Sparse buffer memory bind
			bufferBindsPtr := (*loader.VkSparseBufferMemoryBindInfo)(unsafe.Pointer(val.FieldByName("pBufferBinds").Elem().UnsafeAddr()))
			bufferBindsSlice := ([]loader.VkSparseBufferMemoryBindInfo)(unsafe.Slice(bufferBindsPtr, 1))
			val = reflect.ValueOf(bufferBindsSlice).Index(0)
			bufferHandle := (loader.VkBuffer)(unsafe.Pointer(val.FieldByName("buffer").Elem().UnsafeAddr()))
			require.Equal(t, buffer.Handle(), bufferHandle)

			require.Equal(t, uint64(2), val.FieldByName("bindCount").Uint())
			bindsPtr := (*loader.VkSparseMemoryBind)(unsafe.Pointer(val.FieldByName("pBinds").Elem().UnsafeAddr()))
			bindsSlice := ([]loader.VkSparseMemoryBind)(unsafe.Slice(bindsPtr, 2))

			val = reflect.ValueOf(bindsSlice).Index(0)
			require.Equal(t, uint64(1), val.FieldByName("resourceOffset").Uint())
			require.Equal(t, uint64(3), val.FieldByName("size").Uint())
			require.Equal(t, uint64(5), val.FieldByName("memoryOffset").Uint())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_SPARSE_MEMORY_BIND_METADATA_BIT
			memHandle := (loader.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr()))
			require.Equal(t, memory1.Handle(), memHandle)

			val = reflect.ValueOf(bindsSlice).Index(1)
			require.Equal(t, uint64(7), val.FieldByName("resourceOffset").Uint())
			require.Equal(t, uint64(11), val.FieldByName("size").Uint())
			require.Equal(t, uint64(13), val.FieldByName("memoryOffset").Uint())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			memHandle = (loader.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr()))
			require.Equal(t, memory2.Handle(), memHandle)

			// Sparse image opaque memory bind

			imageOpaqueBindsPtr := (*loader.VkSparseImageOpaqueMemoryBindInfo)(unsafe.Pointer(reflect.ValueOf(bindSlice).Index(0).FieldByName("pImageOpaqueBinds").Elem().UnsafeAddr()))
			imageOpaqueBindsSlice := ([]loader.VkSparseImageOpaqueMemoryBindInfo)(unsafe.Slice(imageOpaqueBindsPtr, 1))
			val = reflect.ValueOf(imageOpaqueBindsSlice).Index(0)
			imageHandle := (loader.VkImage)(unsafe.Pointer(val.FieldByName("image").Elem().UnsafeAddr()))
			require.Equal(t, image1.Handle(), imageHandle)

			require.Equal(t, uint64(1), val.FieldByName("bindCount").Uint())

			val = val.FieldByName("pBinds").Elem()
			require.Equal(t, uint64(17), val.FieldByName("resourceOffset").Uint())
			require.Equal(t, uint64(19), val.FieldByName("size").Uint())
			require.Equal(t, uint64(23), val.FieldByName("memoryOffset").Uint())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			memHandle = (loader.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr()))
			require.Equal(t, memory3.Handle(), memHandle)

			// Sparse image memory bind

			imageBindsPtr := (*loader.VkSparseImageMemoryBindInfo)(unsafe.Pointer(reflect.ValueOf(bindSlice).Index(0).FieldByName("pImageBinds").Elem().UnsafeAddr()))
			imageBindsSlice := ([]loader.VkSparseImageMemoryBindInfo)(unsafe.Slice(imageBindsPtr, 1))
			val = reflect.ValueOf(imageBindsSlice).Index(0)
			imageHandle = (loader.VkImage)(unsafe.Pointer(val.FieldByName("image").Elem().UnsafeAddr()))
			require.Equal(t, image2.Handle(), imageHandle)

			require.Equal(t, uint64(1), val.FieldByName("bindCount").Uint())

			val = val.FieldByName("pBinds").Elem()
			require.Equal(t, uint64(1), val.FieldByName("subresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT
			require.Equal(t, uint64(29), val.FieldByName("subresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(31), val.FieldByName("subresource").FieldByName("arrayLayer").Uint())
			require.Equal(t, int64(37), val.FieldByName("offset").FieldByName("x").Int())
			require.Equal(t, int64(41), val.FieldByName("offset").FieldByName("y").Int())
			require.Equal(t, int64(43), val.FieldByName("offset").FieldByName("z").Int())
			require.Equal(t, uint64(47), val.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(53), val.FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(59), val.FieldByName("extent").FieldByName("depth").Uint())
			require.Equal(t, memory4.Handle(), (loader.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(61), val.FieldByName("memoryOffset").Uint())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_SPARSE_MEMORY_BIND_METADATA_BIT

			return core1_0.VKSuccess, nil
		})

	_, err := driver.QueueBindSparse(queue, nil,
		core1_0.BindSparseInfo{
			WaitSemaphores:   []core.Semaphore{semaphore1},
			SignalSemaphores: []core.Semaphore{semaphore2, semaphore3},
			BufferBinds: []core1_0.SparseBufferMemoryBindInfo{
				{
					Buffer: buffer,
					Binds: []core1_0.SparseMemoryBind{
						{
							ResourceOffset: 1,
							Size:           3,
							Memory:         memory1,
							MemoryOffset:   5,
							Flags:          core1_0.SparseMemoryBindMetadata,
						},
						{
							ResourceOffset: 7,
							Size:           11,
							Memory:         memory2,
							MemoryOffset:   13,
						},
					},
				},
			},
			ImageOpaqueBinds: []core1_0.SparseImageOpaqueMemoryBindInfo{
				{
					Image: image1,
					Binds: []core1_0.SparseMemoryBind{
						{
							ResourceOffset: 17,
							Size:           19,
							Memory:         memory3,
							MemoryOffset:   23,
						},
					},
				},
			},
			ImageBinds: []core1_0.SparseImageMemoryBindInfo{
				{
					Image: image2,
					Binds: []core1_0.SparseImageMemoryBind{
						{
							Subresource: core1_0.ImageSubresource{
								AspectMask: core1_0.ImageAspectColor,
								MipLevel:   29,
								ArrayLayer: 31,
							},
							Offset:       core1_0.Offset3D{37, 41, 43},
							Extent:       core1_0.Extent3D{47, 53, 59},
							Memory:       memory4,
							MemoryOffset: 61,
							Flags:        core1_0.SparseMemoryBindMetadata,
						},
					},
				},
			},
		},
	)
	require.NoError(t, err)
}
