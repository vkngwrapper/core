package impl1_0_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	eventHandle := mocks.NewFakeEventHandle()

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkCreateEvent(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkEventCreateInfo, pAllocator *loader.VkAllocationCallbacks, pEvent *loader.VkEvent) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(10), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EVENT_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			*pEvent = eventHandle
			return core1_0.VKSuccess, nil
		})

	event, _, err := driver.CreateEvent(nil, core1_0.EventCreateInfo{
		Flags: 0,
	})
	require.NoError(t, err)
	require.NotNil(t, event)
	require.Equal(t, eventHandle, event.Handle())
}

func TestVulkanEvent_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	event := mocks.NewDummyEvent(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkSetEvent(device.Handle(), event.Handle()).Return(core1_0.VKSuccess, nil)

	_, err := driver.SetEvent(event)
	require.NoError(t, err)
}

func TestVulkanEvent_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	event := mocks.NewDummyEvent(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkResetEvent(device.Handle(), event.Handle()).Return(core1_0.VKSuccess, nil)

	_, err := driver.ResetEvent(event)
	require.NoError(t, err)
}

func TestVulkanEvent_Status(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	event := mocks.NewDummyEvent(device)

	mockLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, mockLoader)

	mockLoader.EXPECT().VkGetEventStatus(device.Handle(), event.Handle()).Return(core1_0.VKEventReset, nil)

	res, err := driver.GetEventStatus(event)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKEventReset, res)
}
