package core_test

import (
	"github.com/CannibalVox/VKng/core"
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

func TestVulkanFence_Wait(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	fence := mocks.EasyDummyFence(t, loader, device)

	driver.EXPECT().VkWaitForFences(device.Handle(), core.Uint32(1), gomock.Not(nil), core.VkBool32(1), core.Uint64(60000000000)).DoAndReturn(
		func(device core.VkDevice, fenceCount core.Uint32, pFences *core.VkFence, waitAll core.VkBool32, timeout core.Uint64) (core.VkResult, error) {
			fenceSlice := ([]core.VkFence)(unsafe.Slice(pFences, 1))
			require.Equal(t, fence.Handle(), fenceSlice[0])

			return core.VKSuccess, nil
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

	mockDriver.EXPECT().VkResetFences(device.Handle(), core.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, fenceCount core.Uint32, pFence *core.VkFence) (core.VkResult, error) {
			fences := ([]core.VkFence)(unsafe.Slice(pFence, 1))
			require.Equal(t, fence.Handle(), fences[0])

			return core.VKSuccess, nil
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

	mockDriver.EXPECT().VkGetFenceStatus(device.Handle(), fence.Handle()).Return(core.VKNotReady, nil)

	res, err := fence.Status()
	require.NoError(t, err)
	require.Equal(t, core.VKNotReady, res)
}
