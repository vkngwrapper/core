package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanDescriptorPool is an implementation of the DescriptorPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorPool struct {
	core1_1.DescriptorPool
}

// PromoteDescriptorPool accepts a DescriptorPool object from any core version. If provided a descriptor pool that supports
// at least core 1.2, it will return a core1_2.DescriptorPool. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanDescriptorPool, even if it is provided a VulkanDescriptorPool from a higher
// core version. Two Vulkan 1.2 compatible DescriptorPool objects with the same DescriptorPool.Handle will
// return the same interface value when passed to this method.
func PromoteDescriptorPool(descriptorPool core1_0.DescriptorPool) DescriptorPool {
	if descriptorPool == nil {
		return nil
	}
	if !descriptorPool.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promoted, alreadyPromoted := descriptorPool.(DescriptorPool)
	if alreadyPromoted {
		return promoted
	}

	promotedDescriptorPool := core1_1.PromoteDescriptorPool(descriptorPool)

	return descriptorPool.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(descriptorPool.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanDescriptorPool{promotedDescriptorPool}
		}).(DescriptorPool)
}
