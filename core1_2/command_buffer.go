package core1_2

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanCommandBuffer is an implementation of the CommandBuffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanCommandBuffer struct {
	core1_1.CommandBuffer

	DeviceDriver        driver.Driver
	CommandBufferHandle driver.VkCommandBuffer

	CommandCounter *core1_0.CommandCounter
}

// PromoteCommandBuffer accepts a CommandBuffer object from any core version. If provided a command buffer that supports
// at least core 1.2, it will return a core1_2.CommandBuffer. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanCommandBuffer, even if it is provided a VulkanCommandBuffer from a higher
// core version. Two Vulkan 1.2 compatible CommandBuffer objects with the same CommandBuffer.Handle will
// return the same interface value when passed to this method.
func PromoteCommandBuffer(commandBuffer core1_0.CommandBuffer) CommandBuffer {
	if !commandBuffer.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedBuffer := core1_1.PromoteCommandBuffer(commandBuffer)

	return commandBuffer.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(commandBuffer.Handle()),
		driver.Core1_2,
		func() any {
			// The command/dispatch/draw counts should be shared between the various
			// core versions of a command buffer, but if for some reason this isn't a real
			// vulkan command buffer, feel free to just make up some new pointers
			var commandCounter *core1_0.CommandCounter

			promotedBufferImpl, isInternalVulkan := promotedBuffer.(*core1_1.VulkanCommandBuffer)
			if isInternalVulkan {
				baseBuffer, isInternalVulkan := promotedBufferImpl.CommandBuffer.(*core1_0.VulkanCommandBuffer)
				if isInternalVulkan {
					commandCounter = baseBuffer.CommandCounter()
				}
			}

			if commandCounter == nil {
				commandCounter = &core1_0.CommandCounter{}
			}

			return &VulkanCommandBuffer{
				CommandBuffer: promotedBuffer,

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

func (c *VulkanCommandBuffer) CmdBeginRenderPass2(renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin SubpassBeginInfo) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	renderPassBeginPtr, err := common.AllocOptions(arena, renderPassBegin)
	if err != nil {
		return err
	}

	subpassBeginPtr, err := common.AllocOptions(arena, subpassBegin)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdBeginRenderPass2(
		c.CommandBufferHandle,
		(*driver.VkRenderPassBeginInfo)(renderPassBeginPtr),
		(*driver.VkSubpassBeginInfo)(subpassBeginPtr),
	)

	c.CommandCounter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdEndRenderPass2(subpassEnd SubpassEndInfo) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subpassEndPtr, err := common.AllocOptions(arena, subpassEnd)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdEndRenderPass2(
		c.CommandBufferHandle,
		(*driver.VkSubpassEndInfo)(subpassEndPtr),
	)

	c.CommandCounter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdNextSubpass2(subpassBegin SubpassBeginInfo, subpassEnd SubpassEndInfo) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subpassBeginPtr, err := common.AllocOptions(arena, subpassBegin)
	if err != nil {
		return err
	}

	subpassEndPtr, err := common.AllocOptions(arena, subpassEnd)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdNextSubpass2(
		c.CommandBufferHandle,
		(*driver.VkSubpassBeginInfo)(subpassBeginPtr),
		(*driver.VkSubpassEndInfo)(subpassEndPtr),
	)

	c.CommandCounter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdDrawIndexedIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	c.DeviceDriver.VkCmdDrawIndexedIndirectCount(
		c.CommandBufferHandle,
		buffer.Handle(),
		driver.VkDeviceSize(offset),
		countBuffer.Handle(),
		driver.VkDeviceSize(countBufferOffset),
		driver.Uint32(maxDrawCount),
		driver.Uint32(stride),
	)
	c.CommandCounter.CommandCount++
	c.CommandCounter.DrawCallCount++
}

func (c *VulkanCommandBuffer) CmdDrawIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	c.DeviceDriver.VkCmdDrawIndirectCount(
		c.CommandBufferHandle,
		buffer.Handle(),
		driver.VkDeviceSize(offset),
		countBuffer.Handle(),
		driver.VkDeviceSize(countBufferOffset),
		driver.Uint32(maxDrawCount),
		driver.Uint32(stride),
	)
	c.CommandCounter.CommandCount++
	c.CommandCounter.DrawCallCount++
}
