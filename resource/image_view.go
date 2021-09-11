package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type VulkanImageView struct {
	loader *loader.Loader
	handle loader.VkImageView
	device loader.VkDevice
}

func (v *VulkanImageView) Handle() loader.VkImageView {
	return v.handle
}

func (v *VulkanImageView) Destroy() error {
	return v.loader.VkDestroyImageView(v.device, v.handle, nil)
}
