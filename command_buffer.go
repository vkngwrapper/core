package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type vulkanCommandBuffer struct {
	driver driver.Driver
	device driver.VkDevice
	pool   driver.VkCommandPool
	handle driver.VkCommandBuffer
}

func (c *vulkanCommandBuffer) Handle() driver.VkCommandBuffer {
	return c.handle
}

func (c *vulkanCommandBuffer) Begin(o *BeginOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return common.VKErrorUnknown, err
	}

	return c.driver.VkBeginCommandBuffer(c.handle, (*driver.VkCommandBufferBeginInfo)(createInfo))
}

func (c *vulkanCommandBuffer) End() (common.VkResult, error) {
	return c.driver.VkEndCommandBuffer(c.handle)
}

func (c *vulkanCommandBuffer) CmdBeginRenderPass(contents SubpassContents, o *RenderPassBeginOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	c.driver.VkCmdBeginRenderPass(c.handle, (*driver.VkRenderPassBeginInfo)(createInfo), driver.VkSubpassContents(contents))
	return nil
}

func (c *vulkanCommandBuffer) CmdEndRenderPass() {
	c.driver.VkCmdEndRenderPass(c.handle)
}

func (c *vulkanCommandBuffer) CmdBindPipeline(bindPoint common.PipelineBindPoint, pipeline Pipeline) {
	c.driver.VkCmdBindPipeline(c.handle, driver.VkPipelineBindPoint(bindPoint), pipeline.Handle())
}

func (c *vulkanCommandBuffer) CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) {
	c.driver.VkCmdDraw(c.handle, driver.Uint32(vertexCount), driver.Uint32(instanceCount), driver.Uint32(firstVertex), driver.Uint32(firstInstance))
}

func (c *vulkanCommandBuffer) CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) {
	c.driver.VkCmdDrawIndexed(c.handle, driver.Uint32(indexCount), driver.Uint32(instanceCount), driver.Uint32(firstIndex), driver.Int32(vertexOffset), driver.Uint32(firstInstance))
}

func (c *vulkanCommandBuffer) CmdBindVertexBuffers(buffers []Buffer, bufferOffsets []int) {
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

	c.driver.VkCmdBindVertexBuffers(c.handle, driver.Uint32(0), driver.Uint32(bufferCount), bufferArrayPtr, offsetArrayPtr)
}

func (c *vulkanCommandBuffer) CmdBindIndexBuffer(buffer Buffer, offset int, indexType common.IndexType) {
	c.driver.VkCmdBindIndexBuffer(c.handle, buffer.Handle(), driver.VkDeviceSize(offset), driver.VkIndexType(indexType))
}

func (c *vulkanCommandBuffer) CmdBindDescriptorSets(bindPoint common.PipelineBindPoint, layout PipelineLayout, sets []DescriptorSet, dynamicOffsets []int) {
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

	c.driver.VkCmdBindDescriptorSets(c.handle,
		driver.VkPipelineBindPoint(bindPoint),
		layout.Handle(),
		driver.Uint32(0),
		driver.Uint32(setCount),
		(*driver.VkDescriptorSet)(setPtr),
		driver.Uint32(dynamicOffsetCount),
		(*driver.Uint32)(dynamicOffsetPtr))
}

