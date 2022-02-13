package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanFramebuffer struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkFramebuffer
}

func (b *VulkanFramebuffer) Handle() driver.VkFramebuffer {
	return b.handle
}

func (b *VulkanFramebuffer) Destroy(callbacks *driver.AllocationCallbacks) {
	b.driver.VkDestroyFramebuffer(b.device, b.handle, callbacks.Handle())
}
