package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type ImageHandle C.VkImage
type Image struct {
	handle C.VkImage
	device C.VkDevice
}

func CreateFromHandles(handle ImageHandle, device DeviceHandle) *Image {
	return &Image{handle: C.VkImage(handle), device: C.VkDevice(device)}
}

func (i *Image) Handle() ImageHandle {
	return ImageHandle(i.handle)
}
