package impl1_0_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestCommandPoolCreateBasic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	expectedPoolHandle := mocks.NewFakeCommandPoolHandle()

	builder := impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})

	mockDriver.EXPECT().VkCreateCommandPool(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, createInfo *driver.VkCommandPoolCreateInfo, allocator *driver.VkAllocationCallbacks, commandPool *driver.VkCommandPool) (common.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(39), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("flags").Uint()) // VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT
			require.Equal(t, uint64(1), val.FieldByName("queueFamilyIndex").Uint())

			*commandPool = expectedPoolHandle

			return core1_0.VKSuccess, nil
		})

	pool, res, err := device.CreateCommandPool(nil, core1_0.CommandPoolCreateInfo{
		Flags:            core1_0.CommandPoolCreateResetBuffer,
		QueueFamilyIndex: 1,
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NotNil(t, pool)
	require.Equal(t, expectedPoolHandle, pool.Handle())
}

func TestCommandBufferSingleAllocateFree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})

	commandPool := mocks.EasyMockCommandPool(ctrl, device)

	bufferHandle := mocks.NewFakeCommandBufferHandle()

	mockDriver.EXPECT().VkAllocateCommandBuffers(device.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, createInfo *driver.VkCommandBufferAllocateInfo, commandBuffers *driver.VkCommandBuffer) (common.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(40), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("level").Uint()) //VK_COMMAND_BUFFER_LEVEL_PRIMARY
			require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())

			require.Equal(t, commandPool.Handle(), driver.VkCommandPool(unsafe.Pointer(val.FieldByName("commandPool").Elem().UnsafeAddr())))

			bufferSlice := ([]driver.VkCommandBuffer)(unsafe.Slice(commandBuffers, 1))
			bufferSlice[0] = bufferHandle

			return core1_0.VKSuccess, nil
		})

	buffers, res, err := device.AllocateCommandBuffers(core1_0.CommandBufferAllocateInfo{
		CommandPool:        commandPool,
		Level:              core1_0.CommandBufferLevelPrimary,
		CommandBufferCount: 1,
	})

	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Len(t, buffers, 1)
	require.Equal(t, buffers[0].Handle(), bufferHandle)

	mockDriver.EXPECT().VkFreeCommandBuffers(device.Handle(), commandPool.Handle(), driver.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, commandPool driver.VkCommandPool, bufferCount driver.Uint32, buffers *driver.VkCommandBuffer) error {
			slice := ([]driver.VkCommandBuffer)(unsafe.Slice(buffers, 1))
			require.Equal(t, bufferHandle, slice[0])

			return nil
		})

	device.FreeCommandBuffers(buffers)
}

func TestCommandBufferMultiAllocateFree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})

	commandPool := mocks.EasyMockCommandPool(ctrl, device)

	bufferHandles := []driver.VkCommandBuffer{
		mocks.NewFakeCommandBufferHandle(),
		mocks.NewFakeCommandBufferHandle(),
		mocks.NewFakeCommandBufferHandle(),
	}

	mockDriver.EXPECT().VkAllocateCommandBuffers(device.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, createInfo *driver.VkCommandBufferAllocateInfo, commandBuffers *driver.VkCommandBuffer) (common.VkResult, error) {
			val := reflect.ValueOf(*createInfo)
			require.Equal(t, uint64(40), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("level").Uint()) //VK_COMMAND_BUFFER_LEVEL_SECONDARY
			require.Equal(t, uint64(3), val.FieldByName("commandBufferCount").Uint())

			require.Equal(t, commandPool.Handle(), driver.VkCommandPool(unsafe.Pointer(val.FieldByName("commandPool").Elem().UnsafeAddr())))

			bufferSlice := ([]driver.VkCommandBuffer)(unsafe.Slice(commandBuffers, 3))
			bufferSlice[0] = bufferHandles[0]
			bufferSlice[1] = bufferHandles[1]
			bufferSlice[2] = bufferHandles[2]

			return core1_0.VKSuccess, nil
		})

	buffers, res, err := device.AllocateCommandBuffers(core1_0.CommandBufferAllocateInfo{
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

	mockDriver.EXPECT().VkFreeCommandBuffers(device.Handle(), commandPool.Handle(), driver.Uint32(3), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, commandPool driver.VkCommandPool, bufferCount driver.Uint32, buffers *driver.VkCommandBuffer) {
			slice := ([]driver.VkCommandBuffer)(unsafe.Slice(buffers, 3))
			require.Equal(t, bufferHandles[0], slice[0])
			require.Equal(t, bufferHandles[1], slice[1])
			require.Equal(t, bufferHandles[2], slice[2])
		})

	device.FreeCommandBuffers(buffers)
}

func TestVulkanCommandPool_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	mockDevice := mocks.EasyMockDevice(ctrl, mockDriver)

	builder := impl1_0.DeviceObjectBuilderImpl{}
	commandPool := builder.CreateCommandPoolObject(mockDriver, mockDevice.Handle(), mocks.NewFakeCommandPoolHandle(), common.Vulkan1_0)

	mockDriver.EXPECT().VkResetCommandPool(mockDevice.Handle(), commandPool.Handle(),
		driver.VkCommandPoolResetFlags(1), // VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT
	).Return(core1_0.VKSuccess, nil)

	_, err := commandPool.Reset(core1_0.CommandPoolResetReleaseResources)
	require.NoError(t, err)
}
