package internal1_0

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
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanCommandBuffer struct {
	DeviceDriver        driver.Driver
	Device              driver.VkDevice
	CommandPool         driver.VkCommandPool
	CommandBufferHandle driver.VkCommandBuffer

	MaximumAPIVersion common.APIVersion

	CommandBuffer1_1 core1_1.CommandBuffer
}

func (c *VulkanCommandBuffer) Handle() driver.VkCommandBuffer {
	return c.CommandBufferHandle
}

func (c *VulkanCommandBuffer) CommandPoolHandle() driver.VkCommandPool {
	return c.CommandPool
}

func (c *VulkanCommandBuffer) DeviceHandle() driver.VkDevice {
	return c.Device
}

func (c *VulkanCommandBuffer) Driver() driver.Driver {
	return c.DeviceDriver
}

func (c *VulkanCommandBuffer) Core1_1() core1_1.CommandBuffer {
	return c.CommandBuffer1_1
}

func (c *VulkanCommandBuffer) Begin(o core1_0.BeginOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return c.DeviceDriver.VkBeginCommandBuffer(c.CommandBufferHandle, (*driver.VkCommandBufferBeginInfo)(createInfo))
}

func (c *VulkanCommandBuffer) End() (common.VkResult, error) {
	return c.DeviceDriver.VkEndCommandBuffer(c.CommandBufferHandle)
}

func (c *VulkanCommandBuffer) CmdBeginRenderPass(contents common.SubpassContents, o core1_0.RenderPassBeginOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdBeginRenderPass(c.CommandBufferHandle, (*driver.VkRenderPassBeginInfo)(createInfo), driver.VkSubpassContents(contents))
	return nil
}

func (c *VulkanCommandBuffer) CmdEndRenderPass() {
	c.DeviceDriver.VkCmdEndRenderPass(c.CommandBufferHandle)
}

func (c *VulkanCommandBuffer) CmdBindPipeline(bindPoint common.PipelineBindPoint, pipeline core1_0.Pipeline) {
	c.DeviceDriver.VkCmdBindPipeline(c.CommandBufferHandle, driver.VkPipelineBindPoint(bindPoint), pipeline.Handle())
}

func (c *VulkanCommandBuffer) CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) {
	c.DeviceDriver.VkCmdDraw(c.CommandBufferHandle, driver.Uint32(vertexCount), driver.Uint32(instanceCount), driver.Uint32(firstVertex), driver.Uint32(firstInstance))
}

func (c *VulkanCommandBuffer) CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) {
	c.DeviceDriver.VkCmdDrawIndexed(c.CommandBufferHandle, driver.Uint32(indexCount), driver.Uint32(instanceCount), driver.Uint32(firstIndex), driver.Int32(vertexOffset), driver.Uint32(firstInstance))
}

func (c *VulkanCommandBuffer) CmdBindVertexBuffers(buffers []core1_0.Buffer, bufferOffsets []int) {
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

	c.DeviceDriver.VkCmdBindVertexBuffers(c.CommandBufferHandle, driver.Uint32(0), driver.Uint32(bufferCount), bufferArrayPtr, offsetArrayPtr)
}

func (c *VulkanCommandBuffer) CmdBindIndexBuffer(buffer core1_0.Buffer, offset int, indexType common.IndexType) {
	c.DeviceDriver.VkCmdBindIndexBuffer(c.CommandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset), driver.VkIndexType(indexType))
}

func (c *VulkanCommandBuffer) CmdBindDescriptorSets(bindPoint common.PipelineBindPoint, layout core1_0.PipelineLayout, sets []core1_0.DescriptorSet, dynamicOffsets []int) {
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

	c.DeviceDriver.VkCmdBindDescriptorSets(c.CommandBufferHandle,
		driver.VkPipelineBindPoint(bindPoint),
		layout.Handle(),
		driver.Uint32(0),
		driver.Uint32(setCount),
		(*driver.VkDescriptorSet)(setPtr),
		driver.Uint32(dynamicOffsetCount),
		(*driver.Uint32)(dynamicOffsetPtr))
}

