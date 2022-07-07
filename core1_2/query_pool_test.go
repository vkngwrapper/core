package core1_2_test

import (
	"github.com/golang/mock/gomock"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_2"
	"github.com/vkngwrapper/core/driver"
	mock_driver "github.com/vkngwrapper/core/driver/mocks"
	"github.com/vkngwrapper/core/internal/dummies"
	"github.com/vkngwrapper/core/mocks"
	"testing"
)

func TestVulkanQueryPool_ResetQueryPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	queryPool := core1_2.PromoteQueryPool(dummies.EasyDummyQueryPool(coreDriver, device))

	coreDriver.EXPECT().VkResetQueryPool(
		device.Handle(),
		queryPool.Handle(),
		driver.Uint32(1),
		driver.Uint32(3),
	)

	queryPool.Reset(1, 3)
}