func (c *vulkanCommandBuffer) CmdPipelineBarrier(srcStageMask, dstStageMask common.PipelineStages, dependencies common.DependencyFlags, memoryBarriers []*MemoryBarrierOptions, bufferMemoryBarriers []*BufferMemoryBarrierOptions, imageMemoryBarriers []*ImageMemoryBarrierOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	barrierCount := len(memoryBarriers)
	bufferBarrierCount := len(bufferMemoryBarriers)
	imageBarrierCount := len(imageMemoryBarriers)

	var barrierPtr *C.VkMemoryBarrier
	var bufferBarrierPtr *C.VkBufferMemoryBarrier
	var imageBarrierPtr *C.VkImageMemoryBarrier

	if barrierCount > 0 {
		barrierPtr = (*C.VkMemoryBarrier)(arena.Malloc(barrierCount * C.sizeof_struct_VkMemoryBarrier))
		barrierSlice := ([]C.VkMemoryBarrier)(unsafe.Slice(barrierPtr, barrierCount))

		for i := 0; i < barrierCount; i++ {
			next, err := common.AllocNext(arena, memoryBarriers[i])
			if err != nil {
				return err
			}

			err = memoryBarriers[i].populate(&barrierSlice[i], next)
			if err != nil {
				return err
			}
		}
	}

	if bufferBarrierCount > 0 {
		bufferBarrierPtr = (*C.VkBufferMemoryBarrier)(arena.Malloc(bufferBarrierCount * C.sizeof_struct_VkBufferMemoryBarrier))
		bufferBarrierSlice := ([]C.VkBufferMemoryBarrier)(unsafe.Slice(bufferBarrierPtr, bufferBarrierCount))

		for i := 0; i < bufferBarrierCount; i++ {
			next, err := common.AllocNext(arena, bufferMemoryBarriers[i])
			if err != nil {
				return err
			}

			err = bufferMemoryBarriers[i].populate(&bufferBarrierSlice[i], next)
			if err != nil {
				return err
			}
		}
	}

	if imageBarrierCount > 0 {
		imageBarrierPtr = (*C.VkImageMemoryBarrier)(arena.Malloc(imageBarrierCount * C.sizeof_struct_VkImageMemoryBarrier))
		imageBarrierSlice := ([]C.VkImageMemoryBarrier)(unsafe.Slice(imageBarrierPtr, imageBarrierCount))

		for i := 0; i < imageBarrierCount; i++ {
			next, err := common.AllocNext(arena, imageMemoryBarriers[i])
			if err != nil {
				return err
			}

			err = imageMemoryBarriers[i].populate(&imageBarrierSlice[i], next)
			if err != nil {
				return err
			}
		}
	}

	c.driver.VkCmdPipelineBarrier(c.handle, driver.VkPipelineStageFlags(srcStageMask), driver.VkPipelineStageFlags(dstStageMask), driver.VkDependencyFlags(dependencies), driver.Uint32(barrierCount), (*driver.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), driver.Uint32(bufferBarrierCount), (*driver.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), driver.Uint32(imageBarrierCount), (*driver.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	return nil
}

func (c *vulkanCommandBuffer) CmdCopyBufferToImage(buffer Buffer, image Image, layout common.ImageLayout, regions []*BufferImageCopy) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	var regionPtr *C.VkBufferImageCopy

	if regionCount > 0 {
		regionPtr = (*C.VkBufferImageCopy)(arena.Malloc(regionCount * C.sizeof_struct_VkBufferImageCopy))
		regionSlice := ([]C.VkBufferImageCopy)(unsafe.Slice(regionPtr, regionCount))

		for i := 0; i < regionCount; i++ {
			err := regions[i].populate(&regionSlice[i])
			if err != nil {
				return err
			}
		}
	}

	c.driver.VkCmdCopyBufferToImage(c.handle, buffer.Handle(), image.Handle(), driver.VkImageLayout(layout), driver.Uint32(regionCount), (*driver.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))
	return nil
}

func (c *vulkanCommandBuffer) CmdBlitImage(sourceImage Image, sourceImageLayout common.ImageLayout, destinationImage Image, destinationImageLayout common.ImageLayout, regions []*ImageBlit, filter common.Filter) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	regionCount := len(regions)
	regionPtr := (*C.VkImageBlit)(allocator.Malloc(regionCount * C.sizeof_struct_VkImageBlit))
	regionSlice := ([]C.VkImageBlit)(unsafe.Slice(regionPtr, regionCount))

	for i := range regionSlice {
		err := regions[i].Populate(&regionSlice[i])
		if err != nil {
			return err
		}
	}

	c.driver.VkCmdBlitImage(
		c.handle,
		driver.VkImage(sourceImage.Handle()),
		driver.VkImageLayout(sourceImageLayout),
		driver.VkImage(destinationImage.Handle()),
		driver.VkImageLayout(destinationImageLayout),
		driver.Uint32(regionCount),
		(*driver.VkImageBlit)(unsafe.Pointer(regionPtr)),
		driver.VkFilter(filter))
	return nil
}

