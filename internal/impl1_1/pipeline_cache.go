package impl1_1

import (
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
)

// VulkanPipelineCache is an implementation of the PipelineCache interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPipelineCache struct {
	impl1_0.VulkanPipelineCache
}
