package impl1_2

import (
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
)

// VulkanInstance is an implementation of the Instance interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanInstance struct {
	impl1_1.VulkanInstance
}
