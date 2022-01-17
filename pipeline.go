package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	driver3 "github.com/CannibalVox/VKng/core/driver"
)

type vulkanPipeline struct {
	driver driver3.Driver
	device driver3.VkDevice
	handle driver3.VkPipeline
}

func (p *vulkanPipeline) Handle() driver3.VkPipeline {
	return p.handle
}

func (p *vulkanPipeline) Destroy(callbacks *AllocationCallbacks) {
	p.driver.VkDestroyPipeline(p.device, p.handle, callbacks.Handle())
}
