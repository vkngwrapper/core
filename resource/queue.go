package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/loader"
)

type Queue struct {
	loader *loader.Loader
	handle loader.VkQueue
}

func (q *Queue) Handle() loader.VkQueue {
	return q.handle
}

func (q *Queue) Loader() *loader.Loader {
	return q.loader
}

func (q *Queue) WaitForIdle() (loader.VkResult, error) {
	return q.loader.VkQueueWaitIdle(q.handle)
}
