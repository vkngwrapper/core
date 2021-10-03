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

func TestCommandPoolCreateBasic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)

	expectedPoolHandle := mocks.NewFakeCommandPoolHandle()

	device := mocks.EasyMockDevice(ctrl, mockDriver)

	mockDriver.EXPECT().VkCreateCommandPool(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, createInfo *core.VkCommandPoolCreateInfo, allocator *core.VkAllocationCallbacks, commandPool *core.VkCommandPool) (core.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(39), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("flags").Uint()) // VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT
			require.Equal(t, uint64(1), val.FieldByName("queueFamilyIndex").Uint())

			*commandPool = expectedPoolHandle

			return core.VKSuccess, nil
		})

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	graphicsFamily := 1
	pool, res, err := loader.CreateCommandPool(device, &core.CommandPoolOptions{
		Flags:               core.CommandPoolResetBuffer,
		GraphicsQueueFamily: &graphicsFamily,
	})
	require.NoError(t, err)
	require.Equal(t, core.VKSuccess, res)
	require.NotNil(t, pool)
	require.Equal(t, expectedPoolHandle, pool.Handle())
}

func TestCommandPoolNullQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)

	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	pool, res, err := loader.CreateCommandPool(mockDevice, &core.CommandPoolOptions{
		Flags: core.CommandPoolResetBuffer,
	})
	require.Error(t, err)
	require.Equal(t, core.VKErrorUnknown, res)
	require.Nil(t, pool)
}

func TestCommandBufferSingleAllocateFree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	commandPool := mocks.EasyDummyCommandPool(t, loader, mockDevice)

	bufferHandle := mocks.NewFakeCommandBufferHandle()

	mockDriver.EXPECT().VkAllocateCommandBuffers(mockDevice.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, createInfo *core.VkCommandBufferAllocateInfo, commandBuffers *core.VkCommandBuffer) (core.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(40), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("level").Uint()) //VK_COMMAND_BUFFER_LEVEL_PRIMARY
			require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())

			require.Equal(t, commandPool.Handle(), core.VkCommandPool(unsafe.Pointer(val.FieldByName("commandPool").Elem().UnsafeAddr())))

			bufferSlice := ([]core.VkCommandBuffer)(unsafe.Slice(commandBuffers, 1))
			bufferSlice[0] = bufferHandle

			return core.VKSuccess, nil
		})

	buffers, res, err := commandPool.AllocateCommandBuffers(&core.CommandBufferOptions{
		Level:       common.LevelPrimary,
		BufferCount: 1,
	})

	require.NoError(t, err)
	require.Equal(t, core.VKSuccess, res)
	require.Len(t, buffers, 1)
	require.Equal(t, buffers[0].Handle(), bufferHandle)

	mockDriver.EXPECT().VkFreeCommandBuffers(mockDevice.Handle(), commandPool.Handle(), core.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, commandPool core.VkCommandPool, bufferCount core.Uint32, buffers *core.VkCommandBuffer) error {
			slice := ([]core.VkCommandBuffer)(unsafe.Slice(buffers, 1))
			require.Equal(t, bufferHandle, slice[0])

			return nil
		})

	err = commandPool.FreeCommandBuffers(buffers)
	require.NoError(t, err)
}

func TestCommandBufferMultiAllocateFree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)

	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	commandPool := mocks.EasyDummyCommandPool(t, loader, mockDevice)

	bufferHandles := []core.VkCommandBuffer{
		mocks.NewFakeCommandBufferHandle(),
		mocks.NewFakeCommandBufferHandle(),
		mocks.NewFakeCommandBufferHandle(),
	}

	mockDriver.EXPECT().VkAllocateCommandBuffers(mockDevice.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, createInfo *core.VkCommandBufferAllocateInfo, commandBuffers *core.VkCommandBuffer) (core.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(40), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("level").Uint()) //VK_COMMAND_BUFFER_LEVEL_SECONDARY
			require.Equal(t, uint64(3), val.FieldByName("commandBufferCount").Uint())

			require.Equal(t, commandPool.Handle(), core.VkCommandPool(unsafe.Pointer(val.FieldByName("commandPool").Elem().UnsafeAddr())))

			bufferSlice := ([]core.VkCommandBuffer)(unsafe.Slice(commandBuffers, 3))
			bufferSlice[0] = bufferHandles[0]
			bufferSlice[1] = bufferHandles[1]
			bufferSlice[2] = bufferHandles[2]

			return core.VKSuccess, nil
		})

	buffers, res, err := commandPool.AllocateCommandBuffers(&core.CommandBufferOptions{
		Level:       common.LevelSecondary,
		BufferCount: 3,
	})

	require.NoError(t, err)
	require.Equal(t, core.VKSuccess, res)
	require.Len(t, buffers, 3)

	actualHandles := []core.VkCommandBuffer{
		buffers[0].Handle(),
		buffers[1].Handle(),
		buffers[2].Handle(),
	}
	require.ElementsMatch(t, bufferHandles, actualHandles)

	mockDriver.EXPECT().VkFreeCommandBuffers(mockDevice.Handle(), commandPool.Handle(), core.Uint32(3), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, commandPool core.VkCommandPool, bufferCount core.Uint32, buffers *core.VkCommandBuffer) error {
			slice := ([]core.VkCommandBuffer)(unsafe.Slice(buffers, 3))
			require.ElementsMatch(t, bufferHandles, slice)

			return nil
		})

	err = commandPool.FreeCommandBuffers(buffers)
	require.NoError(t, err)
}
