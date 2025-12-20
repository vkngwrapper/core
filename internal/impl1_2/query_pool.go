package impl1_2

import (
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
)

// VulkanQueryPool is an implementation of the QueryPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanQueryPool struct {
	impl1_1.VulkanQueryPool
}

func (q *VulkanQueryPool) Reset(firstQuery, queryCount int) {
	q.Driver().VkResetQueryPool(q.DeviceHandle(), q.Handle(), driver.Uint32(firstQuery), driver.Uint32(queryCount))
}
