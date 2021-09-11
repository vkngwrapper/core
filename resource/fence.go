package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type VulkanFence struct {
	loader *loader.Loader
	device loader.VkDevice
	handle loader.VkFence
}

func (f *VulkanFence) Handle() loader.VkFence {
	return f.handle
}

func (f *VulkanFence) Destroy() error {
	return f.loader.VkDestroyFence(f.device, f.handle, nil)
}
