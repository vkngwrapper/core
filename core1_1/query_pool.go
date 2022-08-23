package core1_1

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanQueryPool is an implementation of the QueryPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanQueryPool struct {
	core1_0.QueryPool
}

// PromoteQueryPool accepts a QueryPool object from any core version. If provided a query pool that supports
// at least core 1.1, it will return a core1_1.QueryPool. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanQueryPool, even if it is provided a VulkanQueryPool from a higher
// core version. Two Vulkan 1.1 compatible QueryPool objects with the same QueryPool.Handle will
// return the same interface value when passed to this method.
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
