package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSemaphore struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkSemaphore
}

func (s *VulkanSemaphore) Handle() driver.VkSemaphore {
	return s.handle
}

func (s *VulkanSemaphore) Destroy(callbacks *driver.AllocationCallbacks) {
	s.driver.VkDestroySemaphore(s.device, s.handle, callbacks.Handle())
}
