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
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) BeginCommandBuffer(commandBuffer core.CommandBuffer, o core1_0.CommandBufferBeginInfo) (common.VkResult, error) {
	if !commandBuffer.Initialized() {
		return core1_0.VKErrorUnknown, errors.New("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return v.LoaderObj.VkBeginCommandBuffer(commandBuffer.Handle(), (*loader.VkCommandBufferBeginInfo)(createInfo))
}

func (v *DeviceVulkanDriver) EndCommandBuffer(commandBuffer core.CommandBuffer) (common.VkResult, error) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	return v.LoaderObj.VkEndCommandBuffer(commandBuffer.Handle())
}

func (v *DeviceVulkanDriver) CmdBeginRenderPass(commandBuffer core.CommandBuffer, contents core1_0.SubpassContents, o core1_0.RenderPassBeginInfo) error {
	if !commandBuffer.Initialized() {
		return errors.New("commandBuffer cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	v.LoaderObj.VkCmdBeginRenderPass(commandBuffer.Handle(), (*loader.VkRenderPassBeginInfo)(createInfo), loader.VkSubpassContents(contents))
	return nil
}

func (v *DeviceVulkanDriver) CmdEndRenderPass(commandBuffer core.CommandBuffer) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdEndRenderPass(commandBuffer.Handle())
}

func (v *DeviceVulkanDriver) CmdBindPipeline(commandBuffer core.CommandBuffer, bindPoint core1_0.PipelineBindPoint, pipeline core.Pipeline) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !pipeline.Initialized() {
		panic("pipeline cannot be uninitialized")
	}

	v.LoaderObj.VkCmdBindPipeline(commandBuffer.Handle(), loader.VkPipelineBindPoint(bindPoint), pipeline.Handle())
}

func (v *DeviceVulkanDriver) CmdDraw(commandBuffer core.CommandBuffer, vertexCount, instanceCount int, firstVertex, firstInstance uint32) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdDraw(commandBuffer.Handle(), loader.Uint32(vertexCount), loader.Uint32(instanceCount), loader.Uint32(firstVertex), loader.Uint32(firstInstance))
}

func (v *DeviceVulkanDriver) CmdDrawIndexed(commandBuffer core.CommandBuffer, indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdDrawIndexed(commandBuffer.Handle(), loader.Uint32(indexCount), loader.Uint32(instanceCount), loader.Uint32(firstIndex), loader.Int32(vertexOffset), loader.Uint32(firstInstance))
}

func (v *DeviceVulkanDriver) CmdBindVertexBuffers(commandBuffer core.CommandBuffer, firstBinding int, buffers []core.Buffer, bufferOffsets []int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	bufferCount := len(buffers)

	bufferArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkBuffer{})))
	offsetArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof(C.VkDeviceSize(0))))

	bufferArrayPtr := (*loader.VkBuffer)(bufferArrayUnsafe)
	offsetArrayPtr := (*loader.VkDeviceSize)(offsetArrayUnsafe)

	bufferArraySlice := ([]loader.VkBuffer)(unsafe.Slice(bufferArrayPtr, bufferCount))
	offsetArraySlice := ([]loader.VkDeviceSize)(unsafe.Slice(offsetArrayPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		if buffers[i].Handle() == 0 {
			panic(fmt.Sprintf("element %d of buffers slice is uninitialized", i))
		}
		bufferArraySlice[i] = buffers[i].Handle()
		offsetArraySlice[i] = loader.VkDeviceSize(bufferOffsets[i])
	}

	v.LoaderObj.VkCmdBindVertexBuffers(commandBuffer.Handle(), loader.Uint32(firstBinding), loader.Uint32(bufferCount), bufferArrayPtr, offsetArrayPtr)
}

