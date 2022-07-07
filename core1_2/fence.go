package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanFence struct {
	core1_1.Fence
}

func PromoteFence(fence core1_0.Fence) Fence {
	if !fence.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedFence := core1_1.PromoteFence(fence)
	return fence.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(fence.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanFence{promotedFence}
		}).(Fence)
}
