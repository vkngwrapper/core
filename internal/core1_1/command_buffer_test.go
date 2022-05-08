package internal1_1_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestVulkanExtension_CmdDispatchBase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	commandBuffer := dummies.EasyDummyCommandBuffer(t, loader, device, commandPool)

	coreDriver.EXPECT().VkCmdDispatchBase(
		commandBuffer.Handle(),
		driver.Uint32(1),
		driver.Uint32(3),
		driver.Uint32(5),
		driver.Uint32(7),
		driver.Uint32(11),
		driver.Uint32(13),
	)

	commandBuffer.Core1_1().CmdDispatchBase(1, 3, 5, 7, 11, 13)
}

func TestVulkanExtension_CmdSetDeviceMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	commandBuffer := dummies.EasyDummyCommandBuffer(t, loader, device, commandPool)

	coreDriver.EXPECT().VkCmdSetDeviceMask(commandBuffer.Handle(), driver.Uint32(3))

	commandBuffer.Core1_1().CmdSetDeviceMask(3)
}

func TestDeviceGroupCommandBufferBeginOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	mockCommandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	commandBuffer := core.CreateCommandBuffer(coreDriver, commandPool.Handle(), device.Handle(), mockCommandBuffer.Handle(), common.Vulkan1_0)

	coreDriver.EXPECT().VkBeginCommandBuffer(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer driver.VkCommandBuffer, pBeginInfo *driver.VkCommandBufferBeginInfo) (common.VkResult, error) {
		val := reflect.ValueOf(pBeginInfo).Elem()

		require.Equal(t, uint64(42), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint())  // VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT
		require.True(t, val.FieldByName("pInheritanceInfo").IsNil())

		next := (*driver.VkDeviceGroupCommandBufferBeginInfo)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceMask").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := commandBuffer.Begin(core1_0.BeginOptions{
		Flags: core1_0.BeginInfoOneTimeSubmit,
		HaveNext: common.HaveNext{Next: core1_1.DeviceGroupCommandBufferBeginOptions{
			DeviceMask: 3,
		}},
	})
	require.NoError(t, err)
}
