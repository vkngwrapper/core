package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) BeginCommandBuffer(commandBuffer types.CommandBuffer, o core1_0.CommandBufferBeginInfo) (common.VkResult, error) {
	if commandBuffer.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.Driver.VkBeginCommandBuffer(commandBuffer.Handle(), (*driver.VkCommandBufferBeginInfo)(createInfo))
}

func (v *Vulkan) EndCommandBuffer(commandBuffer types.CommandBuffer) (common.VkResult, error) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	return v.Driver.VkEndCommandBuffer(commandBuffer.Handle())
}

func (v *Vulkan) CmdBeginRenderPass(commandBuffer types.CommandBuffer, contents core1_0.SubpassContents, o core1_0.RenderPassBeginInfo) error {
	if commandBuffer.Handle() == 0 {
		return errors.New("commandBuffer cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	v.Driver.VkCmdBeginRenderPass(commandBuffer.Handle(), (*driver.VkRenderPassBeginInfo)(createInfo), driver.VkSubpassContents(contents))
	return nil
}

func (v *Vulkan) CmdEndRenderPass(commandBuffer types.CommandBuffer) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdEndRenderPass(commandBuffer.Handle())
}

func (v *Vulkan) CmdBindPipeline(commandBuffer types.CommandBuffer, bindPoint core1_0.PipelineBindPoint, pipeline types.Pipeline) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if pipeline.Handle() == 0 {
		panic("pipeline cannot be uninitialized")
	}

	v.Driver.VkCmdBindPipeline(commandBuffer.Handle(), driver.VkPipelineBindPoint(bindPoint), pipeline.Handle())
}

func (v *Vulkan) CmdDraw(commandBuffer types.CommandBuffer, vertexCount, instanceCount int, firstVertex, firstInstance uint32) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdDraw(commandBuffer.Handle(), driver.Uint32(vertexCount), driver.Uint32(instanceCount), driver.Uint32(firstVertex), driver.Uint32(firstInstance))
}

func (v *Vulkan) CmdDrawIndexed(commandBuffer types.CommandBuffer, indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdDrawIndexed(commandBuffer.Handle(), driver.Uint32(indexCount), driver.Uint32(instanceCount), driver.Uint32(firstIndex), driver.Int32(vertexOffset), driver.Uint32(firstInstance))
}

func (v *Vulkan) CmdBindVertexBuffers(commandBuffer types.CommandBuffer, firstBinding int, buffers []types.Buffer, bufferOffsets []int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

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
		if buffers[i].Handle() == 0 {
			panic(fmt.Sprintf("element %d of buffers slice is uninitialized", i))
		}
		bufferArraySlice[i] = buffers[i].Handle()
		offsetArraySlice[i] = driver.VkDeviceSize(bufferOffsets[i])
	}

	v.Driver.VkCmdBindVertexBuffers(commandBuffer.Handle(), driver.Uint32(firstBinding), driver.Uint32(bufferCount), bufferArrayPtr, offsetArrayPtr)
}

func (v *Vulkan) CmdBindIndexBuffer(commandBuffer types.CommandBuffer, buffer types.Buffer, offset int, indexType core1_0.IndexType) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdBindIndexBuffer(commandBuffer.Handle(), buffer.Handle(), driver.VkDeviceSize(offset), driver.VkIndexType(indexType))
}

