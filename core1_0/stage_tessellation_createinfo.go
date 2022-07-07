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

type PipelineTessellationStateCreateFlags uint32

var pipelineTessellationStateCreateFlagsMapping = common.NewFlagStringMapping[PipelineTessellationStateCreateFlags]()

func (f PipelineTessellationStateCreateFlags) Register(str string) {
	pipelineTessellationStateCreateFlagsMapping.Register(f, str)
}

func (f PipelineTessellationStateCreateFlags) String() string {
	return pipelineTessellationStateCreateFlagsMapping.FlagsToString(f)
}

////

type PipelineTessellationStateCreateInfo struct {
	Flags              PipelineTessellationStateCreateFlags
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
