package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

func (q *VulkanQueue) SubmitToQueue(fence core1_0.Fence, o []core1_0.SubmitOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	submitCount := len(o)
	createInfoPtrUnsafe, err := common.AllocOptionSlice[C.VkSubmitInfo, core1_0.SubmitOptions](arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	var fenceHandle driver.VkFence = nil
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	return q.Driver().VkQueueSubmit(q.Handle(), driver.Uint32(submitCount), (*driver.VkSubmitInfo)(unsafe.Pointer(createInfoPtrUnsafe)), fenceHandle)
}
