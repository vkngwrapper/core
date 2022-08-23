package core1_0

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/driver"
	"time"
	"unsafe"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/core1_0_mocks.go -package mocks

// Buffer represents a linear array of data, which is used for various purposes by binding it
// to a graphics or compute pipeline.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBuffer.html
type Buffer interface {
	// Handle is the internal Vulkan object handle for this Buffer
	Handle() driver.VkBuffer
	// DeviceHandle is the internal Vulkan object handle for the Device this Buffer belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this Buffer
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Buffer. If it is
	// at least vulkan 1.1, core1_1.PromoteBuffer can be used to promote this to a
	// core1_1.Buffer, etc.
	APIVersion() common.APIVersion

	// Destroy deletes this buffer and underlying structures from the device. **Warning**
	// after destruction, this object will still exist, but the Vulkan object handle
	// that backs it will be invalid. Do not call further methods on this object.
	//
	// callbacks - An set of allocation callbacks to control the memory free behavior of this command
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyBuffer.html
	Destroy(callbacks *driver.AllocationCallbacks)
	// MemoryRequirements returns the memory requirements for this Buffer.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements.html
	MemoryRequirements() *MemoryRequirements
	// BindBufferMemory binds DeviceMemory to this Buffer
	//
	// memory - A DeviceMemory object describing the device memory to attach
	//
	// offset - The start offset of the region of memory which is to be bound to the buffer.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory.html
	BindBufferMemory(memory DeviceMemory, offset int) (common.VkResult, error)
}

// BufferView represents a contiguous range of a buffer and a specific format to be used to
// interpret the data.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferView.html
type BufferView interface {
	// Handle is the internal Vulkan object handle for this BufferView
	Handle() driver.VkBufferView
	// DeviceHandle is the internal Vulkan object handle for the Device this BufferView belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the vulkan wrapper driver used by this BufferView
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Buffer. If it is
	// at least vulkan 1.1, core1_1.PromoteBufferView can be used to promote this to a
	// core1_1.BufferView, etc.
	APIVersion() common.APIVersion

	// Destroy deletes this buffer and the underlying structures from the device. **Warning**
	// after destruction, this object will continue to exist, but the Vulkan object handle
	// that backs it will be invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyBufferView.html
	Destroy(callbacks *driver.AllocationCallbacks)
}