func (c *vulkanCommandBuffer) CmdPushConstants(layout PipelineLayout, stageFlags common.ShaderStages, offset int, valueBytes []byte) {
	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	valueBytesPtr := alloc.CBytes(valueBytes)

	c.driver.VkCmdPushConstants(c.handle, layout.Handle(), driver.VkShaderStageFlags(stageFlags), driver.Uint32(offset), driver.Uint32(len(valueBytes)), valueBytesPtr)
}

func (c *vulkanCommandBuffer) CmdSetViewport(viewports []common.Viewport) {
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

	c.driver.VkCmdSetViewport(c.handle, driver.Uint32(0), driver.Uint32(viewportCount), (*driver.VkViewport)(unsafe.Pointer(viewportPtr)))
}

func (c *vulkanCommandBuffer) CmdSetScissor(scissors []common.Rect2D) {
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

	c.driver.VkCmdSetScissor(c.handle, driver.Uint32(0), driver.Uint32(scissorCount), (*driver.VkRect2D)(unsafe.Pointer(scissorPtr)))
}

func (c *vulkanCommandBuffer) CmdNextSubpass(contents SubpassContents) {
	c.driver.VkCmdNextSubpass(c.handle, driver.VkSubpassContents(contents))
}

func (c *vulkanCommandBuffer) CmdWaitEvents(events []Event, srcStageMask common.PipelineStages, dstStageMask common.PipelineStages, memoryBarriers []*MemoryBarrierOptions, bufferMemoryBarriers []*BufferMemoryBarrierOptions, imageMemoryBarriers []*ImageMemoryBarrierOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	eventCount := len(events)
	barrierCount := len(memoryBarriers)
	bufferBarrierCount := len(bufferMemoryBarriers)
	imageBarrierCount := len(imageMemoryBarriers)

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
		barrierPtr = (*C.VkMemoryBarrier)(arena.Malloc(barrierCount * C.sizeof_struct_VkMemoryBarrier))
		barrierSlice := ([]C.VkMemoryBarrier)(unsafe.Slice(barrierPtr, barrierCount))

		for i := 0; i < barrierCount; i++ {
			next, err := common.AllocNext(arena, memoryBarriers[i])
			if err != nil {
				return err
			}

			err = memoryBarriers[i].populate(&barrierSlice[i], next)
			if err != nil {
				return err
			}
		}
	}

	if bufferBarrierCount > 0 {
		bufferBarrierPtr = (*C.VkBufferMemoryBarrier)(arena.Malloc(bufferBarrierCount * C.sizeof_struct_VkBufferMemoryBarrier))
		bufferBarrierSlice := ([]C.VkBufferMemoryBarrier)(unsafe.Slice(bufferBarrierPtr, bufferBarrierCount))

		for i := 0; i < bufferBarrierCount; i++ {
			next, err := common.AllocNext(arena, bufferMemoryBarriers[i])
			if err != nil {
				return err
			}

			err = bufferMemoryBarriers[i].populate(&bufferBarrierSlice[i], next)
			if err != nil {
				return err
			}
		}
	}

	if imageBarrierCount > 0 {
		imageBarrierPtr = (*C.VkImageMemoryBarrier)(arena.Malloc(imageBarrierCount * C.sizeof_struct_VkImageMemoryBarrier))
		imageBarrierSlice := ([]C.VkImageMemoryBarrier)(unsafe.Slice(imageBarrierPtr, imageBarrierCount))

		for i := 0; i < imageBarrierCount; i++ {
			next, err := common.AllocNext(arena, imageMemoryBarriers[i])
			if err != nil {
				return err
			}

			err = imageMemoryBarriers[i].populate(&imageBarrierSlice[i], next)
			if err != nil {
				return err
			}
		}
	}

	c.driver.VkCmdWaitEvents(c.handle, driver.Uint32(eventCount), (*driver.VkEvent)(unsafe.Pointer(eventPtr)), driver.VkPipelineStageFlags(srcStageMask), driver.VkPipelineStageFlags(dstStageMask), driver.Uint32(barrierCount), (*driver.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), driver.Uint32(bufferBarrierCount), (*driver.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), driver.Uint32(imageBarrierCount), (*driver.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	return nil
}

func (c *vulkanCommandBuffer) CmdSetEvent(event Event, stageMask common.PipelineStages) {
	c.driver.VkCmdSetEvent(c.handle, event.Handle(), driver.VkPipelineStageFlags(stageMask))
}

func (c *vulkanCommandBuffer) CmdClearColorImage(image Image, imageLayout common.ImageLayout, color ClearColorValue, ranges []common.ImageSubresourceRange) {
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

	var pColor *C.VkClearColorValue
	if color != nil {
		pColor = (*C.VkClearColorValue)(arena.Malloc(C.sizeof_union_VkClearColorValue))
		color.populateColorUnion(pColor)
	}

	c.driver.VkCmdClearColorImage(c.handle, image.Handle(), driver.VkImageLayout(imageLayout), (*driver.VkClearColorValue)(unsafe.Pointer(pColor)), driver.Uint32(rangeCount), (*driver.VkImageSubresourceRange)(unsafe.Pointer(pRanges)))
}

func (c *vulkanCommandBuffer) CmdResetQueryPool(queryPool QueryPool, startQuery, queryCount int) {
	c.driver.VkCmdResetQueryPool(c.handle, queryPool.Handle(), driver.Uint32(startQuery), driver.Uint32(queryCount))
}

func (c *vulkanCommandBuffer) CmdBeginQuery(queryPool QueryPool, query int, flags common.QueryControlFlags) {
	c.driver.VkCmdBeginQuery(c.handle, queryPool.Handle(), driver.Uint32(query), driver.VkQueryControlFlags(flags))
}

func (c *vulkanCommandBuffer) CmdEndQuery(queryPool QueryPool, query int) {
	c.driver.VkCmdEndQuery(c.handle, queryPool.Handle(), driver.Uint32(query))
}

func (c *vulkanCommandBuffer) CmdCopyQueryPoolResults(queryPool QueryPool, firstQuery, queryCount int, dstBuffer Buffer, dstOffset, stride int, flags common.QueryResultFlags) {
	c.driver.VkCmdCopyQueryPoolResults(c.handle, queryPool.Handle(), driver.Uint32(firstQuery), driver.Uint32(queryCount), dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(stride), driver.VkQueryResultFlags(flags))
}

func (c *vulkanCommandBuffer) CmdExecuteCommands(commandBuffers []CommandBuffer) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	bufferCount := len(commandBuffers)
	commandBufferPtr := (*C.VkCommandBuffer)(arena.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))
	commandBufferSlice := ([]C.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		commandBufferSlice[i] = C.VkCommandBuffer(unsafe.Pointer(commandBuffers[i].Handle()))
	}

	c.driver.VkCmdExecuteCommands(c.handle, driver.Uint32(bufferCount), (*driver.VkCommandBuffer)(unsafe.Pointer(commandBufferPtr)))
}

