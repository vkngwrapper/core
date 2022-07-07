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
