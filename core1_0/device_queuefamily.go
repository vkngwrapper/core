package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type DeviceQueueCreateOptions struct {
	Flags                  common.DeviceQueueCreateFlags
	QueueFamilyIndex       int
	CreatedQueuePriorities []float32

	common.HaveNext
}

func (o DeviceQueueCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDeviceQueueCreateInfo)
	}

	createInfo := (*C.VkDeviceQueueCreateInfo)(preallocatedPointer)

	if len(o.CreatedQueuePriorities) == 0 {
		return nil, errors.Newf("alloc DeviceCreateOptions: queue family %d had no queue priorities", o.QueueFamilyIndex)
	}

	prioritiesPtr := allocator.Malloc(len(o.CreatedQueuePriorities) * int(unsafe.Sizeof(C.float(0))))
	prioritiesArray := ([]C.float)(unsafe.Slice((*C.float)(prioritiesPtr), len(o.CreatedQueuePriorities)))
	for idx, priority := range o.CreatedQueuePriorities {
		prioritiesArray[idx] = C.float(priority)
	}

	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
	createInfo.flags = C.VkDeviceQueueCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.queueCount = C.uint32_t(len(o.CreatedQueuePriorities))
	createInfo.queueFamilyIndex = C.uint32_t(o.QueueFamilyIndex)
	createInfo.pQueuePriorities = (*C.float)(prioritiesPtr)

	return preallocatedPointer, nil
}

func (o DeviceQueueCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkDeviceQueueCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