func (v *Vulkan) CmdBindDescriptorSets(commandBuffer types.CommandBuffer, bindPoint core1_0.PipelineBindPoint, layout types.PipelineLayout, firstSet int, sets []types.DescriptorSet, dynamicOffsets []int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if layout.Handle() == 0 {
		panic("layout cannot be uninitialized")
	}

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
			setSlice[i] = nil

			if sets[i].Handle() != 0 {
				setSlice[i] = (C.VkDescriptorSet)(unsafe.Pointer(sets[i].Handle()))
			}
		}
	}

	if dynamicOffsetCount > 0 {
		dynamicOffsetPtr = arena.Malloc(dynamicOffsetCount * int(unsafe.Sizeof(C.uint32_t(0))))
		dynamicOffsetSlice := ([]C.uint32_t)(unsafe.Slice((*C.uint32_t)(dynamicOffsetPtr), dynamicOffsetCount))

		for i := 0; i < dynamicOffsetCount; i++ {
			dynamicOffsetSlice[i] = (C.uint32_t)(dynamicOffsets[i])
		}
	}

	v.Driver.VkCmdBindDescriptorSets(commandBuffer.Handle(),
		driver.VkPipelineBindPoint(bindPoint),
		layout.Handle(),
		driver.Uint32(firstSet),
		driver.Uint32(setCount),
		(*driver.VkDescriptorSet)(setPtr),
		driver.Uint32(dynamicOffsetCount),
		(*driver.Uint32)(dynamicOffsetPtr))
}

func (v *Vulkan) CmdPipelineBarrier(commandBuffer types.CommandBuffer, srcStageMask, dstStageMask core1_0.PipelineStageFlags, dependencies core1_0.DependencyFlags, memoryBarriers []core1_0.MemoryBarrier, bufferMemoryBarriers []core1_0.BufferMemoryBarrier, imageMemoryBarriers []core1_0.ImageMemoryBarrier) error {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

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
		barrierPtr, err = common.AllocOptionSlice[C.VkMemoryBarrier, core1_0.MemoryBarrier](arena, memoryBarriers)
		if err != nil {
			return err
		}
	}

	if bufferBarrierCount > 0 {
		bufferBarrierPtr, err = common.AllocOptionSlice[C.VkBufferMemoryBarrier, core1_0.BufferMemoryBarrier](arena, bufferMemoryBarriers)
		if err != nil {
			return err
		}
	}

	if imageBarrierCount > 0 {
		imageBarrierPtr, err = common.AllocOptionSlice[C.VkImageMemoryBarrier, core1_0.ImageMemoryBarrier](arena, imageMemoryBarriers)
		if err != nil {
			return err
		}
	}

	v.Driver.VkCmdPipelineBarrier(commandBuffer.Handle(), driver.VkPipelineStageFlags(srcStageMask), driver.VkPipelineStageFlags(dstStageMask), driver.VkDependencyFlags(dependencies), driver.Uint32(barrierCount), (*driver.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), driver.Uint32(bufferBarrierCount), (*driver.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), driver.Uint32(imageBarrierCount), (*driver.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	return nil
}

func (v *Vulkan) CmdCopyBufferToImage(commandBuffer types.CommandBuffer, buffer types.Buffer, image types.Image, layout core1_0.ImageLayout, regions ...core1_0.BufferImageCopy) error {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if buffer.Handle() == 0 {
		panic("buffer cannot be uninitialized")
	}
	if image.Handle() == 0 {
		panic("image cannot be uninitialized")
	}
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

	v.Driver.VkCmdCopyBufferToImage(commandBuffer.Handle(), buffer.Handle(), image.Handle(), driver.VkImageLayout(layout), driver.Uint32(regionCount), (*driver.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))
	return nil
}

func (v *Vulkan) CmdBlitImage(commandBuffer types.CommandBuffer, sourceImage types.Image, sourceImageLayout core1_0.ImageLayout, destinationImage types.Image, destinationImageLayout core1_0.ImageLayout, regions []core1_0.ImageBlit, filter core1_0.Filter) error {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if sourceImage.Handle() == 0 {
		panic("sourceImage must not be uninitialized")
	}

	if destinationImage.Handle() == 0 {
		panic("destinationImage must not be uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	regionCount := len(regions)

	regionPtr, err := common.AllocSlice[C.VkImageBlit, core1_0.ImageBlit](allocator, regions)
	if err != nil {
		return err
	}

	v.Driver.VkCmdBlitImage(
		commandBuffer.Handle(),
		sourceImage.Handle(),
		driver.VkImageLayout(sourceImageLayout),
		destinationImage.Handle(),
		driver.VkImageLayout(destinationImageLayout),
		driver.Uint32(regionCount),
		(*driver.VkImageBlit)(unsafe.Pointer(regionPtr)),
		driver.VkFilter(filter))
	return nil
}

func (v *Vulkan) CmdPushConstants(commandBuffer types.CommandBuffer, layout types.PipelineLayout, stageFlags core1_0.ShaderStageFlags, offset int, valueBytes []byte) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if layout.Handle() == 0 {
		panic("layout cannot be uninitialized")
	}

	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	valueBytesPtr := alloc.CBytes(valueBytes)

	v.Driver.VkCmdPushConstants(commandBuffer.Handle(), layout.Handle(), driver.VkShaderStageFlags(stageFlags), driver.Uint32(offset), driver.Uint32(len(valueBytes)), valueBytesPtr)
}

func (v *Vulkan) CmdSetViewport(commandBuffer types.CommandBuffer, viewports ...core1_0.Viewport) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

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

	v.Driver.VkCmdSetViewport(commandBuffer.Handle(), driver.Uint32(0), driver.Uint32(viewportCount), (*driver.VkViewport)(unsafe.Pointer(viewportPtr)))
}

func (v *Vulkan) CmdSetScissor(commandBuffer types.CommandBuffer, scissors ...core1_0.Rect2D) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

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

	v.Driver.VkCmdSetScissor(commandBuffer.Handle(), driver.Uint32(0), driver.Uint32(scissorCount), (*driver.VkRect2D)(unsafe.Pointer(scissorPtr)))
}

