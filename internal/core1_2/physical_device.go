package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanPhysicalDevice struct {
	core1_1.PhysicalDevice

	InstanceScoped1_2 core1_2.InstanceScopedPhysicalDevice
}

func (p *VulkanPhysicalDevice) InstanceScopedPhysicalDevice1_2() core1_2.InstanceScopedPhysicalDevice {
	return p.InstanceScoped1_2
}

func PromotePhysicalDevice(physicalDevice core1_0.PhysicalDevice) core1_2.PhysicalDevice {
	if !physicalDevice.DeviceAPIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	instanceScoped := PromoteInstanceScopePhysicalDevice(physicalDevice)

	return physicalDevice.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(physicalDevice.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanPhysicalDevice{
				PhysicalDevice: core1_1.PromotePhysicalDevice(physicalDevice),

				InstanceScoped1_2: instanceScoped,
			}
		}).(core1_2.PhysicalDevice)
}

type VulkanInstanceScopedPhysicalDevice struct {
	core1_1.PhysicalDevice
}

func PromoteInstanceScopePhysicalDevice(physicalDevice core1_0.PhysicalDevice) core1_2.InstanceScopedPhysicalDevice {
	if !physicalDevice.InstanceAPIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return physicalDevice.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(physicalDevice.Handle()),
		driver.Core1_2InstanceScope,
		func() any {
			return &VulkanInstanceScopedPhysicalDevice{
				PhysicalDevice: core1_1.PromotePhysicalDevice(physicalDevice),
			}
		}).(core1_1.InstanceScopedPhysicalDevice)
}
