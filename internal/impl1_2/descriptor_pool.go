package impl1_2

import (
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
)

// VulkanDescriptorPool is an implementation of the DescriptorPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorPool struct {
	impl1_1.VulkanDescriptorPool
}
