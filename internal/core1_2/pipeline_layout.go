package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanPipelineLayout struct {
	core1_1.PipelineLayout
}

func PromotePipelineLayout(layout core1_0.PipelineLayout) core1_2.PipelineLayout {
	if !layout.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return layout.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(layout.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanPipelineLayout{core1_1.PromotePipelineLayout(layout)}
		}).(core1_2.PipelineLayout)
}
