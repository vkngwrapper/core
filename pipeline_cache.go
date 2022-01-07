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

func (c *vulkanPipelineCache) CacheData() ([]byte, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	cacheSizePtr := arena.Malloc(int(unsafe.Sizeof(C.size_t(0))))
	cacheSize := (*Size)(cacheSizePtr)

	res, err := c.driver.VkGetPipelineCacheData(c.device, c.handle, cacheSize, nil)
	if err != nil {
		return nil, res, err
	}

	cacheDataPtr := arena.Malloc(int(*cacheSize))

	res, err = c.driver.VkGetPipelineCacheData(c.device, c.handle, cacheSize, cacheDataPtr)
	if err != nil {
		return nil, res, err
	}

	outData := make([]byte, *cacheSize)
	inData := ([]byte)(unsafe.Slice((*byte)(cacheDataPtr), int(*cacheSize)))
	copy(outData, inData)

	return outData, res, nil
}

func (c *vulkanPipelineCache) MergePipelineCaches(srcCaches []PipelineCache) (VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	srcCount := len(srcCaches)
	srcPtr := (*VkPipelineCache)(arena.Malloc(srcCount * int(unsafe.Sizeof([1]VkPipelineCache{}))))
	srcSlice := ([]VkPipelineCache)(unsafe.Slice(srcPtr, srcCount))

	for i := 0; i < srcCount; i++ {
		srcSlice[i] = srcCaches[i].Handle()
	}

	return c.driver.VkMergePipelineCaches(c.device, c.handle, Uint32(srcCount), srcPtr)
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
