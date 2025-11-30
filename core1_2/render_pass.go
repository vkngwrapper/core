package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanRenderPass is an implementation of the RenderPass interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanRenderPass struct {
	core1_1.RenderPass
}

// PromoteRenderPass accepts a RenderPass object from any core version. If provided a render pass that supports
// at least core 1.2, it will return a core1_2.RenderPass. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanRenderPass, even if it is provided a VulkanRenderPass from a higher
// core version. Two Vulkan 1.2 compatible RenderPass objects with the same RenderPass.Handle will
// return the same interface value when passed to this method.
func PromoteRenderPass(renderPass core1_0.RenderPass) RenderPass {
	if renderPass == nil {
		return nil
	}
	if !renderPass.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promoted, alreadyPromoted := renderPass.(RenderPass)
	if alreadyPromoted {
		return promoted
	}

	promotedRenderPass := core1_1.PromoteRenderPass(renderPass)
	return renderPass.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(renderPass.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanRenderPass{promotedRenderPass}
		}).(RenderPass)
}
