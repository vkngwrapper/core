package core1_1

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanSamplerYcbcrConversion is an implementation of the SamplerYcbcrConversion interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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

func (y *VulkanSamplerYcbcrConversion) DeviceHandle() driver.VkDevice {
	return y.Device
}

func (y *VulkanSamplerYcbcrConversion) APIVersion() common.APIVersion {
	return y.MaximumAPIVersion
}

func (y *VulkanSamplerYcbcrConversion) Destroy(allocator *driver.AllocationCallbacks) {
	y.DeviceDriver.VkDestroySamplerYcbcrConversion(y.Device, y.YcbcrHandle, allocator.Handle())
}
