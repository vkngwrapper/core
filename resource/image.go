package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type Image struct {
	handle loader.VkImage
	device loader.VkDevice
}

func CreateFromHandles(handle loader.VkImage, device loader.VkDevice) *Image {
	return &Image{handle: handle, device: device}
}

func (i *Image) Handle() loader.VkImage {
	return i.handle
}
