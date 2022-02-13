package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanPipelineCache struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkPipelineCache
}

func (c *VulkanPipelineCache) Handle() driver.VkPipelineCache {
	return c.handle
}

func (c *VulkanPipelineCache) Destroy(callbacks *driver.AllocationCallbacks) {
	c.driver.VkDestroyPipelineCache(c.device, c.handle, callbacks.Handle())
}

func (c *VulkanPipelineCache) CacheData() ([]byte, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	cacheSizePtr := arena.Malloc(int(unsafe.Sizeof(C.size_t(0))))
	cacheSize := (*driver.Size)(cacheSizePtr)

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

func (c *VulkanPipelineCache) MergePipelineCaches(srcCaches []iface.PipelineCache) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	srcCount := len(srcCaches)
	srcPtr := (*driver.VkPipelineCache)(arena.Malloc(srcCount * int(unsafe.Sizeof([1]driver.VkPipelineCache{}))))
	srcSlice := ([]driver.VkPipelineCache)(unsafe.Slice(srcPtr, srcCount))

	for i := 0; i < srcCount; i++ {
		srcSlice[i] = srcCaches[i].Handle()
	}

	return c.driver.VkMergePipelineCaches(c.device, c.handle, driver.Uint32(srcCount), srcPtr)
}
