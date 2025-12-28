package impl1_2

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
)

type InstanceVulkanDriver struct {
	impl1_1.InstanceVulkanDriver
}

type DeviceVulkanDriver struct {
	impl1_1.DeviceVulkanDriver
}

type CoreVulkanDriver struct {
	InstanceDriverObj core1_0.CoreInstanceDriver
	DeviceVulkanDriver
}

func (c *CoreVulkanDriver) InstanceDriver() core1_0.CoreInstanceDriver {
	return c.InstanceDriverObj
}

var _ core1_2.CoreInstanceDriver = &InstanceVulkanDriver{}
var _ core1_2.DeviceDriver = &DeviceVulkanDriver{}
var _ core1_2.CoreDeviceDriver = &CoreVulkanDriver{}
