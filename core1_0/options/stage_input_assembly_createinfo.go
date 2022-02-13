package options

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

type InputAssemblyOptions struct {
	Topology               common.PrimitiveTopology
	EnablePrimitiveRestart bool

	core.HaveNext
}

func (o InputAssemblyOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineInputAssemblyStateCreateInfo)
	}
	createInfo := (*C.VkPipelineInputAssemblyStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
	createInfo.topology = C.VkPrimitiveTopology(o.Topology)
	createInfo.primitiveRestartEnable = C.VK_FALSE

	if o.EnablePrimitiveRestart {
		createInfo.primitiveRestartEnable = C.VK_TRUE
	}

	return preallocatedPointer, nil
}
