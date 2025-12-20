package impl1_2

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
)

// VulkanCommandBuffer is an implementation of the CommandBuffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanCommandBuffer struct {
	impl1_1.VulkanCommandBuffer
}

func (c *VulkanCommandBuffer) CmdBeginRenderPass2(renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin core1_2.SubpassBeginInfo) error {
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

	c.Driver().VkCmdBeginRenderPass2(
		c.Handle(),
		(*driver.VkRenderPassBeginInfo)(renderPassBeginPtr),
		(*driver.VkSubpassBeginInfo)(subpassBeginPtr),
	)

	c.CommandCounter().CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdEndRenderPass2(subpassEnd core1_2.SubpassEndInfo) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subpassEndPtr, err := common.AllocOptions(arena, subpassEnd)
	if err != nil {
		return err
	}

	c.Driver().VkCmdEndRenderPass2(
		c.Handle(),
		(*driver.VkSubpassEndInfo)(subpassEndPtr),
	)

	c.CommandCounter().CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdNextSubpass2(subpassBegin core1_2.SubpassBeginInfo, subpassEnd core1_2.SubpassEndInfo) error {
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

	c.Driver().VkCmdNextSubpass2(
		c.Handle(),
		(*driver.VkSubpassBeginInfo)(subpassBeginPtr),
		(*driver.VkSubpassEndInfo)(subpassEndPtr),
	)

	c.CommandCounter().CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdDrawIndexedIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	if buffer == nil {
		panic("buffer cannot be nil")
	}
	if countBuffer == nil {
		panic("countBuffer cannot be nil")
	}
	c.Driver().VkCmdDrawIndexedIndirectCount(
		c.Handle(),
		buffer.Handle(),
		driver.VkDeviceSize(offset),
		countBuffer.Handle(),
		driver.VkDeviceSize(countBufferOffset),
		driver.Uint32(maxDrawCount),
		driver.Uint32(stride),
	)
	counter := c.CommandCounter()
	counter.CommandCount++
	counter.DrawCallCount++
}

func (c *VulkanCommandBuffer) CmdDrawIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	if buffer == nil {
		panic("buffer cannot be nil")
	}
	if countBuffer == nil {
		panic("countBuffer cannot be nil")
	}
	c.Driver().VkCmdDrawIndirectCount(
		c.Handle(),
		buffer.Handle(),
		driver.VkDeviceSize(offset),
		countBuffer.Handle(),
		driver.VkDeviceSize(countBufferOffset),
		driver.Uint32(maxDrawCount),
		driver.Uint32(stride),
	)
	counter := c.CommandCounter()
	counter.CommandCount++
	counter.DrawCallCount++
}
