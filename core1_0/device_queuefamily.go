package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// DeviceQueueCreateInfo specifies parameters of a newly-created Device Queue
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html
type DeviceQueueCreateInfo struct {
	// Flags indicates behavior of the Queue objects
	Flags DeviceQueueCreateFlags
	// QueueFamilyIndex indicates the index of the Queue family in which to create the Queue objects
	// on this Device
	QueueFamilyIndex int
	// QueuePriorities is a slice of normalized floating point values, specifying priorities of work
	// that will be submitted to each created queue
	QueuePriorities []float32

	common.NextOptions
}

func (o DeviceQueueCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDeviceQueueCreateInfo)
	}

	createInfo := (*C.VkDeviceQueueCreateInfo)(preallocatedPointer)

	if len(o.QueuePriorities) == 0 {
		return nil, errors.Newf("alloc DeviceCreateInfo: queue family %d had no queue priorities", o.QueueFamilyIndex)
	}

	prioritiesPtr := allocator.Malloc(len(o.QueuePriorities) * int(unsafe.Sizeof(C.float(0))))
	prioritiesArray := ([]C.float)(unsafe.Slice((*C.float)(prioritiesPtr), len(o.QueuePriorities)))
	for idx, priority := range o.QueuePriorities {
		prioritiesArray[idx] = C.float(priority)
	}

	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
	createInfo.flags = C.VkDeviceQueueCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.queueCount = C.uint32_t(len(o.QueuePriorities))
	createInfo.queueFamilyIndex = C.uint32_t(o.QueueFamilyIndex)
	createInfo.pQueuePriorities = (*C.float)(prioritiesPtr)

	return preallocatedPointer, nil
}
