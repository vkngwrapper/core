package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	driver3 "github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type vulkanPipelineCache struct {
	driver driver3.Driver
	device driver3.VkDevice
	handle driver3.VkPipelineCache
}

func (c *vulkanPipelineCache) Handle() driver3.VkPipelineCache {
	return c.handle
}

func (c *vulkanPipelineCache) Destroy(callbacks *AllocationCallbacks) {
	c.driver.VkDestroyPipelineCache(c.device, c.handle, callbacks.Handle())
}

func (c *vulkanPipelineCache) CacheData() ([]byte, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	cacheSizePtr := arena.Malloc(int(unsafe.Sizeof(C.size_t(0))))
	cacheSize := (*driver3.Size)(cacheSizePtr)

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

func (c *vulkanPipelineCache) MergePipelineCaches(srcCaches []PipelineCache) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	srcCount := len(srcCaches)
	srcPtr := (*driver3.VkPipelineCache)(arena.Malloc(srcCount * int(unsafe.Sizeof([1]driver3.VkPipelineCache{}))))
	srcSlice := ([]driver3.VkPipelineCache)(unsafe.Slice(srcPtr, srcCount))

	for i := 0; i < srcCount; i++ {
		srcSlice[i] = srcCaches[i].Handle()
	}

	return c.driver.VkMergePipelineCaches(c.device, c.handle, driver3.Uint32(srcCount), srcPtr)
}

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
