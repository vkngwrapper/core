package core1_2

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanFramebuffer is an implementation of the Framebuffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanFramebuffer struct {
	core1_1.Framebuffer
}

// PromoteFramebuffer accepts a Framebuffer object from any core version. If provided a framebuffer that supports
// at least core 1.2, it will return a core1_2.Framebuffer. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanFramebuffer, even if it is provided a VulkanFramebuffer from a higher
// core version. Two Vulkan 1.2 compatible Framebuffer objects with the same Framebuffer.Handle will
// return the same interface value when passed to this method.
func PromoteFramebuffer(framebuffer core1_0.Framebuffer) Framebuffer {
	if framebuffer == nil {
		return nil
	}
	if !framebuffer.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promoted, alreadyPromoted := framebuffer.(Framebuffer)
	if alreadyPromoted {
		return promoted
	}

	promotedFramebuffer := core1_1.PromoteFramebuffer(framebuffer)
	return framebuffer.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(framebuffer.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanFramebuffer{promotedFramebuffer}
		}).(Framebuffer)
}
