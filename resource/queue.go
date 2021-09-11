package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/loader"
)

type VulkanQueue struct {
	loader *loader.Loader
	handle loader.VkQueue
}

func (q *VulkanQueue) Handle() loader.VkQueue {
	return q.handle
}

func (q *VulkanQueue) Loader() *loader.Loader {
	return q.loader
}

func (q *VulkanQueue) WaitForIdle() (loader.VkResult, error) {
	return q.loader.VkQueueWaitIdle(q.handle)
}
