package impl1_2_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_2"
	"go.uber.org/mock/gomock"
)

func TestVulkanSemaphore_SemaphoreCounterValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_2, []string{})
	semaphore := mocks.NewDummySemaphore(device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_2)
	driver := mocks1_2.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkGetSemaphoreCounterValue(
		device.Handle(),
		semaphore.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		semaphore loader.VkSemaphore,
		pValue *loader.Uint64) (common.VkResult, error) {

		*pValue = loader.Uint64(37)
		return core1_0.VKSuccess, nil
	})

	value, _, err := driver.GetSemaphoreCounterValue(semaphore)
	require.NoError(t, err)
	require.Equal(t, uint64(37), value)
}
