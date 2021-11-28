package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type vulkanFence struct {
	driver Driver
	device VkDevice
	handle VkFence
}

func (f *vulkanFence) Handle() VkFence {
	return f.handle
}

func (f *vulkanFence) Destroy() {
	f.driver.VkDestroyFence(f.device, f.handle, nil)
}
