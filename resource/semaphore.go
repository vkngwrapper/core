package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type SemaphoreHandle C.VkSemaphore
type Semaphore struct {
	device C.VkDevice
	handle C.VkSemaphore
}

func (s *Semaphore) Handle() SemaphoreHandle {
	return SemaphoreHandle(s.handle)
}

func (s *Semaphore) Destroy() {
	C.vkDestroySemaphore(s.device, s.handle, nil)
}