func (c *VulkanCommandBuffer) CmdPipelineBarrier(srcStageMask, dstStageMask common.PipelineStages, dependencies common.DependencyFlags, memoryBarriers []core1_0.MemoryBarrierOptions, bufferMemoryBarriers []core1_0.BufferMemoryBarrierOptions, imageMemoryBarriers []core1_0.ImageMemoryBarrierOptions) error {
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
		barrierPtr, err = common.AllocOptionSlice[C.VkMemoryBarrier, core1_0.MemoryBarrierOptions](arena, memoryBarriers)
		if err != nil {
			return err
		}
	}

	if bufferBarrierCount > 0 {
		bufferBarrierPtr, err = common.AllocOptionSlice[C.VkBufferMemoryBarrier, core1_0.BufferMemoryBarrierOptions](arena, bufferMemoryBarriers)
		if err != nil {
			return err
		}
	}

	if imageBarrierCount > 0 {
		imageBarrierPtr, err = common.AllocOptionSlice[C.VkImageMemoryBarrier, core1_0.ImageMemoryBarrierOptions](arena, imageMemoryBarriers)
		if err != nil {
			return err
		}
	}

	c.DeviceDriver.VkCmdPipelineBarrier(c.CommandBufferHandle, driver.VkPipelineStageFlags(srcStageMask), driver.VkPipelineStageFlags(dstStageMask), driver.VkDependencyFlags(dependencies), driver.Uint32(barrierCount), (*driver.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), driver.Uint32(bufferBarrierCount), (*driver.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), driver.Uint32(imageBarrierCount), (*driver.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	return nil
}

func (c *VulkanCommandBuffer) CmdCopyBufferToImage(buffer core1_0.Buffer, image core1_0.Image, layout common.ImageLayout, regions []core1_0.BufferImageCopy) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	var err error
	var regionPtr *C.VkBufferImageCopy

	if regionCount > 0 {
		regionPtr, err = common.AllocSlice[C.VkBufferImageCopy, core1_0.BufferImageCopy](arena, regions)
		if err != nil {
			return err
		}
	}

	c.DeviceDriver.VkCmdCopyBufferToImage(c.CommandBufferHandle, buffer.Handle(), image.Handle(), driver.VkImageLayout(layout), driver.Uint32(regionCount), (*driver.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))
	return nil
}

func (c *VulkanCommandBuffer) CmdBlitImage(sourceImage core1_0.Image, sourceImageLayout common.ImageLayout, destinationImage core1_0.Image, destinationImageLayout common.ImageLayout, regions []core1_0.ImageBlit, filter common.Filter) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	regionCount := len(regions)

	regionPtr, err := common.AllocSlice[C.VkImageBlit, core1_0.ImageBlit](allocator, regions)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdBlitImage(
		c.CommandBufferHandle,
		driver.VkImage(sourceImage.Handle()),
		driver.VkImageLayout(sourceImageLayout),
		driver.VkImage(destinationImage.Handle()),
		driver.VkImageLayout(destinationImageLayout),
		driver.Uint32(regionCount),
		(*driver.VkImageBlit)(unsafe.Pointer(regionPtr)),
		driver.VkFilter(filter))
	return nil
}

func (c *VulkanCommandBuffer) CmdPushConstants(layout core1_0.PipelineLayout, stageFlags common.ShaderStages, offset int, valueBytes []byte) {
	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	valueBytesPtr := alloc.CBytes(valueBytes)

	c.DeviceDriver.VkCmdPushConstants(c.CommandBufferHandle, layout.Handle(), driver.VkShaderStageFlags(stageFlags), driver.Uint32(offset), driver.Uint32(len(valueBytes)), valueBytesPtr)
}

func (c *VulkanCommandBuffer) CmdSetViewport(viewports []common.Viewport) {
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

	c.DeviceDriver.VkCmdSetViewport(c.CommandBufferHandle, driver.Uint32(0), driver.Uint32(viewportCount), (*driver.VkViewport)(unsafe.Pointer(viewportPtr)))
}

func (c *VulkanCommandBuffer) CmdSetScissor(scissors []common.Rect2D) {
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

	c.DeviceDriver.VkCmdSetScissor(c.CommandBufferHandle, driver.Uint32(0), driver.Uint32(scissorCount), (*driver.VkRect2D)(unsafe.Pointer(scissorPtr)))
}

