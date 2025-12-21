package impl1_2_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_2"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_2"
	"go.uber.org/mock/gomock"
)

func TestVulkanSemaphore_SemaphoreCounterValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_2)
	device := mocks1_2.EasyMockDevice(ctrl, coreDriver)

	builder := &impl1_2.DeviceObjectBuilderImpl{}
	semaphore := builder.CreateSemaphoreObject(coreDriver, device.Handle(), mocks.NewFakeSemaphore(), common.Vulkan1_2).(core1_2.Semaphore)

	coreDriver.EXPECT().VkGetSemaphoreCounterValue(
		device.Handle(),
		semaphore.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		semaphore driver.VkSemaphore,
		pValue *driver.Uint64) (common.VkResult, error) {

		*pValue = driver.Uint64(37)
		return core1_0.VKSuccess, nil
	})

	value, _, err := semaphore.CounterValue()
	require.NoError(t, err)
	require.Equal(t, uint64(37), value)
}
