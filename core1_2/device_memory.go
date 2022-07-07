package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanDeviceMemory struct {
	core1_1.DeviceMemory
}

func PromoteDeviceMemory(deviceMemory core1_0.DeviceMemory) DeviceMemory {
	if !deviceMemory.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedDeviceMemory := core1_1.PromoteDeviceMemory(deviceMemory)

	return deviceMemory.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(deviceMemory.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanDeviceMemory{promotedDeviceMemory}
		}).(core1_1.DeviceMemory)
}
