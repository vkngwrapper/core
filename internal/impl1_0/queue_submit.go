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
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *DeviceVulkanDriver) QueueSubmit(queue types.Queue, fence *types.Fence, o ...core1_0.SubmitInfo) (common.VkResult, error) {
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

	var fenceHandle loader.VkFence
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	return v.LoaderObj.VkQueueSubmit(queue.Handle(), loader.Uint32(submitCount), (*loader.VkSubmitInfo)(unsafe.Pointer(createInfoPtrUnsafe)), fenceHandle)
}
