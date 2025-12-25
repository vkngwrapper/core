package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestCommandPoolCreateBasic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	pool := mocks.NewDummyCommandPool(device)

	mockLoader.EXPECT().VkCreateCommandPool(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, createInfo *loader.VkCommandPoolCreateInfo, allocator *loader.VkAllocationCallbacks, commandPool *loader.VkCommandPool) (common.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(39), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("flags").Uint()) // VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT
			require.Equal(t, uint64(1), val.FieldByName("queueFamilyIndex").Uint())

			*commandPool = pool.Handle()

			return core1_0.VKSuccess, nil
		})

	pool, res, err := driver.CreateCommandPool(device, nil, core1_0.CommandPoolCreateInfo{
		Flags:            core1_0.CommandPoolCreateResetBuffer,
		QueueFamilyIndex: 1,
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NotNil(t, pool)
	require.Equal(t, pool.Handle(), pool.Handle())
}

func TestCommandBufferSingleAllocateFree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	buffer := mocks.NewDummyCommandBuffer(commandPool, device)

	mockLoader.EXPECT().VkAllocateCommandBuffers(device.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, createInfo *loader.VkCommandBufferAllocateInfo, commandBuffers *loader.VkCommandBuffer) (common.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(40), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("level").Uint()) //VK_COMMAND_BUFFER_LEVEL_PRIMARY
			require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())

			require.Equal(t, commandPool.Handle(), loader.VkCommandPool(unsafe.Pointer(val.FieldByName("commandPool").Elem().UnsafeAddr())))

			bufferSlice := ([]loader.VkCommandBuffer)(unsafe.Slice(commandBuffers, 1))
			bufferSlice[0] = buffer.Handle()

			return core1_0.VKSuccess, nil
		})

	buffers, res, err := driver.AllocateCommandBuffers(core1_0.CommandBufferAllocateInfo{
		CommandPool:        commandPool,
		Level:              core1_0.CommandBufferLevelPrimary,
		CommandBufferCount: 1,
	})

	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Len(t, buffers, 1)
	require.Equal(t, buffers[0].Handle(), buffer.Handle())

	mockLoader.EXPECT().VkFreeCommandBuffers(device.Handle(), commandPool.Handle(), loader.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, commandPool loader.VkCommandPool, bufferCount loader.Uint32, buffers *loader.VkCommandBuffer) error {
			slice := ([]loader.VkCommandBuffer)(unsafe.Slice(buffers, 1))
			require.Equal(t, buffer.Handle(), slice[0])

			return nil
		})

	driver.FreeCommandBuffers(buffers...)
}

func TestCommandBufferMultiAllocateFree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)

	bufferHandles := []loader.VkCommandBuffer{
		mocks.NewFakeCommandBufferHandle(),
		mocks.NewFakeCommandBufferHandle(),
		mocks.NewFakeCommandBufferHandle(),
	}

	mockLoader.EXPECT().VkAllocateCommandBuffers(device.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, createInfo *loader.VkCommandBufferAllocateInfo, commandBuffers *loader.VkCommandBuffer) (common.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(40), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("level").Uint()) //VK_COMMAND_BUFFER_LEVEL_SECONDARY
			require.Equal(t, uint64(3), val.FieldByName("commandBufferCount").Uint())

			require.Equal(t, commandPool.Handle(), loader.VkCommandPool(unsafe.Pointer(val.FieldByName("commandPool").Elem().UnsafeAddr())))

			bufferSlice := ([]loader.VkCommandBuffer)(unsafe.Slice(commandBuffers, 3))
			bufferSlice[0] = bufferHandles[0]
			bufferSlice[1] = bufferHandles[1]
			bufferSlice[2] = bufferHandles[2]

			return core1_0.VKSuccess, nil
		})

	buffers, res, err := driver.AllocateCommandBuffers(core1_0.CommandBufferAllocateInfo{
		CommandPool:        commandPool,
		Level:              core1_0.CommandBufferLevelSecondary,
		CommandBufferCount: 3,
	})

	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Len(t, buffers, 3)

	require.Equal(t, bufferHandles[0], buffers[0].Handle())
	require.Equal(t, bufferHandles[1], buffers[1].Handle())
	require.Equal(t, bufferHandles[2], buffers[2].Handle())

	mockLoader.EXPECT().VkFreeCommandBuffers(device.Handle(), commandPool.Handle(), loader.Uint32(3), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, commandPool loader.VkCommandPool, bufferCount loader.Uint32, buffers *loader.VkCommandBuffer) {
			slice := ([]loader.VkCommandBuffer)(unsafe.Slice(buffers, 3))
			require.Equal(t, bufferHandles[0], slice[0])
			require.Equal(t, bufferHandles[1], slice[1])
			require.Equal(t, bufferHandles[2], slice[2])
		})

	driver.FreeCommandBuffers(buffers...)
}

func TestVulkanCommandPool_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)

	mockLoader.EXPECT().VkResetCommandPool(device.Handle(), commandPool.Handle(),
		loader.VkCommandPoolResetFlags(1), // VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT
	).Return(core1_0.VKSuccess, nil)

	_, err := driver.ResetCommandPool(commandPool, core1_0.CommandPoolResetReleaseResources)
	require.NoError(t, err)
}
