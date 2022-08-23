package core1_1_test

import (
	"github.com/golang/mock/gomock"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/core/v2/internal/dummies"
	"github.com/vkngwrapper/core/v2/mocks"
	"testing"
)

func TestVulkanCommandPool_TrimCommandPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := core1_1.PromoteCommandPool(dummies.EasyDummyCommandPool(coreDriver, device))

	coreDriver.EXPECT().VkTrimCommandPool(device.Handle(), commandPool.Handle(), driver.VkCommandPoolTrimFlags(0))

	commandPool.TrimCommandPool(0)
}
