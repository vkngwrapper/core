package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanQueryPool is an implementation of the QueryPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanQueryPool struct {
	core1_1.QueryPool

	DeviceDriver    driver.Driver
	Device          driver.VkDevice
	QueryPoolHandle driver.VkQueryPool
}

// PromoteQueryPool accepts a QueryPool object from any core version. If provided a query pool that supports
// at least core 1.2, it will return a core1_2.QueryPool. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanQueryPool, even if it is provided a VulkanQueryPool from a higher
// core version. Two Vulkan 1.2 compatible QueryPool objects with the same QueryPool.Handle will
// return the same interface value when passed to this method.
func PromoteQueryPool(queryPool core1_0.QueryPool) QueryPool {
	if !queryPool.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedQueryPool := core1_1.PromoteQueryPool(queryPool)
	return queryPool.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(queryPool.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanQueryPool{
				QueryPool: promotedQueryPool,

				DeviceDriver:    queryPool.Driver(),
				Device:          queryPool.DeviceHandle(),
				QueryPoolHandle: queryPool.Handle(),
			}
		}).(QueryPool)
}

func (q *VulkanQueryPool) Reset(firstQuery, queryCount int) {
	q.DeviceDriver.VkResetQueryPool(q.Device, q.QueryPoolHandle, driver.Uint32(firstQuery), driver.Uint32(queryCount))
}
