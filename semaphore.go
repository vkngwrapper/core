package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	driver3 "github.com/CannibalVox/VKng/core/driver"
)

type vulkanSemaphore struct {
	driver driver3.Driver
	device driver3.VkDevice
	handle driver3.VkSemaphore
}

func (s *vulkanSemaphore) Handle() driver3.VkSemaphore {
	return s.handle
}

func (s *vulkanSemaphore) Destroy(callbacks *AllocationCallbacks) {
	s.driver.VkDestroySemaphore(s.device, s.handle, callbacks.Handle())
}
