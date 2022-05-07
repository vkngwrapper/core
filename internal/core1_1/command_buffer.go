package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanCommandBuffer struct {
	DeviceDriver        driver.Driver
	Device              driver.VkDevice
	CommandPool         driver.VkCommandPool
	CommandBufferHandle driver.VkCommandBuffer

	CommandCount  *int
	DispatchCount *int
}

func (c *VulkanCommandBuffer) CmdDispatchBase(baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int) {
	c.DeviceDriver.VkCmdDispatchBase(c.CommandBufferHandle,
		driver.Uint32(baseGroupX),
		driver.Uint32(baseGroupY),
		driver.Uint32(baseGroupZ),
		driver.Uint32(groupCountX),
		driver.Uint32(groupCountY),
		driver.Uint32(groupCountZ))
	*c.CommandCount++
	*c.DispatchCount++
}

func (c *VulkanCommandBuffer) CmdSetDeviceMask(deviceMask uint32) {
	c.DeviceDriver.VkCmdSetDeviceMask(c.CommandBufferHandle, driver.Uint32(deviceMask))
	*c.CommandCount++
}
