package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanPipeline struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkPipeline
}

func (p *VulkanPipeline) Handle() driver.VkPipeline {
	return p.handle
}

func (p *VulkanPipeline) Destroy(callbacks *driver.AllocationCallbacks) {
	p.driver.VkDestroyPipeline(p.device, p.handle, callbacks.Handle())
}
