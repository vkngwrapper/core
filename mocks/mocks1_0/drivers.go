package mocks1_0

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func InternalGlobalDriver(loader loader.Loader) core1_0.GlobalDriver {
	return &impl1_0.GlobalVulkanDriver{
		LoaderObj: loader,
		InstanceDriverFactory: func(global *impl1_0.GlobalVulkanDriver, instance core1_0.Instance) (core1_0.CoreInstanceDriver, error) {
			return InternalCoreInstanceDriver(instance, loader), nil
		},
		DeviceDriverFactory: func(driver core1_0.CoreInstanceDriver, device core1_0.Device) (core1_0.CoreDeviceDriver, error) {
			return InternalCoreDriver(driver.Instance(), device, driver.Loader()), nil
		},
	}
}

func InternalCoreInstanceDriver(instance core1_0.Instance, loader loader.Loader) core1_0.CoreInstanceDriver {
	global := InternalGlobalDriver(loader).(*impl1_0.GlobalVulkanDriver)
	return &impl1_0.InstanceVulkanDriver{
		GlobalVulkanDriver: *global,
		InstanceObj:        instance,
	}
}

func InternalDeviceDriver(device core1_0.Device, loader loader.Loader) core1_0.DeviceDriver {
	return &impl1_0.DeviceVulkanDriver{
		LoaderObj: loader,
		DeviceObj: device,
	}
}

func InternalCoreDriver(instance core1_0.Instance, device core1_0.Device, loader loader.Loader) core1_0.CoreDeviceDriver {
	instanceDriver := InternalCoreInstanceDriver(instance, loader)
	deviceDriver := InternalDeviceDriver(device, loader).(*impl1_0.DeviceVulkanDriver)
	return &impl1_0.CoreVulkanDriver{
		InstanceDriverObj:  instanceDriver,
		DeviceVulkanDriver: *deviceDriver,
	}
}
