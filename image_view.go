package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	driver3 "github.com/CannibalVox/VKng/core/driver"
)

type vulkanImageView struct {
	driver driver3.Driver
	handle driver3.VkImageView
	device driver3.VkDevice
}

func (v *vulkanImageView) Handle() driver3.VkImageView {
	return v.handle
}

func (v *vulkanImageView) Destroy(callbacks *AllocationCallbacks) {
	v.driver.VkDestroyImageView(v.device, v.handle, callbacks.Handle())
}
