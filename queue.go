package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
)

type QueueHandle C.VkQueue
type Queue struct {
	handle QueueHandle
}

func (q *Queue) Handle() QueueHandle {
	return q.handle
}

func (q *Queue) WaitForIdle() (VKng.Result, error) {
	res := VKng.Result(C.vkQueueWaitIdle(q.handle))
	return res, res.ToError()
}
