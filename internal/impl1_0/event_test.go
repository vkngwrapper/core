package impl1_0_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"go.uber.org/mock/gomock"
)

func TestVulkanLoader1_0_CreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	device := builder.CreateDeviceObject(mockDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0, []string{})
	eventHandle := mocks.NewFakeEventHandle()

	mockDriver.EXPECT().VkCreateEvent(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkEventCreateInfo, pAllocator *driver.VkAllocationCallbacks, pEvent *driver.VkEvent) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(10), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EVENT_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			*pEvent = eventHandle
			return core1_0.VKSuccess, nil
		})

	event, _, err := device.CreateEvent(nil, core1_0.EventCreateInfo{
		Flags: 0,
	})
	require.NoError(t, err)
	require.NotNil(t, event)
	require.Equal(t, eventHandle, event.Handle())
}

func TestVulkanEvent_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks1_0.EasyMockDevice(ctrl, driver)
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	event := builder.CreateEventObject(driver, device.Handle(), mocks.NewFakeEventHandle(), common.Vulkan1_0)

	driver.EXPECT().VkSetEvent(device.Handle(), event.Handle()).Return(core1_0.VKSuccess, nil)

	_, err := event.Set()
	require.NoError(t, err)
}

func TestVulkanEvent_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks1_0.EasyMockDevice(ctrl, driver)
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	event := builder.CreateEventObject(driver, device.Handle(), mocks.NewFakeEventHandle(), common.Vulkan1_0)

	driver.EXPECT().VkResetEvent(device.Handle(), event.Handle()).Return(core1_0.VKSuccess, nil)

	_, err := event.Reset()
	require.NoError(t, err)
}

func TestVulkanEvent_Status(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks1_0.EasyMockDevice(ctrl, driver)
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	event := builder.CreateEventObject(driver, device.Handle(), mocks.NewFakeEventHandle(), common.Vulkan1_0)

	driver.EXPECT().VkGetEventStatus(device.Handle(), event.Handle()).Return(core1_0.VKEventReset, nil)

	res, err := event.Status()
	require.NoError(t, err)
	require.Equal(t, core1_0.VKEventReset, res)
}
