package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanQueryPool struct {
	core1_1.QueryPool

	DeviceDriver    driver.Driver
	Device          driver.VkDevice
	QueryPoolHandle driver.VkQueryPool
}

func PromoteQueryPool(queryPool core1_0.QueryPool) core1_2.QueryPool {
	if !queryPool.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return queryPool.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(queryPool.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanQueryPool{
				QueryPool: core1_1.PromoteQueryPool(queryPool),

				DeviceDriver:    queryPool.Driver(),
				Device:          queryPool.DeviceHandle(),
				QueryPoolHandle: queryPool.Handle(),
			}
		}).(core1_2.QueryPool)
}

func (q *VulkanQueryPool) ResetQueryPool(firstQuery, queryCount int) {
	q.DeviceDriver.VkResetQueryPool(q.Device, q.QueryPoolHandle, driver.Uint32(firstQuery), driver.Uint32(queryCount))
}
