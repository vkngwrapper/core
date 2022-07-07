package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanQueryPool struct {
	core1_0.QueryPool
}

func PromoteQueryPool(queryPool core1_0.QueryPool) QueryPool {
	if !queryPool.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return queryPool.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(queryPool.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanQueryPool{queryPool}
		}).(QueryPool)
}
