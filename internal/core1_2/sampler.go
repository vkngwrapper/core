package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSampler struct {
	core1_1.Sampler
}

func PromoteSampler(sampler core1_0.Sampler) core1_2.Sampler {
	if !sampler.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return sampler.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(sampler.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanSampler{core1_1.PromoteSampler(sampler)}
		}).(core1_2.Sampler)
}
