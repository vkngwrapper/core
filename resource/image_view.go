package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type ImageView struct {
	loader *loader.Loader
	handle loader.VkImageView
	device loader.VkDevice
}

func (v *ImageView) Handle() loader.VkImageView {
	return v.handle
}

func (v *ImageView) Destroy() error {
	return v.loader.VkDestroyImageView(v.device, v.handle, nil)
}