type ClearAttachment struct {
	AspectMask      common.ImageAspectFlags
	ColorAttachment int
	ClearValue      ClearValue
}

type ClearRect struct {
	Rect           common.Rect2D
	BaseArrayLayer int
	LayerCount     int
}

func (c *vulkanCommandBuffer) CmdClearAttachments(attachments []ClearAttachment, rects []ClearRect) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	attachmentCount := len(attachments)
	attachmentsPtr := (*C.VkClearAttachment)(arena.Malloc(attachmentCount * C.sizeof_struct_VkClearAttachment))
	attachmentsSlice := ([]C.VkClearAttachment)(unsafe.Slice(attachmentsPtr, attachmentCount))

	for i := 0; i < attachmentCount; i++ {
		attachmentsSlice[i].aspectMask = C.VkImageAspectFlags(attachments[i].AspectMask)
		attachmentsSlice[i].colorAttachment = C.uint32_t(attachments[i].ColorAttachment)
		attachments[i].ClearValue.populateValueUnion(&attachmentsSlice[i].clearValue)
	}

	rectsCount := len(rects)
	rectsPtr := (*C.VkClearRect)(arena.Malloc(rectsCount * C.sizeof_struct_VkClearRect))
	rectsSlice := ([]C.VkClearRect)(unsafe.Slice(rectsPtr, rectsCount))

	for i := 0; i < rectsCount; i++ {
		rectsSlice[i].baseArrayLayer = C.uint32_t(rects[i].BaseArrayLayer)
		rectsSlice[i].layerCount = C.uint32_t(rects[i].LayerCount)
		rectsSlice[i].rect.extent.width = C.uint32_t(rects[i].Rect.Extent.Width)
		rectsSlice[i].rect.extent.height = C.uint32_t(rects[i].Rect.Extent.Height)
		rectsSlice[i].rect.offset.x = C.int32_t(rects[i].Rect.Offset.X)
		rectsSlice[i].rect.offset.y = C.int32_t(rects[i].Rect.Offset.Y)
	}

	c.driver.VkCmdClearAttachments(c.handle, driver.Uint32(attachmentCount), (*driver.VkClearAttachment)(unsafe.Pointer(attachmentsPtr)), driver.Uint32(rectsCount), (*driver.VkClearRect)(unsafe.Pointer(rectsPtr)))
}

