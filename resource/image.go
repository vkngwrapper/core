package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type VulkanImage struct {
	handle loader.VkImage
	device loader.VkDevice
}

func CreateFromHandles(handle loader.VkImage, device loader.VkDevice) Image {
	return &VulkanImage{handle: handle, device: device}
}

func (i *VulkanImage) Handle() loader.VkImage {
	return i.handle
}
