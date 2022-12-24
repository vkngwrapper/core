package core1_1

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanDescriptorSetLayout is an implementation of the DescriptorSetLayout interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorSetLayout struct {
	core1_0.DescriptorSetLayout
}

// PromoteDescriptorSetLayout accepts a DescriptorSetLayout object from any core version. If provided a descriptor set layout that supports
// at least core 1.1, it will return a core1_1.DescriptorSetLayout. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanDescriptorSetLayout, even if it is provided a VulkanDescriptorSetLayout from a higher
// core version. Two Vulkan 1.1 compatible DescriptorSetLayout objects with the same DescriptorSetLayout.Handle will
// return the same interface value when passed to this method.
func PromoteDescriptorSetLayout(layout core1_0.DescriptorSetLayout) DescriptorSetLayout {
	if layout == nil {
		return nil
	}

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
