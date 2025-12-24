package impl1_0_test

import (
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateFence(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)

	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	fenceHandle := mocks.NewFakeFenceHandle()

	mockDriver.EXPECT().VkCreateFence(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkFenceCreateInfo, pAllocator *loader.VkAllocationCallbacks, pFence *loader.VkFence) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(8), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_FENCE_CREATE_SIGNALED_BIT

			*pFence = fenceHandle
			return core1_0.VKSuccess, nil
		})

	fence, _, err := device.CreateFence(nil, core1_0.FenceCreateInfo{
		Flags: core1_0.FenceCreateSignaled,
	})
	require.NoError(t, err)
	require.NotNil(t, fence)
	require.Equal(t, fenceHandle, fence.Handle())
}

func TestVulkanFence_Wait(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)

	device := mocks1_0.EasyMockDevice(ctrl, mockDriver)

	builder := &impl1_0.DeviceObjectBuilderImpl{}
	fence := builder.CreateFenceObject(mockDriver, device.Handle(), mocks.NewFakeFenceHandle(), common.Vulkan1_0)

	mockDriver.EXPECT().VkWaitForFences(device.Handle(), loader.Uint32(1), gomock.Not(nil), loader.VkBool32(1), loader.Uint64(60000000000)).DoAndReturn(
		func(device loader.VkDevice, fenceCount loader.Uint32, pFences *loader.VkFence, waitAll loader.VkBool32, timeout loader.Uint64) (common.VkResult, error) {
			fenceSlice := ([]loader.VkFence)(unsafe.Slice(pFences, 1))
			require.Equal(t, fence.Handle(), fenceSlice[0])

			return core1_0.VKSuccess, nil
		})

	_, err := fence.Wait(time.Minute)
	require.NoError(t, err)
}

func TestVulkanFence_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)

	device := mocks1_0.EasyMockDevice(ctrl, mockDriver)
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	fence := builder.CreateFenceObject(mockDriver, device.Handle(), mocks.NewFakeFenceHandle(), common.Vulkan1_0)

	mockDriver.EXPECT().VkResetFences(device.Handle(), loader.Uint32(1), gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, fenceCount loader.Uint32, pFence *loader.VkFence) (common.VkResult, error) {
			fences := ([]loader.VkFence)(unsafe.Slice(pFence, 1))
			require.Equal(t, fence.Handle(), fences[0])

			return core1_0.VKSuccess, nil
		})

	_, err := fence.Reset()
	require.NoError(t, err)
}

func TestVulkanFence_Status(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks1_0.EasyMockDevice(ctrl, mockDriver)
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	fence := builder.CreateFenceObject(mockDriver, device.Handle(), mocks.NewFakeFenceHandle(), common.Vulkan1_0)

	mockDriver.EXPECT().VkGetFenceStatus(device.Handle(), fence.Handle()).Return(core1_0.VKNotReady, nil)

	res, err := fence.Status()
	require.NoError(t, err)
	require.Equal(t, core1_0.VKNotReady, res)
}
