package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanPipelineLayout struct {
	core1_1.PipelineLayout
}

func PromotePipelineLayout(layout core1_0.PipelineLayout) PipelineLayout {
	if !layout.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedLayout := core1_1.PromotePipelineLayout(layout)
	return layout.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(layout.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanPipelineLayout{promotedLayout}
		}).(PipelineLayout)
}
