package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSamplerYcbcrConversion struct {
	DeviceDriver driver.Driver
	Device       driver.VkDevice
	YcbcrHandle  driver.VkSamplerYcbcrConversion

	MaximumAPIVersion common.APIVersion
}

func (y *VulkanSamplerYcbcrConversion) Handle() driver.VkSamplerYcbcrConversion {
	return y.YcbcrHandle
}

func (y *VulkanSamplerYcbcrConversion) Driver() driver.Driver {
	return y.DeviceDriver
}

func (y *VulkanSamplerYcbcrConversion) APIVersion() common.APIVersion {
	return y.MaximumAPIVersion
}

func (y *VulkanSamplerYcbcrConversion) Destroy(allocator *driver.AllocationCallbacks) {
	y.DeviceDriver.VkDestroySamplerYcbcrConversion(y.Device, y.YcbcrHandle, allocator.Handle())
}
