package impl1_0

import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanQueryPool is an implementation of the QueryPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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

func (p *VulkanQueryPool) DeviceHandle() driver.VkDevice {
	return p.Device
}

func (p *VulkanQueryPool) APIVersion() common.APIVersion {
	return p.MaximumAPIVersion
}

func (p *VulkanQueryPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.DeviceDriver.VkDestroyQueryPool(p.Device, p.QueryPoolHandle, callbacks.Handle())
}

func (p *VulkanQueryPool) PopulateResults(firstQuery, queryCount int, results []byte, resultStride int, flags core1_0.QueryResultFlags) (common.VkResult, error) {
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
