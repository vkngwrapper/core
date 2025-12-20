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
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanCommandBuffer is an implementation of the CommandBuffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanCommandBuffer struct {
	DeviceDriver        driver.Driver
	Device              driver.VkDevice
	CommandPool         driver.VkCommandPool
	CommandBufferHandle driver.VkCommandBuffer

	Counter *core1_0.CommandCounter

	MaximumAPIVersion common.APIVersion
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

func (c *VulkanCommandBuffer) APIVersion() common.APIVersion {
	return c.MaximumAPIVersion
}

func (c *VulkanCommandBuffer) CommandCounter() *core1_0.CommandCounter {
	return c.Counter
}

func (c *VulkanCommandBuffer) CommandsRecorded() int {
	return c.Counter.CommandCount
}

func (c *VulkanCommandBuffer) DrawsRecorded() int {
	return c.Counter.DrawCallCount
}

func (c *VulkanCommandBuffer) DispatchesRecorded() int {
	return c.Counter.DispatchCount
}

func (c *VulkanCommandBuffer) Begin(o core1_0.CommandBufferBeginInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	res, err := c.DeviceDriver.VkBeginCommandBuffer(c.CommandBufferHandle, (*driver.VkCommandBufferBeginInfo)(createInfo))
	if err == nil {
		c.Counter.CommandCount = 0
		c.Counter.DrawCallCount = 0
		c.Counter.DispatchCount = 0
	}

	return res, err
}

func (c *VulkanCommandBuffer) End() (common.VkResult, error) {
	return c.DeviceDriver.VkEndCommandBuffer(c.CommandBufferHandle)
}

func (c *VulkanCommandBuffer) CmdBeginRenderPass(contents core1_0.SubpassContents, o core1_0.RenderPassBeginInfo) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdBeginRenderPass(c.CommandBufferHandle, (*driver.VkRenderPassBeginInfo)(createInfo), driver.VkSubpassContents(contents))
	c.Counter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdEndRenderPass() {
	c.DeviceDriver.VkCmdEndRenderPass(c.CommandBufferHandle)
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdBindPipeline(bindPoint core1_0.PipelineBindPoint, pipeline core1_0.Pipeline) {
	if pipeline == nil {
		panic("pipeline cannot be nil")
	}

	c.DeviceDriver.VkCmdBindPipeline(c.CommandBufferHandle, driver.VkPipelineBindPoint(bindPoint), pipeline.Handle())
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) {
	c.DeviceDriver.VkCmdDraw(c.CommandBufferHandle, driver.Uint32(vertexCount), driver.Uint32(instanceCount), driver.Uint32(firstVertex), driver.Uint32(firstInstance))
	c.Counter.CommandCount++
	c.Counter.DrawCallCount++
}

func (c *VulkanCommandBuffer) CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) {
	c.DeviceDriver.VkCmdDrawIndexed(c.CommandBufferHandle, driver.Uint32(indexCount), driver.Uint32(instanceCount), driver.Uint32(firstIndex), driver.Int32(vertexOffset), driver.Uint32(firstInstance))
	c.Counter.CommandCount++
	c.Counter.DrawCallCount++
}

func (c *VulkanCommandBuffer) CmdBindVertexBuffers(firstBinding int, buffers []core1_0.Buffer, bufferOffsets []int) {
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
		if buffers[i] == nil {
			panic(fmt.Sprintf("element %d of buffers slice is nil", i))
		}
		bufferArraySlice[i] = buffers[i].Handle()
		offsetArraySlice[i] = driver.VkDeviceSize(bufferOffsets[i])
	}

	c.DeviceDriver.VkCmdBindVertexBuffers(c.CommandBufferHandle, driver.Uint32(firstBinding), driver.Uint32(bufferCount), bufferArrayPtr, offsetArrayPtr)
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdBindIndexBuffer(buffer core1_0.Buffer, offset int, indexType core1_0.IndexType) {
	c.DeviceDriver.VkCmdBindIndexBuffer(c.CommandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset), driver.VkIndexType(indexType))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdBindDescriptorSets(bindPoint core1_0.PipelineBindPoint, layout core1_0.PipelineLayout, firstSet int, sets []core1_0.DescriptorSet, dynamicOffsets []int) {
	if layout == nil {
		panic("layout cannot be nil")
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

			if sets[i] != nil {
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

	c.DeviceDriver.VkCmdBindDescriptorSets(c.CommandBufferHandle,
		driver.VkPipelineBindPoint(bindPoint),
		layout.Handle(),
		driver.Uint32(firstSet),
		driver.Uint32(setCount),
		(*driver.VkDescriptorSet)(setPtr),
		driver.Uint32(dynamicOffsetCount),
		(*driver.Uint32)(dynamicOffsetPtr))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdPipelineBarrier(srcStageMask, dstStageMask core1_0.PipelineStageFlags, dependencies core1_0.DependencyFlags, memoryBarriers []core1_0.MemoryBarrier, bufferMemoryBarriers []core1_0.BufferMemoryBarrier, imageMemoryBarriers []core1_0.ImageMemoryBarrier) error {
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

	c.DeviceDriver.VkCmdPipelineBarrier(c.CommandBufferHandle, driver.VkPipelineStageFlags(srcStageMask), driver.VkPipelineStageFlags(dstStageMask), driver.VkDependencyFlags(dependencies), driver.Uint32(barrierCount), (*driver.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), driver.Uint32(bufferBarrierCount), (*driver.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), driver.Uint32(imageBarrierCount), (*driver.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	c.Counter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdCopyBufferToImage(buffer core1_0.Buffer, image core1_0.Image, layout core1_0.ImageLayout, regions []core1_0.BufferImageCopy) error {
	if buffer == nil {
		panic("buffer cannot be nil")
	}
	if image == nil {
		panic("image cannot be nil")
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

	c.DeviceDriver.VkCmdCopyBufferToImage(c.CommandBufferHandle, buffer.Handle(), image.Handle(), driver.VkImageLayout(layout), driver.Uint32(regionCount), (*driver.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))
	c.Counter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdBlitImage(sourceImage core1_0.Image, sourceImageLayout core1_0.ImageLayout, destinationImage core1_0.Image, destinationImageLayout core1_0.ImageLayout, regions []core1_0.ImageBlit, filter core1_0.Filter) error {

	if sourceImage == nil {
		panic("sourceImage must not be nil")
	}

	if destinationImage == nil {
		panic("destinationImage must not be nil")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	regionCount := len(regions)

	regionPtr, err := common.AllocSlice[C.VkImageBlit, core1_0.ImageBlit](allocator, regions)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdBlitImage(
		c.CommandBufferHandle,
		sourceImage.Handle(),
		driver.VkImageLayout(sourceImageLayout),
		destinationImage.Handle(),
		driver.VkImageLayout(destinationImageLayout),
		driver.Uint32(regionCount),
		(*driver.VkImageBlit)(unsafe.Pointer(regionPtr)),
		driver.VkFilter(filter))
	c.Counter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdPushConstants(layout core1_0.PipelineLayout, stageFlags core1_0.ShaderStageFlags, offset int, valueBytes []byte) {
	if layout == nil {
		panic("layout cannot be nil")
	}

	alloc := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(alloc)

	valueBytesPtr := alloc.CBytes(valueBytes)

	c.DeviceDriver.VkCmdPushConstants(c.CommandBufferHandle, layout.Handle(), driver.VkShaderStageFlags(stageFlags), driver.Uint32(offset), driver.Uint32(len(valueBytes)), valueBytesPtr)
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetViewport(viewports []core1_0.Viewport) {
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
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetScissor(scissors []core1_0.Rect2D) {
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
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdNextSubpass(contents core1_0.SubpassContents) {
	c.DeviceDriver.VkCmdNextSubpass(c.CommandBufferHandle, driver.VkSubpassContents(contents))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdWaitEvents(events []core1_0.Event, srcStageMask core1_0.PipelineStageFlags, dstStageMask core1_0.PipelineStageFlags, memoryBarriers []core1_0.MemoryBarrier, bufferMemoryBarriers []core1_0.BufferMemoryBarrier, imageMemoryBarriers []core1_0.ImageMemoryBarrier) error {
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
			if events[i] == nil {
				panic(fmt.Sprintf("element %d of the events slice was nil", i))
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

	c.DeviceDriver.VkCmdWaitEvents(c.CommandBufferHandle, driver.Uint32(eventCount), (*driver.VkEvent)(unsafe.Pointer(eventPtr)), driver.VkPipelineStageFlags(srcStageMask), driver.VkPipelineStageFlags(dstStageMask), driver.Uint32(barrierCount), (*driver.VkMemoryBarrier)(unsafe.Pointer(barrierPtr)), driver.Uint32(bufferBarrierCount), (*driver.VkBufferMemoryBarrier)(unsafe.Pointer(bufferBarrierPtr)), driver.Uint32(imageBarrierCount), (*driver.VkImageMemoryBarrier)(unsafe.Pointer(imageBarrierPtr)))
	c.Counter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdSetEvent(event core1_0.Event, stageMask core1_0.PipelineStageFlags) {
	if event == nil {
		panic("event cannot be nil")
	}

	c.DeviceDriver.VkCmdSetEvent(c.CommandBufferHandle, event.Handle(), driver.VkPipelineStageFlags(stageMask))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdClearColorImage(image core1_0.Image, imageLayout core1_0.ImageLayout, color core1_0.ClearColorValue, ranges []core1_0.ImageSubresourceRange) {
	if image == nil {
		panic("image cannot be nil")
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

	c.DeviceDriver.VkCmdClearColorImage(c.CommandBufferHandle, image.Handle(), driver.VkImageLayout(imageLayout), (*driver.VkClearColorValue)(pColor), driver.Uint32(rangeCount), (*driver.VkImageSubresourceRange)(unsafe.Pointer(pRanges)))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdResetQueryPool(queryPool core1_0.QueryPool, startQuery, queryCount int) {
	if queryPool == nil {
		panic("queryPool cannot be nil")
	}

	c.DeviceDriver.VkCmdResetQueryPool(c.CommandBufferHandle, queryPool.Handle(), driver.Uint32(startQuery), driver.Uint32(queryCount))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdBeginQuery(queryPool core1_0.QueryPool, query int, flags core1_0.QueryControlFlags) {
	if queryPool == nil {
		panic("queryPool cannot be nil")
	}

	c.DeviceDriver.VkCmdBeginQuery(c.CommandBufferHandle, queryPool.Handle(), driver.Uint32(query), driver.VkQueryControlFlags(flags))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdEndQuery(queryPool core1_0.QueryPool, query int) {
	if queryPool == nil {
		panic("queryPool cannot be nil")
	}

	c.DeviceDriver.VkCmdEndQuery(c.CommandBufferHandle, queryPool.Handle(), driver.Uint32(query))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdCopyQueryPoolResults(queryPool core1_0.QueryPool, firstQuery, queryCount int, dstBuffer core1_0.Buffer, dstOffset, stride int, flags core1_0.QueryResultFlags) {
	if queryPool == nil {
		panic("queryPool cannot be nil")
	}
	if dstBuffer == nil {
		panic("dstBuffer cannot be nil")
	}
	c.DeviceDriver.VkCmdCopyQueryPoolResults(c.CommandBufferHandle, queryPool.Handle(), driver.Uint32(firstQuery), driver.Uint32(queryCount), dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(stride), driver.VkQueryResultFlags(flags))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdExecuteCommands(commandBuffers []core1_0.CommandBuffer) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	bufferCount := len(commandBuffers)
	commandBufferPtr := (*C.VkCommandBuffer)(arena.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))
	commandBufferSlice := ([]C.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, bufferCount))

	var addToDrawCount int
	var addToDispatchCount int
	for i := 0; i < bufferCount; i++ {
		if commandBuffers[i] == nil {
			panic(fmt.Sprintf("element %d of the commandBuffers slice was nil", i))
		}
		commandBufferSlice[i] = C.VkCommandBuffer(unsafe.Pointer(commandBuffers[i].Handle()))
		addToDrawCount += commandBuffers[i].DrawsRecorded()
		addToDispatchCount += commandBuffers[i].DispatchesRecorded()
	}

	c.DeviceDriver.VkCmdExecuteCommands(c.CommandBufferHandle, driver.Uint32(bufferCount), (*driver.VkCommandBuffer)(unsafe.Pointer(commandBufferPtr)))
	c.Counter.CommandCount++
	c.Counter.DrawCallCount += addToDrawCount
	c.Counter.DispatchCount += addToDispatchCount
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
	c.Counter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdClearDepthStencilImage(image core1_0.Image, imageLayout core1_0.ImageLayout, depthStencil *core1_0.ClearValueDepthStencil, ranges []core1_0.ImageSubresourceRange) {
	if image == nil {
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

	c.DeviceDriver.VkCmdClearDepthStencilImage(c.CommandBufferHandle, image.Handle(), driver.VkImageLayout(imageLayout), (*driver.VkClearDepthStencilValue)(unsafe.Pointer(depthStencilPtr)), driver.Uint32(rangeCount), (*driver.VkImageSubresourceRange)(unsafe.Pointer(rangePtr)))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdCopyImageToBuffer(srcImage core1_0.Image, srcImageLayout core1_0.ImageLayout, dstBuffer core1_0.Buffer, regions []core1_0.BufferImageCopy) error {
	if srcImage == nil {
		panic("srcImage cannot be nil")
	}
	if dstBuffer == nil {
		panic("dstBuffer cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionPtr, err := common.AllocSlice[C.VkBufferImageCopy, core1_0.BufferImageCopy](arena, regions)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdCopyImageToBuffer(c.CommandBufferHandle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstBuffer.Handle(), driver.Uint32(regionCount), (*driver.VkBufferImageCopy)(unsafe.Pointer(regionPtr)))
	c.Counter.CommandCount++
	return nil
}

func (c *VulkanCommandBuffer) CmdDispatch(groupCountX, groupCountY, groupCountZ int) {
	c.DeviceDriver.VkCmdDispatch(c.CommandBufferHandle, driver.Uint32(groupCountX), driver.Uint32(groupCountY), driver.Uint32(groupCountZ))
	c.Counter.CommandCount++
	c.Counter.DispatchCount++
}

func (c *VulkanCommandBuffer) CmdDispatchIndirect(buffer core1_0.Buffer, offset int) {
	if buffer == nil {
		panic("buffer cannot be nil")
	}
	c.DeviceDriver.VkCmdDispatchIndirect(c.CommandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset))
	c.Counter.CommandCount++
	c.Counter.DispatchCount++
}

func (c *VulkanCommandBuffer) CmdDrawIndexedIndirect(buffer core1_0.Buffer, offset int, drawCount, stride int) {
	if buffer == nil {
		panic("buffer cannot be nil")
	}
	c.DeviceDriver.VkCmdDrawIndexedIndirect(c.CommandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset), driver.Uint32(drawCount), driver.Uint32(stride))
	c.Counter.CommandCount++
	c.Counter.DrawCallCount++
}

func (c *VulkanCommandBuffer) CmdDrawIndirect(buffer core1_0.Buffer, offset int, drawCount, stride int) {
	if buffer == nil {
		panic("buffer cannot be nil")
	}
	c.DeviceDriver.VkCmdDrawIndirect(c.CommandBufferHandle, buffer.Handle(), driver.VkDeviceSize(offset), driver.Uint32(drawCount), driver.Uint32(stride))
	c.Counter.CommandCount++
	c.Counter.DrawCallCount++
}

func (c *VulkanCommandBuffer) CmdFillBuffer(dstBuffer core1_0.Buffer, dstOffset int, size int, data uint32) {
	if dstBuffer == nil {
		panic("dstBuffer cannot be nil")
	}
	c.DeviceDriver.VkCmdFillBuffer(c.CommandBufferHandle, dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(size), driver.Uint32(data))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdResetEvent(event core1_0.Event, stageMask core1_0.PipelineStageFlags) {
	if event == nil {
		panic("event cannot be nil")
	}
	c.DeviceDriver.VkCmdResetEvent(c.CommandBufferHandle, event.Handle(), driver.VkPipelineStageFlags(stageMask))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdResolveImage(srcImage core1_0.Image, srcImageLayout core1_0.ImageLayout, dstImage core1_0.Image, dstImageLayout core1_0.ImageLayout, regions []core1_0.ImageResolve) error {
	if srcImage == nil {
		panic("srcImage cannot be nil")
	}
	if dstImage == nil {
		panic("dstImage cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	regionCount := len(regions)
	regionsPtr, err := common.AllocSlice[C.VkImageResolve, core1_0.ImageResolve](arena, regions)
	if err != nil {
		return err
	}

	c.DeviceDriver.VkCmdResolveImage(c.CommandBufferHandle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstImage.Handle(), driver.VkImageLayout(dstImageLayout), driver.Uint32(regionCount), (*driver.VkImageResolve)(unsafe.Pointer(regionsPtr)))
	c.Counter.CommandCount++
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
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetDepthBias(depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32) {
	c.DeviceDriver.VkCmdSetDepthBias(c.CommandBufferHandle, driver.Float(depthBiasConstantFactor), driver.Float(depthBiasClamp), driver.Float(depthBiasSlopeFactor))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetDepthBounds(min, max float32) {
	c.DeviceDriver.VkCmdSetDepthBounds(c.CommandBufferHandle, driver.Float(min), driver.Float(max))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetLineWidth(lineWidth float32) {
	c.DeviceDriver.VkCmdSetLineWidth(c.CommandBufferHandle, driver.Float(lineWidth))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetStencilCompareMask(faceMask core1_0.StencilFaceFlags, compareMask uint32) {
	c.DeviceDriver.VkCmdSetStencilCompareMask(c.CommandBufferHandle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(compareMask))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetStencilReference(faceMask core1_0.StencilFaceFlags, reference uint32) {
	c.DeviceDriver.VkCmdSetStencilReference(c.CommandBufferHandle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(reference))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdSetStencilWriteMask(faceMask core1_0.StencilFaceFlags, writeMask uint32) {
	c.DeviceDriver.VkCmdSetStencilWriteMask(c.CommandBufferHandle, driver.VkStencilFaceFlags(faceMask), driver.Uint32(writeMask))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdUpdateBuffer(dstBuffer core1_0.Buffer, dstOffset int, dataSize int, data []byte) {
	if dstBuffer == nil {
		panic("dstBuffer cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	size := len(data)
	dataPtr := arena.Malloc(size)
	dataSlice := ([]byte)(unsafe.Slice((*byte)(dataPtr), size))
	copy(dataSlice, data)

	c.DeviceDriver.VkCmdUpdateBuffer(c.CommandBufferHandle, dstBuffer.Handle(), driver.VkDeviceSize(dstOffset), driver.VkDeviceSize(dataSize), dataPtr)
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) CmdWriteTimestamp(pipelineStage core1_0.PipelineStageFlags, queryPool core1_0.QueryPool, query int) {
	if queryPool == nil {
		panic("queryPool cannot be nil")
	}

	c.DeviceDriver.VkCmdWriteTimestamp(c.CommandBufferHandle, driver.VkPipelineStageFlags(pipelineStage), queryPool.Handle(), driver.Uint32(query))
	c.Counter.CommandCount++
}

func (c *VulkanCommandBuffer) Reset(flags core1_0.CommandBufferResetFlags) (common.VkResult, error) {
	return c.DeviceDriver.VkResetCommandBuffer(c.CommandBufferHandle, driver.VkCommandBufferResetFlags(flags))
}

func (c *VulkanCommandBuffer) Free() {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	vkCommandBuffer := (*driver.VkCommandBuffer)(arena.Malloc(int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))
	commandBufferSlice := ([]driver.VkCommandBuffer)(unsafe.Slice(vkCommandBuffer, 1))
	commandBufferSlice[0] = c.CommandBufferHandle

	c.DeviceDriver.VkFreeCommandBuffers(c.Device, c.CommandPool, 1, vkCommandBuffer)
}