// CommandBuffer is an object used to record commands which can be subsequently submitted to
// a device queue for execution.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html
type CommandBuffer interface {
	// Handle is the internal Vulkan object handle for this CommandBuffer
	Handle() driver.VkCommandBuffer
	// Driver is the vulkan wrapper driver used by this CommandBuffer
	Driver() driver.Driver
	// DeviceHandle is the internal Vulkan object handle for the Device this CommandBuffer belongs to
	DeviceHandle() driver.VkDevice
	// CommandPoolHandle is the internal Vulkan object handle for the CommandPool used to allocate
	// this CommandBuffer
	CommandPoolHandle() driver.VkCommandPool
	// APIVersion is the maximum Vulkan API version supported by this CommandBuffer. If it is at
	// least vulkan 1.1, core1_1.PromoteCommandBuffer can be used to promote this to a core1_1.CommandBuffer,
	// etc.
	APIVersion() common.APIVersion

	// Free frees this command buffer and usually returns the underlying memory to the CommandPool
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFreeCommandBuffers.html
	Free()
	// Begin starts recording on this CommandBuffer
	//
	// o - Defines additional information about how the CommandBuffer begins recording
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkBeginCommandBuffer.html
	Begin(o CommandBufferBeginInfo) (common.VkResult, error)
	// End finishes recording on this command buffer
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEndCommandBuffer.html
	End() (common.VkResult, error)
	// Reset returns this CommandBuffer to its initial state
	//
	// flags - Options controlling the reset operation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetCommandBuffer.html
	Reset(flags CommandBufferResetFlags) (common.VkResult, error)
	// CommandsRecorded returns the number of commands recorded to this CommandBuffer since the last time
	// Begin was called
	CommandsRecorded() int
	// DrawsRecorded returns the number of draw commands recorded to this CommandBuffer since the last time
	// Begin was called
	DrawsRecorded() int
	// DispatchesRecorded returns the number of dispatch commands recorded to this CommandBuffer since
	// the last time Begin was called
	DispatchesRecorded() int

	// CmdBeginRenderPass begins a new RenderPass
	//
	// contents - Specifies how the commands in the first subpass will be provided
	//
	// o - Specifies the RenderPass to begin an instance of, and the Framebuffer the instance uses
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html
	CmdBeginRenderPass(contents SubpassContents, o RenderPassBeginInfo) error
	// CmdEndRenderPass ends the current RenderPass
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass.html
	CmdEndRenderPass()
	// CmdBindPipeline binds a pipeline object to this CommandBuffer
	//
	// bindPoint - Specifies to which bind point the Pipeline is bound
	//
	// pipeline - The Pipeline to be bound
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBindPipeline.html
	CmdBindPipeline(bindPoint PipelineBindPoint, pipeline Pipeline)
	// CmdDraw draws primitives without indexing the vertices
	//
	// vertexCount - The number of vertices to draw
	//
	// instanceCount - The number of instances to draw
	//
	// firstVertex - The index of the first vertex to draw
	//
	// firstInstance - The instance ID of the first instance to draw
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDraw.html
	CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32)
	// CmdDrawIndexed draws primitives with indexed vertices
	//
	// indexCount - The number of vertices to draw
	//
	// instanceCount - The number of instances to draw
	//
	// firstIndex - The base index within the index Buffer
	//
	// vertexOffset - The value added to the vertex index before indexing into the vertex Buffer
	//
	// firstInstance - The instance ID of the first instance to draw
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexed.html
	CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32)
	// CmdBindVertexBuffers binds vertex Buffers to this CommandBuffer
	//
	// firstBinding - The index of the first input binding whose state is updated by the command
	//
	// buffers - A slice of Buffer objects
	//
	// bufferOffsets - A slice of Buffer offsets
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers.html
	CmdBindVertexBuffers(firstBinding int, buffers []Buffer, bufferOffsets []int)
	// CmdBindIndexBuffer binds an index Buffer to this CommandBuffer
	//
	// buffer - The Buffer being bound
	//
	// offset - The starting offset in bytes within Buffer, used in index Buffer address calculations
	//
	// indexType - Specifies the size of the indices
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer.html
	CmdBindIndexBuffer(buffer Buffer, offset int, indexType IndexType)
	// CmdCopyBuffer copies data between Buffer regions
	//
	// srcBuffer - The source Buffer
	//
	// dstBuffer - The destination Buffer
	//
	// copyRegions - A slice of structures specifying the regions to copy
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBuffer.html
	CmdCopyBuffer(srcBuffer Buffer, dstBuffer Buffer, copyRegions []BufferCopy) error
	// CmdBindDescriptorSets binds DescriptorSets to this CommandBuffer
	//
	// bindPoint - Indicates the type of the pipeline that will use the descriptors
	//
	// layout - A PipelineLayout object used to program the bindings
	//
	// firstSet - The set number of the first DescriptorSet to be bound
	//
	// sets - A slice of DescriptorSet objects describing the DescriptorSets to bind
	//
	// dynamicOffsets - A slice of values specifying dynamic offsets
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html
	CmdBindDescriptorSets(bindPoint PipelineBindPoint, layout PipelineLayout, firstSet int, sets []DescriptorSet, dynamicOffsets []int)
	// CmdPipelineBarrier inserts a memory dependency into the recorded commands
	//
	// srcStageMask - Specifies the source stages
	//
	// dstStageMask - Specifies the destination stages
	//
	// dependencies - Specifies how execution and memory dependencies are formed
	//
	// memoryBarriers - A slice of MemoryBarrier structures
	//
	// bufferMemoryBarriers - A slice of BufferMemoryBarrier structures
	//
	// imageMemoryBarriers - A slice of ImageMemoryBarrier structures
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier.html
	CmdPipelineBarrier(srcStageMask, dstStageMask PipelineStageFlags, dependencies DependencyFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) error
	// CmdCopyBufferToImage copies data from a Buffer to an Image
	//
	// buffer - The source buffer
	//
	// image - The destination Image
	//
	// layout - The layout of the destination Image subresources for the copy
	//
	// regions - A slice of BufferImageCopy structures specifying the regions to copy
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage.html
	CmdCopyBufferToImage(buffer Buffer, image Image, layout ImageLayout, regions []BufferImageCopy) error
	// CmdBlitImage copies regions of an Image, potentially performing format conversion
	//
	// sourceImage - The source Image
	//
	// sourceImageLayout - The layout of the source Image subresources for the blit
	//
	// destinationImage - The destination Image
	//
	// destinationImageLayout - The layout of the destination Image subresources for the blit
	//
	// regions - A slice of ImageBlit structures specifying the regions to blit
	//
	// filter - Specifies the filter to apply if the blits require scaling
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBlitImage.html
	CmdBlitImage(sourceImage Image, sourceImageLayout ImageLayout, destinationImage Image, destinationImageLayout ImageLayout, regions []ImageBlit, filter Filter) error
	// CmdPushConstants updates the values of push constants
	//
	// layout - The pipeline layout used to program the push constant updates
	//
	// stageFlags - Specifies the shader stages that will use the push constants in the updated range
	//
	// offset - The start offset of the push constant range to update, in units of bytes
	//
	// valueBytes - A slice of bytes containing the new push constant values
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdPushConstants.html
	CmdPushConstants(layout PipelineLayout, stageFlags ShaderStageFlags, offset int, valueBytes []byte)
	// CmdSetViewport sets the viewport dynamically for a CommandBuffer
	//
	// viewports - A slice of Viewport structures specifying viewport parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewport.html
	CmdSetViewport(viewports []Viewport)
	// CmdSetScissor sets scissor rectangles dynamically for a CommandBuffer
	//
	// scissors - A slice of Rect2D structures specifying scissor rectangles
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissor.html
	CmdSetScissor(scissors []Rect2D)
	// CmdCopyImage copies data between Images
	//
	// srcImage - The source Image
	//
	// srcImageLayout - The current layout of the source Image subresource
	//
	// dstImage - The destination Image
	//
	// dstImageLayout - The current layout of the destination Image subresource
	//
	// regions - A slice of ImageCopy structures specifying the regions to copy
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImage.html
	CmdCopyImage(srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regions []ImageCopy) error
	// CmdNextSubpass transitions to the next subpass of a RenderPass
	//
	// contents - Specifies how the commands in the next subpass will be provided
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass.html
	CmdNextSubpass(contents SubpassContents)
	// CmdWaitEvents waits for one or more events and inserts a set of memory
	//
	// events - A slice of Event objects to wait on
	//
	// srcStageMask - Specifies the source stage mask
	//
	// dstStageMask - Specifies the destination stage mask
	//
	// memoryBarriers - A slice of MemoryBarrier structures
	//
	// bufferMemoryBarriers - A slice of BufferMemoryBarrier structures
	//
	// imageMemoryBarriers - A slice of ImageMemoryBarrier structures
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html
	CmdWaitEvents(events []Event, srcStageMask PipelineStageFlags, dstStageMask PipelineStageFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) error
	// CmdSetEvent sets an Event object to the signaled state
	//
	// event - The Event that will be signaled
	//
	// stageMask - Specifies teh source stage mask used to determine the first synchronization scope
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent.html
	CmdSetEvent(event Event, stageMask PipelineStageFlags)
	// CmdClearColorImage clears regions of a color Image
	//
	// image - The Image to be cleared
	//
	// imageLayout - Specifies the current layout of the Image subresource ranges to be cleared
	//
	// color - A ClearColorValue containing the values that the Image subresource ranges will be cleared to
	//
	// ranges - A slice of ImageSubresourceRange structures describing a range of mipmap levels, array layers,
	// and aspects to be cleared.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdClearColorImage.html
	CmdClearColorImage(image Image, imageLayout ImageLayout, color ClearColorValue, ranges []ImageSubresourceRange)
	// CmdResetQueryPool resets queries in a QueryPool
	//
	// queryPool - The QueryPool managing the queries being reset
	//
	// startQuery - The initial query index to reset
	//
	// queryCount - The number of queries to reset
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdResetQueryPool.html
	CmdResetQueryPool(queryPool QueryPool, startQuery, queryCount int)
	// CmdBeginQuery begins a query
	//
	// queryPool - The QueryPool that will manage the results of the query
	//
	// query - The query index within the QueryPool that will contain the results
	//
	// flags - Specifies constraints on the types of queries that can be performed
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQuery.html
	CmdBeginQuery(queryPool QueryPool, query int, flags QueryControlFlags)
	// CmdEndQuery ends a query
	//
	// queryPool - The QueryPool that is managing the results of the query
	//
	// query - The query index within the QueryPool where the result is stored
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdEndQuery.html
	CmdEndQuery(queryPool QueryPool, query int)
	// CmdCopyQueryPoolResults copies the results of queries in a QueryPool to a Buffer object
	//
	// queryPool - The QueryPool managing the queries containing the desired results
	//
	// firstQuery - The initial query index
	//
	// queryCount - The number of queries
	//
	// dstBuffer - A Buffer object that will receive the results of the copy command
	//
	// dstOffset - An offset into the destination Buffer
	//
	// stride - The stride in bytes between the results for individual queries within the destination Buffer
	//
	// flags - Specifies how and when results are returned
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdCopyQueryPoolResults.html
	CmdCopyQueryPoolResults(queryPool QueryPool, firstQuery, queryCount int, dstBuffer Buffer, dstOffset, stride int, flags QueryResultFlags)
	// CmdExecuteCommands executes a secondary CommandBuffer from a primary CommandBuffer
	//
	// commandBuffers - A slice of CommandBuffer objects, which are recorded to execute in the primary CommandBuffer
	// in the order they are listed in the slice
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteCommands.html
	CmdExecuteCommands(commandBuffers []CommandBuffer)
	// CmdClearAttachments clears regions within bound Framebuffer attachments
	//
	// attachments - A slice of ClearAttachment structures defining the attachments to clear and the clear values to use.
	//
	// rects - A slice of ClearRect structures defining regions within each selected attachment to clear
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdClearAttachments.html
	CmdClearAttachments(attachments []ClearAttachment, rects []ClearRect) error
	// CmdClearDepthStencilImage fills regions of a combined depth/stencil image
	//
	// image - The Image to be cleared
	//
	// imageLayout - Specifies the current layout of the Image subresource ranges to be cleared
	//
	// depthStencil - Contains the values that the depth and stencil images will be cleared to
	//
	// ranges - A slice of ImageSubrsourceRange structures describing a range of mipmap levels, array layers,
	// and aspects to be cleared
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdClearDepthStencilImage.html
	CmdClearDepthStencilImage(image Image, imageLayout ImageLayout, depthStencil *ClearValueDepthStencil, ranges []ImageSubresourceRange)
	// CmdCopyImageToBuffer copies image data into a buffer
	//
	// srcImage - The source Image
	//
	// srcImageLayout - The layout of the source Image subresources for the copy
	//
	// dstBuffer - The desination Buffer
	//
	// regions - A slice of BufferImageCopy structures specifying the regions to copy
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImageToBuffer.html
	CmdCopyImageToBuffer(srcImage Image, srcImageLayout ImageLayout, dstBuffer Buffer, regions []BufferImageCopy) error
	// CmdDispatch dispatches compute work items
	//
	// groupCountX - the number of local workgroups to dispatch in the X dimension
	//
	// groupCountY - the number of local workgroups to dispatch in the Y dimension
	//
	// groupCountZ - the number of local workgroups to dispatch in the Z dimension
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDispatch.html
	CmdDispatch(groupCountX, groupCountY, groupCountZ int)
	// CmdDispatchIndirect dispatches compute work items with indirect parameters
	//
	// buffer - The Buffer containing dispatch parameters
	//
	// offset - The byte offset into the Buffer where parameters begin
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchIndirect.html
	CmdDispatchIndirect(buffer Buffer, offset int)
	// CmdDrawIndexedIndirect draws primitives with indirect parameters and indexed vertices
	//
	// buffer - The Buffer containing draw parameters
	//
	// offset - The byte offset into the Buffer where parameters begin
	//
	// drawCount - The number of draws to execute, which can be zero
	//
	// stride - The byte stride between successive sets of draw parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirect.html
	CmdDrawIndexedIndirect(buffer Buffer, offset int, drawCount, stride int)
	// CmdDrawIndirect draws primitives with indirect parameters
	//
	// buffer - The buffer containing draw parameters
	//
	// offset - The byte offset into the Buffer where parameters begin
	//
	// drawCount - The number of draws to execute, which can be zero
	//
	// stride - The byte stride between successive sets of draw parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirect.html
	CmdDrawIndirect(buffer Buffer, offset int, drawCount, stride int)
	// CmdFillBuffer fills a region of a buffer with a fixed value
	//
	// dstBuffer - The Buffer to be filled
	//
	// dstOffset - The byte offset into the Buffer at which to start filling, must be a multiple of 4
	//
	// size - The number of bytes to fill
	//
	// data - The 4-byte word written repeatedly to the Buffer to fill size bytes of data.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdFillBuffer.html
	CmdFillBuffer(dstBuffer Buffer, dstOffset int, size int, data uint32)
	// CmdResetEvent resets an Event object to non-signaled state
	//
	// event - The Event that will be unsignaled
	//
	// stageMask - Specifies the source stage mask used to determine when the Event is unsignaled
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdResetEvent.html
	CmdResetEvent(event Event, stageMask PipelineStageFlags)
	// CmdResolveImage resolves regions of an Image
	//
	// srcImage - The source Image
	//
	// srcImageLayout - The layout of the source Image subresources for the resolve
	//
	// dstImage - The destination Image
	//
	// dstImageLayout - The layout of the destination Image subresources for the resolve
	//
	// regions - A slice of ImageResolve structure specifying the regions to resolve
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdResolveImage.html
	CmdResolveImage(srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regions []ImageResolve) error
	// CmdSetBlendConstants sets the values of the blend constants
	//
	// blendConstants - An array of four values specifying the R, G, B, and A components of the blend
	// color used in blending, depending on the blend factor
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetBlendConstants.html
	CmdSetBlendConstants(blendConstants [4]float32)
	// CmdSetDepthBias sets depth bias factors and clamp dynamically for the CommandBuffer
	//
	// depthBiasConstantFactor - The scalar factor controlling the constant depth value added to each fragment
	//
	// depthBiasClamp - The maximum (or minimum) depth bias of a fragment
	//
	// depthBiasSlopeFactor - The scalar factor applied to a fragment's slope in depth bias calculations
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBias.html
	CmdSetDepthBias(depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32)
	// CmdSetDepthBounds sets depth bounds range dynamically for the CommandBuffer
	//
	// min - The minimum depth bound
	//
	// max - The maximum depth bound
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBounds.html
	CmdSetDepthBounds(min, max float32)
	// CmdSetLineWidth sets line width dynamically for the CommandBuffer
	//
	// lineWidth - The width of rasterized line segments
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineWidth.html
	CmdSetLineWidth(lineWidth float32)
	// CmdSetStencilCompareMask sets the stencil compare mask dynamically for the CommandBuffer
	//
	// faceMask - Specifies the set of stencil state for which to update the compare mask
	//
	// compareMask - The new value to use as the stencil compare mask
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilCompareMask.html
	CmdSetStencilCompareMask(faceMask StencilFaceFlags, compareMask uint32)
	// CmdSetStencilReference sets stencil reference value dynamically for the CommandBuffer
	//
	// faceMask - Specifies the set of stencil state for which to update the reference value
	//
	// reference - The new value to use as the stencil reference value
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilReference.html
	CmdSetStencilReference(faceMask StencilFaceFlags, reference uint32)
	// CmdSetStencilWriteMask sets the stencil write mask dynamically for the CommandBuffer
	//
	// faceMask - Specifies the set of stencil state for which to update the write mask
	//
	// reference - The new value to use as the stencil write mask
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilWriteMask.html
	CmdSetStencilWriteMask(faceMask StencilFaceFlags, writeMask uint32)
	// CmdUpdateBuffer updates a buffer's contents from host memory
	//
	// dstBuffer - The Buffer to be updated
	//
	// dstOffset - The byte offset into the Buffer to start updating, must be a multiple of 4
	//
	// dataSize - The number of bytes to update, must be a multiple of 4
	//
	// data - The source data for the buffer update, must be at least dataSize bytes in size
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdUpdateBuffer.html
	CmdUpdateBuffer(dstBuffer Buffer, dstOffset int, dataSize int, data []byte)
	// CmdWriteTimestamp writes a device timestamp into a query object
	//
	// pipelineStage - Specifies a stage of the pipeline
	//
	// queryPool - The QueryPool that will manage the timestamp
	//
	// query - The query within the QueryPool that will contain the timestamp
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp.html
	CmdWriteTimestamp(pipelineStage PipelineStageFlags, queryPool QueryPool, query int)
}