func (v *Vulkan) CmdNextSubpass(commandBuffer types.CommandBuffer, contents core1_0.SubpassContents) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdNextSubpass(commandBuffer.Handle(), driver.VkSubpassContents(contents))
}

func (v *Vulkan) CmdWaitEvents(commandBuffer types.CommandBuffer, events []types.Event, srcStageMask core1_0.PipelineStageFlags, dstStageMask core1_0.PipelineStageFlags, memoryBarriers []core1_0.MemoryBarrier, bufferMemoryBarriers []core1_0.BufferMemoryBarrier, imageMemoryBarriers []core1_0.ImageMemoryBarrier) error {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

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
			if events[i].Handle() == 0 {
				panic(fmt.Sprintf("element %d of the events slice was uninitialized", i))
			}
			eventSlice[i] = C.VkEvent(unsafe.Pointer(events[i].Handle()))
		}
	}

	if barrierCount > 0 {
		barrierPtr, err = common.AllocOptionSlice[C.VkMemoryBarrier, core1_0.MemoryBarrier](arena, memoryBarriers)
		if err != nil {
			return err
		}
	}

	if bufferBarrierCount > 0 {
		bufferBarrierPtr, err = common.AllocOptionSlice[C.VkBufferMemoryBarrier, core1_0.BufferMemoryBarrier](arena, bufferMemoryBarriers)
		if err != nil {
			return err
		}
	}

	if imageBarrierCount > 0 {
		imageBarrierPtr, err = common.AllocOptionSlice[C.VkImageMemoryBarrier, core1_0.ImageMemoryBarrier](arena, imageMemoryBarriers)
		if err != nil {
			return err
		}
	}

	v.Driver.VkCmdWaitEvents(commandBuffer.Handle(), driver.Uint32(eventCount), (*driver.VkEvent)(unsafe.Pointer(eventPtr)), driver.VkPipelineStageFlags(srcStageMask), driver.VkPipelineStageFlags(dstStageMask), driver.Uint32(barrierCount), (*driver.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), driver.Uint32(bufferBarrierCount), (*driver.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), driver.Uint32(imageBarrierCount), (*driver.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	return nil
}

func (v *Vulkan) CmdSetEvent(commandBuffer types.CommandBuffer, event types.Event, stageMask core1_0.PipelineStageFlags) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if event.Handle() == 0 {
		panic("event cannot be uninitialized")
	}

	v.Driver.VkCmdSetEvent(commandBuffer.Handle(), event.Handle(), driver.VkPipelineStageFlags(stageMask))
}

