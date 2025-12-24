package impl1_2

import (
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
	"github.com/vkngwrapper/core/v3/loader"
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

func NewInstanceDriver(loader loader.Loader) *InstanceVulkanDriver {
	return &InstanceVulkanDriver{
		InstanceVulkanDriver: impl1_1.InstanceVulkanDriver{
			InstanceVulkanDriver: impl1_0.InstanceVulkanDriver{
				GlobalVulkanDriver: impl1_0.GlobalVulkanDriver{LoaderObj: loader},
			}},
	}
}

func NewDeviceDriver(loader loader.Loader) *DeviceVulkanDriver {
	return &DeviceVulkanDriver{
		DeviceVulkanDriver: impl1_1.DeviceVulkanDriver{
			DeviceVulkanDriver: impl1_0.DeviceVulkanDriver{
				LoaderObj: loader,
			},
		},
	}
}

func NewCoreDriver(loader loader.Loader) *CoreVulkanDriver {
	return &CoreVulkanDriver{
		InstanceVulkanDriver{
			InstanceVulkanDriver: impl1_1.InstanceVulkanDriver{
				InstanceVulkanDriver: impl1_0.InstanceVulkanDriver{
					GlobalVulkanDriver: impl1_0.GlobalVulkanDriver{
						LoaderObj: loader,
					},
				},
			},
		},
		DeviceVulkanDriver{
			DeviceVulkanDriver: impl1_1.DeviceVulkanDriver{
				DeviceVulkanDriver: impl1_0.DeviceVulkanDriver{
					LoaderObj: loader,
				},
			},
		},
	}
}
