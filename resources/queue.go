package resources

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/loader"
)

type vulkanQueue struct {
	loader loader.Loader
	handle loader.VkQueue
}

func (q *vulkanQueue) Handle() loader.VkQueue {
	return q.handle
}

func (q *vulkanQueue) Loader() loader.Loader {
	return q.loader
}

func (q *vulkanQueue) WaitForIdle() (loader.VkResult, error) {
	return q.loader.VkQueueWaitIdle(q.handle)
}
