package core1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanFramebuffer struct {
	core1_0.Framebuffer
}

func PromoteFramebuffer(framebuffer core1_0.Framebuffer) Framebuffer {
	if !framebuffer.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return framebuffer.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(framebuffer.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanFramebuffer{framebuffer}
		}).(Framebuffer)
}
