package internal1_0

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

type VulkanQueue struct {
	DeviceDriver driver.Driver
	QueueHandle  driver.VkQueue

	MaximumAPIVersion common.APIVersion
}

func (q *VulkanQueue) Handle() driver.VkQueue {
	return q.QueueHandle
}

func (q *VulkanQueue) Driver() driver.Driver {
	return q.DeviceDriver
}

func (q *VulkanQueue) WaitForIdle() (common.VkResult, error) {
	return q.DeviceDriver.VkQueueWaitIdle(q.QueueHandle)
}

func (q *VulkanQueue) BindSparse(fence core1_0.Fence, bindInfos []core1_0.BindSparseOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	var fenceHandle driver.VkFence
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	bindInfoCount := len(bindInfos)
	bindInfoPtr, err := common.AllocOptionSlice[C.VkBindSparseInfo, core1_0.BindSparseOptions](arena, bindInfos)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return q.DeviceDriver.VkQueueBindSparse(q.QueueHandle, driver.Uint32(bindInfoCount), (*driver.VkBindSparseInfo)(unsafe.Pointer(bindInfoPtr)), fenceHandle)
}
