package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanCommandBuffer struct {
	core1_0.CommandBuffer

	DeviceDriver        driver.Driver
	CommandBufferHandle driver.VkCommandBuffer

	CommandCounter *core1_0.CommandCounter
}

func PromoteCommandBuffer(commandBuffer core1_0.CommandBuffer) CommandBuffer {
	if !commandBuffer.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return commandBuffer.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(commandBuffer.Handle()),
		driver.Core1_1,
		func() any {
			// The command/dispatch/draw counts should be shared between the various
			// core versions of a command buffer, but if for some reason this isn't a real
			// vulkan command buffer, feel free to just make up some new pointers
			var commandCounter *core1_0.CommandCounter
			rootCommandBuffer, isInternalVulkan := commandBuffer.(*core1_0.VulkanCommandBuffer)
			if isInternalVulkan {
				commandCounter = rootCommandBuffer.CommandCounter()
			} else {
				commandCounter = &core1_0.CommandCounter{}
			}

			return &VulkanCommandBuffer{
				CommandBuffer: commandBuffer,

				DeviceDriver:        commandBuffer.Driver(),
				CommandBufferHandle: commandBuffer.Handle(),

				CommandCounter: commandCounter,
			}
		}).(CommandBuffer)
}

func PromoteCommandBufferSlice(commandBuffers []core1_0.CommandBuffer) []CommandBuffer {
	outBuffers := make([]CommandBuffer, len(commandBuffers))

	for i := 0; i < len(commandBuffers); i++ {
		outBuffers[i] = PromoteCommandBuffer(commandBuffers[i])

		if outBuffers[i] == nil {
			return nil
		}
	}

	return outBuffers
}

func (c *VulkanCommandBuffer) CmdDispatchBase(baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int) {
	c.DeviceDriver.VkCmdDispatchBase(c.CommandBufferHandle,
		driver.Uint32(baseGroupX),
		driver.Uint32(baseGroupY),
		driver.Uint32(baseGroupZ),
		driver.Uint32(groupCountX),
		driver.Uint32(groupCountY),
		driver.Uint32(groupCountZ))
	c.CommandCounter.CommandCount++
	c.CommandCounter.DispatchCount++
}

func (c *VulkanCommandBuffer) CmdSetDeviceMask(deviceMask uint32) {
	c.DeviceDriver.VkCmdSetDeviceMask(c.CommandBufferHandle, driver.Uint32(deviceMask))
	c.CommandCounter.CommandCount++
}
