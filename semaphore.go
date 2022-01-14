package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type vulkanSemaphore struct {
	driver Driver
	device VkDevice
	handle VkSemaphore
}

func (s *vulkanSemaphore) Handle() VkSemaphore {
	return s.handle
}

func (s *vulkanSemaphore) Destroy(callbacks *AllocationCallbacks) {
	s.driver.VkDestroySemaphore(s.device, s.handle, callbacks.Handle())
}
