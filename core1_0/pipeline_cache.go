package core1_0

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

type PipelineCacheFlags int32

const (
	PipelineCacheExternallySynchronized PipelineCacheFlags = C.VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT_EXT
)

var pipelineCacheFlagsToString = map[PipelineCacheFlags]string{
	PipelineCacheExternallySynchronized: "Externally Synchronized",
}

func (f PipelineCacheFlags) String() string {
	return common.FlagsToString(f, pipelineCacheFlagsToString)
}

type PipelineCacheOptions struct {
	Flags       PipelineCacheFlags
	InitialData []byte

	core.HaveNext
}

func (o *PipelineCacheOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
