package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanFence struct {
	core1_1.Fence
}

func PromoteFence(fence core1_0.Fence) core1_2.Fence {
	if !fence.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return fence.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(fence.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanFence{core1_1.PromoteFence(fence)}
		}).(core1_2.Fence)
}
