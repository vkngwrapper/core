package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanRenderPass struct {
	core1_1.RenderPass
}

func PromoteRenderPass(renderPass core1_0.RenderPass) RenderPass {
	if !renderPass.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedRenderPass := core1_1.PromoteRenderPass(renderPass)
	return renderPass.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(renderPass.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanRenderPass{promotedRenderPass}
		}).(RenderPass)
}
