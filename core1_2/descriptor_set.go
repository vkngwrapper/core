package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanDescriptorSet is an implementation of the DescriptorSet interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorSet struct {
	core1_1.DescriptorSet
}

// PromoteDescriptorSet accepts a DescriptorSet object from any core version. If provided a descriptor set that supports
// at least core 1.2, it will return a core1_2.DescriptorSet. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanDescriptorSet, even if it is provided a VulkanDescriptorSet from a higher
// core version. Two Vulkan 1.2 compatible DescriptorSet objects with the same DescriptorSet.Handle will
// return the same interface value when passed to this method.
func PromoteDescriptorSet(set core1_0.DescriptorSet) DescriptorSet {
	if set == nil {
		return nil
	}
	if !set.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promoted, alreadyPromoted := set.(DescriptorSet)
	if alreadyPromoted {
		return promoted
	}

	return set.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(set.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanDescriptorSet{set}
		}).(DescriptorSet)
}

func PromoteDescriptorSetSlice(sets []core1_0.DescriptorSet) []DescriptorSet {
	for i := 0; i < len(sets); i++ {
		if sets[i].APIVersion() < common.Vulkan1_2 {
			return nil
		}
	}

	outSets := make([]DescriptorSet, len(sets))
	for i := 0; i < len(sets); i++ {
		outSets[i] = PromoteDescriptorSet(sets[i])

		if outSets[i] == nil {
			return nil
		}
	}

	return outSets
}