// CommandPool is an opaque object that CommandBuffer memory is allocated from
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html
type CommandPool interface {
	// Handle is the internal Vulkan object handle for this CommandPool
	Handle() driver.VkCommandPool
	// DeviceHandle is the internal Vulkan object handle for the Device this CommandPool belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this CommandPool
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this CommandPool. If it is
	// at least Vulkan 1.1, core1_1.PromoteCommandPool can be used to promote this to a core1_1.CommandPool,
	// etc.
	APIVersion() common.APIVersion

	// Destroy destroys the CommandPool object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs
	// it will be invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyCommandPool.html
	Destroy(callbacks *driver.AllocationCallbacks)
	// Reset resets the CommandPool, recycling all the resources from all the CommandBuffer objects
	// allocated from the CommandPool back to the CommandPool.  All CommandBuffer objects that
	// have been allocated from the CommandPool are put in the initial state.
	//
	// flags - Controls the reset operation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetCommandPool.html
	Reset(flags CommandPoolResetFlags) (common.VkResult, error)
}

// DescriptorPool maintains a pool of descriptors, from which DescriptorSet objects are allocated.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html
type DescriptorPool interface {
	// Handle is the internal Vulkan object handle for this DescriptorPool
	Handle() driver.VkDescriptorPool
	// DeviceHandle is the internal Vulkan object handle for the Device this DescriptorPool belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this DescriptorPool
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this DescriptorPool. If it is at
	// least Vulkan 1.1, core1_1.PromoteDescriptorPool can be used to promote this to a core1_1.DescriptorPool,
	// etc.
	APIVersion() common.APIVersion

	// Destroy destroys the DescriptorPool object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorPool.html
	Destroy(callbacks *driver.AllocationCallbacks)
	// Reset resets the DescriptorPool and recycles all of the resources from all of the DescriptorSet
	// objects allocated from the DescriptorPool back to the DescriptorPool, and the DescriptorSet
	// objects are implicitly freed.
	//
	// flags - Reserved (always 0)
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetDescriptorPool.html
	Reset(flags DescriptorPoolResetFlags) (common.VkResult, error)
}

