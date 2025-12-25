package impl1_2_test

import (
	"testing"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/internal/impl1_2"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanQueryPool_ResetQueryPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := impl1_2.NewDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	queryPool := mocks.NewDummyQueryPool(device)

	coreLoader.EXPECT().VkResetQueryPool(
		device.Handle(),
		queryPool.Handle(),
		loader.Uint32(1),
		loader.Uint32(3),
	)

	driver.ResetQueryPool(queryPool, 1, 3)
}
