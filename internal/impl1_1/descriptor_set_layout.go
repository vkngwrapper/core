package impl1_1

import (
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
)

// VulkanDescriptorSetLayout is an implementation of the DescriptorSetLayout interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorSetLayout struct {
	impl1_0.VulkanDescriptorSetLayout
}
