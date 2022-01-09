package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanQueue_WaitForIdle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	queue := mocks.EasyDummyQueue(device)

	driver.EXPECT().VkQueueWaitIdle(mocks.Exactly(queue.Handle())).Return(core.VKSuccess, nil)

	_, err = queue.WaitForIdle()
	require.NoError(t, err)
}

func TestVulkanQueue_BindSparse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	queue := mocks.EasyDummyQueue(device)

	semaphore1 := mocks.EasyMockSemaphore(ctrl)
	semaphore2 := mocks.EasyMockSemaphore(ctrl)
	semaphore3 := mocks.EasyMockSemaphore(ctrl)

	buffer := mocks.EasyMockBuffer(ctrl)
	image1 := mocks.EasyMockImage(ctrl)
	image2 := mocks.EasyMockImage(ctrl)
	memory1 := mocks.EasyMockDeviceMemory(ctrl)
	memory2 := mocks.EasyMockDeviceMemory(ctrl)
	memory3 := mocks.EasyMockDeviceMemory(ctrl)
	memory4 := mocks.EasyMockDeviceMemory(ctrl)

	driver.EXPECT().VkQueueBindSparse(mocks.Exactly(queue.Handle()), core.Uint32(1), gomock.Not(nil), nil).DoAndReturn(
		func(queue core.VkQueue, bindInfoCount core.Uint32, pBindInfo *core.VkBindSparseInfo, fence core.VkFence) (core.VkResult, error) {
			bindSlice := ([]core.VkBindSparseInfo)(unsafe.Slice(pBindInfo, 1))
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
			waitSemaphorePtr := (*core.VkSemaphore)(unsafe.Pointer(val.FieldByName("pWaitSemaphores").Elem().UnsafeAddr()))
			waitSemaphoreSlice := ([]core.VkSemaphore)(unsafe.Slice(waitSemaphorePtr, 1))

			require.Same(t, semaphore1.Handle(), waitSemaphoreSlice[0])

			// Signal Semaphores
			signalSemaphorePtr := (*core.VkSemaphore)(unsafe.Pointer(val.FieldByName("pSignalSemaphores").Elem().UnsafeAddr()))
			signalSemaphoreSlice := ([]core.VkSemaphore)(unsafe.Slice(signalSemaphorePtr, 2))

			require.Same(t, semaphore2.Handle(), signalSemaphoreSlice[0])
			require.Same(t, semaphore3.Handle(), signalSemaphoreSlice[1])

			// Sparse buffer memory bind
			bufferBindsPtr := (*core.VkSparseBufferMemoryBindInfo)(unsafe.Pointer(val.FieldByName("pBufferBinds").Elem().UnsafeAddr()))
			bufferBindsSlice := ([]core.VkSparseBufferMemoryBindInfo)(unsafe.Slice(bufferBindsPtr, 1))
			val = reflect.ValueOf(bufferBindsSlice).Index(0)
			bufferHandle := (core.VkBuffer)(unsafe.Pointer(val.FieldByName("buffer").Elem().UnsafeAddr()))
			require.Same(t, buffer.Handle(), bufferHandle)

			require.Equal(t, uint64(2), val.FieldByName("bindCount").Uint())
			bindsPtr := (*core.VkSparseMemoryBind)(unsafe.Pointer(val.FieldByName("pBinds").Elem().UnsafeAddr()))
			bindsSlice := ([]core.VkSparseMemoryBind)(unsafe.Slice(bindsPtr, 2))

			val = reflect.ValueOf(bindsSlice).Index(0)
			require.Equal(t, uint64(1), val.FieldByName("resourceOffset").Uint())
			require.Equal(t, uint64(3), val.FieldByName("size").Uint())
			require.Equal(t, uint64(5), val.FieldByName("memoryOffset").Uint())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_SPARSE_MEMORY_BIND_METADATA_BIT
			memHandle := (core.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr()))
			require.Same(t, memory1.Handle(), memHandle)

			val = reflect.ValueOf(bindsSlice).Index(1)
			require.Equal(t, uint64(7), val.FieldByName("resourceOffset").Uint())
			require.Equal(t, uint64(11), val.FieldByName("size").Uint())
			require.Equal(t, uint64(13), val.FieldByName("memoryOffset").Uint())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			memHandle = (core.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr()))
			require.Same(t, memory2.Handle(), memHandle)

			// Sparse image opaque memory bind

			imageOpaqueBindsPtr := (*core.VkSparseImageOpaqueMemoryBindInfo)(unsafe.Pointer(reflect.ValueOf(bindSlice).Index(0).FieldByName("pImageOpaqueBinds").Elem().UnsafeAddr()))
			imageOpaqueBindsSlice := ([]core.VkSparseImageOpaqueMemoryBindInfo)(unsafe.Slice(imageOpaqueBindsPtr, 1))
			val = reflect.ValueOf(imageOpaqueBindsSlice).Index(0)
			imageHandle := (core.VkImage)(unsafe.Pointer(val.FieldByName("image").Elem().UnsafeAddr()))
			require.Same(t, image1.Handle(), imageHandle)

			require.Equal(t, uint64(1), val.FieldByName("bindCount").Uint())

			val = val.FieldByName("pBinds").Elem()
			require.Equal(t, uint64(17), val.FieldByName("resourceOffset").Uint())
			require.Equal(t, uint64(19), val.FieldByName("size").Uint())
			require.Equal(t, uint64(23), val.FieldByName("memoryOffset").Uint())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			memHandle = (core.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr()))
			require.Same(t, memory3.Handle(), memHandle)

			// Sparse image memory bind

			imageBindsPtr := (*core.VkSparseImageMemoryBindInfo)(unsafe.Pointer(reflect.ValueOf(bindSlice).Index(0).FieldByName("pImageBinds").Elem().UnsafeAddr()))
			imageBindsSlice := ([]core.VkSparseImageMemoryBindInfo)(unsafe.Slice(imageBindsPtr, 1))
			val = reflect.ValueOf(imageBindsSlice).Index(0)
			imageHandle = (core.VkImage)(unsafe.Pointer(val.FieldByName("image").Elem().UnsafeAddr()))
			require.Same(t, image2.Handle(), imageHandle)

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
			require.Same(t, memory4.Handle(), (core.VkDeviceMemory)(unsafe.Pointer(val.FieldByName("memory").Elem().UnsafeAddr())))
			require.Equal(t, uint64(61), val.FieldByName("memoryOffset").Uint())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_SPARSE_MEMORY_BIND_METADATA_BIT

			return core.VKSuccess, nil
		})

	_, err = queue.BindSparse(nil, []*core.BindSparseOptions{
		{
			WaitSemaphores:   []core.Semaphore{semaphore1},
			SignalSemaphores: []core.Semaphore{semaphore2, semaphore3},
			BufferBinds: []core.SparseBufferMemoryBindInfo{
				{
					Buffer: buffer,
					Binds: []core.SparseMemoryBind{
						{
							ResourceOffset: 1,
							Size:           3,
							Memory:         memory1,
							MemoryOffset:   5,
							Flags:          core.SparseMemoryBindMetadata,
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
			ImageOpaqueBinds: []core.SparseImageOpaqueMemoryBindInfo{
				{
					Image: image1,
					Binds: []core.SparseMemoryBind{
						{
							ResourceOffset: 17,
							Size:           19,
							Memory:         memory3,
							MemoryOffset:   23,
						},
					},
				},
			},
			ImageBinds: []core.SparseImageMemoryBindInfo{
				{
					Image: image2,
					Binds: []core.SparseImageMemoryBind{
						{
							Subresource: common.ImageSubresource{
								AspectMask: common.AspectColor,
								MipLevel:   29,
								ArrayLayer: 31,
							},
							Offset:       common.Offset3D{37, 41, 43},
							Extent:       common.Extent3D{47, 53, 59},
							Memory:       memory4,
							MemoryOffset: 61,
							Flags:        core.SparseMemoryBindMetadata,
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)
}
