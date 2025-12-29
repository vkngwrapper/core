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
)

func (v *DeviceVulkanDriver) QueueWaitIdle(queue core1_0.Queue) (common.VkResult, error) {
	if !queue.Initialized() {
		return core1_0.VKErrorUnknown, fmt.Errorf("queue is uninitialized")
	}
	return v.LoaderObj.VkQueueWaitIdle(queue.Handle())
}

func (v *DeviceVulkanDriver) QueueBindSparse(queue core1_0.Queue, fence *core1_0.Fence, bindInfos ...core1_0.BindSparseInfo) (common.VkResult, error) {
	if !queue.Initialized() {
		return core1_0.VKErrorUnknown, fmt.Errorf("queue is uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	var fenceHandle loader.VkFence
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	bindInfoCount := len(bindInfos)
	bindInfoPtr, err := common.AllocOptionSlice[C.VkBindSparseInfo, core1_0.BindSparseInfo](arena, bindInfos)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.LoaderObj.VkQueueBindSparse(queue.Handle(), loader.Uint32(bindInfoCount), (*loader.VkBindSparseInfo)(unsafe.Pointer(bindInfoPtr)), fenceHandle)
}
