package core1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDeviceMemory struct {
	core1_0.DeviceMemory
}

func PromoteDeviceMemory(deviceMemory core1_0.DeviceMemory) DeviceMemory {
	if !deviceMemory.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return deviceMemory.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(deviceMemory.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanDeviceMemory{deviceMemory}
		}).(DeviceMemory)
}
