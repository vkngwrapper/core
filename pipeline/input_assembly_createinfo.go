package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type InputAssemblyOptions struct {
	Topology               core.PrimitiveTopology
	EnablePrimitiveRestart bool

	core.HaveNext
}

func (o *InputAssemblyOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineInputAssemblyStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineInputAssemblyStateCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
	createInfo.topology = C.VkPrimitiveTopology(o.Topology)
	createInfo.primitiveRestartEnable = C.VK_FALSE

	if o.EnablePrimitiveRestart {
		createInfo.primitiveRestartEnable = C.VK_TRUE
	}

	return unsafe.Pointer(createInfo), nil
}
