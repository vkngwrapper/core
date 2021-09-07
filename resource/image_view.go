package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type ImageViewHandle C.VkImageView
type ImageView struct {
	handle C.VkImageView
	device C.VkDevice
}

func (v *ImageView) Handle() ImageViewHandle {
	return ImageViewHandle(v.handle)
}

func (v *ImageView) Destroy() {
	C.vkDestroyImageView(v.device, v.handle, nil)
}
