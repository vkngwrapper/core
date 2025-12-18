package core1_1

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanDeviceMemory is an implementation of the DeviceMemory interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDeviceMemory struct {
	core1_0.DeviceMemory
}

// PromoteDeviceMemory accepts a DeviceMemory object from any core version. If provided a device memory that supports
// at least core 1.1, it will return a core1_1.DeviceMemory. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanDeviceMemory, even if it is provided a VulkanDeviceMemory from a higher
// core version. Two Vulkan 1.1 compatible DeviceMemory objects with the same DeviceMemory.Handle will
// return the same interface value when passed to this method.
func PromoteDeviceMemory(deviceMemory core1_0.DeviceMemory) DeviceMemory {
	if deviceMemory == nil {
		return nil
	}
	if !deviceMemory.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	promoted, alreadyPromoted := deviceMemory.(DeviceMemory)
	if alreadyPromoted {
		return promoted
	}

	return deviceMemory.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(deviceMemory.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanDeviceMemory{deviceMemory}
		}).(DeviceMemory)
}
