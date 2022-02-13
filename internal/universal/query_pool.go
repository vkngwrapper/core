package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanQueryPool struct {
	driver driver.Driver
	handle driver.VkQueryPool
	device driver.VkDevice
}

func (p *VulkanQueryPool) Handle() driver.VkQueryPool {
	return p.handle
}

func (p *VulkanQueryPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.driver.VkDestroyQueryPool(p.device, p.handle, callbacks.Handle())
}

func (p *VulkanQueryPool) PopulateResults(firstQuery, queryCount int, resultSize, resultStride int, flags core1_0.QueryResultFlags) ([]byte, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outBuffer := make([]byte, resultSize)

	inPointer := arena.Malloc(resultSize)

	res, err := p.driver.VkGetQueryPoolResults(p.device, p.handle, driver.Uint32(firstQuery), driver.Uint32(queryCount), driver.Size(resultSize), inPointer, driver.VkDeviceSize(resultStride), driver.VkQueryResultFlags(flags))
	if err != nil {
		return nil, res, err
	}

	inBuffer := ([]byte)(unsafe.Slice((*byte)(inPointer), resultSize))
	copy(outBuffer, inBuffer)

	return outBuffer, res, nil
}
