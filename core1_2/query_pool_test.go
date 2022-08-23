package core1_2_test

import (
	"github.com/golang/mock/gomock"
	"github.com/vkngwrapper/core/v2/internal/dummies"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_2"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
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
