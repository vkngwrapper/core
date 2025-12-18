package core1_1

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanFramebuffer is an implementation of the Framebuffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanFramebuffer struct {
	core1_0.Framebuffer
}

// PromoteFramebuffer accepts a Framebuffer object from any core version. If provided a framebuffer that supports
// at least core 1.1, it will return a core1_1.Framebuffer. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanFramebuffer, even if it is provided a VulkanFramebuffer from a higher
// core version. Two Vulkan 1.1 compatible Framebuffer objects with the same Framebuffer.Handle will
// return the same interface value when passed to this method.
func PromoteFramebuffer(framebuffer core1_0.Framebuffer) Framebuffer {
	if framebuffer == nil {
		return nil
	}
	if !framebuffer.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	promoted, alreadyPromoted := framebuffer.(Framebuffer)
	if alreadyPromoted {
		return promoted
	}

	return framebuffer.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(framebuffer.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanFramebuffer{framebuffer}
		}).(Framebuffer)
}
