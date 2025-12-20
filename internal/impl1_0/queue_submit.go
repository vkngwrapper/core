package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

func (q *VulkanQueue) Submit(fence core1_0.Fence, o []core1_0.SubmitInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	submitCount := len(o)
	createInfoPtrUnsafe, err := common.AllocOptionSlice[C.VkSubmitInfo, core1_0.SubmitInfo](arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	var fenceHandle driver.VkFence
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	return q.Driver().VkQueueSubmit(q.Handle(), driver.Uint32(submitCount), (*driver.VkSubmitInfo)(unsafe.Pointer(createInfoPtrUnsafe)), fenceHandle)
}
