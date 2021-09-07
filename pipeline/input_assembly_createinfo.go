package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type InputAssemblyOptions struct {
	Topology               core.PrimitiveTopology
	EnablePrimitiveRestart bool

	Next core.Options
}

func (o *InputAssemblyOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineInputAssemblyStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineInputAssemblyStateCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.topology = C.VkPrimitiveTopology(o.Topology)
	createInfo.primitiveRestartEnable = C.VK_FALSE

	if o.EnablePrimitiveRestart {
		createInfo.primitiveRestartEnable = C.VK_TRUE
	}

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