func (v *DeviceVulkanDriver) CmdBindIndexBuffer(commandBuffer core.CommandBuffer, buffer core.Buffer, offset int, indexType core1_0.IndexType) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdBindIndexBuffer(commandBuffer.Handle(), buffer.Handle(), loader.VkDeviceSize(offset), loader.VkIndexType(indexType))
}

func (v *DeviceVulkanDriver) CmdBindDescriptorSets(commandBuffer core.CommandBuffer, bindPoint core1_0.PipelineBindPoint, layout core.PipelineLayout, firstSet int, sets []core.DescriptorSet, dynamicOffsets []int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !layout.Initialized() {
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

	v.LoaderObj.VkCmdBindDescriptorSets(commandBuffer.Handle(),
		loader.VkPipelineBindPoint(bindPoint),
		layout.Handle(),
		loader.Uint32(firstSet),
		loader.Uint32(setCount),
		(*loader.VkDescriptorSet)(setPtr),
		loader.Uint32(dynamicOffsetCount),
		(*loader.Uint32)(dynamicOffsetPtr))
}

func (v *DeviceVulkanDriver) CmdPipelineBarrier(commandBuffer core.CommandBuffer, srcStageMask, dstStageMask core1_0.PipelineStageFlags, dependencies core1_0.DependencyFlags, memoryBarriers []core1_0.MemoryBarrier, bufferMemoryBarriers []core1_0.BufferMemoryBarrier, imageMemoryBarriers []core1_0.ImageMemoryBarrier) error {
	if !commandBuffer.Initialized() {
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

	v.LoaderObj.VkCmdPipelineBarrier(commandBuffer.Handle(), loader.VkPipelineStageFlags(srcStageMask), loader.VkPipelineStageFlags(dstStageMask), loader.VkDependencyFlags(dependencies), loader.Uint32(barrierCount), (*loader.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), loader.Uint32(bufferBarrierCount), (*loader.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), loader.Uint32(imageBarrierCount), (*loader.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	return nil
}

func (v *DeviceVulkanDriver) CmdCopyBufferToImage(commandBuffer core.CommandBuffer, buffer core.Buffer, image core.Image, layout core1_0.ImageLayout, regions ...core1_0.BufferImageCopy) error {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !buffer.Initialized() {
		panic("buffer cannot be uninitialized")
	}
	if !image.Initialized() {
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

	v.LoaderObj.VkCmdCopyBufferToImage(commandBuffer.Handle(), buffer.Handle(), image.Handle(), loader.VkImageLayout(layout), loader.Uint32(regionCount), (*loader.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))
	return nil
}

func (v *DeviceVulkanDriver) CmdBlitImage(commandBuffer core.CommandBuffer, sourceImage core.Image, sourceImageLayout core1_0.ImageLayout, destinationImage core.Image, destinationImageLayout core1_0.ImageLayout, regions []core1_0.ImageBlit, filter core1_0.Filter) error {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !sourceImage.Initialized() {
		panic("sourceImage must not be uninitialized")
	}

	if !destinationImage.Initialized() {
		panic("destinationImage must not be uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	regionCount := len(regions)

	regionPtr, err := common.AllocSlice[C.VkImageBlit, core1_0.ImageBlit](allocator, regions)
	if err != nil {
		return err
	}

	v.LoaderObj.VkCmdBlitImage(
		commandBuffer.Handle(),
		sourceImage.Handle(),
		loader.VkImageLayout(sourceImageLayout),
		destinationImage.Handle(),
		loader.VkImageLayout(destinationImageLayout),
		loader.Uint32(regionCount),
		(*loader.VkImageBlit)(unsafe.Pointer(regionPtr)),
		loader.VkFilter(filter))
	return nil
}

func (v *DeviceVulkanDriver) CmdPushConstants(commandBuffer core.CommandBuffer, layout core.PipelineLayout, stageFlags core1_0.ShaderStageFlags, offset int, valueBytes []byte) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !layout.Initialized() {
		panic("layout cannot be uninitialized")
	}

	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	valueBytesPtr := alloc.CBytes(valueBytes)

	v.LoaderObj.VkCmdPushConstants(commandBuffer.Handle(), layout.Handle(), loader.VkShaderStageFlags(stageFlags), loader.Uint32(offset), loader.Uint32(len(valueBytes)), valueBytesPtr)
}

func (v *DeviceVulkanDriver) CmdSetViewport(commandBuffer core.CommandBuffer, viewports ...core1_0.Viewport) {
	if !commandBuffer.Initialized() {
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

	v.LoaderObj.VkCmdSetViewport(commandBuffer.Handle(), loader.Uint32(0), loader.Uint32(viewportCount), (*loader.VkViewport)(unsafe.Pointer(viewportPtr)))
}

func (v *DeviceVulkanDriver) CmdSetScissor(commandBuffer core.CommandBuffer, scissors ...core1_0.Rect2D) {
	if !commandBuffer.Initialized() {
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

	v.LoaderObj.VkCmdSetScissor(commandBuffer.Handle(), loader.Uint32(0), loader.Uint32(scissorCount), (*loader.VkRect2D)(unsafe.Pointer(scissorPtr)))
}

func (v *DeviceVulkanDriver) CmdNextSubpass(commandBuffer core.CommandBuffer, contents core1_0.SubpassContents) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdNextSubpass(commandBuffer.Handle(), loader.VkSubpassContents(contents))
}

func (v *DeviceVulkanDriver) CmdWaitEvents(commandBuffer core.CommandBuffer, events []core.Event, srcStageMask core1_0.PipelineStageFlags, dstStageMask core1_0.PipelineStageFlags, memoryBarriers []core1_0.MemoryBarrier, bufferMemoryBarriers []core1_0.BufferMemoryBarrier, imageMemoryBarriers []core1_0.ImageMemoryBarrier) error {
	if !commandBuffer.Initialized() {
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

	v.LoaderObj.VkCmdWaitEvents(commandBuffer.Handle(), loader.Uint32(eventCount), (*loader.VkEvent)(unsafe.Pointer(eventPtr)), loader.VkPipelineStageFlags(srcStageMask), loader.VkPipelineStageFlags(dstStageMask), loader.Uint32(barrierCount), (*loader.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), loader.Uint32(bufferBarrierCount), (*loader.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), loader.Uint32(imageBarrierCount), (*loader.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	return nil
}

func (v *DeviceVulkanDriver) CmdSetEvent(commandBuffer core.CommandBuffer, event core.Event, stageMask core1_0.PipelineStageFlags) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !event.Initialized() {
		panic("event cannot be uninitialized")
	}

	v.LoaderObj.VkCmdSetEvent(commandBuffer.Handle(), event.Handle(), loader.VkPipelineStageFlags(stageMask))
}

func (v *DeviceVulkanDriver) CmdClearColorImage(commandBuffer core.CommandBuffer, image core.Image, imageLayout core1_0.ImageLayout, color core1_0.ClearColorValue, ranges ...core1_0.ImageSubresourceRange) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !image.Initialized() {
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

	v.LoaderObj.VkCmdClearColorImage(commandBuffer.Handle(), image.Handle(), loader.VkImageLayout(imageLayout), (*loader.VkClearColorValue)(pColor), loader.Uint32(rangeCount), (*loader.VkImageSubresourceRange)(unsafe.Pointer(pRanges)))
}

func (v *DeviceVulkanDriver) CmdResetQueryPool(commandBuffer core.CommandBuffer, queryPool core.QueryPool, startQuery, queryCount int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !queryPool.Initialized() {
		panic("queryPool cannot be uninitialized")
	}

	v.LoaderObj.VkCmdResetQueryPool(commandBuffer.Handle(), queryPool.Handle(), loader.Uint32(startQuery), loader.Uint32(queryCount))
}

func (v *DeviceVulkanDriver) CmdBeginQuery(commandBuffer core.CommandBuffer, queryPool core.QueryPool, query int, flags core1_0.QueryControlFlags) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !queryPool.Initialized() {
		panic("queryPool cannot be uninitialized")
	}

	v.LoaderObj.VkCmdBeginQuery(commandBuffer.Handle(), queryPool.Handle(), loader.Uint32(query), loader.VkQueryControlFlags(flags))
}

func (v *DeviceVulkanDriver) CmdEndQuery(commandBuffer core.CommandBuffer, queryPool core.QueryPool, query int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !queryPool.Initialized() {
		panic("queryPool cannot be uninitialized")
	}

	v.LoaderObj.VkCmdEndQuery(commandBuffer.Handle(), queryPool.Handle(), loader.Uint32(query))
}

func (v *DeviceVulkanDriver) CmdCopyQueryPoolResults(commandBuffer core.CommandBuffer, queryPool core.QueryPool, firstQuery, queryCount int, dstBuffer core.Buffer, dstOffset, stride int, flags core1_0.QueryResultFlags) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !queryPool.Initialized() {
		panic("queryPool cannot be uninitialized")
	}
	if !dstBuffer.Initialized() {
		panic("dstBuffer cannot be uninitialized")
	}
	v.LoaderObj.VkCmdCopyQueryPoolResults(commandBuffer.Handle(), queryPool.Handle(), loader.Uint32(firstQuery), loader.Uint32(queryCount), dstBuffer.Handle(), loader.VkDeviceSize(dstOffset), loader.VkDeviceSize(stride), loader.VkQueryResultFlags(flags))
}

func (v *DeviceVulkanDriver) CmdExecuteCommands(commandBuffer core.CommandBuffer, commandBuffers ...core.CommandBuffer) {
	if !commandBuffer.Initialized() {
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

	v.LoaderObj.VkCmdExecuteCommands(commandBuffer.Handle(), loader.Uint32(bufferCount), (*loader.VkCommandBuffer)(unsafe.Pointer(commandBufferPtr)))
}

func (v *DeviceVulkanDriver) CmdClearAttachments(commandBuffer core.CommandBuffer, attachments []core1_0.ClearAttachment, rects []core1_0.ClearRect) error {
	if !commandBuffer.Initialized() {
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

	v.LoaderObj.VkCmdClearAttachments(commandBuffer.Handle(), loader.Uint32(attachmentCount), (*loader.VkClearAttachment)(unsafe.Pointer(attachmentsPtr)), loader.Uint32(rectsCount), (*loader.VkClearRect)(unsafe.Pointer(rectsPtr)))
	return nil
}

func (v *DeviceVulkanDriver) CmdClearDepthStencilImage(commandBuffer core.CommandBuffer, image core.Image, imageLayout core1_0.ImageLayout, depthStencil *core1_0.ClearValueDepthStencil, ranges ...core1_0.ImageSubresourceRange) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !image.Initialized() {
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

	v.LoaderObj.VkCmdClearDepthStencilImage(commandBuffer.Handle(), image.Handle(), loader.VkImageLayout(imageLayout), (*loader.VkClearDepthStencilValue)(unsafe.Pointer(depthStencilPtr)), loader.Uint32(rangeCount), (*loader.VkImageSubresourceRange)(unsafe.Pointer(rangePtr)))
}

func (v *DeviceVulkanDriver) CmdCopyImageToBuffer(commandBuffer core.CommandBuffer, srcImage core.Image, srcImageLayout core1_0.ImageLayout, dstBuffer core.Buffer, regions ...core1_0.BufferImageCopy) error {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !srcImage.Initialized() {
		panic("srcImage cannot be uninitailized")
	}
	if !dstBuffer.Initialized() {
		panic("dstBuffer cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionPtr, err := common.AllocSlice[C.VkBufferImageCopy, core1_0.BufferImageCopy](arena, regions)
	if err != nil {
		return err
	}

	v.LoaderObj.VkCmdCopyImageToBuffer(commandBuffer.Handle(), srcImage.Handle(), loader.VkImageLayout(srcImageLayout), dstBuffer.Handle(), loader.Uint32(regionCount), (*loader.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))
	return nil
}

func (v *DeviceVulkanDriver) CmdDispatch(commandBuffer core.CommandBuffer, groupCountX, groupCountY, groupCountZ int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdDispatch(commandBuffer.Handle(), loader.Uint32(groupCountX), loader.Uint32(groupCountY), loader.Uint32(groupCountZ))
}

func (v *DeviceVulkanDriver) CmdDispatchIndirect(commandBuffer core.CommandBuffer, buffer core.Buffer, offset int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !buffer.Initialized() {
		panic("buffer cannot be uninitialized")
	}
	v.LoaderObj.VkCmdDispatchIndirect(commandBuffer.Handle(), buffer.Handle(), loader.VkDeviceSize(offset))
}

func (v *DeviceVulkanDriver) CmdDrawIndexedIndirect(commandBuffer core.CommandBuffer, buffer core.Buffer, offset int, drawCount, stride int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !buffer.Initialized() {
		panic("buffer cannot be uninitialized")
	}
	v.LoaderObj.VkCmdDrawIndexedIndirect(commandBuffer.Handle(), buffer.Handle(), loader.VkDeviceSize(offset), loader.Uint32(drawCount), loader.Uint32(stride))
}

func (v *DeviceVulkanDriver) CmdDrawIndirect(commandBuffer core.CommandBuffer, buffer core.Buffer, offset int, drawCount, stride int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !buffer.Initialized() {
		panic("buffer cannot be uninitialized")
	}
	v.LoaderObj.VkCmdDrawIndirect(commandBuffer.Handle(), buffer.Handle(), loader.VkDeviceSize(offset), loader.Uint32(drawCount), loader.Uint32(stride))
}

func (v *DeviceVulkanDriver) CmdFillBuffer(commandBuffer core.CommandBuffer, dstBuffer core.Buffer, dstOffset int, size int, data uint32) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !dstBuffer.Initialized() {
		panic("dstBuffer cannot be uninitialized")
	}
	v.LoaderObj.VkCmdFillBuffer(commandBuffer.Handle(), dstBuffer.Handle(), loader.VkDeviceSize(dstOffset), loader.VkDeviceSize(size), loader.Uint32(data))
}

func (v *DeviceVulkanDriver) CmdResetEvent(commandBuffer core.CommandBuffer, event core.Event, stageMask core1_0.PipelineStageFlags) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !event.Initialized() {
		panic("event cannot be uninitialized")
	}
	v.LoaderObj.VkCmdResetEvent(commandBuffer.Handle(), event.Handle(), loader.VkPipelineStageFlags(stageMask))
}

func (v *DeviceVulkanDriver) CmdResolveImage(commandBuffer core.CommandBuffer, srcImage core.Image, srcImageLayout core1_0.ImageLayout, dstImage core.Image, dstImageLayout core1_0.ImageLayout, regions ...core1_0.ImageResolve) error {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !srcImage.Initialized() {
		panic("srcImage cannot be uninitialized")
	}
	if !dstImage.Initialized() {
		panic("dstImage cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionsPtr, err := common.AllocSlice[C.VkImageResolve, core1_0.ImageResolve](arena, regions)
	if err != nil {
		return err
	}

	v.LoaderObj.VkCmdResolveImage(commandBuffer.Handle(), srcImage.Handle(), loader.VkImageLayout(srcImageLayout), dstImage.Handle(), loader.VkImageLayout(dstImageLayout), loader.Uint32(regionCount), (*loader.VkImageResolve)(unsafe.Pointer(regionsPtr)))
	return nil
}

func (v *DeviceVulkanDriver) CmdSetBlendConstants(commandBuffer core.CommandBuffer, blendConstants [4]float32) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	constsPtr := (*C.float)(arena.Malloc(16))
	constsSlice := ([]C.float)(unsafe.Slice(constsPtr, 4))

	for i := 0; i < 4; i++ {
		constsSlice[i] = C.float(blendConstants[i])
	}

	v.LoaderObj.VkCmdSetBlendConstants(commandBuffer.Handle(), (*loader.Float)(constsPtr))
}

func (v *DeviceVulkanDriver) CmdSetDepthBias(commandBuffer core.CommandBuffer, depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdSetDepthBias(commandBuffer.Handle(), loader.Float(depthBiasConstantFactor), loader.Float(depthBiasClamp), loader.Float(depthBiasSlopeFactor))
}

func (v *DeviceVulkanDriver) CmdSetDepthBounds(commandBuffer core.CommandBuffer, min, max float32) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdSetDepthBounds(commandBuffer.Handle(), loader.Float(min), loader.Float(max))
}

func (v *DeviceVulkanDriver) CmdSetLineWidth(commandBuffer core.CommandBuffer, lineWidth float32) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdSetLineWidth(commandBuffer.Handle(), loader.Float(lineWidth))
}

func (v *DeviceVulkanDriver) CmdSetStencilCompareMask(commandBuffer core.CommandBuffer, faceMask core1_0.StencilFaceFlags, compareMask uint32) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdSetStencilCompareMask(commandBuffer.Handle(), loader.VkStencilFaceFlags(faceMask), loader.Uint32(compareMask))
}

func (v *DeviceVulkanDriver) CmdSetStencilReference(commandBuffer core.CommandBuffer, faceMask core1_0.StencilFaceFlags, reference uint32) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdSetStencilReference(commandBuffer.Handle(), loader.VkStencilFaceFlags(faceMask), loader.Uint32(reference))
}

func (v *DeviceVulkanDriver) CmdSetStencilWriteMask(commandBuffer core.CommandBuffer, faceMask core1_0.StencilFaceFlags, writeMask uint32) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.LoaderObj.VkCmdSetStencilWriteMask(commandBuffer.Handle(), loader.VkStencilFaceFlags(faceMask), loader.Uint32(writeMask))
}

func (v *DeviceVulkanDriver) CmdUpdateBuffer(commandBuffer core.CommandBuffer, dstBuffer core.Buffer, dstOffset int, dataSize int, data []byte) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !dstBuffer.Initialized() {
		panic("dstBuffer cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	size := len(data)
	dataPtr := arena.Malloc(size)
	dataSlice := ([]byte)(unsafe.Slice((*byte)(dataPtr), size))
	copy(dataSlice, data)

	v.LoaderObj.VkCmdUpdateBuffer(commandBuffer.Handle(), dstBuffer.Handle(), loader.VkDeviceSize(dstOffset), loader.VkDeviceSize(dataSize), dataPtr)
}

func (v *DeviceVulkanDriver) CmdWriteTimestamp(commandBuffer core.CommandBuffer, pipelineStage core1_0.PipelineStageFlags, queryPool core.QueryPool, query int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !queryPool.Initialized() {
		panic("queryPool cannot be uninitialized")
	}

	v.LoaderObj.VkCmdWriteTimestamp(commandBuffer.Handle(), loader.VkPipelineStageFlags(pipelineStage), queryPool.Handle(), loader.Uint32(query))
}

func (v *DeviceVulkanDriver) ResetCommandBuffer(commandBuffer core.CommandBuffer, flags core1_0.CommandBufferResetFlags) (common.VkResult, error) {
	if !commandBuffer.Initialized() {
		return core1_0.VKErrorUnknown, errors.New("commandBuffer cannot be uninitialized")
	}

	return v.LoaderObj.VkResetCommandBuffer(commandBuffer.Handle(), loader.VkCommandBufferResetFlags(flags))
}
