package mocks1_2

import (
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/internal/impl1_2"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
)

func InternalCoreInstanceDriver(loader loader.Loader) core1_2.CoreInstanceDriver {
	core11 := mocks1_1.InternalCoreInstanceDriver(loader).(*impl1_1.InstanceVulkanDriver)
	return &impl1_2.InstanceVulkanDriver{
		InstanceVulkanDriver: *core11,
	}
}

func InternalDeviceDriver(loader loader.Loader) core1_2.DeviceDriver {
	core10 := mocks1_1.InternalDeviceDriver(loader).(*impl1_1.DeviceVulkanDriver)
	return &impl1_2.DeviceVulkanDriver{
		DeviceVulkanDriver: *core10,
	}
}

func InternalCoreDriver(loader loader.Loader) core1_2.CoreDeviceDriver {
	instance := InternalCoreInstanceDriver(loader).(*impl1_2.InstanceVulkanDriver)
	device := InternalDeviceDriver(loader).(*impl1_2.DeviceVulkanDriver)
	return &impl1_2.CoreVulkanDriver{
		InstanceVulkanDriver: *instance,
		DeviceVulkanDriver:   *device,
	}
}
