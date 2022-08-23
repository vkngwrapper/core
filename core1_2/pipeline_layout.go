package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanPipelineLayout is an implementation of the PipelineLayout interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPipelineLayout struct {
	core1_1.PipelineLayout
}

// PromotePipelineLayout accepts a PipelineLayout object from any core version. If provided a pipeline cache that supports
// at least core 1.2, it will return a core1_2.PipelineLayout. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanPipelineLayout, even if it is provided a VulkanPipelineLayout from a higher
// core version. Two Vulkan 1.2 compatible PipelineLayout objects with the same PipelineLayout.Handle will
// return the same interface value when passed to this method.
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
