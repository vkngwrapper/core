package impl1_2

import (
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
)

// VulkanDescriptorUpdateTemplate is an implementation of the DescriptorUpdateTemplate interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorUpdateTemplate struct {
	impl1_1.VulkanDescriptorUpdateTemplate
}
