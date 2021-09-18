package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type vulkanImageView struct {
	driver Driver
	handle VkImageView
	device VkDevice
}

func (v *vulkanImageView) Handle() VkImageView {
	return v.handle
}

func (v *vulkanImageView) Destroy() error {
	return v.driver.VkDestroyImageView(v.device, v.handle, nil)
}
