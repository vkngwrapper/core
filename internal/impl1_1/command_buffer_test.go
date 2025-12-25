package impl1_1_test

import (
	"testing"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestCommandBuffer_CmdDispatchBase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := impl1_1.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	coreLoader.EXPECT().VkCmdDispatchBase(
		commandBuffer.Handle(),
		loader.Uint32(1),
		loader.Uint32(3),
		loader.Uint32(5),
		loader.Uint32(7),
		loader.Uint32(11),
		loader.Uint32(13),
	)

	driver.CmdDispatchBase(commandBuffer, 1, 3, 5, 7, 11, 13)
}

func TestCommandBuffer_CmdSetDeviceMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := impl1_1.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	coreLoader.EXPECT().VkCmdSetDeviceMask(commandBuffer.Handle(), loader.Uint32(3))

	driver.CmdSetDeviceMask(commandBuffer, 3)
}
