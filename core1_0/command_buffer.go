package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
	"unsafe"
)

type VulkanCommandBuffer struct {
	deviceDriver        driver.Driver
	device              driver.VkDevice
	commandPool         driver.VkCommandPool
	commandBufferHandle driver.VkCommandBuffer

	commandCounter *CommandCounter

	maximumAPIVersion common.APIVersion
}

func (c *VulkanCommandBuffer) Handle() driver.VkCommandBuffer {
	return c.commandBufferHandle
}

func (c *VulkanCommandBuffer) CommandPoolHandle() driver.VkCommandPool {
	return c.commandPool
}

func (c *VulkanCommandBuffer) DeviceHandle() driver.VkDevice {
	return c.device
}

func (c *VulkanCommandBuffer) Driver() driver.Driver {
	return c.deviceDriver
}

func (c *VulkanCommandBuffer) APIVersion() common.APIVersion {
	return c.maximumAPIVersion
}

func (c *VulkanCommandBuffer) CommandCounter() *CommandCounter {
	return c.commandCounter
}

func (c *VulkanCommandBuffer) CommandsRecorded() int {
	return c.commandCounter.CommandCount
}

func (c *VulkanCommandBuffer) DrawsRecorded() int {
	return c.commandCounter.DrawCallCount
}

func (c *VulkanCommandBuffer) DispatchesRecorded() int {
	return c.commandCounter.DispatchCount
}

func (c *VulkanCommandBuffer) Begin(o CommandBufferBeginInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return VKErrorUnknown, err
	}

	res, err := c.deviceDriver.VkBeginCommandBuffer(c.commandBufferHandle, (*driver.VkCommandBufferBeginInfo)(createInfo))
	if err == nil {
		c.commandCounter.CommandCount = 0
		c.commandCounter.DrawCallCount = 0
		c.commandCounter.DispatchCount = 0
	}

	return res, err
}

func (c *VulkanCommandBuffer) End() (common.VkResult, error) {
	return c.deviceDriver.VkEndCommandBuffer(c.commandBufferHandle)
}

func (c *VulkanCommandBuffer) CmdBeginRenderPass(contents SubpassContents, o RenderPassBeginInfo) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	c.deviceDriver.VkCmdBeginRenderPass(c.commandBufferHandle, (*driver.VkRenderPassBeginInfo)(createInfo), driver.VkSubpassContents(contents))
	c.commandCounter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdEndRenderPass() {
	c.deviceDriver.VkCmdEndRenderPass(c.commandBufferHandle)
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdBindPipeline(bindPoint PipelineBindPoint, pipeline Pipeline) {
	c.deviceDriver.VkCmdBindPipeline(c.commandBufferHandle, driver.VkPipelineBindPoint(bindPoint), pipeline.Handle())
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) {
	c.deviceDriver.VkCmdDraw(c.commandBufferHandle, driver.Uint32(vertexCount), driver.Uint32(instanceCount), driver.Uint32(firstVertex), driver.Uint32(firstInstance))
	c.commandCounter.CommandCount++
	c.commandCounter.DrawCallCount++
}

func (c *VulkanCommandBuffer) CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) {
	c.deviceDriver.VkCmdDrawIndexed(c.commandBufferHandle, driver.Uint32(indexCount), driver.Uint32(instanceCount), driver.Uint32(firstIndex), driver.Int32(vertexOffset), driver.Uint32(firstInstance))
	c.commandCounter.CommandCount++
	c.commandCounter.DrawCallCount++
}

