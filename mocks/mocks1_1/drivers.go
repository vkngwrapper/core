package mocks1_1

import (
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
)

func InternalCoreInstanceDriver(loader loader.Loader) core1_1.CoreInstanceDriver {
	core10 := mocks1_0.InternalCoreInstanceDriver(loader).(*impl1_0.InstanceVulkanDriver)
	return &impl1_1.InstanceVulkanDriver{
		InstanceVulkanDriver: *core10,
	}
}

func InternalDeviceDriver(loader loader.Loader) core1_1.DeviceDriver {
	core10 := mocks1_0.InternalDeviceDriver(loader).(*impl1_0.DeviceVulkanDriver)
	return &impl1_1.DeviceVulkanDriver{
		DeviceVulkanDriver: *core10,
	}
}

func InternalCoreDriver(loader loader.Loader) core1_1.CoreDeviceDriver {
	instance := InternalCoreInstanceDriver(loader).(*impl1_1.InstanceVulkanDriver)
	device := InternalDeviceDriver(loader).(*impl1_1.DeviceVulkanDriver)
	return &impl1_1.CoreVulkanDriver{
		InstanceVulkanDriver: *instance,
		DeviceVulkanDriver:   *device,
	}
}
