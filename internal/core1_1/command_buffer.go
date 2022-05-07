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
}

func (cb *VulkanCommandBuffer) CmdDispatchBase(baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int) {
	cb.DeviceDriver.VkCmdDispatchBase(cb.CommandBufferHandle,
		driver.Uint32(baseGroupX),
		driver.Uint32(baseGroupY),
		driver.Uint32(baseGroupZ),
		driver.Uint32(groupCountX),
		driver.Uint32(groupCountY),
		driver.Uint32(groupCountZ))
}

func (cb *VulkanCommandBuffer) CmdSetDeviceMask(deviceMask uint32) {
	cb.DeviceDriver.VkCmdSetDeviceMask(cb.CommandBufferHandle, driver.Uint32(deviceMask))
}
