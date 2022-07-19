package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

// VulkanFence is an implementation of the Fence interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanFence struct {
	core1_0.Fence
}

// PromoteFence accepts a Fence object from any core version. If provided a fence that supports
// at least core 1.1, it will return a core1_1.Fence. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanFence, even if it is provided a VulkanFence from a higher
// core version. Two Vulkan 1.1 compatible Fence objects with the same Fence.Handle will
// return the same interface value when passed to this method.
func PromoteFence(fence core1_0.Fence) Fence {
	if !fence.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return fence.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(fence.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanFence{fence}
		}).(Fence)
}
