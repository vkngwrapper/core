package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanSamplerYcbcrConversion is an implementation of the SamplerYcbcrConversion interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanSamplerYcbcrConversion struct {
	core1_1.SamplerYcbcrConversion
}

// PromoteSamplerYcbcrConversion accepts a SamplerYcbcrConversion object from any core version. If provided a sampler ycbcr conversion that supports
// at least core 1.2, it will return a core1_2.SamplerYcbcrConversion. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanSamplerYcbcrConversion, even if it is provided a VulkanSamplerYcbcrConversion from a higher
// core version. Two Vulkan 1.2 compatible SamplerYcbcrConversion objects with the same SamplerYcbcrConversion.Handle will
// return the same interface value when passed to this method.
func PromoteSamplerYcbcrConversion(ycbcr core1_1.SamplerYcbcrConversion) SamplerYcbcrConversion {
	if ycbcr == nil {
		return nil
	}
	if !ycbcr.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return ycbcr.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(ycbcr.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanSamplerYcbcrConversion{ycbcr}
		}).(SamplerYcbcrConversion)
}
