package impl1_2

import (
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
)

// VulkanFence is an implementation of the Fence interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanFence struct {
	impl1_1.VulkanFence
}
