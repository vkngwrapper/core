package core

import (
	"unsafe"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/internal/impl1_2"
	"github.com/vkngwrapper/core/v3/loader"
)

func CreateDriverFromProcAddr(procAddr unsafe.Pointer) (core1_0.GlobalDriver, error) {
	loaderObj, err := loader.CreateLoaderFromProcAddr(procAddr)
	if err != nil {
		return nil, err
	}

	return &impl1_0.GlobalVulkanDriver{
		LoaderObj:             loaderObj,
		InstanceDriverFactory: buildInstanceDriver,
		DeviceDriverFactory:   buildDeviceDriver,
	}, nil
}

func buildInstanceDriver(driver *impl1_0.GlobalVulkanDriver, instance core1_0.Instance) (core1_0.CoreInstanceDriver, error) {
	loaderObj, err := driver.LoaderObj.CreateInstanceLoader(instance.Handle())
	if err != nil {
		return nil, err
	}

	switch instance.APIVersion() {
	case common.Vulkan1_2:
		return &impl1_2.InstanceVulkanDriver{
			InstanceVulkanDriver: impl1_1.InstanceVulkanDriver{
				InstanceVulkanDriver: impl1_0.InstanceVulkanDriver{
					GlobalVulkanDriver: impl1_0.GlobalVulkanDriver{
						LoaderObj:             loaderObj,
						InstanceDriverFactory: buildInstanceDriver,
						DeviceDriverFactory:   buildDeviceDriver,
					},
					InstanceObj: instance,
				},
			},
		}, nil
	case common.Vulkan1_1:
		return &impl1_1.InstanceVulkanDriver{
			InstanceVulkanDriver: impl1_0.InstanceVulkanDriver{
				GlobalVulkanDriver: impl1_0.GlobalVulkanDriver{
					LoaderObj:             loaderObj,
					InstanceDriverFactory: buildInstanceDriver,
					DeviceDriverFactory:   buildDeviceDriver,
				},
				InstanceObj: instance,
			},
		}, nil
	default:
		return &impl1_0.InstanceVulkanDriver{
			GlobalVulkanDriver: impl1_0.GlobalVulkanDriver{
				LoaderObj:             loaderObj,
				InstanceDriverFactory: buildInstanceDriver,
				DeviceDriverFactory:   buildDeviceDriver,
			},
			InstanceObj: instance,
		}, nil
	}
}

func buildDeviceDriver(driver core1_0.CoreInstanceDriver, device core1_0.Device) (core1_0.CoreDeviceDriver, error) {
	loaderObj, err := driver.Loader().CreateDeviceLoader(device.Handle())
	if err != nil {
		return nil, err
	}

	switch device.APIVersion() {
	case common.Vulkan1_2:
		return &impl1_2.CoreVulkanDriver{
			InstanceDriverObj: driver,
			DeviceVulkanDriver: impl1_2.DeviceVulkanDriver{
				DeviceVulkanDriver: impl1_1.DeviceVulkanDriver{
					DeviceVulkanDriver: impl1_0.DeviceVulkanDriver{
						LoaderObj: loaderObj,
						DeviceObj: device,
					},
				},
			},
		}, nil
	case common.Vulkan1_1:
		return &impl1_1.CoreVulkanDriver{
			InstanceDriverObj: driver,
			DeviceVulkanDriver: impl1_1.DeviceVulkanDriver{
				DeviceVulkanDriver: impl1_0.DeviceVulkanDriver{
					LoaderObj: loaderObj,
					DeviceObj: device,
				},
			},
		}, nil
	default:
		return &impl1_0.CoreVulkanDriver{
			InstanceDriverObj: driver,
			DeviceVulkanDriver: impl1_0.DeviceVulkanDriver{
				LoaderObj: loaderObj,
				DeviceObj: device,
			},
		}, nil
	}
}
