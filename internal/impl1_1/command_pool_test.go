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

func TestVulkanCommandPool_TrimCommandPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_1)
	device := mocks1_1.EasyMockDevice(ctrl, coreDriver)
	builder := &impl1_1.DeviceObjectBuilderImpl{}
	commandPool := builder.CreateCommandPoolObject(coreDriver, device.Handle(), mocks.NewFakeCommandPoolHandle(), common.Vulkan1_1).(core1_1.CommandPool)

	coreDriver.EXPECT().VkTrimCommandPool(device.Handle(), commandPool.Handle(), loader.VkCommandPoolTrimFlags(0))

	commandPool.TrimCommandPool(0)
}
