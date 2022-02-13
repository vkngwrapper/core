package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanImageView struct {
	driver driver.Driver
	handle driver.VkImageView
	device driver.VkDevice
}

func (v *VulkanImageView) Handle() driver.VkImageView {
	return v.handle
}

func (v *VulkanImageView) Destroy(callbacks *driver.AllocationCallbacks) {
	v.driver.VkDestroyImageView(v.device, v.handle, callbacks.Handle())
}
