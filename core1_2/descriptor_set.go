package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorSet struct {
	core1_1.DescriptorSet
}

func PromoteDescriptorSet(set core1_0.DescriptorSet) DescriptorSet {
	if !set.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return set.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(set.Handle()),
		driver.Core1_2,
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
