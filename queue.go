package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type vulkanQueue struct {
	driver Driver
	handle VkQueue
}

func (q *vulkanQueue) Handle() VkQueue {
	return q.handle
}

func (q *vulkanQueue) Driver() Driver {
	return q.driver
}

func (q *vulkanQueue) WaitForIdle() (VkResult, error) {
	return q.driver.VkQueueWaitIdle(q.handle)
}
