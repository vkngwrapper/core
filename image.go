package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type vulkanImage struct {
	handle VkImage
	device VkDevice
}

func CreateFromHandles(handle VkImage, device VkDevice) Image {
	return &vulkanImage{handle: handle, device: device}
}

func (i *vulkanImage) Handle() VkImage {
	return i.handle
}
