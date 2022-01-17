package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestVulkanLoader1_0_CreateSemaphore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	semaphoreHandle := mocks.NewFakeSemaphore()

	mockDriver.EXPECT().VkCreateSemaphore(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkSemaphoreCreateInfo, pAllocator *driver.VkAllocationCallbacks, pSemaphore *driver.VkSemaphore) (common.VkResult, error) {
			*pSemaphore = semaphoreHandle
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(9), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			return common.VKSuccess, nil
		})

	semaphore, _, err := loader.CreateSemaphore(device, nil, &core.SemaphoreOptions{})
	require.NoError(t, err)
	require.NotNil(t, semaphore)
	require.Same(t, semaphoreHandle, semaphore.Handle())
}
