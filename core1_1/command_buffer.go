package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanCommandBuffer is an implementation of the CommandBuffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanCommandBuffer struct {
	core1_0.CommandBuffer

	DeviceDriver        driver.Driver
	CommandBufferHandle driver.VkCommandBuffer

	CommandCounter *core1_0.CommandCounter
}

// PromoteCommandBuffer accepts a CommandBuffer object from any core version. If provided a command buffer that supports
// at least core 1.1, it will return a core1_1.CommandBuffer. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanCommandBuffer, even if it is provided a VulkanCommandBuffer from a higher
// core version. Two Vulkan 1.1 compatible CommandBuffer objects with the same CommandBuffer.Handle will
// return the same interface value when passed to this method.
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

// PromoteCommandBufferSlice accepts a slice of CommandBuffer objects from any core version.
// If provided a descriptor set that supports at least core 1.1, it will return a core1_1.CommandBuffer.
// Otherwise, it will left out of the returned slice. This method will always return a
// core1_1.VulkanCommandBuffer, even if it is provided a VulkanCommandBuffer from a higher core version. Two
// Vulkan 1.1 compatible CommandBuffer objects with the same CommandBuffer.Handle will return the same interface
// value when passed to this method.
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