func (c *VulkanCommandBuffer) CmdBindVertexBuffers(buffers []Buffer, bufferOffsets []int) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	bufferCount := len(buffers)

	bufferArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkBuffer{})))
	offsetArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof(C.VkDeviceSize(0))))

	bufferArrayPtr := (*driver.VkBuffer)(bufferArrayUnsafe)
	offsetArrayPtr := (*driver.VkDeviceSize)(offsetArrayUnsafe)

	bufferArraySlice := ([]driver.VkBuffer)(unsafe.Slice(bufferArrayPtr, bufferCount))
	offsetArraySlice := ([]driver.VkDeviceSize)(unsafe.Slice(offsetArrayPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		bufferArraySlice[i] = buffers[i].Handle()
		offsetArraySlice[i] = driver.VkDeviceSize(bufferOffsets[i])
	}

	c.deviceDriver.VkCmdBindVertexBuffers(c.commandBufferHandle, driver.Uint32(0), driver.Uint32(bufferCount), bufferArrayPtr, offsetArrayPtr)
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdBindIndexBuffer(buffer Buffer, offset int, indexType IndexType) {
	c.deviceDriver.VkCmdBindIndexBuffer(c.commandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset), driver.VkIndexType(indexType))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdBindDescriptorSets(bindPoint PipelineBindPoint, layout PipelineLayout, sets []DescriptorSet, dynamicOffsets []int) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	setCount := len(sets)
	dynamicOffsetCount := len(dynamicOffsets)

	var setPtr unsafe.Pointer
	var dynamicOffsetPtr unsafe.Pointer

	if setCount > 0 {
		setPtr = arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{})))
		setSlice := ([]C.VkDescriptorSet)(unsafe.Slice((*C.VkDescriptorSet)(setPtr), setCount))
		for i := 0; i < setCount; i++ {
			setSlice[i] = (C.VkDescriptorSet)(unsafe.Pointer(sets[i].Handle()))
		}
	}

	if dynamicOffsetCount > 0 {
		dynamicOffsetPtr = arena.Malloc(dynamicOffsetCount * int(unsafe.Sizeof(C.uint32_t(0))))
		dynamicOffsetSlice := ([]C.uint32_t)(unsafe.Slice((*C.uint32_t)(dynamicOffsetPtr), dynamicOffsetCount))

		for i := 0; i < dynamicOffsetCount; i++ {
			dynamicOffsetSlice[i] = (C.uint32_t)(dynamicOffsets[i])
		}
	}

	c.deviceDriver.VkCmdBindDescriptorSets(c.commandBufferHandle,
		driver.VkPipelineBindPoint(bindPoint),
		layout.Handle(),
		driver.Uint32(0),
		driver.Uint32(setCount),
		(*driver.VkDescriptorSet)(setPtr),
		driver.Uint32(dynamicOffsetCount),
		(*driver.Uint32)(dynamicOffsetPtr))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdPipelineBarrier(srcStageMask, dstStageMask PipelineStageFlags, dependencies DependencyFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	barrierCount := len(memoryBarriers)
	bufferBarrierCount := len(bufferMemoryBarriers)
	imageBarrierCount := len(imageMemoryBarriers)

	var err error
	var barrierPtr *C.VkMemoryBarrier
	var bufferBarrierPtr *C.VkBufferMemoryBarrier
	var imageBarrierPtr *C.VkImageMemoryBarrier

	if barrierCount > 0 {
		barrierPtr, err = common.AllocOptionSlice[C.VkMemoryBarrier, MemoryBarrier](arena, memoryBarriers)
		if err != nil {
			return err
		}
	}

	if bufferBarrierCount > 0 {
		bufferBarrierPtr, err = common.AllocOptionSlice[C.VkBufferMemoryBarrier, BufferMemoryBarrier](arena, bufferMemoryBarriers)
		if err != nil {
			return err
		}
	}

	if imageBarrierCount > 0 {
		imageBarrierPtr, err = common.AllocOptionSlice[C.VkImageMemoryBarrier, ImageMemoryBarrier](arena, imageMemoryBarriers)
		if err != nil {
			return err
		}
	}

	c.deviceDriver.VkCmdPipelineBarrier(c.commandBufferHandle, driver.VkPipelineStageFlags(srcStageMask), driver.VkPipelineStageFlags(dstStageMask), driver.VkDependencyFlags(dependencies), driver.Uint32(barrierCount), (*driver.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), driver.Uint32(bufferBarrierCount), (*driver.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), driver.Uint32(imageBarrierCount), (*driver.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	c.commandCounter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdCopyBufferToImage(buffer Buffer, image Image, layout ImageLayout, regions []BufferImageCopy) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	var err error
	var regionPtr *C.VkBufferImageCopy

	if regionCount > 0 {
		regionPtr, err = common.AllocSlice[C.VkBufferImageCopy, BufferImageCopy](arena, regions)
		if err != nil {
			return err
		}
	}

	c.deviceDriver.VkCmdCopyBufferToImage(c.commandBufferHandle, buffer.Handle(), image.Handle(), driver.VkImageLayout(layout), driver.Uint32(regionCount), (*driver.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))
	c.commandCounter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdBlitImage(sourceImage Image, sourceImageLayout ImageLayout, destinationImage Image, destinationImageLayout ImageLayout, regions []ImageBlit, filter Filter) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	regionCount := len(regions)

	regionPtr, err := common.AllocSlice[C.VkImageBlit, ImageBlit](allocator, regions)
	if err != nil {
		return err
	}

	c.deviceDriver.VkCmdBlitImage(
		c.commandBufferHandle,
		driver.VkImage(sourceImage.Handle()),
		driver.VkImageLayout(sourceImageLayout),
		driver.VkImage(destinationImage.Handle()),
		driver.VkImageLayout(destinationImageLayout),
		driver.Uint32(regionCount),
		(*driver.VkImageBlit)(unsafe.Pointer(regionPtr)),
		driver.VkFilter(filter))
	c.commandCounter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdPushConstants(layout PipelineLayout, stageFlags ShaderStageFlags, offset int, valueBytes []byte) {
	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	valueBytesPtr := alloc.CBytes(valueBytes)

	c.deviceDriver.VkCmdPushConstants(c.commandBufferHandle, layout.Handle(), driver.VkShaderStageFlags(stageFlags), driver.Uint32(offset), driver.Uint32(len(valueBytes)), valueBytesPtr)
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetViewport(viewports []Viewport) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	viewportCount := len(viewports)
	var viewportPtr *C.VkViewport

	if viewportCount > 0 {
		viewportPtr = (*C.VkViewport)(allocator.Malloc(viewportCount * C.sizeof_struct_VkViewport))
		viewportSlice := ([]C.VkViewport)(unsafe.Slice(viewportPtr, viewportCount))

		for i := 0; i < viewportCount; i++ {
			viewport := viewports[i]
			viewportSlice[i].x = C.float(viewport.X)
			viewportSlice[i].y = C.float(viewport.Y)
			viewportSlice[i].width = C.float(viewport.Width)
			viewportSlice[i].height = C.float(viewport.Height)
			viewportSlice[i].minDepth = C.float(viewport.MinDepth)
			viewportSlice[i].maxDepth = C.float(viewport.MaxDepth)
		}
	}

	c.deviceDriver.VkCmdSetViewport(c.commandBufferHandle, driver.Uint32(0), driver.Uint32(viewportCount), (*driver.VkViewport)(unsafe.Pointer(viewportPtr)))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetScissor(scissors []Rect2D) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	scissorCount := len(scissors)
	var scissorPtr *C.VkRect2D

	if scissorCount > 0 {
		scissorPtr = (*C.VkRect2D)(allocator.Malloc(scissorCount * C.sizeof_struct_VkRect2D))
		scissorSlice := ([]C.VkRect2D)(unsafe.Slice(scissorPtr, scissorCount))

		for i := 0; i < scissorCount; i++ {
			scissor := scissors[i]
			scissorSlice[i].offset.x = C.int32_t(scissor.Offset.X)
			scissorSlice[i].offset.y = C.int32_t(scissor.Offset.Y)
			scissorSlice[i].extent.width = C.uint32_t(scissor.Extent.Width)
			scissorSlice[i].extent.height = C.uint32_t(scissor.Extent.Height)
		}
	}

	c.deviceDriver.VkCmdSetScissor(c.commandBufferHandle, driver.Uint32(0), driver.Uint32(scissorCount), (*driver.VkRect2D)(unsafe.Pointer(scissorPtr)))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdNextSubpass(contents SubpassContents) {
	c.deviceDriver.VkCmdNextSubpass(c.commandBufferHandle, driver.VkSubpassContents(contents))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdWaitEvents(events []Event, srcStageMask PipelineStageFlags, dstStageMask PipelineStageFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	eventCount := len(events)
	barrierCount := len(memoryBarriers)
	bufferBarrierCount := len(bufferMemoryBarriers)
	imageBarrierCount := len(imageMemoryBarriers)

	var err error
	var eventPtr *C.VkEvent
	var barrierPtr *C.VkMemoryBarrier
	var bufferBarrierPtr *C.VkBufferMemoryBarrier
	var imageBarrierPtr *C.VkImageMemoryBarrier

	if eventCount > 0 {
		eventPtr = (*C.VkEvent)(arena.Malloc(eventCount * int(unsafe.Sizeof([1]C.VkEvent{}))))
		eventSlice := ([]C.VkEvent)(unsafe.Slice(eventPtr, eventCount))

		for i := 0; i < eventCount; i++ {
			eventSlice[i] = C.VkEvent(unsafe.Pointer(events[i].Handle()))
		}
	}

	if barrierCount > 0 {
		barrierPtr, err = common.AllocOptionSlice[C.VkMemoryBarrier, MemoryBarrier](arena, memoryBarriers)
		if err != nil {
			return err
		}
	}

	if bufferBarrierCount > 0 {
		bufferBarrierPtr, err = common.AllocOptionSlice[C.VkBufferMemoryBarrier, BufferMemoryBarrier](arena, bufferMemoryBarriers)
		if err != nil {
			return err
		}
	}

	if imageBarrierCount > 0 {
		imageBarrierPtr, err = common.AllocOptionSlice[C.VkImageMemoryBarrier, ImageMemoryBarrier](arena, imageMemoryBarriers)
		if err != nil {
			return err
		}
	}

	c.deviceDriver.VkCmdWaitEvents(c.commandBufferHandle, driver.Uint32(eventCount), (*driver.VkEvent)(unsafe.Pointer(eventPtr)), driver.VkPipelineStageFlags(srcStageMask), driver.VkPipelineStageFlags(dstStageMask), driver.Uint32(barrierCount), (*driver.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), driver.Uint32(bufferBarrierCount), (*driver.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), driver.Uint32(imageBarrierCount), (*driver.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	c.commandCounter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdSetEvent(event Event, stageMask PipelineStageFlags) {
	c.deviceDriver.VkCmdSetEvent(c.commandBufferHandle, event.Handle(), driver.VkPipelineStageFlags(stageMask))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdClearColorImage(image Image, imageLayout ImageLayout, color ClearColorValue, ranges []ImageSubresourceRange) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	var pRanges *C.VkImageSubresourceRange

	if rangeCount > 0 {
		pRanges = (*C.VkImageSubresourceRange)(arena.Malloc(rangeCount * C.sizeof_struct_VkImageSubresourceRange))
		rangeSlice := ([]C.VkImageSubresourceRange)(unsafe.Slice(pRanges, rangeCount))

		for rangeIndex, oneRange := range ranges {
			rangeSlice[rangeIndex].aspectMask = C.VkImageAspectFlags(oneRange.AspectMask)
			rangeSlice[rangeIndex].baseMipLevel = C.uint32_t(oneRange.BaseMipLevel)
			rangeSlice[rangeIndex].levelCount = C.uint32_t(oneRange.LevelCount)
			rangeSlice[rangeIndex].baseArrayLayer = C.uint32_t(oneRange.BaseArrayLayer)
			rangeSlice[rangeIndex].layerCount = C.uint32_t(oneRange.LayerCount)
		}
	}

	var pColor unsafe.Pointer
	if color != nil {
		pColor = arena.Malloc(C.sizeof_union_VkClearColorValue)
		color.PopulateColorUnion(pColor)
	}

	c.deviceDriver.VkCmdClearColorImage(c.commandBufferHandle, image.Handle(), driver.VkImageLayout(imageLayout), (*driver.VkClearColorValue)(pColor), driver.Uint32(rangeCount), (*driver.VkImageSubresourceRange)(unsafe.Pointer(pRanges)))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdResetQueryPool(queryPool QueryPool, startQuery, queryCount int) {
	c.deviceDriver.VkCmdResetQueryPool(c.commandBufferHandle, queryPool.Handle(), driver.Uint32(startQuery), driver.Uint32(queryCount))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdBeginQuery(queryPool QueryPool, query int, flags QueryControlFlags) {
	c.deviceDriver.VkCmdBeginQuery(c.commandBufferHandle, queryPool.Handle(), driver.Uint32(query), driver.VkQueryControlFlags(flags))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdEndQuery(queryPool QueryPool, query int) {
	c.deviceDriver.VkCmdEndQuery(c.commandBufferHandle, queryPool.Handle(), driver.Uint32(query))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdCopyQueryPoolResults(queryPool QueryPool, firstQuery, queryCount int, dstBuffer Buffer, dstOffset, stride int, flags QueryResultFlags) {
	c.deviceDriver.VkCmdCopyQueryPoolResults(c.commandBufferHandle, queryPool.Handle(), driver.Uint32(firstQuery), driver.Uint32(queryCount), dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(stride), driver.VkQueryResultFlags(flags))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdExecuteCommands(commandBuffers []CommandBuffer) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	bufferCount := len(commandBuffers)
	commandBufferPtr := (*C.VkCommandBuffer)(arena.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))
	commandBufferSlice := ([]C.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, bufferCount))

	var addToDrawCount int
	var addToDispatchCount int
	for i := 0; i < bufferCount; i++ {
		commandBufferSlice[i] = C.VkCommandBuffer(unsafe.Pointer(commandBuffers[i].Handle()))
		addToDrawCount += commandBuffers[i].DrawsRecorded()
		addToDispatchCount += commandBuffers[i].DispatchesRecorded()
	}

	c.deviceDriver.VkCmdExecuteCommands(c.commandBufferHandle, driver.Uint32(bufferCount), (*driver.VkCommandBuffer)(unsafe.Pointer(commandBufferPtr)))
	c.commandCounter.CommandCount++
	c.commandCounter.DrawCallCount += addToDrawCount
	c.commandCounter.DispatchCount += addToDispatchCount
}

func (c *VulkanCommandBuffer) CmdClearAttachments(attachments []ClearAttachment, rects []ClearRect) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	attachmentCount := len(attachments)
	attachmentsPtr, err := common.AllocSlice[C.VkClearAttachment, ClearAttachment](arena, attachments)
	if err != nil {
		return err
	}

	rectsCount := len(rects)
	rectsPtr, err := common.AllocSlice[C.VkClearRect, ClearRect](arena, rects)
	if err != nil {
		return err
	}

	c.deviceDriver.VkCmdClearAttachments(c.commandBufferHandle, driver.Uint32(attachmentCount), (*driver.VkClearAttachment)(unsafe.Pointer(attachmentsPtr)), driver.Uint32(rectsCount), (*driver.VkClearRect)(unsafe.Pointer(rectsPtr)))
	c.commandCounter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdClearDepthStencilImage(image Image, imageLayout ImageLayout, depthStencil *ClearValueDepthStencil, ranges []ImageSubresourceRange) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	rangePtr := (*C.VkImageSubresourceRange)(arena.Malloc(rangeCount * C.sizeof_struct_VkImageSubresourceRange))
	rangeSlice := ([]C.VkImageSubresourceRange)(unsafe.Slice(rangePtr, rangeCount))

	for i := 0; i < rangeCount; i++ {
		rangeSlice[i].aspectMask = C.VkImageAspectFlags(ranges[i].AspectMask)
		rangeSlice[i].baseMipLevel = C.uint32_t(ranges[i].BaseMipLevel)
		rangeSlice[i].levelCount = C.uint32_t(ranges[i].LevelCount)
		rangeSlice[i].baseArrayLayer = C.uint32_t(ranges[i].BaseArrayLayer)
		rangeSlice[i].layerCount = C.uint32_t(ranges[i].LayerCount)
	}

	depthStencilPtr := (*C.VkClearDepthStencilValue)(arena.Malloc(C.sizeof_struct_VkClearDepthStencilValue))
	depthStencilPtr.depth = C.float(depthStencil.Depth)
	depthStencilPtr.stencil = C.uint32_t(depthStencil.Stencil)

	c.deviceDriver.VkCmdClearDepthStencilImage(c.commandBufferHandle, image.Handle(), driver.VkImageLayout(imageLayout), (*driver.VkClearDepthStencilValue)(unsafe.Pointer(depthStencilPtr)), driver.Uint32(rangeCount), (*driver.VkImageSubresourceRange)(unsafe.Pointer(rangePtr)))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdCopyImageToBuffer(srcImage Image, srcImageLayout ImageLayout, dstBuffer Buffer, regions []BufferImageCopy) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionPtr, err := common.AllocSlice[C.VkBufferImageCopy, BufferImageCopy](arena, regions)
	if err != nil {
		return err
	}

	c.deviceDriver.VkCmdCopyImageToBuffer(c.commandBufferHandle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstBuffer.Handle(), driver.Uint32(regionCount), (*driver.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))
	c.commandCounter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdDispatch(groupCountX, groupCountY, groupCountZ int) {
	c.deviceDriver.VkCmdDispatch(c.commandBufferHandle, driver.Uint32(groupCountX), driver.Uint32(groupCountY), driver.Uint32(groupCountZ))
	c.commandCounter.CommandCount++
	c.commandCounter.DispatchCount++
}

func (c *VulkanCommandBuffer) CmdDispatchIndirect(buffer Buffer, offset int) {
	c.deviceDriver.VkCmdDispatchIndirect(c.commandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset))
	c.commandCounter.CommandCount++
	c.commandCounter.DispatchCount++
}

func (c *VulkanCommandBuffer) CmdDrawIndexedIndirect(buffer Buffer, offset int, drawCount, stride int) {
	c.deviceDriver.VkCmdDrawIndexedIndirect(c.commandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset), driver.Uint32(drawCount), driver.Uint32(stride))
	c.commandCounter.CommandCount++
	c.commandCounter.DrawCallCount++
}

func (c *VulkanCommandBuffer) CmdDrawIndirect(buffer Buffer, offset int, drawCount, stride int) {
	c.deviceDriver.VkCmdDrawIndirect(c.commandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset), driver.Uint32(drawCount), driver.Uint32(stride))
	c.commandCounter.CommandCount++
	c.commandCounter.DrawCallCount++
}

func (c *VulkanCommandBuffer) CmdFillBuffer(dstBuffer Buffer, dstOffset int, size int, data uint32) {
	c.deviceDriver.VkCmdFillBuffer(c.commandBufferHandle, dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(size), driver.Uint32(data))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdResetEvent(event Event, stageMask PipelineStageFlags) {
	c.deviceDriver.VkCmdResetEvent(c.commandBufferHandle, event.Handle(), driver.VkPipelineStageFlags(stageMask))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdResolveImage(srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regions []ImageResolve) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionsPtr, err := common.AllocSlice[C.VkImageResolve, ImageResolve](arena, regions)
	if err != nil {
		return err
	}

	c.deviceDriver.VkCmdResolveImage(c.commandBufferHandle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstImage.Handle(), driver.VkImageLayout(dstImageLayout), driver.Uint32(regionCount), (*driver.VkImageResolve)(unsafe.Pointer(regionsPtr)))
	c.commandCounter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdSetBlendConstants(blendConstants [4]float32) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	constsPtr := (*C.float)(arena.Malloc(16))
	constsSlice := ([]C.float)(unsafe.Slice(constsPtr, 4))

	for i := 0; i < 4; i++ {
		constsSlice[i] = C.float(blendConstants[i])
	}

	c.deviceDriver.VkCmdSetBlendConstants(c.commandBufferHandle, (*driver.Float)(constsPtr))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetDepthBias(depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32) {
	c.deviceDriver.VkCmdSetDepthBias(c.commandBufferHandle, driver.Float(depthBiasConstantFactor), driver.Float(depthBiasClamp), driver.Float(depthBiasSlopeFactor))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetDepthBounds(min, max float32) {
	c.deviceDriver.VkCmdSetDepthBounds(c.commandBufferHandle, driver.Float(min), driver.Float(max))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetLineWidth(lineWidth float32) {
	c.deviceDriver.VkCmdSetLineWidth(c.commandBufferHandle, driver.Float(lineWidth))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetStencilCompareMask(faceMask StencilFaceFlags, compareMask uint32) {
	c.deviceDriver.VkCmdSetStencilCompareMask(c.commandBufferHandle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(compareMask))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetStencilReference(faceMask StencilFaceFlags, reference uint32) {
	c.deviceDriver.VkCmdSetStencilReference(c.commandBufferHandle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(reference))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetStencilWriteMask(faceMask StencilFaceFlags, writeMask uint32) {
	c.deviceDriver.VkCmdSetStencilWriteMask(c.commandBufferHandle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(writeMask))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdUpdateBuffer(dstBuffer Buffer, dstOffset int, dataSize int, data []byte) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	size := len(data)
	dataPtr := arena.Malloc(size)
	dataSlice := ([]byte)(unsafe.Slice((*byte)(dataPtr), size))
	copy(dataSlice, data)

	c.deviceDriver.VkCmdUpdateBuffer(c.commandBufferHandle, dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(dataSize), dataPtr)
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdWriteTimestamp(pipelineStage PipelineStageFlags, queryPool QueryPool, query int) {
	c.deviceDriver.VkCmdWriteTimestamp(c.commandBufferHandle, driver.VkPipelineStageFlags(pipelineStage), queryPool.Handle(), driver.Uint32(query))
	c.commandCounter.CommandCount++
}

func (c *VulkanCommandBuffer) Reset(flags CommandBufferResetFlags) (common.VkResult, error) {
	return c.deviceDriver.VkResetCommandBuffer(c.commandBufferHandle, driver.VkCommandBufferResetFlags(flags))
}

func (c *VulkanCommandBuffer) Free() {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	vkCommandBuffer := (*driver.VkCommandBuffer)(arena.Malloc(int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))
	commandBufferSlice := ([]driver.VkCommandBuffer)(unsafe.Slice(vkCommandBuffer, 1))
	commandBufferSlice[0] = c.commandBufferHandle

	c.deviceDriver.VkFreeCommandBuffers(c.device, c.commandPool, 1, vkCommandBuffer)
	c.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(c.commandBufferHandle))
}
