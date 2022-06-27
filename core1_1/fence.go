package core1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
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
