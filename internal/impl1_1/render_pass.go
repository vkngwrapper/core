package impl1_1

import (
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
)

// VulkanRenderPass is an implementation of the RenderPass interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanRenderPass struct {
	impl1_0.VulkanRenderPass
}
