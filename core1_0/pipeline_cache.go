package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
	"unsafe"
)

// VulkanPipelineCache is an implementation of the PipelineCache interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPipelineCache struct {
	deviceDriver        driver.Driver
	device              driver.VkDevice
	pipelineCacheHandle driver.VkPipelineCache

	maximumAPIVersion common.APIVersion
}

func (c *VulkanPipelineCache) Handle() driver.VkPipelineCache {
	return c.pipelineCacheHandle
}

func (c *VulkanPipelineCache) DeviceHandle() driver.VkDevice {
	return c.device
}

func (c *VulkanPipelineCache) Driver() driver.Driver {
	return c.deviceDriver
}

func (c *VulkanPipelineCache) APIVersion() common.APIVersion {
	return c.maximumAPIVersion
}

func (c *VulkanPipelineCache) Destroy(callbacks *driver.AllocationCallbacks) {
	c.deviceDriver.VkDestroyPipelineCache(c.device, c.pipelineCacheHandle, callbacks.Handle())
	c.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(c.pipelineCacheHandle))
}

func (c *VulkanPipelineCache) CacheData() ([]byte, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	cacheSizePtr := arena.Malloc(int(unsafe.Sizeof(C.size_t(0))))
	cacheSize := (*driver.Size)(cacheSizePtr)

	res, err := c.deviceDriver.VkGetPipelineCacheData(c.device, c.pipelineCacheHandle, cacheSize, nil)
	if err != nil {
		return nil, res, err
	}

	cacheDataPtr := arena.Malloc(int(*cacheSize))

	res, err = c.deviceDriver.VkGetPipelineCacheData(c.device, c.pipelineCacheHandle, cacheSize, cacheDataPtr)
	if err != nil {
		return nil, res, err
	}

	outData := make([]byte, *cacheSize)
	inData := ([]byte)(unsafe.Slice((*byte)(cacheDataPtr), int(*cacheSize)))
	copy(outData, inData)

	return outData, res, nil
}

func (c *VulkanPipelineCache) MergePipelineCaches(srcCaches []PipelineCache) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	srcCount := len(srcCaches)
	srcPtr := (*driver.VkPipelineCache)(arena.Malloc(srcCount * int(unsafe.Sizeof([1]driver.VkPipelineCache{}))))
	srcSlice := ([]driver.VkPipelineCache)(unsafe.Slice(srcPtr, srcCount))

	for i := 0; i < srcCount; i++ {
		srcSlice[i] = srcCaches[i].Handle()
	}

	return c.deviceDriver.VkMergePipelineCaches(c.device, c.pipelineCacheHandle, driver.Uint32(srcCount), srcPtr)
}
