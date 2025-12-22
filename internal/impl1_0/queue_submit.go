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

func (v *Vulkan) Submit(queue types.Queue, fence types.Fence, o []core1_0.SubmitInfo) (common.VkResult, error) {
	if queue.Handle() == 0 {
		return core1_0.VKErrorUnknown, fmt.Errorf("queue is uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	submitCount := len(o)
	createInfoPtrUnsafe, err := common.AllocOptionSlice[C.VkSubmitInfo, core1_0.SubmitInfo](arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	var fenceHandle driver.VkFence
	if fence.Handle() != 0 {
		fenceHandle = fence.Handle()
	}

	return v.Driver.VkQueueSubmit(queue.Handle(), driver.Uint32(submitCount), (*driver.VkSubmitInfo)(unsafe.Pointer(createInfoPtrUnsafe)), fenceHandle)
}
