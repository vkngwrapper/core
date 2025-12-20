package impl1_2

import (
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
)

// VulkanSampler is an implementation of the Sampler interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanSampler struct {
	impl1_1.VulkanSampler
}
