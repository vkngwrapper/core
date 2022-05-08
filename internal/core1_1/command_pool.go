package internal1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanCommandPool struct {
	DeviceDriver      driver.Driver
	CommandPoolHandle driver.VkCommandPool
	DeviceHandle      driver.VkDevice
}

func (p *VulkanCommandPool) TrimCommandPool(flags core1_1.CommandPoolTrimFlags) {
	p.DeviceDriver.VkTrimCommandPool(p.DeviceHandle,
		p.CommandPoolHandle,
		driver.VkCommandPoolTrimFlags(flags),
	)
}
