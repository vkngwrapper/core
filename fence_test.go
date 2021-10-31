package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestVulkanLoader1_0_CreateFence(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	fenceHandle := mocks.NewFakeFenceHandle()

	driver.EXPECT().VkCreateFence(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkFenceCreateInfo, pAllocator *core.VkAllocationCallbacks, pFence *core.VkFence) (core.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(8), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_FENCE_CREATE_SIGNALED_BIT

			*pFence = fenceHandle
			return core.VKSuccess, nil
		})

	fence, _, err := loader.CreateFence(device, &core.FenceOptions{
		Flags: core.FenceSignaled,
	})
	require.NoError(t, err)
	require.NotNil(t, fence)
	require.Equal(t, fenceHandle, fence.Handle())
}