// DescriptorSet is an opaque object allocated from a DescriptorPool
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetDescriptorPool.html
type DescriptorSet interface {
	// Handle is the internal Vulkan object handle for this DescriptorPool
	Handle() driver.VkDescriptorSet
	// DescriptorPoolHandle is the internal Vulkan object handle for the DescriptorPool this DescriptorSet
	// was allocated from
	DescriptorPoolHandle() driver.VkDescriptorPool
	// DeviceHandle is the internal Vulkan object handle for the Device this DescriptorSet belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this DescriptorSet
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this DescriptorSet. If it is at least
	// Vulkan 1.1, core1_1.PromoteDescriptorSet can be used to promote this to a core1_1.DescriptorSet,
	// etc.
	APIVersion() common.APIVersion

	// Free frees this DescriptorSet
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFreeDescriptorSets.html
	Free() (common.VkResult, error)
}

// DescriptorSetLayout is a group of zero or more descriptor bindings definitions.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html
type DescriptorSetLayout interface {
	// Handle is the internal Vulkan object handle for this DescriptorSetLayout
	Handle() driver.VkDescriptorSetLayout
	// DeviceHandle is the internal Vulkan object handle for the Device this DescriptorSetLayout belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this DescriptorSetLayout
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this DescriptorSetLayout. If it is at
	// least Vulkan 1.1, core1_1.PromoteDescriptorSetLayout can be used to promote this to a
	// core1_1.DescriptorSetLayout, etc.
	APIVersion() common.APIVersion

	// Destroy destroys the DescriptorSetLayout object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorSetLayout.html
	Destroy(callbacks *driver.AllocationCallbacks)
}

// DeviceMemory represents a block of memory on the device
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html
type DeviceMemory interface {
	// Handle is the internal Vulkan object handle for this DeviceMemory
	Handle() driver.VkDeviceMemory
	// DeviceHandle is the internal Vulkan object handle for the Device this DeviceMemory belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this DeviceMemory
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this DeviceMemory. If it is at least
	// Vulkan 1.1, core1_1.PromoteDeviceMemory can be used to promote this to a core1_1.DeviceMemory,
	// etc.
	APIVersion() common.APIVersion

	// Map maps a memory object into application address space
	//
	// offset - A zero-based byte offset from the beginning of the memory object
	//
	// size - The size of the memory range to map, or -1 to map from offset to the end of the
	// allocation
	//
	// flags - Reserved (always 0)
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkMapMemory.html
	Map(offset int, size int, flags MemoryMapFlags) (unsafe.Pointer, common.VkResult, error)
	// Unmap unmaps a previously-mapped memory object
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUnmapMemory.html
	Unmap()
	// Free frees the DeviceMemory. **Warning** after freeing, this object will continue to exist,
	// but the Vulkan object handle that backs it will be invalid. Do not call further methods on
	// this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFreeMemory.html
	Free(callbacks *driver.AllocationCallbacks)
	// Commitment returns the current number of bytes currently committed to this DeviceMemory
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryCommitment.html
	Commitment() int
	// FlushAll flushes all mapped memory ranges in this DeviceMemory
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFlushMappedMemoryRanges.html
	FlushAll() (common.VkResult, error)
	// InvalidateAll invalidates all mapped memory ranges in this DeviceMemory
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkInvalidateMappedMemoryRanges.html
	InvalidateAll() (common.VkResult, error)
}

