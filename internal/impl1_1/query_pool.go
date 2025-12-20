package impl1_1

import (
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
)

// VulkanQueryPool is an implementation of the QueryPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanQueryPool struct {
	impl1_0.VulkanQueryPool
}
