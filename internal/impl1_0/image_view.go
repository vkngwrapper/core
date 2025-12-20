package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanImageView is an implementation of the ImageView interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanImageView struct {
	DeviceDriver    driver.Driver
	ImageViewHandle driver.VkImageView
	Device          driver.VkDevice

	MaximumAPIVersion common.APIVersion
}

func (v *VulkanImageView) Handle() driver.VkImageView {
	return v.ImageViewHandle
}

func (v *VulkanImageView) Driver() driver.Driver {
	return v.DeviceDriver
}

func (v *VulkanImageView) DeviceHandle() driver.VkDevice {
	return v.Device
}

func (v *VulkanImageView) APIVersion() common.APIVersion {
	return v.MaximumAPIVersion
}

func (v *VulkanImageView) Destroy(callbacks *driver.AllocationCallbacks) {
	v.DeviceDriver.VkDestroyImageView(v.Device, v.ImageViewHandle, callbacks.Handle())
}
