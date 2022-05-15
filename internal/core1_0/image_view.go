package internal1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

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

func (v *VulkanImageView) APIVersion() common.APIVersion {
	return v.MaximumAPIVersion
}

func (v *VulkanImageView) Destroy(callbacks *driver.AllocationCallbacks) {
	v.DeviceDriver.VkDestroyImageView(v.Device, v.ImageViewHandle, callbacks.Handle())
	v.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(v.ImageViewHandle))
}
