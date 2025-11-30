package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanPhysicalDevice is an implementation of the PhysicalDevice interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPhysicalDevice struct {
	core1_1.PhysicalDevice

	InstanceScoped1_2 InstanceScopedPhysicalDevice
}

func (p *VulkanPhysicalDevice) InstanceScopedPhysicalDevice1_2() InstanceScopedPhysicalDevice {
	return p.InstanceScoped1_2
}

// PromotePhysicalDevice accepts a PhysicalDevice object from any core version. If provided a physical device that supports
// at least core 1.2 for its device-scoped functionality, it will return a core1_2.PhysicalDevice. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanPhysicalDevice, even if it is provided a VulkanPhysicalDevice from a higher
// core version. Two Vulkan 1.2 compatible PhysicalDevice objects with the same PhysicalDevice.Handle will
// return the same interface value when passed to this method.
func PromotePhysicalDevice(physicalDevice core1_0.PhysicalDevice) PhysicalDevice {
	if physicalDevice == nil {
		return nil
	}
	if !physicalDevice.DeviceAPIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promoted, alreadyPromoted := physicalDevice.(PhysicalDevice)
	if alreadyPromoted {
		return promoted
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

// VulkanInstanceScopedPhysicalDevice is an implementation of the InstanceScopedPhysicalDevice interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanInstanceScopedPhysicalDevice struct {
	core1_1.InstanceScopedPhysicalDevice
}

// PromoteInstanceScopedPhysicalDevice accepts a InstanceScopedPhysicalDevice object from any core version. If provided an instance-scoped physical device that supports
// at least core 1.2 for its instance-scoped functionality, it will return a core1_2.InstanceScopedPhysicalDevice. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanInstanceScopedPhysicalDevice, even if it is provided a VulkanInstanceScopedPhysicalDevice from a higher
// core version. Two Vulkan 1.2 compatible InstanceScopedPhysicalDevice objects with the same InstanceScopedPhysicalDevice.Handle will
// return the same interface value when passed to this method.
func PromoteInstanceScopedPhysicalDevice(physicalDevice core1_0.PhysicalDevice) InstanceScopedPhysicalDevice {
	if physicalDevice == nil {
		return nil
	}
	if !physicalDevice.InstanceAPIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promoted, alreadyPromoted := physicalDevice.(PhysicalDevice)
	if alreadyPromoted {
		return promoted.InstanceScopedPhysicalDevice1_2()
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
