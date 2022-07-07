package core1_0_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	mock_driver "github.com/vkngwrapper/core/driver/mocks"
	internal_mocks "github.com/vkngwrapper/core/internal/dummies"
	"github.com/vkngwrapper/core/mocks"
	"reflect"
	"testing"
)

func TestVulkanLoader1_0_CreateSemaphore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)
	semaphoreHandle := mocks.NewFakeSemaphore()

	mockDriver.EXPECT().VkCreateSemaphore(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkSemaphoreCreateInfo, pAllocator *driver.VkAllocationCallbacks, pSemaphore *driver.VkSemaphore) (common.VkResult, error) {
			*pSemaphore = semaphoreHandle
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(9), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			return core1_0.VKSuccess, nil
		})

	semaphore, _, err := device.CreateSemaphore(nil, core1_0.SemaphoreCreateInfo{})
	require.NoError(t, err)
	require.NotNil(t, semaphore)
	require.Equal(t, semaphoreHandle, semaphore.Handle())
}
