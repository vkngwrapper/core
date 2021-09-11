package pipeline

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type TessellationOptions struct {
	PatchControlPoints uint32

	Next core.Options
}

func (o *TessellationOptions) AllocForC(allocator *cgoparam.Allocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineTessellationStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineTessellationStateCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.patchControlPoints = C.uint32_t(o.PatchControlPoints)

	var err error
	var next unsafe.Pointer
	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}

	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