func (c *vulkanCommandBuffer) CmdClearDepthStencilImage(image Image, imageLayout common.ImageLayout, depthStencil *ClearValueDepthStencil, ranges []common.ImageSubresourceRange) {
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

	c.driver.VkCmdClearDepthStencilImage(c.handle, image.Handle(), driver.VkImageLayout(imageLayout), (*driver.VkClearDepthStencilValue)(unsafe.Pointer(depthStencilPtr)), driver.Uint32(rangeCount), (*driver.VkImageSubresourceRange)(unsafe.Pointer(rangePtr)))
}

func (c *vulkanCommandBuffer) CmdCopyImageToBuffer(srcImage Image, srcImageLayout common.ImageLayout, dstBuffer Buffer, regions []BufferImageCopy) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionPtr := (*C.VkBufferImageCopy)(arena.Malloc(regionCount * C.sizeof_struct_VkBufferImageCopy))
	regionSlice := ([]C.VkBufferImageCopy)(unsafe.Slice(regionPtr, regionCount))

	for i := 0; i < regionCount; i++ {
		err := regions[i].populate(&regionSlice[i])
		if err != nil {
			return err
		}
	}

	c.driver.VkCmdCopyImageToBuffer(c.handle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstBuffer.Handle(), driver.Uint32(regionCount), (*driver.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))

	return nil
}

func (c *vulkanCommandBuffer) CmdDispatch(groupCountX, groupCountY, groupCountZ int) {
	c.driver.VkCmdDispatch(c.handle, driver.Uint32(groupCountX), driver.Uint32(groupCountY), driver.Uint32(groupCountZ))
}

func (c *vulkanCommandBuffer) CmdDispatchIndirect(buffer Buffer, offset int) {
	c.driver.VkCmdDispatchIndirect(c.handle, buffer.Handle(), driver.VkDeviceSize(offset))
}

func (c *vulkanCommandBuffer) CmdDrawIndexedIndirect(buffer Buffer, offset int, drawCount, stride int) {
	c.driver.VkCmdDrawIndexedIndirect(c.handle, buffer.Handle(), driver.VkDeviceSize(offset), driver.Uint32(drawCount), driver.Uint32(stride))
}

