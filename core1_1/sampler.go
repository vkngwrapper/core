package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanSampler struct {
	core1_0.Sampler
}

func PromoteSampler(sampler core1_0.Sampler) Sampler {
	if !sampler.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return sampler.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(sampler.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanSampler{sampler}
		}).(Sampler)
}
