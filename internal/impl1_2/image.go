package impl1_2

import (
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
)

// VulkanImage is an implementation of the Image interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanImage struct {
	impl1_1.VulkanImage
}
