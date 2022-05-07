package core1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSamplerYcbcrConversion struct {
	Driver      driver.Driver
	Device      driver.VkDevice
	YcbcrHandle driver.VkSamplerYcbcrConversion

	MaximumAPIVersion common.APIVersion
}

func (y *VulkanSamplerYcbcrConversion) Handle() driver.VkSamplerYcbcrConversion {
	return y.YcbcrHandle
}

func (y *VulkanSamplerYcbcrConversion) Destroy(allocator *driver.AllocationCallbacks) {
	y.Driver.VkDestroySamplerYcbcrConversion(y.Device, y.YcbcrHandle, allocator.Handle())
}
