package impl1_1_test

import (
	"testing"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestCommandBuffer_CmdDispatchBase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	builder := &impl1_1.DeviceObjectBuilderImpl{}
	commandBuffer := builder.CreateCommandBufferObject(coreDriver, commandPool.Handle(), device.Handle(), mocks.NewFakeCommandBufferHandle(), common.Vulkan1_1).(core1_1.CommandBuffer)

	coreDriver.EXPECT().VkCmdDispatchBase(
		commandBuffer.Handle(),
		driver.Uint32(1),
		driver.Uint32(3),
		driver.Uint32(5),
		driver.Uint32(7),
		driver.Uint32(11),
		driver.Uint32(13),
	)

	commandBuffer.CmdDispatchBase(1, 3, 5, 7, 11, 13)
}

func TestCommandBuffer_CmdSetDeviceMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)
	builder := &impl1_1.DeviceObjectBuilderImpl{}
	commandBuffer := builder.CreateCommandBufferObject(coreDriver, commandPool.Handle(), device.Handle(), mocks.NewFakeCommandBufferHandle(), common.Vulkan1_1).(core1_1.CommandBuffer)

	coreDriver.EXPECT().VkCmdSetDeviceMask(commandBuffer.Handle(), driver.Uint32(3))

	commandBuffer.CmdSetDeviceMask(3)
}
