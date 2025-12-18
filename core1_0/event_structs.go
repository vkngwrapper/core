package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

// EventCreateInfo specifies parameters of a newly-created Event
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateInfo.html
type EventCreateInfo struct {
	// Flags defines additional creation parameters
	Flags EventCreateFlags

	common.NextOptions
}

func (o EventCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkEventCreateInfo)
	}
	createInfo := (*C.VkEventCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_EVENT_CREATE_INFO
	createInfo.flags = C.VkEventCreateFlags(o.Flags)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
