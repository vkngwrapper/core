package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanQueryPool struct {
	Driver          driver.Driver
	QueryPoolHandle driver.VkQueryPool
	Device          driver.VkDevice

	MaximumAPIVersion common.APIVersion

	QueryPool1_1 core1_1.QueryPool
}

func (p *VulkanQueryPool) Handle() driver.VkQueryPool {
	return p.QueryPoolHandle
}

func (p *VulkanQueryPool) Core1_1() core1_1.QueryPool {
	return p.QueryPool1_1
}

func (p *VulkanQueryPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.Driver.VkDestroyQueryPool(p.Device, p.QueryPoolHandle, callbacks.Handle())
	p.Driver.ObjectStore().Delete(driver.VulkanHandle(p.QueryPoolHandle), p)
}

func (p *VulkanQueryPool) PopulateResults(firstQuery, queryCount int, results []byte, resultStride int, flags common.QueryResultFlags) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	resultSize := len(results)

	inPointer := arena.Malloc(resultSize)

	res, err := p.Driver.VkGetQueryPoolResults(p.Device, p.QueryPoolHandle, driver.Uint32(firstQuery), driver.Uint32(queryCount), driver.Size(resultSize), inPointer, driver.VkDeviceSize(resultStride), driver.VkQueryResultFlags(flags))
	if err != nil {
		return res, err
	}

	inBuffer := ([]byte)(unsafe.Slice((*byte)(inPointer), resultSize))
	copy(results, inBuffer)

	return res, nil
}
