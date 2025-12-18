package core1_2_test

import (
	"testing"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/dummies"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
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
