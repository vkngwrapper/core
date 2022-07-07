package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanDescriptorSetLayout struct {
	core1_1.DescriptorSetLayout
}

func PromoteDescriptorSetLayout(layout core1_0.DescriptorSetLayout) DescriptorSetLayout {
	if !layout.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedDescriptorSetLayout := core1_1.PromoteDescriptorSetLayout(layout)
	return layout.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(layout.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanDescriptorSetLayout{promotedDescriptorSetLayout}
		}).(DescriptorSetLayout)
}
