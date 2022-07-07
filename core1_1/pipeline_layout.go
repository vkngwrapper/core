package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanPipelineLayout struct {
	core1_0.PipelineLayout
}

func PromotePipelineLayout(layout core1_0.PipelineLayout) PipelineLayout {
	if !layout.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return layout.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(layout.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanPipelineLayout{layout}
		}).(PipelineLayout)
}
