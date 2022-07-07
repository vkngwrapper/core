package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanDescriptorSet struct {
	core1_0.DescriptorSet
}

func PromoteDescriptorSet(set core1_0.DescriptorSet) DescriptorSet {
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