func (c *VulkanCommandBuffer) CmdNextSubpass(contents common.SubpassContents) {
	c.DeviceDriver.VkCmdNextSubpass(c.CommandBufferHandle, driver.VkSubpassContents(contents))
}

func (c *VulkanCommandBuffer) CmdWaitEvents(events []core1_0.Event, srcStageMask common.PipelineStages, dstStageMask common.PipelineStages, memoryBarriers []core1_0.MemoryBarrierOptions, bufferMemoryBarriers []core1_0.BufferMemoryBarrierOptions, imageMemoryBarriers []core1_0.ImageMemoryBarrierOptions) error {
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
		barrierPtr, err = common.AllocOptionSlice[C.VkMemoryBarrier, core1_0.MemoryBarrierOptions](arena, memoryBarriers)
		if err != nil {
			return err
		}
	}

	if bufferBarrierCount > 0 {
		bufferBarrierPtr, err = common.AllocOptionSlice[C.VkBufferMemoryBarrier, core1_0.BufferMemoryBarrierOptions](arena, bufferMemoryBarriers)
		if err != nil {
			return err
		}
	}

	if imageBarrierCount > 0 {
		imageBarrierPtr, err = common.AllocOptionSlice[C.VkImageMemoryBarrier, core1_0.ImageMemoryBarrierOptions](arena, imageMemoryBarriers)
		if err != nil {
			return err
		}
	}

	c.DeviceDriver.VkCmdWaitEvents(c.CommandBufferHandle, driver.Uint32(eventCount), (*driver.VkEvent)(unsafe.Pointer(eventPtr)), driver.VkPipelineStageFlags(srcStageMask), driver.VkPipelineStageFlags(dstStageMask), driver.Uint32(barrierCount), (*driver.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), driver.Uint32(bufferBarrierCount), (*driver.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), driver.Uint32(imageBarrierCount), (*driver.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	return nil
}

func (c *VulkanCommandBuffer) CmdSetEvent(event core1_0.Event, stageMask common.PipelineStages) {
	c.DeviceDriver.VkCmdSetEvent(c.CommandBufferHandle, event.Handle(), driver.VkPipelineStageFlags(stageMask))
}

func (c *VulkanCommandBuffer) CmdClearColorImage(image core1_0.Image, imageLayout common.ImageLayout, color common.ClearColorValue, ranges []common.ImageSubresourceRange) {
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

	c.DeviceDriver.VkCmdClearColorImage(c.CommandBufferHandle, image.Handle(), driver.VkImageLayout(imageLayout), (*driver.VkClearColorValue)(pColor), driver.Uint32(rangeCount), (*driver.VkImageSubresourceRange)(unsafe.Pointer(pRanges)))
}

func (c *VulkanCommandBuffer) CmdResetQueryPool(queryPool core1_0.QueryPool, startQuery, queryCount int) {
	c.DeviceDriver.VkCmdResetQueryPool(c.CommandBufferHandle, queryPool.Handle(), driver.Uint32(startQuery), driver.Uint32(queryCount))
}

func (c *VulkanCommandBuffer) CmdBeginQuery(queryPool core1_0.QueryPool, query int, flags common.QueryControlFlags) {
	c.DeviceDriver.VkCmdBeginQuery(c.CommandBufferHandle, queryPool.Handle(), driver.Uint32(query), driver.VkQueryControlFlags(flags))
}

func (c *VulkanCommandBuffer) CmdEndQuery(queryPool core1_0.QueryPool, query int) {
	c.DeviceDriver.VkCmdEndQuery(c.CommandBufferHandle, queryPool.Handle(), driver.Uint32(query))
}

func (c *VulkanCommandBuffer) CmdCopyQueryPoolResults(queryPool core1_0.QueryPool, firstQuery, queryCount int, dstBuffer core1_0.Buffer, dstOffset, stride int, flags common.QueryResultFlags) {
	c.DeviceDriver.VkCmdCopyQueryPoolResults(c.CommandBufferHandle, queryPool.Handle(), driver.Uint32(firstQuery), driver.Uint32(queryCount), dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(stride), driver.VkQueryResultFlags(flags))
}

func (c *VulkanCommandBuffer) CmdExecuteCommands(commandBuffers []core1_0.CommandBuffer) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	bufferCount := len(commandBuffers)
	commandBufferPtr := (*C.VkCommandBuffer)(arena.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))
	commandBufferSlice := ([]C.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		commandBufferSlice[i] = C.VkCommandBuffer(unsafe.Pointer(commandBuffers[i].Handle()))
	}

	c.DeviceDriver.VkCmdExecuteCommands(c.CommandBufferHandle, driver.Uint32(bufferCount), (*driver.VkCommandBuffer)(unsafe.Pointer(commandBufferPtr)))
}

