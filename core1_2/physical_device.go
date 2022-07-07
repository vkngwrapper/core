package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanPhysicalDevice struct {
	core1_1.PhysicalDevice

	InstanceScoped1_2 InstanceScopedPhysicalDevice
}

func (p *VulkanPhysicalDevice) InstanceScopedPhysicalDevice1_2() InstanceScopedPhysicalDevice {
	return p.InstanceScoped1_2
}

func PromotePhysicalDevice(physicalDevice core1_0.PhysicalDevice) PhysicalDevice {
	if !physicalDevice.DeviceAPIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	instanceScoped := PromoteInstanceScopedPhysicalDevice(physicalDevice)
	promotedPhysicalDevice := core1_1.PromotePhysicalDevice(physicalDevice)

	return physicalDevice.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(physicalDevice.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanPhysicalDevice{
				PhysicalDevice: promotedPhysicalDevice,

				InstanceScoped1_2: instanceScoped,
			}
		}).(PhysicalDevice)
}

type VulkanInstanceScopedPhysicalDevice struct {
	core1_1.InstanceScopedPhysicalDevice
}

func PromoteInstanceScopedPhysicalDevice(physicalDevice core1_0.PhysicalDevice) InstanceScopedPhysicalDevice {
	if !physicalDevice.InstanceAPIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedPhysicalDevice := core1_1.PromoteInstanceScopedPhysicalDevice(physicalDevice)
	return physicalDevice.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(physicalDevice.Handle()),
		driver.Core1_2InstanceScope,
		func() any {
			return &VulkanInstanceScopedPhysicalDevice{
				InstanceScopedPhysicalDevice: promotedPhysicalDevice,
			}
		}).(InstanceScopedPhysicalDevice)
}
