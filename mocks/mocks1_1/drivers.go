package mocks1_1

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
)

func InternalCoreInstanceDriver(instance core.Instance, loader loader.Loader) core1_1.CoreInstanceDriver {
	core10 := mocks1_0.InternalCoreInstanceDriver(instance, loader).(*impl1_0.InstanceVulkanDriver)
	return &impl1_1.InstanceVulkanDriver{
		InstanceVulkanDriver: *core10,
	}
}

func InternalDeviceDriver(device core.Device, loader loader.Loader) core1_1.DeviceDriver {
	core10 := mocks1_0.InternalDeviceDriver(device, loader).(*impl1_0.DeviceVulkanDriver)
	return &impl1_1.DeviceVulkanDriver{
		DeviceVulkanDriver: *core10,
	}
}

func InternalCoreDriver(instance core.Instance, device core.Device, loader loader.Loader) core1_1.CoreDeviceDriver {
	instanceDriver := InternalCoreInstanceDriver(instance, loader)
	deviceDriver := InternalDeviceDriver(device, loader).(*impl1_1.DeviceVulkanDriver)
	return &impl1_1.CoreVulkanDriver{
		InstanceDriverObj:  instanceDriver,
		DeviceVulkanDriver: *deviceDriver,
	}
}
