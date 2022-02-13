package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanBufferView struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkBufferView
}

func (v *VulkanBufferView) Handle() driver.VkBufferView {
	return v.handle
}

func (v *VulkanBufferView) Destroy(callbacks *driver.AllocationCallbacks) {
	v.driver.VkDestroyBufferView(v.device, v.handle, callbacks.Handle())
}