func (v *Vulkan) CmdClearColorImage(commandBuffer types.CommandBuffer, image types.Image, imageLayout core1_0.ImageLayout, color core1_0.ClearColorValue, ranges ...core1_0.ImageSubresourceRange) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if image.Handle() == 0 {
		panic("image cannot be uninitialized")
	}

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

	v.Driver.VkCmdClearColorImage(commandBuffer.Handle(), image.Handle(), driver.VkImageLayout(imageLayout), (*driver.VkClearColorValue)(pColor), driver.Uint32(rangeCount), (*driver.VkImageSubresourceRange)(unsafe.Pointer(pRanges)))
}

func (v *Vulkan) CmdResetQueryPool(commandBuffer types.CommandBuffer, queryPool types.QueryPool, startQuery, queryCount int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if queryPool.Handle() == 0 {
		panic("queryPool cannot be uninitialized")
	}

	v.Driver.VkCmdResetQueryPool(commandBuffer.Handle(), queryPool.Handle(), driver.Uint32(startQuery), driver.Uint32(queryCount))
}

func (v *Vulkan) CmdBeginQuery(commandBuffer types.CommandBuffer, queryPool types.QueryPool, query int, flags core1_0.QueryControlFlags) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if queryPool.Handle() == 0 {
		panic("queryPool cannot be uninitialized")
	}

	v.Driver.VkCmdBeginQuery(commandBuffer.Handle(), queryPool.Handle(), driver.Uint32(query), driver.VkQueryControlFlags(flags))
}

func (v *Vulkan) CmdEndQuery(commandBuffer types.CommandBuffer, queryPool types.QueryPool, query int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if queryPool.Handle() == 0 {
		panic("queryPool cannot be uninitialized")
	}

	v.Driver.VkCmdEndQuery(commandBuffer.Handle(), queryPool.Handle(), driver.Uint32(query))
}

func (v *Vulkan) CmdCopyQueryPoolResults(commandBuffer types.CommandBuffer, queryPool types.QueryPool, firstQuery, queryCount int, dstBuffer types.Buffer, dstOffset, stride int, flags core1_0.QueryResultFlags) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if queryPool.Handle() == 0 {
		panic("queryPool cannot be uninitialized")
	}
	if dstBuffer.Handle() == 0 {
		panic("dstBuffer cannot be uninitialized")
	}
	v.Driver.VkCmdCopyQueryPoolResults(commandBuffer.Handle(), queryPool.Handle(), driver.Uint32(firstQuery), driver.Uint32(queryCount), dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(stride), driver.VkQueryResultFlags(flags))
}

func (v *Vulkan) CmdExecuteCommands(commandBuffer types.CommandBuffer, commandBuffers ...types.CommandBuffer) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	bufferCount := len(commandBuffers)
	commandBufferPtr := (*C.VkCommandBuffer)(arena.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))
	commandBufferSlice := ([]C.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		if commandBuffers[i].Handle() == 0 {
			panic(fmt.Sprintf("element %d of the commandBuffers slice was uninitialized", i))
		}
		commandBufferSlice[i] = C.VkCommandBuffer(unsafe.Pointer(commandBuffers[i].Handle()))
	}

	v.Driver.VkCmdExecuteCommands(commandBuffer.Handle(), driver.Uint32(bufferCount), (*driver.VkCommandBuffer)(unsafe.Pointer(commandBufferPtr)))
}

func (v *Vulkan) CmdClearAttachments(commandBuffer types.CommandBuffer, attachments []core1_0.ClearAttachment, rects []core1_0.ClearRect) error {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

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

	v.Driver.VkCmdClearAttachments(commandBuffer.Handle(), driver.Uint32(attachmentCount), (*driver.VkClearAttachment)(unsafe.Pointer(attachmentsPtr)), driver.Uint32(rectsCount), (*driver.VkClearRect)(unsafe.Pointer(rectsPtr)))
	return nil
}

