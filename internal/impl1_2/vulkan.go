package impl1_2

import (
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
	InstanceVulkanDriver
	DeviceVulkanDriver
}

var _ core1_2.CoreInstanceDriver = &InstanceVulkanDriver{}
var _ core1_2.DeviceDriver = &DeviceVulkanDriver{}
var _ core1_2.CoreDeviceDriver = &CoreVulkanDriver{}
