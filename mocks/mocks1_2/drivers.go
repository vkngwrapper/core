package mocks1_2

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/internal/impl1_2"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
)

func InternalCoreInstanceDriver(instance core.Instance, loader loader.Loader) core1_2.CoreInstanceDriver {
	core11 := mocks1_1.InternalCoreInstanceDriver(instance, loader).(*impl1_1.InstanceVulkanDriver)
	return &impl1_2.InstanceVulkanDriver{
		InstanceVulkanDriver: *core11,
	}
}

func InternalDeviceDriver(device core.Device, loader loader.Loader) core1_2.DeviceDriver {
	core10 := mocks1_1.InternalDeviceDriver(device, loader).(*impl1_1.DeviceVulkanDriver)
	return &impl1_2.DeviceVulkanDriver{
		DeviceVulkanDriver: *core10,
	}
}

func InternalCoreDriver(instance core.Instance, device core.Device, loader loader.Loader) core1_2.CoreDeviceDriver {
	instanceDriver := InternalCoreInstanceDriver(instance, loader).(*impl1_2.InstanceVulkanDriver)
	deviceDriver := InternalDeviceDriver(device, loader).(*impl1_2.DeviceVulkanDriver)
	return &impl1_2.CoreVulkanDriver{
		InstanceDriverObj:  instanceDriver,
		DeviceVulkanDriver: *deviceDriver,
	}
}