func (v *Vulkan) CmdClearDepthStencilImage(commandBuffer types.CommandBuffer, image types.Image, imageLayout core1_0.ImageLayout, depthStencil *core1_0.ClearValueDepthStencil, ranges ...core1_0.ImageSubresourceRange) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if image.Handle() == 0 {
		panic("image cannot be nil")
	}
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

	v.Driver.VkCmdClearDepthStencilImage(commandBuffer.Handle(), image.Handle(), driver.VkImageLayout(imageLayout), (*driver.VkClearDepthStencilValue)(unsafe.Pointer(depthStencilPtr)), driver.Uint32(rangeCount), (*driver.VkImageSubresourceRange)(unsafe.Pointer(rangePtr)))
}

func (v *Vulkan) CmdCopyImageToBuffer(commandBuffer types.CommandBuffer, srcImage types.Image, srcImageLayout core1_0.ImageLayout, dstBuffer types.Buffer, regions ...core1_0.BufferImageCopy) error {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if srcImage.Handle() == 0 {
		panic("srcImage cannot be uninitailized")
	}
	if dstBuffer.Handle() == 0 {
		panic("dstBuffer cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionPtr, err := common.AllocSlice[C.VkBufferImageCopy, core1_0.BufferImageCopy](arena, regions)
	if err != nil {
		return err
	}

	v.Driver.VkCmdCopyImageToBuffer(commandBuffer.Handle(), srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstBuffer.Handle(), driver.Uint32(regionCount), (*driver.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))
	return nil
}

func (v *Vulkan) CmdDispatch(commandBuffer types.CommandBuffer, groupCountX, groupCountY, groupCountZ int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdDispatch(commandBuffer.Handle(), driver.Uint32(groupCountX), driver.Uint32(groupCountY), driver.Uint32(groupCountZ))
}

func (v *Vulkan) CmdDispatchIndirect(commandBuffer types.CommandBuffer, buffer types.Buffer, offset int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if buffer.Handle() == 0 {
		panic("buffer cannot be uninitialized")
	}
	v.Driver.VkCmdDispatchIndirect(commandBuffer.Handle(), buffer.Handle(), driver.VkDeviceSize(offset))
}

func (v *Vulkan) CmdDrawIndexedIndirect(commandBuffer types.CommandBuffer, buffer types.Buffer, offset int, drawCount, stride int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if buffer.Handle() == 0 {
		panic("buffer cannot be uninitialized")
	}
	v.Driver.VkCmdDrawIndexedIndirect(commandBuffer.Handle(), buffer.Handle(), driver.VkDeviceSize(offset), driver.Uint32(drawCount), driver.Uint32(stride))
}

func (v *Vulkan) CmdDrawIndirect(commandBuffer types.CommandBuffer, buffer types.Buffer, offset int, drawCount, stride int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if buffer.Handle() == 0 {
		panic("buffer cannot be uninitialized")
	}
	v.Driver.VkCmdDrawIndirect(commandBuffer.Handle(), buffer.Handle(), driver.VkDeviceSize(offset), driver.Uint32(drawCount), driver.Uint32(stride))
}

func (v *Vulkan) CmdFillBuffer(commandBuffer types.CommandBuffer, dstBuffer types.Buffer, dstOffset int, size int, data uint32) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if dstBuffer.Handle() == 0 {
		panic("dstBuffer cannot be uninitialized")
	}
	v.Driver.VkCmdFillBuffer(commandBuffer.Handle(), dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(size), driver.Uint32(data))
}

func (v *Vulkan) CmdResetEvent(commandBuffer types.CommandBuffer, event types.Event, stageMask core1_0.PipelineStageFlags) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if event.Handle() == 0 {
		panic("event cannot be uninitialized")
	}
	v.Driver.VkCmdResetEvent(commandBuffer.Handle(), event.Handle(), driver.VkPipelineStageFlags(stageMask))
}

