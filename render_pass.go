package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type vulkanRenderPass struct {
	driver Driver
	device VkDevice
	handle VkRenderPass
}

func (p *vulkanRenderPass) Handle() VkRenderPass {
	return p.handle
}

func (p *vulkanRenderPass) Destroy() error {
	return p.driver.VkDestroyRenderPass(p.device, p.handle, nil)
}
