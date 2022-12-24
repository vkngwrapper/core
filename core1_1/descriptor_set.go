package core1_1

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanDescriptorSet is an implementation of the DescriptorSet interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorSet struct {
	core1_0.DescriptorSet
}

// PromoteDescriptorSet accepts a DescriptorSet object from any core version. If provided a descriptor set that supports
// at least core 1.1, it will return a core1_1.DescriptorSet. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanDescriptorSet, even if it is provided a VulkanDescriptorSet from a higher
// core version. Two Vulkan 1.1 compatible DescriptorSet objects with the same DescriptorSet.Handle will
// return the same interface value when passed to this method.
func PromoteDescriptorSet(set core1_0.DescriptorSet) DescriptorSet {
	if set == nil {
		return nil
	}

	if !set.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return set.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(set.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanDescriptorSet{set}
		}).(DescriptorSet)
}

// PromoteDescriptorSetSlice accepts a slice of DescriptorSet objects from any core version.
// If provided a descriptor set that supports at least core 1.1, it will return a core1_1.DescriptorSet.
// Otherwise, it will left out of the returned slice. This method will always return a
// core1_1.VulkanDescriptorSet, even if it is provided a VulkanDescriptorSet from a higher core version. Two
// Vulkan 1.1 compatible DescriptorSet objects with the same DescriptorSet.Handle will return the same interface
// value when passed to this method.
func PromoteDescriptorSetSlice(sets []core1_0.DescriptorSet) []DescriptorSet {
	outSets := make([]DescriptorSet, len(sets))
	for i := 0; i < len(sets); i++ {
		outSets[i] = PromoteDescriptorSet(sets[i])

		if outSets[i] == nil {
			return nil
		}
	}

	return outSets
}
