package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestVulkanLoader1_0_CreateSemaphore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	semaphoreHandle := mocks.NewFakeSemaphore()

	driver.EXPECT().VkCreateSemaphore(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkSemaphoreCreateInfo, pAllocator *core.VkAllocationCallbacks, pSemaphore *core.VkSemaphore) (core.VkResult, error) {
			*pSemaphore = semaphoreHandle
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(9), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			return core.VKSuccess, nil
		})

	semaphore, _, err := loader.CreateSemaphore(device, &core.SemaphoreOptions{})
	require.NoError(t, err)
	require.NotNil(t, semaphore)
	require.Equal(t, semaphoreHandle, semaphore.Handle())
}
