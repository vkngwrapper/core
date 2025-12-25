package impl1_0

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

type InstanceDriverFactory func(global *GlobalVulkanDriver, instance core.Instance) (core1_0.CoreInstanceDriver, error)
type DeviceDriverFactory func(global *GlobalVulkanDriver, device core.Device) (core1_0.CoreDeviceDriver, error)

type GlobalVulkanDriver struct {
	LoaderObj             loader.Loader
	InstanceDriverFactory InstanceDriverFactory
	DeviceDriverFactory   DeviceDriverFactory
}

func (l *GlobalVulkanDriver) Loader() loader.Loader {
	return l.LoaderObj
}

func (l *GlobalVulkanDriver) BuildInstanceDriver(instance core.Instance) (core1_0.CoreInstanceDriver, error) {
	return l.InstanceDriverFactory(l, instance)
}

type InstanceVulkanDriver struct {
	GlobalVulkanDriver
}

func (l *InstanceVulkanDriver) BuildDeviceDriver(device core.Device) (core1_0.CoreDeviceDriver, error) {
	return l.DeviceDriverFactory(&l.GlobalVulkanDriver, device)
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
