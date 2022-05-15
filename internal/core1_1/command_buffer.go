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
	internal1_0 "github.com/CannibalVox/VKng/core/internal/core1_0"
)

type VulkanCommandBuffer struct {
	core1_0.CommandBuffer

	DeviceDriver        driver.Driver
	CommandBufferHandle driver.VkCommandBuffer

	CommandCount  *int
	DispatchCount *int
}

func PromoteCommandBuffer(commandBuffer core1_0.CommandBuffer) core1_1.CommandBuffer {
	if !commandBuffer.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return commandBuffer.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(commandBuffer.Handle()),
		driver.Core1_1,
		func() any {
			var commandCount int
			var dispatchCount int

			// The command/dispatch/draw pointers should be shared between the various
			// core versions of a command buffer, but if for some reason this isn't a real
			// vulkan command buffer, feel free to just make up some new pointers
			commandCountPtr := &commandCount
			dispatchCountPtr := &dispatchCount
			rootCommandBuffer, isInternalVulkan := commandBuffer.(*internal1_0.VulkanCommandBuffer)
			if isInternalVulkan {
				commandCountPtr = rootCommandBuffer.CommandCount
				dispatchCountPtr = rootCommandBuffer.DispatchCount
			}

			return &VulkanCommandBuffer{
				CommandBuffer: commandBuffer,

				DeviceDriver:        commandBuffer.Driver(),
				CommandBufferHandle: commandBuffer.Handle(),

				CommandCount:  commandCountPtr,
				DispatchCount: dispatchCountPtr,
			}
		}).(core1_1.CommandBuffer)
}

func PromoteCommandBufferSlice(commandBuffers []core1_0.CommandBuffer) []core1_1.CommandBuffer {
	outBuffers := make([]core1_1.CommandBuffer, len(commandBuffers))

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
	*c.CommandCount++
	*c.DispatchCount++
}

func (c *VulkanCommandBuffer) CmdSetDeviceMask(deviceMask uint32) {
	c.DeviceDriver.VkCmdSetDeviceMask(c.CommandBufferHandle, driver.Uint32(deviceMask))
	*c.CommandCount++
}
