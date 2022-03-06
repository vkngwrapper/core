package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type EventFlags int32

const (
	EventDeviceOnlyKHR EventFlags = C.VK_EVENT_CREATE_DEVICE_ONLY_BIT_KHR
)

var eventFlagsToString = map[EventFlags]string{
	EventDeviceOnlyKHR: "Device Only (Khronos Extension)",
}

func (f EventFlags) String() string {
	return common.FlagsToString(f, eventFlagsToString)
}

type EventOptions struct {
	Flags EventFlags

	core.HaveNext
}

func (o *EventOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkEventCreateInfo)
	}
	createInfo := (*C.VkEventCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_EVENT_CREATE_INFO
	createInfo.flags = C.VkEventCreateFlags(o.Flags)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
