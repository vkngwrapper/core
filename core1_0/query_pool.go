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
	deviceDriver    driver.Driver
	queryPoolHandle driver.VkQueryPool
	device          driver.VkDevice

	maximumAPIVersion common.APIVersion
}

func (p *VulkanQueryPool) Handle() driver.VkQueryPool {
	return p.queryPoolHandle
}

func (p *VulkanQueryPool) Driver() driver.Driver {
	return p.deviceDriver
}

func (p *VulkanQueryPool) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p *VulkanQueryPool) APIVersion() common.APIVersion {
	return p.maximumAPIVersion
}

func (p *VulkanQueryPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.deviceDriver.VkDestroyQueryPool(p.device, p.queryPoolHandle, callbacks.Handle())
	p.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(p.queryPoolHandle))
}

func (p *VulkanQueryPool) PopulateResults(firstQuery, queryCount int, results []byte, resultStride int, flags common.QueryResultFlags) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	resultSize := len(results)

	inPointer := arena.Malloc(resultSize)

	res, err := p.deviceDriver.VkGetQueryPoolResults(p.device, p.queryPoolHandle, driver.Uint32(firstQuery), driver.Uint32(queryCount), driver.Size(resultSize), inPointer, driver.VkDeviceSize(resultStride), driver.VkQueryResultFlags(flags))
	if err != nil {
		return res, err
	}

	inBuffer := ([]byte)(unsafe.Slice((*byte)(inPointer), resultSize))
	copy(results, inBuffer)

	return res, nil
}