// Device represents a logical device on the host
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDevice.html
type Device interface {
	// Handle is the internal Vulkan object handle for this Device
	Handle() driver.VkDevice
	// Driver is the Vulkan wrapper drive used by this Device
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Device. If it is at least
	// Vulkan 1.1, core1_1.PromoteDevice can be used to promote this to a core1_1.Device, etc.
	APIVersion() common.APIVersion

	// IsDeviceExtensionActive will return true if a Device extension with the provided name was
	// activated on Device creation
	//
	// extensionName - The name of the extension to query
	IsDeviceExtensionActive(extensionName string) bool

	// CreateBuffer creates a new Buffer object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Buffer
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateBuffer.html
	CreateBuffer(allocationCallbacks *driver.AllocationCallbacks, o BufferCreateInfo) (Buffer, common.VkResult, error)
	// CreateBufferView creates a new BufferView object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the BufferView
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateBufferView.html
	CreateBufferView(allocationCallbacks *driver.AllocationCallbacks, o BufferViewCreateInfo) (BufferView, common.VkResult, error)
	// CreateCommandPool creates a new CommandPool object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the CommandPool
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateCommandPool.html
	CreateCommandPool(allocationCallbacks *driver.AllocationCallbacks, o CommandPoolCreateInfo) (CommandPool, common.VkResult, error)
	// CreateDescriptorPool creates a new DescriptorPool object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the DescriptorPool
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorPool.html
	CreateDescriptorPool(allocationCallbacks *driver.AllocationCallbacks, o DescriptorPoolCreateInfo) (DescriptorPool, common.VkResult, error)
	// CreateDescriptorSetLayout creates a new DescriptorSetLayout object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the DescriptorSetLayout
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorSetLayout.html
	CreateDescriptorSetLayout(allocationCallbacks *driver.AllocationCallbacks, o DescriptorSetLayoutCreateInfo) (DescriptorSetLayout, common.VkResult, error)
	// CreateEvent creates a new Event object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Event
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateEvent.html
	CreateEvent(allocationCallbacks *driver.AllocationCallbacks, options EventCreateInfo) (Event, common.VkResult, error)
	// CreateFence creates a new Fence object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Fence
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateFence.html
	CreateFence(allocationCallbacks *driver.AllocationCallbacks, o FenceCreateInfo) (Fence, common.VkResult, error)
	// CreateFramebuffer creates a new Framebuffer object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Framebuffer
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateFramebuffer.html
	CreateFramebuffer(allocationCallbacks *driver.AllocationCallbacks, o FramebufferCreateInfo) (Framebuffer, common.VkResult, error)
	// CreateGraphicsPipelines creates a slice of new Pipeline objects which can be used for drawing graphics
	//
	// pipelineCache - A PipelineCache object which can be used to accelerate pipeline creation
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - A slice of GraphicsPipelineCreateInfo structures containing parameters affecting the creation of the Pipeline objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateGraphicsPipelines.html
	CreateGraphicsPipelines(pipelineCache PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []GraphicsPipelineCreateInfo) ([]Pipeline, common.VkResult, error)
	// CreateComputePipelines creates a slice of new Pipeline objects which can be used for dispatching compute workloads
	//
	// pipelineCache - A PipelineCache object which can be used to accelerate pipeline creation
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - A slice of ComputePipelineCreateInfo structures containing parameters affecting the creation of the Pipeline objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateComputePipelines.html
	CreateComputePipelines(pipelineCache PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []ComputePipelineCreateInfo) ([]Pipeline, common.VkResult, error)
	// CreateImage creates a new Image object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Image
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html
	CreateImage(allocationCallbacks *driver.AllocationCallbacks, options ImageCreateInfo) (Image, common.VkResult, error)
	// CreateImageView creates a new ImageView object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the ImageView
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateImageView.html
	CreateImageView(allocationCallbacks *driver.AllocationCallbacks, o ImageViewCreateInfo) (ImageView, common.VkResult, error)
	// CreatePipelineCache creates a new PipelineCache object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the PipelineCache
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreatePipelineCache.html
	CreatePipelineCache(allocationCallbacks *driver.AllocationCallbacks, o PipelineCacheCreateInfo) (PipelineCache, common.VkResult, error)
	// CreatePipelineLayout creates a new PipelineLayout object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the PipelineLayout
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreatePipelineLayout.html
	CreatePipelineLayout(allocationCallbacks *driver.AllocationCallbacks, o PipelineLayoutCreateInfo) (PipelineLayout, common.VkResult, error)
	// CreateQueryPool creates a new QueryPool object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the QueryPool
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateQueryPool.html
	CreateQueryPool(allocationCallbacks *driver.AllocationCallbacks, o QueryPoolCreateInfo) (QueryPool, common.VkResult, error)
	// CreateRenderPass creates a new RenderPass object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the RenderPass
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass.html
	CreateRenderPass(allocationCallbacks *driver.AllocationCallbacks, o RenderPassCreateInfo) (RenderPass, common.VkResult, error)
	// CreateSampler creates a new Sampler object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Sampler
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateSampler.html
	CreateSampler(allocationCallbacks *driver.AllocationCallbacks, o SamplerCreateInfo) (Sampler, common.VkResult, error)
	// CreateSemaphore creates a new Semaphore object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the Semaphore
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateSemaphore.html
	CreateSemaphore(allocationCallbacks *driver.AllocationCallbacks, o SemaphoreCreateInfo) (Semaphore, common.VkResult, error)
	// CreateShaderModule creates a new ShaderModule object
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Parameters affecting the creation of the ShaderModule
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateShaderModule.html
	CreateShaderModule(allocationCallbacks *driver.AllocationCallbacks, o ShaderModuleCreateInfo) (ShaderModule, common.VkResult, error)

	// GetQueue gets a Queue object from the Device
	//
	// queueFamilyIndex - The index of the queue family to which the Queue belongs
	//
	// queueIndex - The index within this queue family of the Queue to retrieve
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetDeviceQueue.html
	GetQueue(queueFamilyIndex int, queueIndex int) Queue
	// AllocateMemory allocates DeviceMemory
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// o - Describes the parameters of the allocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkAllocateMemory.html
	AllocateMemory(allocationCallbacks *driver.AllocationCallbacks, o MemoryAllocateInfo) (DeviceMemory, common.VkResult, error)
	// FreeMemory frees DeviceMemory. **Warning** after freeing, the DeviceMemory object will continue to
	// exist, but the Vulkan object handle that backs it will be invalid. Do not call further methods
	// on the DeviceMemory object.
	//
	// deviceMemory - The DeviceMemory object to be freed
	//
	// allocationCallbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFreeMemory.html
	FreeMemory(deviceMemory DeviceMemory, allocationCallbacks *driver.AllocationCallbacks)

	// AllocateCommandBuffers allocates CommandBuffer objects from an existing CommandPool
	//
	// o - Describes parameters of the allocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkAllocateCommandBuffers.html
	AllocateCommandBuffers(o CommandBufferAllocateInfo) ([]CommandBuffer, common.VkResult, error)
	// FreeCommandBuffers frees CommandBuffer objects. **Warning** after freeing, these objects will
	// continue to exist, but the Vulkan object handles that back them will be invalid. Do not call
	// further methods on these objects.
	//
	// buffers - A slice of CommandBuffer objects to free
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFreeCommandBuffers.html
	FreeCommandBuffers(buffers []CommandBuffer)
	// AllocateDescriptorSets allocates one or more DescriptorSet objects from a DescriptorPool
	//
	// o - Describes the parameters of the allocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkAllocateDescriptorSets.html
	AllocateDescriptorSets(o DescriptorSetAllocateInfo) ([]DescriptorSet, common.VkResult, error)
	// FreeDescriptorSets frees one or more DescriptorSet objects. **Warning** after freeing, these objects
	// will continue to exist, but the Vulkan object handles that back them will be invalid. Do not call
	// further methods on these objects.
	//
	// sets - A slice of DescriptorSet objects to free
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFreeDescriptorSets.html
	FreeDescriptorSets(sets []DescriptorSet) (common.VkResult, error)

	// Destroy destroys a logical Device object.  **Warning** after destruction, this object will continue
	// to exist, but the Vulkan object handle that backs it will be invalid. Do not call further methods
	// on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyDevice.html
	Destroy(callbacks *driver.AllocationCallbacks)
	// WaitIdle waits for the Device to become idle
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDeviceWaitIdle.html
	WaitIdle() (common.VkResult, error)
	// WaitForFences waits for one or more Fence objects to become signaled
	//
	// waitForAll - If true, then the call will wait until all fences in `fences` are signaled. If
	// false, the call will wait until any fence in `fences` is signaled.
	//
	// timeout - How long to wait before returning VKTimeout. May be common.NoTimeout to wait indefinitely.
	// The timeout is adjusted to the closest value allowed by the implementation timeout accuracy,
	// which may be substantially longer than the requested timeout.
	//
	// fences - A slice of Fence objects to wait for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkWaitForFences.html
	WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (common.VkResult, error)
	// ResetFences resets one or more objects to the unsignaled state
	//
	// fences - A slice of Fence objects to reset
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetFences.html
	ResetFences(fences []Fence) (common.VkResult, error)
	// UpdateDescriptorSets updates the contents of one or more DescriptorSet objects
	//
	// writes - A slice of WriteDescriptorSet structures describing the DescriptorSet objects to
	// write to
	//
	// copies - A slice of CopyDescriptorSet structures describing the DescriptorSet objects to
	// copy between
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSets.html
	UpdateDescriptorSets(writes []WriteDescriptorSet, copies []CopyDescriptorSet) error
	// FlushMappedMemoryRanges flushes one or more mapped memory ranges
	//
	// ranges - A slice of MappedMemoryRange structures describing the memory ranges to flush
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkFlushMappedMemoryRanges.html
	FlushMappedMemoryRanges(ranges []MappedMemoryRange) (common.VkResult, error)
	// InvalidateMappedMemoryRanges invalidates one or more mapped memory ranges
	//
	// ranges - A slice of MappedMemoryRange structures describing the memory ranges to invalidate
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkInvalidateMappedMemoryRanges.html
	InvalidateMappedMemoryRanges(ranges []MappedMemoryRange) (common.VkResult, error)
}