func (c *vulkanCommandBuffer) CmdDrawIndirect(buffer Buffer, offset int, drawCount, stride int) {
	c.driver.VkCmdDrawIndirect(c.handle, buffer.Handle(), driver.VkDeviceSize(offset), driver.Uint32(drawCount), driver.Uint32(stride))
}

func (c *vulkanCommandBuffer) CmdFillBuffer(dstBuffer Buffer, dstOffset int, size int, data uint32) {
	c.driver.VkCmdFillBuffer(c.handle, dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(size), driver.Uint32(data))
}

func (c *vulkanCommandBuffer) CmdResetEvent(event Event, stageMask common.PipelineStages) {
	c.driver.VkCmdResetEvent(c.handle, event.Handle(), driver.VkPipelineStageFlags(stageMask))
}

type ImageResolve struct {
	SrcSubresource common.ImageSubresourceLayers
	SrcOffset      common.Offset3D
	DstSubresource common.ImageSubresourceLayers
	DstOffset      common.Offset3D
	Extent         common.Extent3D
}

func (c *vulkanCommandBuffer) CmdResolveImage(srcImage Image, srcImageLayout common.ImageLayout, dstImage Image, dstImageLayout common.ImageLayout, regions []ImageResolve) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionsPtr := (*C.VkImageResolve)(arena.Malloc(regionCount * C.sizeof_struct_VkImageResolve))
	regionSlice := ([]C.VkImageResolve)(unsafe.Slice(regionsPtr, regionCount))

	for i := 0; i < regionCount; i++ {
		regionSlice[i].srcSubresource.aspectMask = C.VkImageAspectFlags(regions[i].SrcSubresource.AspectMask)
		regionSlice[i].srcSubresource.mipLevel = C.uint32_t(regions[i].SrcSubresource.MipLevel)
		regionSlice[i].srcSubresource.baseArrayLayer = C.uint32_t(regions[i].SrcSubresource.BaseArrayLayer)
		regionSlice[i].srcSubresource.layerCount = C.uint32_t(regions[i].SrcSubresource.LayerCount)

		regionSlice[i].srcOffset.x = C.int32_t(regions[i].SrcOffset.X)
		regionSlice[i].srcOffset.y = C.int32_t(regions[i].SrcOffset.Y)
		regionSlice[i].srcOffset.z = C.int32_t(regions[i].SrcOffset.Z)

		regionSlice[i].dstSubresource.aspectMask = C.VkImageAspectFlags(regions[i].DstSubresource.AspectMask)
		regionSlice[i].dstSubresource.mipLevel = C.uint32_t(regions[i].DstSubresource.MipLevel)
		regionSlice[i].dstSubresource.baseArrayLayer = C.uint32_t(regions[i].DstSubresource.BaseArrayLayer)
		regionSlice[i].dstSubresource.layerCount = C.uint32_t(regions[i].DstSubresource.LayerCount)

		regionSlice[i].dstOffset.x = C.int32_t(regions[i].DstOffset.X)
		regionSlice[i].dstOffset.y = C.int32_t(regions[i].DstOffset.Y)
		regionSlice[i].dstOffset.z = C.int32_t(regions[i].DstOffset.Z)

		regionSlice[i].extent.width = C.uint32_t(regions[i].Extent.Width)
		regionSlice[i].extent.height = C.uint32_t(regions[i].Extent.Height)
		regionSlice[i].extent.depth = C.uint32_t(regions[i].Extent.Depth)
	}

	c.driver.VkCmdResolveImage(c.handle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstImage.Handle(), driver.VkImageLayout(dstImageLayout), driver.Uint32(regionCount), (*driver.VkImageResolve)(unsafe.Pointer(regionsPtr)))
}

