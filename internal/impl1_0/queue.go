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

// VulkanQueue is an implementation of the Queue interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanQueue struct {
	DeviceDriver driver.Driver
	Device       driver.VkDevice
	QueueHandle  driver.VkQueue

	MaximumAPIVersion common.APIVersion
}

func (q *VulkanQueue) Handle() driver.VkQueue {
	return q.QueueHandle
}

func (q *VulkanQueue) Driver() driver.Driver {
	return q.DeviceDriver
}

func (q *VulkanQueue) DeviceHandle() driver.VkDevice {
	return q.Device
}

func (q *VulkanQueue) APIVersion() common.APIVersion {
	return q.MaximumAPIVersion
}

func (q *VulkanQueue) WaitIdle() (common.VkResult, error) {
	return q.DeviceDriver.VkQueueWaitIdle(q.QueueHandle)
}

func (q *VulkanQueue) BindSparse(fence core1_0.Fence, bindInfos []core1_0.BindSparseInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	var fenceHandle driver.VkFence
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	bindInfoCount := len(bindInfos)
	bindInfoPtr, err := common.AllocOptionSlice[C.VkBindSparseInfo, core1_0.BindSparseInfo](arena, bindInfos)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return q.DeviceDriver.VkQueueBindSparse(q.QueueHandle, driver.Uint32(bindInfoCount), (*driver.VkBindSparseInfo)(unsafe.Pointer(bindInfoPtr)), fenceHandle)
}