func (c *VulkanCommandBuffer) CmdClearAttachments(attachments []core1_0.ClearAttachment, rects []core1_0.ClearRect) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	attachmentCount := len(attachments)
	attachmentsPtr, err := common.AllocSlice[C.VkClearAttachment, core1_0.ClearAttachment](arena, attachments)
	if err != nil {
		return err
	}

	rectsCount := len(rects)
	rectsPtr, err := common.AllocSlice[C.VkClearRect, core1_0.ClearRect](arena, rects)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdClearAttachments(c.CommandBufferHandle, driver.Uint32(attachmentCount), (*driver.VkClearAttachment)(unsafe.Pointer(attachmentsPtr)), driver.Uint32(rectsCount), (*driver.VkClearRect)(unsafe.Pointer(rectsPtr)))
	return nil
}

func (c *VulkanCommandBuffer) CmdClearDepthStencilImage(image core1_0.Image, imageLayout common.ImageLayout, depthStencil *common.ClearValueDepthStencil, ranges []common.ImageSubresourceRange) {
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

	c.DeviceDriver.VkCmdClearDepthStencilImage(c.CommandBufferHandle, image.Handle(), driver.VkImageLayout(imageLayout), (*driver.VkClearDepthStencilValue)(unsafe.Pointer(depthStencilPtr)), driver.Uint32(rangeCount), (*driver.VkImageSubresourceRange)(unsafe.Pointer(rangePtr)))
}

func (c *VulkanCommandBuffer) CmdCopyImageToBuffer(srcImage core1_0.Image, srcImageLayout common.ImageLayout, dstBuffer core1_0.Buffer, regions []core1_0.BufferImageCopy) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionPtr, err := common.AllocSlice[C.VkBufferImageCopy, core1_0.BufferImageCopy](arena, regions)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdCopyImageToBuffer(c.CommandBufferHandle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstBuffer.Handle(), driver.Uint32(regionCount), (*driver.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))

	return nil
}

func (c *VulkanCommandBuffer) CmdDispatch(groupCountX, groupCountY, groupCountZ int) {
	c.DeviceDriver.VkCmdDispatch(c.CommandBufferHandle, driver.Uint32(groupCountX), driver.Uint32(groupCountY), driver.Uint32(groupCountZ))
}

func (c *VulkanCommandBuffer) CmdDispatchIndirect(buffer core1_0.Buffer, offset int) {
	c.DeviceDriver.VkCmdDispatchIndirect(c.CommandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset))
}

func (c *VulkanCommandBuffer) CmdDrawIndexedIndirect(buffer core1_0.Buffer, offset int, drawCount, stride int) {
	c.DeviceDriver.VkCmdDrawIndexedIndirect(c.CommandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset), driver.Uint32(drawCount), driver.Uint32(stride))
}

func (c *VulkanCommandBuffer) CmdDrawIndirect(buffer core1_0.Buffer, offset int, drawCount, stride int) {
	c.DeviceDriver.VkCmdDrawIndirect(c.CommandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset), driver.Uint32(drawCount), driver.Uint32(stride))
}

func (c *VulkanCommandBuffer) CmdFillBuffer(dstBuffer core1_0.Buffer, dstOffset int, size int, data uint32) {
	c.DeviceDriver.VkCmdFillBuffer(c.CommandBufferHandle, dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(size), driver.Uint32(data))
}