func (c *vulkanCommandBuffer) CmdSetBlendConstants(blendConstants [4]float32) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	constsPtr := (*C.float)(arena.Malloc(16))
	constsSlice := ([]C.float)(unsafe.Slice(constsPtr, 4))

	for i := 0; i < 4; i++ {
		constsSlice[i] = C.float(blendConstants[i])
	}

	c.driver.VkCmdSetBlendConstants(c.handle, (*driver.Float)(constsPtr))
}

func (c *vulkanCommandBuffer) CmdSetDepthBias(depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32) {
	c.driver.VkCmdSetDepthBias(c.handle, driver.Float(depthBiasConstantFactor), driver.Float(depthBiasClamp), driver.Float(depthBiasSlopeFactor))
}

func (c *vulkanCommandBuffer) CmdSetDepthBounds(min, max float32) {
	c.driver.VkCmdSetDepthBounds(c.handle, driver.Float(min), driver.Float(max))
}

func (c *vulkanCommandBuffer) CmdSetLineWidth(lineWidth float32) {
	c.driver.VkCmdSetLineWidth(c.handle, driver.Float(lineWidth))
}

func (c *vulkanCommandBuffer) CmdSetStencilCompareMask(faceMask common.StencilFaces, compareMask uint32) {
	c.driver.VkCmdSetStencilCompareMask(c.handle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(compareMask))
}

func (c *vulkanCommandBuffer) CmdSetStencilReference(faceMask common.StencilFaces, reference uint32) {
	c.driver.VkCmdSetStencilReference(c.handle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(reference))
}

func (c *vulkanCommandBuffer) CmdSetStencilWriteMask(faceMask common.StencilFaces, writeMask uint32) {
	c.driver.VkCmdSetStencilWriteMask(c.handle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(writeMask))
}

func (c *vulkanCommandBuffer) CmdUpdateBuffer(dstBuffer Buffer, dstOffset int, dataSize int, data []byte) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	size := len(data)
	dataPtr := arena.Malloc(size)
	dataSlice := ([]byte)(unsafe.Slice((*byte)(dataPtr), size))
	copy(dataSlice, data)

	c.driver.VkCmdUpdateBuffer(c.handle, dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(dataSize), dataPtr)
}

func (c *vulkanCommandBuffer) CmdWriteTimestamp(pipelineStage common.PipelineStages, queryPool QueryPool, query int) {
	c.driver.VkCmdWriteTimestamp(c.handle, driver.VkPipelineStageFlags(pipelineStage), queryPool.Handle(), driver.Uint32(query))
}

type CommandBufferResetFlags int32

const (
	ResetReleaseResources CommandBufferResetFlags = C.VK_COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT
)

var resetFlagsToString = map[CommandBufferResetFlags]string{
	ResetReleaseResources: "Reset Release Resources",
}

func (f CommandBufferResetFlags) String() string {
	return common.FlagsToString(f, resetFlagsToString)
}

func (c *vulkanCommandBuffer) Reset(flags CommandBufferResetFlags) (common.VkResult, error) {
	return c.driver.VkResetCommandBuffer(c.handle, driver.VkCommandBufferResetFlags(flags))
}

type CommandBufferOptions struct {
	Level       common.CommandBufferLevel
	BufferCount int
	commandPool CommandPool

	common.HaveNext
}

func (o *CommandBufferOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Level == common.LevelUnset {
		return nil, errors.New("attempted to create command buffers without setting Level")
	}
	if o.BufferCount == 0 {
		return nil, errors.New("attempted to create 0 command buffers")
	}

	createInfo := (*C.VkCommandBufferAllocateInfo)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkCommandBufferAllocateInfo{}))))
	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	createInfo.pNext = next

	createInfo.level = C.VkCommandBufferLevel(o.Level)
	createInfo.commandBufferCount = C.uint32_t(o.BufferCount)
	createInfo.commandPool = C.VkCommandPool(unsafe.Pointer(o.commandPool.Handle()))

	return unsafe.Pointer(createInfo), nil
}

func (o *CommandBufferOptions) MustBeRootOptions() bool {
	return true
}
