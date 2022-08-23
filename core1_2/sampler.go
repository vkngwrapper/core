package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanSampler is an implementation of the Sampler interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanSampler struct {
	core1_1.Sampler
}

// PromoteSampler accepts a Sampler object from any core version. If provided a sampler that supports
// at least core 1.2, it will return a core1_2.Sampler. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanSampler, even if it is provided a VulkanSampler from a higher
// core version. Two Vulkan 1.2 compatible Sampler objects with the same Sampler.Handle will
// return the same interface value when passed to this method.
func PromoteSampler(sampler core1_0.Sampler) Sampler {
	if !sampler.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedSampler := core1_1.PromoteSampler(sampler)
	return sampler.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(sampler.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanSampler{promotedSampler}
		}).(Sampler)
}
