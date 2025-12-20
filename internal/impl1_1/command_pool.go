package impl1_1

import (
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
)

// VulkanCommandPool is an implementation of the CommandPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanCommandPool struct {
	impl1_0.VulkanCommandPool
}

func (p *VulkanCommandPool) TrimCommandPool(flags core1_1.CommandPoolTrimFlags) {
	p.Driver().VkTrimCommandPool(p.DeviceHandle(),
		p.Handle(),
		driver.VkCommandPoolTrimFlags(flags),
	)
}