func (c *VulkanCommandBuffer) CmdResetEvent(event core1_0.Event, stageMask common.PipelineStages) {
	c.DeviceDriver.VkCmdResetEvent(c.CommandBufferHandle, event.Handle(), driver.VkPipelineStageFlags(stageMask))
}

func (c *VulkanCommandBuffer) CmdResolveImage(srcImage core1_0.Image, srcImageLayout common.ImageLayout, dstImage core1_0.Image, dstImageLayout common.ImageLayout, regions []core1_0.ImageResolve) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionsPtr, err := common.AllocSlice[C.VkImageResolve, core1_0.ImageResolve](arena, regions)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdResolveImage(c.CommandBufferHandle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstImage.Handle(), driver.VkImageLayout(dstImageLayout), driver.Uint32(regionCount), (*driver.VkImageResolve)(unsafe.Pointer(regionsPtr)))
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

	c.DeviceDriver.VkCmdSetBlendConstants(c.CommandBufferHandle, (*driver.Float)(constsPtr))
}

func (c *VulkanCommandBuffer) CmdSetDepthBias(depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32) {
	c.DeviceDriver.VkCmdSetDepthBias(c.CommandBufferHandle, driver.Float(depthBiasConstantFactor), driver.Float(depthBiasClamp), driver.Float(depthBiasSlopeFactor))
}

func (c *VulkanCommandBuffer) CmdSetDepthBounds(min, max float32) {
	c.DeviceDriver.VkCmdSetDepthBounds(c.CommandBufferHandle, driver.Float(min), driver.Float(max))
}

func (c *VulkanCommandBuffer) CmdSetLineWidth(lineWidth float32) {
	c.DeviceDriver.VkCmdSetLineWidth(c.CommandBufferHandle, driver.Float(lineWidth))
}

func (c *VulkanCommandBuffer) CmdSetStencilCompareMask(faceMask common.StencilFaces, compareMask uint32) {
	c.DeviceDriver.VkCmdSetStencilCompareMask(c.CommandBufferHandle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(compareMask))
}

func (c *VulkanCommandBuffer) CmdSetStencilReference(faceMask common.StencilFaces, reference uint32) {
	c.DeviceDriver.VkCmdSetStencilReference(c.CommandBufferHandle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(reference))
}

func (c *VulkanCommandBuffer) CmdSetStencilWriteMask(faceMask common.StencilFaces, writeMask uint32) {
	c.DeviceDriver.VkCmdSetStencilWriteMask(c.CommandBufferHandle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(writeMask))
}

func (c *VulkanCommandBuffer) CmdUpdateBuffer(dstBuffer core1_0.Buffer, dstOffset int, dataSize int, data []byte) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	size := len(data)
	dataPtr := arena.Malloc(size)
	dataSlice := ([]byte)(unsafe.Slice((*byte)(dataPtr), size))
	copy(dataSlice, data)

	c.DeviceDriver.VkCmdUpdateBuffer(c.CommandBufferHandle, dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(dataSize), dataPtr)
}

func (c *VulkanCommandBuffer) CmdWriteTimestamp(pipelineStage common.PipelineStages, queryPool core1_0.QueryPool, query int) {
	c.DeviceDriver.VkCmdWriteTimestamp(c.CommandBufferHandle, driver.VkPipelineStageFlags(pipelineStage), queryPool.Handle(), driver.Uint32(query))
}

func (c *VulkanCommandBuffer) Reset(flags common.CommandBufferResetFlags) (common.VkResult, error) {
	return c.DeviceDriver.VkResetCommandBuffer(c.CommandBufferHandle, driver.VkCommandBufferResetFlags(flags))
}

func (c *VulkanCommandBuffer) Free() {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	vkCommandBuffer := (*driver.VkCommandBuffer)(arena.Malloc(int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))
	commandBufferSlice := ([]driver.VkCommandBuffer)(unsafe.Slice(vkCommandBuffer, 1))
	commandBufferSlice[0] = c.CommandBufferHandle

	c.DeviceDriver.VkFreeCommandBuffers(c.Device, c.CommandPool, 1, vkCommandBuffer)
	c.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(c.CommandBufferHandle), c)
}
