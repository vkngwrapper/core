package internal1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanCommandPool struct {
	core1_0.CommandPool

	DeviceDriver      driver.Driver
	CommandPoolHandle driver.VkCommandPool
	Device            driver.VkDevice
}

func (p *VulkanCommandPool) TrimCommandPool(flags core1_1.CommandPoolTrimFlags) {
	p.DeviceDriver.VkTrimCommandPool(p.Device,
		p.CommandPoolHandle,
		driver.VkCommandPoolTrimFlags(flags),
	)
}

func PromoteCommandPool(commandPool core1_0.CommandPool) core1_1.CommandPool {
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
		}).(core1_1.CommandPool)
}
