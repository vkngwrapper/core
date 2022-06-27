package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSamplerYcbcrConversion struct {
	core1_1.SamplerYcbcrConversion
}

func PromoteSamplerYcbcrConversion(ycbcr core1_1.SamplerYcbcrConversion) SamplerYcbcrConversion {
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
