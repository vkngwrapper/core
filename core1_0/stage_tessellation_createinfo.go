package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// PipelineTessellationStateCreateFlags is reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateFlags.html
type PipelineTessellationStateCreateFlags uint32

var pipelineTessellationStateCreateFlagsMapping = common.NewFlagStringMapping[PipelineTessellationStateCreateFlags]()

func (f PipelineTessellationStateCreateFlags) Register(str string) {
	pipelineTessellationStateCreateFlagsMapping.Register(f, str)
}

func (f PipelineTessellationStateCreateFlags) String() string {
	return pipelineTessellationStateCreateFlagsMapping.FlagsToString(f)
}

////

// PipelineTessellationStateCreateInfo specifies parameters of a newly-created Pipeline tessellation
// state
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html
type PipelineTessellationStateCreateInfo struct {
	// Flags is reserved for future use
	Flags PipelineTessellationStateCreateFlags
	// PatchControlPoints is the number of control points per patch
	PatchControlPoints uint32

	common.NextOptions
}

func (o PipelineTessellationStateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatePointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatePointer == nil {
		preallocatePointer = allocator.Malloc(C.sizeof_struct_VkPipelineTessellationStateCreateInfo)
	}
	createInfo := (*C.VkPipelineTessellationStateCreateInfo)(preallocatePointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO
	createInfo.flags = C.VkPipelineTessellationStateCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.patchControlPoints = C.uint32_t(o.PatchControlPoints)

	return preallocatePointer, nil
}
