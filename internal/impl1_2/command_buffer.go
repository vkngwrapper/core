package impl1_2

import (
	"fmt"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) CmdBeginRenderPass2(commandBuffer core.CommandBuffer, renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin core1_2.SubpassBeginInfo) error {
	if commandBuffer.Handle() == 0 {
		return fmt.Errorf("commandBuffer cannot be uninitialized")
	}
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

	v.LoaderObj.VkCmdBeginRenderPass2(
		commandBuffer.Handle(),
		(*loader.VkRenderPassBeginInfo)(renderPassBeginPtr),
		(*loader.VkSubpassBeginInfo)(subpassBeginPtr),
	)

	return nil
}

func (v *DeviceVulkanDriver) CmdEndRenderPass2(commandBuffer core.CommandBuffer, subpassEnd core1_2.SubpassEndInfo) error {
	if commandBuffer.Handle() == 0 {
		return fmt.Errorf("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subpassEndPtr, err := common.AllocOptions(arena, subpassEnd)
	if err != nil {
		return err
	}

	v.LoaderObj.VkCmdEndRenderPass2(
		commandBuffer.Handle(),
		(*loader.VkSubpassEndInfo)(subpassEndPtr),
	)

	return nil
}

func (v *DeviceVulkanDriver) CmdNextSubpass2(commandBuffer core.CommandBuffer, subpassBegin core1_2.SubpassBeginInfo, subpassEnd core1_2.SubpassEndInfo) error {
	if commandBuffer.Handle() == 0 {
		return fmt.Errorf("commandBuffer cannot be uninitialized")
	}

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

	v.LoaderObj.VkCmdNextSubpass2(
		commandBuffer.Handle(),
		(*loader.VkSubpassBeginInfo)(subpassBeginPtr),
		(*loader.VkSubpassEndInfo)(subpassEndPtr),
	)

	return nil
}

func (v *DeviceVulkanDriver) CmdDrawIndexedIndirectCount(commandBuffer core.CommandBuffer, buffer core.Buffer, offset uint64, countBuffer core.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}
	if buffer.Handle() == 0 {
		panic("buffer cannot be nil")
	}
	if countBuffer.Handle() == 0 {
		panic("countBuffer cannot be nil")
	}
	v.LoaderObj.VkCmdDrawIndexedIndirectCount(
		commandBuffer.Handle(),
		buffer.Handle(),
		loader.VkDeviceSize(offset),
		countBuffer.Handle(),
		loader.VkDeviceSize(countBufferOffset),
		loader.Uint32(maxDrawCount),
		loader.Uint32(stride),
	)
}

func (v *DeviceVulkanDriver) CmdDrawIndirectCount(commandBuffer core.CommandBuffer, buffer core.Buffer, offset uint64, countBuffer core.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}
	if buffer.Handle() == 0 {
		panic("buffer cannot be uninitialized")
	}
	if countBuffer.Handle() == 0 {
		panic("countBuffer cannot be uninitialized")
	}
	v.LoaderObj.VkCmdDrawIndirectCount(
		commandBuffer.Handle(),
		buffer.Handle(),
		loader.VkDeviceSize(offset),
		countBuffer.Handle(),
		loader.VkDeviceSize(countBufferOffset),
		loader.Uint32(maxDrawCount),
		loader.Uint32(stride),
	)
}
