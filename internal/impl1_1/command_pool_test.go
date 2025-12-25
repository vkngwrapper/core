package impl1_1_test

import (
	"testing"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"go.uber.org/mock/gomock"
)

func TestVulkanCommandPool_TrimCommandPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_1)
	driver := mocks1_1.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	commandPool := mocks.NewDummyCommandPool(device)

	coreLoader.EXPECT().VkTrimCommandPool(device.Handle(), commandPool.Handle(), loader.VkCommandPoolTrimFlags(0))

	driver.TrimCommandPool(commandPool, 0)
}
