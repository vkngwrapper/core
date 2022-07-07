package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
	"unsafe"
)

func (q *VulkanQueue) Submit(fence Fence, o []SubmitInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	submitCount := len(o)
	createInfoPtrUnsafe, err := common.AllocOptionSlice[C.VkSubmitInfo, SubmitInfo](arena, o)
	if err != nil {
		return VKErrorUnknown, err
	}

	var fenceHandle driver.VkFence
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	return q.Driver().VkQueueSubmit(q.Handle(), driver.Uint32(submitCount), (*driver.VkSubmitInfo)(unsafe.Pointer(createInfoPtrUnsafe)), fenceHandle)
}
