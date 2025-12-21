package mocks1_1

import (
	unsafe "unsafe"

	common "github.com/vkngwrapper/core/v3/common"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	gomock "go.uber.org/mock/gomock"
)

func MockRig(ctrl *gomock.Controller, deviceVersion common.APIVersion, instanceExtensions []string, deviceExtensions []string) (*MockInstance, *MockPhysicalDevice, *MockDevice) {
	driver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	driver.EXPECT().LoadProcAddr(gomock.Any()).AnyTimes().Return(unsafe.Pointer(nil))

	instance := NewMockInstance(ctrl)
	instance.EXPECT().Handle().Return(mocks.NewFakeInstanceHandle()).AnyTimes()
	instance.EXPECT().Driver().Return(driver).AnyTimes()
	instance.EXPECT().APIVersion().Return(common.Vulkan1_1).AnyTimes()
	instance.EXPECT().IsInstanceExtensionActive(gomock.Any()).AnyTimes().DoAndReturn(
		func(extension string) bool {
			for _, ext := range instanceExtensions {
				if ext == extension {
					return true
				}
			}

			return false
		},
	)

	physDeviceHandle := mocks.NewFakePhysicalDeviceHandle()

	physicalDevice := NewMockPhysicalDevice(ctrl)
	physicalDevice.EXPECT().Handle().Return(physDeviceHandle).AnyTimes()
	physicalDevice.EXPECT().Driver().Return(driver).AnyTimes()
	physicalDevice.EXPECT().DeviceAPIVersion().Return(deviceVersion).AnyTimes()
	physicalDevice.EXPECT().InstanceAPIVersion().Return(common.Vulkan1_1).AnyTimes()

	device := NewMockDevice(ctrl)
	device.EXPECT().Handle().Return(mocks.NewFakeDeviceHandle()).AnyTimes()
	device.EXPECT().Driver().AnyTimes().Return(driver)
	device.EXPECT().APIVersion().AnyTimes().Return(deviceVersion)
	device.EXPECT().IsDeviceExtensionActive(gomock.Any()).AnyTimes().DoAndReturn(
		func(extension string) bool {
			for _, ext := range deviceExtensions {
				if ext == extension {
					return true
				}
			}

			return false
		},
	)

	return instance, physicalDevice, device
}
