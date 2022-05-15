package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanQueryPool struct {
	core1_0.QueryPool
}

func PromoteQueryPool(queryPool core1_0.QueryPool) core1_1.QueryPool {
	if !queryPool.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return queryPool.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(queryPool.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanQueryPool{queryPool}
		}).(core1_1.QueryPool)
}
