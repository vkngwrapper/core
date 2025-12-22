package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) QueueWaitIdle(queue types.Queue) (common.VkResult, error) {
	if queue.Handle() == 0 {
		return core1_0.VKErrorUnknown, fmt.Errorf("queue is uninitialized")
	}
	return v.Driver.VkQueueWaitIdle(queue.Handle())
}

func (v *Vulkan) QueueBindSparse(queue types.Queue, fence types.Fence, bindInfos []core1_0.BindSparseInfo) (common.VkResult, error) {
	if queue.Handle() == 0 {
		return core1_0.VKErrorUnknown, fmt.Errorf("queue is uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	var fenceHandle driver.VkFence
	if fence.Handle() != 0 {
		fenceHandle = fence.Handle()
	}

	bindInfoCount := len(bindInfos)
	bindInfoPtr, err := common.AllocOptionSlice[C.VkBindSparseInfo, core1_0.BindSparseInfo](arena, bindInfos)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.Driver.VkQueueBindSparse(queue.Handle(), driver.Uint32(bindInfoCount), (*driver.VkBindSparseInfo)(unsafe.Pointer(bindInfoPtr)), fenceHandle)
}
