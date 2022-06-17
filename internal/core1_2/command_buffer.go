package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
	internal1_1 "github.com/CannibalVox/VKng/core/internal/core1_1"
	"github.com/CannibalVox/cgoparam"
)

type VulkanCommandBuffer struct {
	core1_1.CommandBuffer

	DeviceDriver        driver.Driver
	CommandBufferHandle driver.VkCommandBuffer

	CommandCount  *int
	DispatchCount *int
	DrawCallCount *int
}

func PromoteCommandBuffer(commandBuffer core1_0.CommandBuffer) core1_2.CommandBuffer {
	if !commandBuffer.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedBuffer := core1_1.PromoteCommandBuffer(commandBuffer)

	return commandBuffer.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(commandBuffer.Handle()),
		driver.Core1_2,
		func() any {
			var commandCount int
			var drawCount int
			var dispatchCount int

			// The command/dispatch/draw pointers should be shared between the various
			// core versions of a command buffer, but if for some reason this isn't a real
			// vulkan command buffer, feel free to just make up some new pointers
			commandCountPtr := &commandCount
			drawCountPtr := &drawCount
			dispatchCountPtr := &dispatchCount

			promotedBufferImpl, isInternalVulkan := promotedBuffer.(*internal1_1.VulkanCommandBuffer)
			if isInternalVulkan {
				commandCountPtr = promotedBufferImpl.CommandCount
				drawCountPtr = promotedBufferImpl.DrawCallCount
				dispatchCountPtr = promotedBufferImpl.DispatchCount
			}

			return &VulkanCommandBuffer{
				CommandBuffer: promotedBuffer,

				DeviceDriver:        commandBuffer.Driver(),
				CommandBufferHandle: commandBuffer.Handle(),

				CommandCount:  commandCountPtr,
				DrawCallCount: drawCountPtr,
				DispatchCount: dispatchCountPtr,
			}
		}).(core1_2.CommandBuffer)
}

func PromoteCommandBufferSlice(commandBuffers []core1_0.CommandBuffer) []core1_2.CommandBuffer {
	outBuffers := make([]core1_2.CommandBuffer, len(commandBuffers))

	for i := 0; i < len(commandBuffers); i++ {
		outBuffers[i] = PromoteCommandBuffer(commandBuffers[i])

		if outBuffers[i] == nil {
			return nil
		}
	}

	return outBuffers
}

func (c *VulkanCommandBuffer) CmdBeginRenderPass2(renderPassBegin core1_0.RenderPassBeginOptions, subpassBegin core1_2.SubpassBeginOptions) error {
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

	*c.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdEndRenderPass2(subpassEnd core1_2.SubpassEndOptions) error {
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

	*c.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdNextSubpass2(subpassBegin core1_2.SubpassBeginOptions, subpassEnd core1_2.SubpassEndOptions) error {
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

	*c.CommandCount++
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
	*c.DrawCallCount++
	*c.CommandCount++
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
	*c.DrawCallCount++
	*c.CommandCount++
}
