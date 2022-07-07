package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanFence struct {
	core1_0.Fence
}

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
