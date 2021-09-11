package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type vulkanFence struct {
	loader *loader.Loader
	device loader.VkDevice
	handle loader.VkFence
}

func (f *vulkanFence) Handle() loader.VkFence {
	return f.handle
}

func (f *vulkanFence) Destroy() error {
	return f.loader.VkDestroyFence(f.device, f.handle, nil)
}
