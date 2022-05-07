package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DeviceQueueOptions struct {
	Flags            common.DeviceQueueCreateFlags
	QueueFamilyIndex int
	QueueIndex       int

	common.HaveNext
}

func (o DeviceQueueOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDeviceQueueInfo2)
	}

	info := (*C.VkDeviceQueueInfo2)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2
	info.pNext = next
	info.flags = C.VkDeviceQueueCreateFlags(o.Flags)
	info.queueFamilyIndex = C.uint32_t(o.QueueFamilyIndex)
	info.queueIndex = C.uint32_t(o.QueueIndex)

	return preallocatedPointer, nil
}

func (o DeviceQueueOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDeviceQueueInfo2)(cDataPointer)
	return info.pNext, nil
}
