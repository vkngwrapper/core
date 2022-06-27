package core1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorSetLayout struct {
	core1_0.DescriptorSetLayout
}

func PromoteDescriptorSetLayout(layout core1_0.DescriptorSetLayout) DescriptorSetLayout {
	if !layout.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return layout.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(layout.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanDescriptorSetLayout{layout}
		}).(DescriptorSetLayout)
}
