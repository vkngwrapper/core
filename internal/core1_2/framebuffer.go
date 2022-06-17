package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanFramebuffer struct {
	core1_1.Framebuffer
}

func PromoteFramebuffer(framebuffer core1_0.Framebuffer) core1_2.Framebuffer {
	if !framebuffer.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return framebuffer.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(framebuffer.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanFramebuffer{core1_1.PromoteFramebuffer(framebuffer)}
		}).(core1_2.Framebuffer)
}
