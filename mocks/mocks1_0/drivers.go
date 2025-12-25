package mocks1_0

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func InternalGlobalDriver(loader loader.Loader) core1_0.GlobalDriver {
	return &impl1_0.GlobalVulkanDriver{
		LoaderObj: loader,
		InstanceDriverFactory: func(global *impl1_0.GlobalVulkanDriver, instance core.Instance) (core1_0.CoreInstanceDriver, error) {
			return InternalCoreInstanceDriver(loader), nil
		},
		DeviceDriverFactory: func(global *impl1_0.GlobalVulkanDriver, device core.Device) (core1_0.CoreDeviceDriver, error) {
			return InternalCoreDriver(loader), nil
		},
	}
}

func InternalCoreInstanceDriver(loader loader.Loader) core1_0.CoreInstanceDriver {
	global := InternalGlobalDriver(loader).(*impl1_0.GlobalVulkanDriver)
	return &impl1_0.InstanceVulkanDriver{
		GlobalVulkanDriver: *global,
	}
}

func InternalDeviceDriver(loader loader.Loader) core1_0.DeviceDriver {
	return &impl1_0.DeviceVulkanDriver{
		LoaderObj: loader,
	}
}

func InternalCoreDriver(loader loader.Loader) core1_0.CoreDeviceDriver {
	instance := InternalCoreInstanceDriver(loader).(*impl1_0.InstanceVulkanDriver)
	device := InternalDeviceDriver(loader).(*impl1_0.DeviceVulkanDriver)
	return &impl1_0.CoreVulkanDriver{
		InstanceVulkanDriver: *instance,
		DeviceVulkanDriver:   *device,
	}
}