func (v *Vulkan) CmdResolveImage(commandBuffer types.CommandBuffer, srcImage types.Image, srcImageLayout core1_0.ImageLayout, dstImage types.Image, dstImageLayout core1_0.ImageLayout, regions ...core1_0.ImageResolve) error {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if srcImage.Handle() == 0 {
		panic("srcImage cannot be uninitialized")
	}
	if dstImage.Handle() == 0 {
		panic("dstImage cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionsPtr, err := common.AllocSlice[C.VkImageResolve, core1_0.ImageResolve](arena, regions)
	if err != nil {
		return err
	}

	v.Driver.VkCmdResolveImage(commandBuffer.Handle(), srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstImage.Handle(), driver.VkImageLayout(dstImageLayout), driver.Uint32(regionCount), (*driver.VkImageResolve)(unsafe.Pointer(regionsPtr)))
	return nil
}

func (v *Vulkan) CmdSetBlendConstants(commandBuffer types.CommandBuffer, blendConstants [4]float32) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	constsPtr := (*C.float)(arena.Malloc(16))
	constsSlice := ([]C.float)(unsafe.Slice(constsPtr, 4))

	for i := 0; i < 4; i++ {
		constsSlice[i] = C.float(blendConstants[i])
	}

	v.Driver.VkCmdSetBlendConstants(commandBuffer.Handle(), (*driver.Float)(constsPtr))
}

func (v *Vulkan) CmdSetDepthBias(commandBuffer types.CommandBuffer, depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdSetDepthBias(commandBuffer.Handle(), driver.Float(depthBiasConstantFactor), driver.Float(depthBiasClamp), driver.Float(depthBiasSlopeFactor))
}

func (v *Vulkan) CmdSetDepthBounds(commandBuffer types.CommandBuffer, min, max float32) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdSetDepthBounds(commandBuffer.Handle(), driver.Float(min), driver.Float(max))
}

func (v *Vulkan) CmdSetLineWidth(commandBuffer types.CommandBuffer, lineWidth float32) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdSetLineWidth(commandBuffer.Handle(), driver.Float(lineWidth))
}

func (v *Vulkan) CmdSetStencilCompareMask(commandBuffer types.CommandBuffer, faceMask core1_0.StencilFaceFlags, compareMask uint32) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdSetStencilCompareMask(commandBuffer.Handle(), driver.VkStencilFaceFlags(faceMask), driver.Uint32(compareMask))
}

func (v *Vulkan) CmdSetStencilReference(commandBuffer types.CommandBuffer, faceMask core1_0.StencilFaceFlags, reference uint32) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdSetStencilReference(commandBuffer.Handle(), driver.VkStencilFaceFlags(faceMask), driver.Uint32(reference))
}

func (v *Vulkan) CmdSetStencilWriteMask(commandBuffer types.CommandBuffer, faceMask core1_0.StencilFaceFlags, writeMask uint32) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	v.Driver.VkCmdSetStencilWriteMask(commandBuffer.Handle(), driver.VkStencilFaceFlags(faceMask), driver.Uint32(writeMask))
}

func (v *Vulkan) CmdUpdateBuffer(commandBuffer types.CommandBuffer, dstBuffer types.Buffer, dstOffset int, dataSize int, data []byte) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if dstBuffer.Handle() == 0 {
		panic("dstBuffer cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	size := len(data)
	dataPtr := arena.Malloc(size)
	dataSlice := ([]byte)(unsafe.Slice((*byte)(dataPtr), size))
	copy(dataSlice, data)

	v.Driver.VkCmdUpdateBuffer(commandBuffer.Handle(), dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(dataSize), dataPtr)
}

func (v *Vulkan) CmdWriteTimestamp(commandBuffer types.CommandBuffer, pipelineStage core1_0.PipelineStageFlags, queryPool types.QueryPool, query int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}

	if queryPool.Handle() == 0 {
		panic("queryPool cannot be uninitialized")
	}

	v.Driver.VkCmdWriteTimestamp(commandBuffer.Handle(), driver.VkPipelineStageFlags(pipelineStage), queryPool.Handle(), driver.Uint32(query))
}

func (v *Vulkan) ResetCommandBuffer(commandBuffer types.CommandBuffer, flags core1_0.CommandBufferResetFlags) (common.VkResult, error) {
	if commandBuffer.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("commandBuffer cannot be uninitialized")
	}

	return v.Driver.VkResetCommandBuffer(commandBuffer.Handle(), driver.VkCommandBufferResetFlags(flags))
}
