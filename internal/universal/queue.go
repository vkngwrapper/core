package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0/options"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanQueue struct {
	driver driver.Driver
	handle driver.VkQueue
}

func (q *VulkanQueue) Handle() driver.VkQueue {
	return q.handle
}

func (q *VulkanQueue) Driver() driver.Driver {
	return q.driver
}

func (q *VulkanQueue) WaitForIdle() (common.VkResult, error) {
	return q.driver.VkQueueWaitIdle(q.handle)
}

func (q *VulkanQueue) BindSparse(fence iface.Fence, bindInfos []options.BindSparseOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	var fenceHandle driver.VkFence
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	bindInfoCount := len(bindInfos)
	bindInfoPtr, err := core.AllocOptionSlice[C.VkBindSparseInfo, options.BindSparseOptions](arena, bindInfos)
	if err != nil {
		return common.VKErrorUnknown, err
	}

	return q.driver.VkQueueBindSparse(q.handle, driver.Uint32(bindInfoCount), (*driver.VkBindSparseInfo)(unsafe.Pointer(bindInfoPtr)), fenceHandle)
}
