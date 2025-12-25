package impl1_0_test

import (
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateFence(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	fenceHandle := mocks.NewFakeFenceHandle()

	mockLoader.EXPECT().VkCreateFence(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkFenceCreateInfo, pAllocator *loader.VkAllocationCallbacks, pFence *loader.VkFence) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(8), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_FENCE_CREATE_SIGNALED_BIT

			*pFence = fenceHandle
			return core1_0.VKSuccess, nil
		})

	fence, _, err := driver.CreateFence(device, nil, core1_0.FenceCreateInfo{
		Flags: core1_0.FenceCreateSignaled,
	})
	require.NoError(t, err)
	require.NotNil(t, fence)
	require.Equal(t, fenceHandle, fence.Handle())
}

func TestVulkanFence_Wait(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	mockLoader.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	driver := mocks1_0.InternalDeviceDriver(mockLoader)

	fence := mocks.NewDummyFence(device)

	mockLoader.EXPECT().VkWaitForFences(device.Handle(), loader.Uint32(1), gomock.Not(nil), loader.VkBool32(1), loader.Uint64(60000000000)).DoAndReturn(
		func(device loader.VkDevice, fenceCount loader.Uint32, pFences *loader.VkFence, waitAll loader.VkBool32, timeout loader.Uint64) (common.VkResult, error) {
			fenceSlice := ([]loader.VkFence)(unsafe.Slice(pFences, 1))
			require.Equal(t, fence.Handle(), fenceSlice[0])

			return core1_0.VKSuccess, nil
		})

	_, err := driver.WaitForFences(true, time.Minute, fence)
	require.NoError(t, err)
}

func TestVulkanFence_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	mockLoader.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	fence := mocks.NewDummyFence(device)

	mockLoader.EXPECT().VkResetFences(device.Handle(), loader.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, fenceCount loader.Uint32, pFence *loader.VkFence) (common.VkResult, error) {
			fences := ([]loader.VkFence)(unsafe.Slice(pFence, 1))
			require.Equal(t, fence.Handle(), fences[0])

			return core1_0.VKSuccess, nil
		})

	_, err := driver.ResetFences(fence)
	require.NoError(t, err)
}

func TestVulkanFence_Status(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(mockLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	fence := mocks.NewDummyFence(device)

	mockLoader.EXPECT().VkGetFenceStatus(device.Handle(), fence.Handle()).Return(core1_0.VKNotReady, nil)

	res, err := driver.GetFenceStatus(fence)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKNotReady, res)
}