// Event is a synchronization primitive that can be used to insert fine-grained dependencies between
// commands submitted to the same queue, or between the host and a queue.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkEvent.html
type Event interface {
	// Handle is the internal Vulkan object handle for this Event
	Handle() driver.VkEvent
	// DeviceHandle is the internal Vulkan object handle for the Device this Event belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this Event
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Event. If it is at least Vulkan
	// 1.1, core1_1.PromoteEvent can be used to promote this to a core1_1.Event, etc.
	APIVersion() common.APIVersion

	// Destroy destroys the Event and the underlying structures. **Warning** after destruction,
	// this object will continue to exist, but the Vulkan object handle that backs it will be invalid.
	// Do not call further methods on this object.
	Destroy(callbacks *driver.AllocationCallbacks)
	// Set sets this Event to the signaled state
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkSetEvent.html
	Set() (common.VkResult, error)
	// Reset sets this Event to the unsignaled state
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetEvent.html
	Reset() (common.VkResult, error)
	// Status retrieves the status of this Event
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetEventStatus.html
	Status() (common.VkResult, error)
}

// Fence is a synchronization primitive that can be used to insert a dependency from a queue to
// the host.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkFence.html
type Fence interface {
	// Handle is the internal Vulkan object handle for this Fence
	Handle() driver.VkFence
	// DeviceHandle is the internal Vulkan object handle for the Device this Fence belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this Fence
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Fence. If it is at least
	// Vulkan 1.1, core1_1.PromoteFence can be used to promote this to a core1_1.Fence, etc.
	APIVersion() common.APIVersion

	// Destroy destroys this Fence object and the underlying structures.  **Warning** after destruction,
	// this object will continue to exist, but the Vulkan object handle that backs it will be invalid.
	// Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocatoin
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyFence.html
	Destroy(callbacks *driver.AllocationCallbacks)
	// Wait waits for this fence to become signaled
	//
	// timeout - How long to wait before returning VKTimeout. May be common.NoTimeout to wait indefinitely.
	// The timeout is adjusted to the closest value allowed by the implementation timeout accuracy,
	// which may be substantially longer than the requested timeout.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkWaitForFences.html
	Wait(timeout time.Duration) (common.VkResult, error)
	// Reset resets this Fence object to the unsignaled state
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetFences.html
	Reset() (common.VkResult, error)
	// Status returns the status of this Fence
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetFenceStatus.html
	Status() (common.VkResult, error)
}

// Framebuffer represents a collection of specific memory attachments that a RenderPass uses
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html
type Framebuffer interface {
	// Handle is the internal Vulkan object handle for this Framebuffer
	Handle() driver.VkFramebuffer
	// DeviceHandle is the internal Vulkan object handle for the Device this Framebuffer belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this Framebuffer
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Framebuffer. If it is at least
	// Vulkan 1.1, core1_1.PromoteFramebuffer can be used to promote this to a core1_1.Framebuffer, etc.
	APIVersion() common.APIVersion

	// Destroy destroys this Framebuffer object and the underlying structures. **Warning** after destruction,
	// this object will continue to exist, but the Vulkan object handle that backs it will be invalid. Do
	// not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyFramebuffer.html
	Destroy(callbacks *driver.AllocationCallbacks)
}

// Image represents multidimensional arrays of data which can be used for various purposes.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImage.html
type Image interface {
	// Handle is the internal Vulkan object handle for this Image
	Handle() driver.VkImage
	// DeviceHandle is the internal Vulkan object handle for the Device this Image belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this Image
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Image. If it is at least Vulkan
	// 1.1, core1_1.PromoteImage can be used to promote this to a core1_1.Image, etc.
	APIVersion() common.APIVersion

	// Destroy destroys this Image object and the underlying structures. **Warning** after destruction,
	// this object will continue to exist, but the Vulkan object handle that backs it will be invalid.
	// Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyImage.html
	Destroy(callbacks *driver.AllocationCallbacks)
	// MemoryRequirements returns the memory requirements for this Image
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements.html
	MemoryRequirements() *MemoryRequirements
	// BindImageMemory binds a DeviceMemory object to this Image object
	//
	// memory - Describes the DeviceMemory to attach
	//
	// offset - The start offset of the region of memory which is to be bound to the image.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory.html
	BindImageMemory(memory DeviceMemory, offset int) (common.VkResult, error)
	// SubresourceLayout retrieves information about an Image subresource
	//
	// subresource - Selects a specific subresource from the Image
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout.html
	SubresourceLayout(subresource *ImageSubresource) *SubresourceLayout
	// SparseMemoryRequirements queries the memory requirements for a sparse image
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetImageSparseMemoryRequirements.html
	SparseMemoryRequirements() []SparseImageMemoryRequirements
}

// ImageView represents contiguous ranges of Image subresources and contains additional metadata
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImageView.html
type ImageView interface {
	// Handle is the internal Vulkan object handle for this ImageView
	Handle() driver.VkImageView
	// DeviceHandle is the internal Vulkan object handle for the Device this ImageView belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this ImageView
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this ImageView. If it is at least Vulkan
	// 1.1, core1_1.PromoteImageView can be used to promote this to a core1_1.ImageView, etc.
	APIVersion() common.APIVersion

	// Destroy destroys the ImageView object and the underlying structures. **Warning** after destruction,
	// this object will continue to exist, but the Vulkan object handle that backs it will be invalid. Do not
	// call further methods on this object.
	Destroy(callbacks *driver.AllocationCallbacks)
}

