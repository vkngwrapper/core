package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanImageView struct {
	Driver          driver.Driver
	ImageViewHandle driver.VkImageView
	Device          driver.VkDevice

	MaximumAPIVersion common.APIVersion

	ImageView1_1 core1_1.ImageView
}

func (v *VulkanImageView) Handle() driver.VkImageView {
	return v.ImageViewHandle
}

func (v *VulkanImageView) Core1_1() core1_1.ImageView {
	return v.ImageView1_1
}

func (v *VulkanImageView) Destroy(callbacks *driver.AllocationCallbacks) {
	v.Driver.VkDestroyImageView(v.Device, v.ImageViewHandle, callbacks.Handle())
}
