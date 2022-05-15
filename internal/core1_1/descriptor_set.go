package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorSet struct {
	core1_0.DescriptorSet
}

func PromoteDescriptorSet(set core1_0.DescriptorSet) core1_1.DescriptorSet {
	if !set.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return set.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(set.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanDescriptorSet{set}
		}).(core1_1.DescriptorSet)
}

func PromoteDescriptorSetSlice(sets []core1_0.DescriptorSet) []core1_1.DescriptorSet {
	outSets := make([]core1_1.DescriptorSet, len(sets))
	for i := 0; i < len(sets); i++ {
		outSets[i] = PromoteDescriptorSet(sets[i])

		if outSets[i] == nil {
			return nil
		}
	}

	return outSets
}
