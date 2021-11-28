package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"strings"
	"unsafe"
)

type vulkanPipelineCache struct {
	driver Driver
	device VkDevice
	handle VkPipelineCache
}

func (c *vulkanPipelineCache) Handle() VkPipelineCache {
	return c.handle
}

func (c *vulkanPipelineCache) Destroy() {
	c.driver.VkDestroyPipelineCache(c.device, c.handle, nil)
}

type PipelineCacheFlags int32

const (
	PipelineCacheExternallySynchronized PipelineCacheFlags = C.VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT_EXT
)

var pipelineCacheFlagsToString = map[PipelineCacheFlags]string{
	PipelineCacheExternallySynchronized: "Externally Synchronized",
}

func (f PipelineCacheFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		checkBit := PipelineCacheFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := pipelineCacheFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type PipelineCacheOptions struct {
	Flags       PipelineCacheFlags
	InitialData []byte

	common.HaveNext
}

func (o *PipelineCacheOptions) AllocForC(alloc *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineCacheCreateInfo)(alloc.Malloc(C.sizeof_struct_VkPipelineCacheCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkPipelineCacheCreateFlags(o.Flags)

	initialSize := len(o.InitialData)
	createInfo.initialDataSize = C.size_t(initialSize)
	createInfo.pInitialData = nil

	if initialSize > 0 {
		createInfo.pInitialData = alloc.CBytes(o.InitialData)
	}

	return unsafe.Pointer(createInfo), nil
}
