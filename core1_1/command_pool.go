package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

// VulkanCommandPool is an implementation of the CommandPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanCommandPool struct {
	core1_0.CommandPool

	DeviceDriver      driver.Driver
	CommandPoolHandle driver.VkCommandPool
	Device            driver.VkDevice
}

func (p *VulkanCommandPool) TrimCommandPool(flags CommandPoolTrimFlags) {
	p.DeviceDriver.VkTrimCommandPool(p.Device,
		p.CommandPoolHandle,
		driver.VkCommandPoolTrimFlags(flags),
	)
}

// PromoteCommandPool accepts a CommandPool object from any core version. If provided a command pool that supports
// at least core 1.1, it will return a core1_1.CommandPool. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanCommandPool, even if it is provided a VulkanCommandPool from a higher
// core version. Two Vulkan 1.1 compatible CommandPool objects with the same CommandPool.Handle will
// return the same interface value when passed to this method.
func PromoteCommandPool(commandPool core1_0.CommandPool) CommandPool {
	if !commandPool.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return commandPool.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(commandPool.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanCommandPool{
				CommandPool: commandPool,

				DeviceDriver:      commandPool.Driver(),
				Device:            commandPool.DeviceHandle(),
				CommandPoolHandle: commandPool.Handle(),
			}
		}).(CommandPool)
}
