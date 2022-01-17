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
	"time"
	"unsafe"
)

func TestVulkanLoader1_0_CreateFence(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	fenceHandle := mocks.NewFakeFenceHandle()

	mockDriver.EXPECT().VkCreateFence(mocks.Exactly(device.Handle()), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkFenceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pFence *driver.VkFence) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(8), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_FENCE_CREATE_SIGNALED_BIT

			*pFence = fenceHandle
			return common.VKSuccess, nil
		})

	fence, _, err := loader.CreateFence(device, nil, &core.FenceOptions{
		Flags: core.FenceSignaled,
	})
	require.NoError(t, err)
	require.NotNil(t, fence)
	require.Same(t, fenceHandle, fence.Handle())
}

func TestVulkanFence_Wait(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	fence := mocks.EasyDummyFence(t, loader, device)

	mockDriver.EXPECT().VkWaitForFences(mocks.Exactly(device.Handle()), driver.Uint32(1), gomock.Not(nil), driver.VkBool32(1), driver.Uint64(60000000000)).DoAndReturn(
		func(device driver.VkDevice, fenceCount driver.Uint32, pFences *driver.VkFence, waitAll driver.VkBool32, timeout driver.Uint64) (common.VkResult, error) {
			fenceSlice := ([]driver.VkFence)(unsafe.Slice(pFences, 1))
			require.Same(t, fence.Handle(), fenceSlice[0])

			return common.VKSuccess, nil
		})

	_, err = fence.Wait(time.Minute)
	require.NoError(t, err)
}

func TestVulkanFence_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	fence := mocks.EasyDummyFence(t, loader, device)

	mockDriver.EXPECT().VkResetFences(mocks.Exactly(device.Handle()), driver.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, fenceCount driver.Uint32, pFence *driver.VkFence) (common.VkResult, error) {
			fences := ([]driver.VkFence)(unsafe.Slice(pFence, 1))
			require.Same(t, fence.Handle(), fences[0])

			return common.VKSuccess, nil
		})

	_, err = fence.Reset()
	require.NoError(t, err)
}

func TestVulkanFence_Status(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(mockDriver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	fence := mocks.EasyDummyFence(t, loader, device)

	mockDriver.EXPECT().VkGetFenceStatus(mocks.Exactly(device.Handle()), mocks.Exactly(fence.Handle())).Return(common.VKNotReady, nil)

	res, err := fence.Status()
	require.NoError(t, err)
	require.Equal(t, common.VKNotReady, res)
}
