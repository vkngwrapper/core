package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

// VulkanCommandPool is an implementation of the CommandPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanCommandPool struct {
	core1_1.CommandPool
}

// PromoteCommandPool accepts a CommandPool object from any core version. If provided a command pool that supports
// at least core 1.2, it will return a core1_2.CommandPool. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanCommandPool, even if it is provided a VulkanCommandPool from a higher
// core version. Two Vulkan 1.2 compatible CommandPool objects with the same CommandPool.Handle will
// return the same interface value when passed to this method.
func PromoteCommandPool(commandPool core1_0.CommandPool) CommandPool {
	if !commandPool.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedCommandPool := core1_1.PromoteCommandPool(commandPool)

	return commandPool.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(commandPool.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanCommandPool{
				CommandPool: promotedCommandPool,
			}
		}).(CommandPool)
}
