package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
)

type QueueHandle C.VkQueue
type Queue struct {
	handle QueueHandle
}

func (q *Queue) Handle() QueueHandle {
	return q.handle
}

func (q *Queue) WaitForIdle() (core.Result, error) {
	res := core.Result(C.vkQueueWaitIdle(q.handle))
	return res, res.ToError()
}
