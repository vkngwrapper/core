package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanDescriptorSetLayout is an implementation of the DescriptorSetLayout interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorSetLayout struct {
	core1_1.DescriptorSetLayout
}

// PromoteDescriptorSetLayout accepts a DescriptorSetLayout object from any core version. If provided a descriptor set layout that supports
// at least core 1.2, it will return a core1_2.DescriptorSetLayout. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanDescriptorSetLayout, even if it is provided a VulkanDescriptorSetLayout from a higher
// core version. Two Vulkan 1.2 compatible DescriptorSetLayout objects with the same DescriptorSetLayout.Handle will
// return the same interface value when passed to this method.
func PromoteDescriptorSetLayout(layout core1_0.DescriptorSetLayout) DescriptorSetLayout {
	if layout == nil {
		return nil
	}
	if !layout.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promoted, alreadyPromoted := layout.(DescriptorSetLayout)
	if alreadyPromoted {
		return promoted
	}

	promotedDescriptorSetLayout := core1_1.PromoteDescriptorSetLayout(layout)
	return layout.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(layout.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanDescriptorSetLayout{promotedDescriptorSetLayout}
		}).(DescriptorSetLayout)
}
