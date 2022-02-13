package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanPipelineLayout struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkPipelineLayout
}

func (l *VulkanPipelineLayout) Handle() driver.VkPipelineLayout {
	return l.handle
}

func (l *VulkanPipelineLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	l.driver.VkDestroyPipelineLayout(l.device, l.handle, callbacks.Handle())
}
