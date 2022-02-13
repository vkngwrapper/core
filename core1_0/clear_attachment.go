package core1_0

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ClearAttachment struct {
	AspectMask      common.ImageAspectFlags
	ColorAttachment int
	ClearValue      core.ClearValue
}

func (c ClearAttachment) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkClearAttachment)
	}

	clearAttachment := (*C.VkClearAttachment)(preallocatedPointer)
	clearAttachment.aspectMask = C.VkImageAspectFlags(c.AspectMask)
	clearAttachment.colorAttachment = C.uint32_t(c.ColorAttachment)
	c.ClearValue.PopulateValueUnion(unsafe.Pointer(&clearAttachment.clearValue))

	return preallocatedPointer, nil
}
