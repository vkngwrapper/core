package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type FenceHandle C.VkFence
type Fence struct {
	device C.VkDevice
	handle C.VkFence
}

func (f *Fence) Handle() FenceHandle {
	return FenceHandle(f.handle)
}

func (f *Fence) Destroy() {
	C.vkDestroyFence(f.device, f.handle, nil)
}