// Instance stores per-application state for Vulkan
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkInstance.html
type Instance interface {
	// Handle is the internal Vulkan object handle for this Instance
	Handle() driver.VkInstance
	// Driver ist he Vulkan wrapper driver used by this Instance
	Driver() driver.Driver
	// APIVersion is the maximum VUlkan API supported by this Instance. If it is at least Vulkan 1.1,
	// core1_1.PromoteInstance can be used to promote this to a core1_1.Instance, etc.
	APIVersion() common.APIVersion

	// IsInstanceExtensionActive will return true if an Instance extension with the provided name was
	// activated on Instance creation
	//
	// extensionName - THe name of the extension to query
	IsInstanceExtensionActive(extensionName string) bool
	// EnumeratePhysicalDevices enumerates the physical devices accessible to this Instance
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDevices.html
	EnumeratePhysicalDevices() ([]PhysicalDevice, common.VkResult, error)

	// Destroy destroys the Instance object and the underlying structures. **Warning** after destruction,
	// this object will continue to exist, but the Vulkan object handle that backs it will be invalid.
	// Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyInstance.html
	Destroy(callbacks *driver.AllocationCallbacks)
}

// PhysicalDevice represents a single complete implementation of Vulkan available to the host, of which
// there are a finite number.
//
// PhysicalDevice objects are unusual in that they exist between the Instance and (logical) Device level.
// As a result, PhysicalDevices are the only object that can be extended by both Instance and Device
// extensions. As a result, there are some unusual cases in which a higher core version may be available
// for some PhysicalDevice functionality but not others. In order to represent this, physical devices
// are split into two objects at core1.1+, the PhysicalDevice and the "instance-scoped" PhysicalDevice.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html
type PhysicalDevice interface {
	// Handle is the internal Vulkan object handle for this PhysicalDevice
	Handle() driver.VkPhysicalDevice
	// Driver is the Vulkan wrapper driver used by this PhysicalDevice
	Driver() driver.Driver
	// InstanceAPIVersion is the maximum Vulkan API version supported by instance-scoped functionality
	// on this PhysicalDevice. This is usually the same as DeviceAPIVersion, but in some rare cases, it
	// may be higher. If it is at least Vulkan 1.1, core1_1.PromoteInstanceScopedPhysicalDevice can
	// be used to promote this to a core1_1.InstanceScopedPhysicalDevice, etc.
	InstanceAPIVersion() common.APIVersion
	// DeviceAPIVersion is the maximum Vulkan API version supported by device-scoped functionality on this
	// PhysicalDevice. This represents the highest API version supported by ALL functionality on this
	// PhysicalDevice. If it is at least Vulkan 1.1, core1_1.PromotePhysicalDevice can be used to promote
	// this to a core1_1.PhysicalDevice, etc.
	DeviceAPIVersion() common.APIVersion

	// CreateDevice creates a new logical device as a connection to this PhysicalDevice
	//
	// allocationCallbacks - Controls host memory allocation
	//
	// options - Parameters affecting the creation of the Device
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateDevice.html
	CreateDevice(allocationCallbacks *driver.AllocationCallbacks, options DeviceCreateInfo) (Device, common.VkResult, error)

	// QueueFamilyProperties reports properties of the queues of this PhysicalDevice
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html
	QueueFamilyProperties() []*QueueFamilyProperties

	// Properties returns properties of this PhysicalDevice
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties.html
	Properties() (*PhysicalDeviceProperties, error)
	// Features reports capabilities of this PhysicalDevice
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures.html
	Features() *PhysicalDeviceFeatures
	// EnumerateDeviceExtensionProperties returns properties of available PhysicalDevice extensions
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceExtensionProperties.html
	EnumerateDeviceExtensionProperties() (map[string]*ExtensionProperties, common.VkResult, error)
	// EnumerateDeviceExtensionPropertiesForLayer returns properties of available PhysicalDevice extensions
	// for the specifies layer
	//
	// layerName - Name of the layer to retrieve extensions from
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceExtensionProperties.html
	EnumerateDeviceExtensionPropertiesForLayer(layerName string) (map[string]*ExtensionProperties, common.VkResult, error)
	// EnumerateDeviceLayerProperties returns properties of available PhysicalDevice layers
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceLayerProperties.html
	EnumerateDeviceLayerProperties() (map[string]*LayerProperties, common.VkResult, error)
	// MemoryProperties reports memory information for this PhysicalDevice
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties.html
	MemoryProperties() *PhysicalDeviceMemoryProperties
	// FormatProperties lists this PhysicalDevice object's format capabilities
	//
	// format - The format whose properties are queried
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html
	FormatProperties(format Format) *FormatProperties
	// ImageFormatProperties lists this PhysicalDevice object's image format capabilities
	//
	// format - Specifies the Image format
	//
	// imageType - Specifies the Image type
	//
	// tiling - Specifies the Image tiling
	//
	// usages - Specifies the intended usage of the Image
	//
	// flags - Specifies additional parmeters of the Image
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html
	ImageFormatProperties(format Format, imageType ImageType, tiling ImageTiling, usages ImageUsageFlags, flags ImageCreateFlags) (*ImageFormatProperties, common.VkResult, error)
	// SparseImageFormatProperties retrieves properties of an image format applied to sparse images
	//
	// format - The Image format
	//
	// imageType - The dimensionality of the Image
	//
	// samples - Specifies the number of samples per texel
	//
	// usages - Describes the intended usage of the Image
	//
	// tiling - The tiling arrangement of the texel blocks in memory
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties.html
	SparseImageFormatProperties(format Format, imageType ImageType, samples SampleCountFlags, usages ImageUsageFlags, tiling ImageTiling) []SparseImageFormatProperties
}

// Pipeline represents compute, ray tracing, and graphics pipelines
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipeline.html
type Pipeline interface {
	// Handle is the internal Vulkan object handle for this Pipeline
	Handle() driver.VkPipeline
	// DeviceHandle is the internal Vulkan object handle for the Device this Pipeline belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this Pipeline
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Pipeline. If it is at least Vulkan
	// 1.1, core1_1.PromotePipeline can be used to promote this to a core1_1.Pipeline, etc.
	APIVersion() common.APIVersion

	// Destroy destroys the Pipeline object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will be
	// invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyPipeline.html
	Destroy(callbacks *driver.AllocationCallbacks)
}

