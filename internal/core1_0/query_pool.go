package internal1_0

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
	DeviceDriver    driver.Driver
	QueryPoolHandle driver.VkQueryPool
	Device          driver.VkDevice

	MaximumAPIVersion common.APIVersion
}

func (p *VulkanQueryPool) Handle() driver.VkQueryPool {
	return p.QueryPoolHandle
}

func (p *VulkanQueryPool) Driver() driver.Driver {
	return p.DeviceDriver
}

func (p *VulkanQueryPool) APIVersion() common.APIVersion {
	return p.MaximumAPIVersion
}

func (p *VulkanQueryPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.DeviceDriver.VkDestroyQueryPool(p.Device, p.QueryPoolHandle, callbacks.Handle())
	p.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(p.QueryPoolHandle))
}

func (p *VulkanQueryPool) PopulateResults(firstQuery, queryCount int, results []byte, resultStride int, flags common.QueryResultFlags) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	resultSize := len(results)

	inPointer := arena.Malloc(resultSize)

	res, err := p.DeviceDriver.VkGetQueryPoolResults(p.Device, p.QueryPoolHandle, driver.Uint32(firstQuery), driver.Uint32(queryCount), driver.Size(resultSize), inPointer, driver.VkDeviceSize(resultStride), driver.VkQueryResultFlags(flags))
	if err != nil {
		return res, err
	}

	inBuffer := ([]byte)(unsafe.Slice((*byte)(inPointer), resultSize))
	copy(results, inBuffer)

	return res, nil
}
