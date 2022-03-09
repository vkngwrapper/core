package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanQueryPool struct {
	Driver          driver.Driver
	QueryPoolHandle driver.VkQueryPool
	Device          driver.VkDevice

	MaximumAPIVersion common.APIVersion
}

func (p *VulkanQueryPool) Handle() driver.VkQueryPool {
	return p.QueryPoolHandle
}

func (p *VulkanQueryPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.Driver.VkDestroyQueryPool(p.Device, p.QueryPoolHandle, callbacks.Handle())
}

func (p *VulkanQueryPool) PopulateResults(firstQuery, queryCount int, resultSize, resultStride int, flags common.QueryResultFlags) ([]byte, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outBuffer := make([]byte, resultSize)

	inPointer := arena.Malloc(resultSize)

	res, err := p.Driver.VkGetQueryPoolResults(p.Device, p.QueryPoolHandle, driver.Uint32(firstQuery), driver.Uint32(queryCount), driver.Size(resultSize), inPointer, driver.VkDeviceSize(resultStride), driver.VkQueryResultFlags(flags))
	if err != nil {
		return nil, res, err
	}

	inBuffer := ([]byte)(unsafe.Slice((*byte)(inPointer), resultSize))
	copy(outBuffer, inBuffer)

	return outBuffer, res, nil
}
