package core1_0

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
	Driver          driver.Driver
	ImageViewHandle driver.VkImageView
	Device          driver.VkDevice

	MaximumAPIVersion common.APIVersion
}

func (v *VulkanImageView) Handle() driver.VkImageView {
	return v.ImageViewHandle
}

func (v *VulkanImageView) Destroy(callbacks *driver.AllocationCallbacks) {
	v.Driver.VkDestroyImageView(v.Device, v.ImageViewHandle, callbacks.Handle())
}
