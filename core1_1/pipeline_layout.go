package core1_1

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanPipelineLayout is an implementation of the PipelineLayout interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPipelineLayout struct {
	core1_0.PipelineLayout
}

// PromotePipelineLayout accepts a PipelineLayout object from any core version. If provided a pipeline layout that supports
// at least core 1.1, it will return a core1_1.PipelineLayout. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanPipelineLayout, even if it is provided a VulkanPipelineLayout from a higher
// core version. Two Vulkan 1.1 compatible PipelineLayout objects with the same PipelineLayout.Handle will
// return the same interface value when passed to this method.
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
