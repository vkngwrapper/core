package impl1_0

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

type GlobalVulkanDriver struct {
	LoaderObj loader.Loader
}

func (l *GlobalVulkanDriver) Loader() loader.Loader {
	return l.LoaderObj
}

type InstanceVulkanDriver struct {
	GlobalVulkanDriver
}

type DeviceVulkanDriver struct {
	LoaderObj loader.Loader
}

func (v *DeviceVulkanDriver) Loader() loader.Loader {
	return v.LoaderObj
}

type CoreVulkanDriver struct {
	InstanceVulkanDriver
	DeviceVulkanDriver
}

var _ core1_0.GlobalDriver = &GlobalVulkanDriver{}
var _ core1_0.CoreInstanceDriver = &InstanceVulkanDriver{}
var _ core1_0.DeviceDriver = &DeviceVulkanDriver{}
var _ core1_0.CoreDeviceDriver = &CoreVulkanDriver{}

func NewGlobalDriver(loader loader.Loader) *GlobalVulkanDriver {
	return &GlobalVulkanDriver{LoaderObj: loader}
}

func NewInstanceDriver(loader loader.Loader) *InstanceVulkanDriver {
	return &InstanceVulkanDriver{GlobalVulkanDriver{LoaderObj: loader}}
}

func NewDeviceDriver(loader loader.Loader) *DeviceVulkanDriver {
	return &DeviceVulkanDriver{LoaderObj: loader}
}

func NewCoreDriver(loader loader.Loader) *CoreVulkanDriver {
	return &CoreVulkanDriver{
		InstanceVulkanDriver{GlobalVulkanDriver{LoaderObj: loader}},
		DeviceVulkanDriver{LoaderObj: loader},
	}
}
