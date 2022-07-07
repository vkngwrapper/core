package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
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
