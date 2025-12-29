package impl1_0

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

type InstanceDriverFactory func(global *GlobalVulkanDriver, instance core1_0.Instance) (core1_0.CoreInstanceDriver, error)
type DeviceDriverFactory func(instance core1_0.CoreInstanceDriver, device core1_0.Device) (core1_0.CoreDeviceDriver, error)

type GlobalVulkanDriver struct {
	LoaderObj             loader.Loader
	InstanceDriverFactory InstanceDriverFactory
	DeviceDriverFactory   DeviceDriverFactory
}

func (l *GlobalVulkanDriver) Loader() loader.Loader {
	return l.LoaderObj
}

func (l *GlobalVulkanDriver) BuildInstanceDriver(instance core1_0.Instance) (core1_0.CoreInstanceDriver, error) {
	return l.InstanceDriverFactory(l, instance)
}

type InstanceVulkanDriver struct {
	GlobalVulkanDriver
	InstanceObj core1_0.Instance
}

func (l *InstanceVulkanDriver) BuildDeviceDriver(device core1_0.Device) (core1_0.CoreDeviceDriver, error) {
	return l.DeviceDriverFactory(l, device)
}

func (l *InstanceVulkanDriver) Instance() core1_0.Instance {
	return l.InstanceObj
}

type DeviceVulkanDriver struct {
	LoaderObj loader.Loader
	DeviceObj core1_0.Device
}

func (v *DeviceVulkanDriver) Loader() loader.Loader {
	return v.LoaderObj
}

func (v *DeviceVulkanDriver) Device() core1_0.Device {
	return v.DeviceObj
}

type CoreVulkanDriver struct {
	InstanceDriverObj core1_0.CoreInstanceDriver
	DeviceVulkanDriver
}

func (c *CoreVulkanDriver) InstanceDriver() core1_0.CoreInstanceDriver {
	return c.InstanceDriverObj
}

var _ core1_0.GlobalDriver = &GlobalVulkanDriver{}
var _ core1_0.CoreInstanceDriver = &InstanceVulkanDriver{}
var _ core1_0.DeviceDriver = &DeviceVulkanDriver{}
var _ core1_0.CoreDeviceDriver = &CoreVulkanDriver{}
