package impl1_1

import (
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
)

type InstanceVulkanDriver struct {
	impl1_0.InstanceVulkanDriver
}

type DeviceVulkanDriver struct {
	impl1_0.DeviceVulkanDriver
}

type CoreVulkanDriver struct {
	InstanceVulkanDriver
	DeviceVulkanDriver
}

var _ core1_1.CoreInstanceDriver = &InstanceVulkanDriver{}
var _ core1_1.DeviceDriver = &DeviceVulkanDriver{}
var _ core1_1.CoreDeviceDriver = &CoreVulkanDriver{}
