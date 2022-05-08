package core1_0

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

type TessellationStateOptions struct {
	PatchControlPoints uint32

	common.HaveNext
}

func (o TessellationStateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatePointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatePointer == nil {
		preallocatePointer = allocator.Malloc(C.sizeof_struct_VkPipelineTessellationStateCreateInfo)
	}
	createInfo := (*C.VkPipelineTessellationStateCreateInfo)(preallocatePointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
	createInfo.patchControlPoints = C.uint32_t(o.PatchControlPoints)

	return preallocatePointer, nil
}

func (o TessellationStateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPipelineTessellationStateCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
