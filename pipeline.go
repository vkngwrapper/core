package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type vulkanPipeline struct {
	driver Driver
	device VkDevice
	handle VkPipeline
}

func (p *vulkanPipeline) Handle() VkPipeline {
	return p.handle
}

func (p *vulkanPipeline) Destroy() error {
	return p.driver.VkDestroyPipeline(p.device, p.handle, nil)
}
