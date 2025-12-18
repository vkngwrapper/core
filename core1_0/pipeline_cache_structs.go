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

// PipelineCacheCreateInfo specifies parameters of a newly-created PipelineCache
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateInfo.html
type PipelineCacheCreateInfo struct {
	// Flags specifies the behavior of the PipelineCache
	Flags PipelineCacheCreateFlags
	// InitialData contains previously-retrieved PipelineCache data
	InitialData []byte

	common.NextOptions
}

func (o PipelineCacheCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineCacheCreateInfo)
	}
	createInfo := (*C.VkPipelineCacheCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkPipelineCacheCreateFlags(o.Flags)

	initialSize := len(o.InitialData)
	createInfo.initialDataSize = C.size_t(initialSize)
	createInfo.pInitialData = nil

	if initialSize > 0 {
		createInfo.pInitialData = allocator.CBytes(o.InitialData)
	}

	return preallocatedPointer, nil
}
