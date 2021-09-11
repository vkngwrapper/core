package resources

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type vulkanImage struct {
	handle loader.VkImage
	device loader.VkDevice
}

func CreateFromHandles(handle loader.VkImage, device loader.VkDevice) Image {
	return &vulkanImage{handle: handle, device: device}
}

func (i *vulkanImage) Handle() loader.VkImage {
	return i.handle
}