// PipelineCache allows the result of Pipeline construction to be reused between Pipeline objects
// and between runs of an application.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html
type PipelineCache interface {
	// Handle is the internal Vulkan object handle for this PipelineCache
	Handle() driver.VkPipelineCache
	// DeviceHandle is the internal Vulkan object handle for the Device this PipelineCache belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this PipelineCache
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this PipelineCache. If it is at least
	// Vulkan 1.1, core1_1.PromotePipelineCache can be used to promote this to a core1_1.PipelineCache,
	// etc.
	APIVersion() common.APIVersion

	// Destroy destroys the PipelineCache object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyPipelineCache.html
	Destroy(callbacks *driver.AllocationCallbacks)
	// CacheData gets the data store from this PipelineCache
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPipelineCacheData.html
	CacheData() ([]byte, common.VkResult, error)
	// MergePipelineCaches combines the data stores of multiple PipelineCache object into this one
	//
	// srcCaches - A slice of PipelineCache objects which will be merged into this PipelineCache
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkMergePipelineCaches.html
	MergePipelineCaches(srcCaches []PipelineCache) (common.VkResult, error)
}

// PipelineLayout provides access to descriptor sets to Pipeline objects by combining zero or more
// descriptor sets and zero or more push constant ranges.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html
type PipelineLayout interface {
	// Handle is the internal Vulkan object handle for this PipelineLayout
	Handle() driver.VkPipelineLayout
	// DeviceHandle is the internal Vulkan object handle for the Device this PipelineLayout belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this PipelineLayout
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this PipelineLayout. If it is at least
	// Vulkan 1.1, core1_1.PromotePipelineLayout can be used to promote this to a core1_1.PipelineLayout,
	// etc.
	APIVersion() common.APIVersion

	// Destroy destroys the PipelineLayout object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will be
	// invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyPipelineLayout.html
	Destroy(callbacks *driver.AllocationCallbacks)
}

// QueryPool is a collection of a specific number of queries of a particular type.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html
type QueryPool interface {
	// Handle is the internal Vulkan object handle for this QueryPool
	Handle() driver.VkQueryPool
	// DeviceHandle is the internal Vulkan object handle for the Device this QueryPool belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this QueryPool
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this QueryPool. If it is at least
	// Vulkan 1.1, core1_1.PromoteQueryPool can be used to promote this to a core1_1.QueryPool, etc.
	APIVersion() common.APIVersion

	// Destroy destroys the QueryPool object and the underlying structures. **Warning** after destruction,
	// this object will continue to exist, but the Vulkan object handle that backs it will be invalid.
	// Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyQueryPool.html
	Destroy(callbacks *driver.AllocationCallbacks)
	// PopulateResults retrieves the status and results for a set of queries, and populates those results
	// into a preallocated byte array
	//
	// firstQuery - The initial query index
	//
	// queryCount - The number of queries to read
	//
	// results - A user-allocated slice of bytes where the results will be written
	//
	// resultStride - The stride in bytes between results for individual queries
	//
	// flags - Specifies how and when results are returned
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetQueryPoolResults.html
	PopulateResults(firstQuery, queryCount int, results []byte, resultStride int, flags QueryResultFlags) (common.VkResult, error)
}

// Queue represents a Device resource on which work is performed
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkQueue.html
type Queue interface {
	// Handle is the internal Vulkan object handle for this Queue
	Handle() driver.VkQueue
	// DeviceHandle is the internal Vulkan object handle for the Device this Queue belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this Queue
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Queue. If it is at least Vulkan 1.1,
	// core1_1.PromoteQueue can be used to promote this to a core1_1.Queue, etc.
	APIVersion() common.APIVersion

	// WaitIdle waits for this Queue to become idle
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkQueueWaitIdle.html
	WaitIdle() (common.VkResult, error)
	// Submit submits a sequence of Semaphore or CommandBuffer objects to this queue
	//
	// fence - An optional Fence object to be signaled once all submitted CommandBuffer objects have
	// completed execution.
	//
	// o - A slice of SubmitInfo structures, each specifying a CommandBuffer submission batch
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit.html
	Submit(fence Fence, o []SubmitInfo) (common.VkResult, error)
	// BindSparse binds DeviceMemory to a sparse resource object
	//
	// fence - An optional Fence object to be signaled.
	//
	// bindInfos - A slice of BindSparseInfo structures, each speicfying a sparse binding submission batch
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkQueueBindSparse.html
	BindSparse(fence Fence, bindInfos []BindSparseInfo) (common.VkResult, error)
}

// RenderPass represents a collection of attachments, subpasses, and dependencies between the subpasses
// and describes how the attachments are used over the course of the subpasses
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html
type RenderPass interface {
	// Handle is the internal Vulkan object handle for this RenderPass
	Handle() driver.VkRenderPass
	// DeviceHandle is the internal Vulkan object handle for the Device this RenderPass belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this RenderPass
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this RenderPass. If it is at least Vulkan
	// 1.1, core1_1.PromoteRenderPass can be used to promote this to a core1_1.RenderPass, etc.
	APIVersion() common.APIVersion

	// Destroy destroys the RenderPass object and the underlying structures.  **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods on this object.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyRenderPass.html
	Destroy(callbacks *driver.AllocationCallbacks)
	// RenderAreaGranularity returns the granularity for optimal render area
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetRenderAreaGranularity.html
	RenderAreaGranularity() Extent2D
}

// Sampler represents the state of an Image sampler, which is used by the implementation to read Image data
// and apply filtering and other transformations for the shader.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSampler.html
type Sampler interface {
	// Handle is the internal Vulkan object handle for this Sampler
	Handle() driver.VkSampler
	// DeviceHandle is the internal Vulkan object handle for the Device this Sampler belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this Sampler
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Sampler. If it is at least Vulkan
	// 1.1, core1_1.PromoteSampler can be used to promote this to a core1_1.Sampler, etc.
	APIVersion() common.APIVersion

	// Destroy destroys the Sampler object and the underlying structures. **Warning** after destruction,
	// this object will continue to exist, but the Vulkan objec thandle that backs it will be invalid.
	// Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroySampler.html
	Destroy(callbacks *driver.AllocationCallbacks)
}

// Semaphore is a synchronization primitive that can be used to insert a dependency between Queue operations
// or between a Queue operation and the host.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html
type Semaphore interface {
	// Handle is the internal Vulkan object handle for this Semaphore
	Handle() driver.VkSemaphore
	// DeviceHandle is the internal Vulkan object handle for the Device this Semaphore belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this Semaphore
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Semaphore. If it is at least
	// Vulkan 1.1, core1_1.PromoteSemaphore can be used to promote this to a core1_1.Semaphore
	APIVersion() common.APIVersion

	// Destroy destroys the Semaphore object and the underlying structures. **Warning** after destruciton,
	// this object will continue to exist, but the Vulkan object handle that backs it will be invalid.
	// Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroySemaphore.html
	Destroy(callbacks *driver.AllocationCallbacks)
}

// ShaderModule objects contain shader code and one or more entry points.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html
type ShaderModule interface {
	// Handle is the internal Vulkan object handle for this ShaderModule
	Handle() driver.VkShaderModule
	// DeviceHandle is the internal Vulkan object handle for the Device this ShaderModule belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this ShaderModule
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this ShaderModule. If it is at least
	// Vulkan 1.1, core1_1.PromoteShaderModule can be used to promote this to a core1_1.ShaderModule, etc.
	APIVersion() common.APIVersion

	// Destroy destroys the ShaderModule object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyShaderModule.html
	Destroy(callbacks *driver.AllocationCallbacks)
}
