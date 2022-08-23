package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/driver"
	"unsafe"
)

// VulkanQueue is an implementation of the Queue interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanQueue struct {
	deviceDriver driver.Driver
	device       driver.VkDevice
	queueHandle  driver.VkQueue

	maximumAPIVersion common.APIVersion
}

func (q *VulkanQueue) Handle() driver.VkQueue {
	return q.queueHandle
}

func (q *VulkanQueue) Driver() driver.Driver {
	return q.deviceDriver
}

func (q *VulkanQueue) DeviceHandle() driver.VkDevice {
	return q.device
}

func (q *VulkanQueue) APIVersion() common.APIVersion {
	return q.maximumAPIVersion
}

func (q *VulkanQueue) WaitIdle() (common.VkResult, error) {
	return q.deviceDriver.VkQueueWaitIdle(q.queueHandle)
}

func (q *VulkanQueue) BindSparse(fence Fence, bindInfos []BindSparseInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	var fenceHandle driver.VkFence
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	bindInfoCount := len(bindInfos)
	bindInfoPtr, err := common.AllocOptionSlice[C.VkBindSparseInfo, BindSparseInfo](arena, bindInfos)
	if err != nil {
		return VKErrorUnknown, err
	}

	return q.deviceDriver.VkQueueBindSparse(q.queueHandle, driver.Uint32(bindInfoCount), (*driver.VkBindSparseInfo)(unsafe.Pointer(bindInfoPtr)), fenceHandle)
}
