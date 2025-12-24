package impl1_1_test

import (
	"testing"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestCommandBuffer_CmdDispatchBase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_1)
	device := mocks1_1.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks1_1.EasyMockCommandPool(ctrl, device)
	builder := &impl1_1.DeviceObjectBuilderImpl{}
	commandBuffer := builder.CreateCommandBufferObject(coreDriver, commandPool.Handle(), device.Handle(), mocks.NewFakeCommandBufferHandle(), common.Vulkan1_1).(core1_1.CommandBuffer)

	coreDriver.EXPECT().VkCmdDispatchBase(
		commandBuffer.Handle(),
		loader.Uint32(1),
		loader.Uint32(3),
		loader.Uint32(5),
		loader.Uint32(7),
		loader.Uint32(11),
		loader.Uint32(13),
	)

	commandBuffer.CmdDispatchBase(1, 3, 5, 7, 11, 13)
}

func TestCommandBuffer_CmdSetDeviceMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_1)
	device := mocks1_1.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks1_1.EasyMockCommandPool(ctrl, device)
	builder := &impl1_1.DeviceObjectBuilderImpl{}
	commandBuffer := builder.CreateCommandBufferObject(coreDriver, commandPool.Handle(), device.Handle(), mocks.NewFakeCommandBufferHandle(), common.Vulkan1_1).(core1_1.CommandBuffer)

	coreDriver.EXPECT().VkCmdSetDeviceMask(commandBuffer.Handle(), loader.Uint32(3))

	commandBuffer.CmdSetDeviceMask(3)
}
