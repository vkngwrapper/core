package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSampler struct {
	core1_0.Sampler
}

func PromoteSampler(sampler core1_0.Sampler) core1_1.Sampler {
	if !sampler.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return sampler.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(sampler.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanSampler{sampler}
		}).(core1_1.Sampler)
}
