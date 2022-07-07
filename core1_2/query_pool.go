package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanQueryPool struct {
	core1_1.QueryPool

	DeviceDriver    driver.Driver
	Device          driver.VkDevice
	QueryPoolHandle driver.VkQueryPool
}

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
