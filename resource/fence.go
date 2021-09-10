package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/loader"

type Fence struct {
	loader *loader.Loader
	device loader.VkDevice
	handle loader.VkFence
}

func (f *Fence) Handle() loader.VkFence {
	return f.handle
}

func (f *Fence) Destroy() error {
	return f.loader.VkDestroyFence(f.device, f.handle, nil)
}
