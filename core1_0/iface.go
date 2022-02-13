package core1_0

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0/options"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/iface"
	"time"
	"unsafe"
)

type Buffer interface {
	iface.Buffer
	MemoryRequirements() *MemoryRequirements
	BindBufferMemory(memory iface.DeviceMemory, offset int) (common.VkResult, error)
}

type BufferView interface {
	iface.BufferView
}

type CommandBuffer interface {
	iface.CommandBuffer
	Free()

	Begin(o *options.BeginOptions) (common.VkResult, error)
	End() (common.VkResult, error)
	Reset(flags core.CommandBufferResetFlags) (common.VkResult, error)

	CmdBeginRenderPass(contents core.SubpassContents, o *options.RenderPassBeginOptions) error
	CmdEndRenderPass()
	CmdBindPipeline(bindPoint common.PipelineBindPoint, pipeline iface.Pipeline)
	CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32)
	CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32)
	CmdBindVertexBuffers(buffers []iface.Buffer, bufferOffsets []int)
	CmdBindIndexBuffer(buffer iface.Buffer, offset int, indexType common.IndexType)
	CmdCopyBuffer(srcBuffer iface.Buffer, dstBuffer iface.Buffer, copyRegions []BufferCopy) error
	CmdBindDescriptorSets(bindPoint common.PipelineBindPoint, layout iface.PipelineLayout, sets []iface.DescriptorSet, dynamicOffsets []int)
	CmdPipelineBarrier(srcStageMask, dstStageMask common.PipelineStages, dependencies common.DependencyFlags, memoryBarriers []options.MemoryBarrierOptions, bufferMemoryBarriers []options.BufferMemoryBarrierOptions, imageMemoryBarriers []options.ImageMemoryBarrierOptions) error
	CmdCopyBufferToImage(buffer iface.Buffer, image iface.Image, layout common.ImageLayout, regions []BufferImageCopy) error
	CmdBlitImage(sourceImage iface.Image, sourceImageLayout common.ImageLayout, destinationImage iface.Image, destinationImageLayout common.ImageLayout, regions []ImageBlit, filter common.Filter) error
	CmdPushConstants(layout iface.PipelineLayout, stageFlags common.ShaderStages, offset int, valueBytes []byte)
	CmdSetViewport(viewports []common.Viewport)
	CmdSetScissor(scissors []common.Rect2D)
	CmdCopyImage(srcImage iface.Image, srcImageLayout common.ImageLayout, dstImage iface.Image, dstImageLayout common.ImageLayout, regions []ImageCopy) error
	CmdNextSubpass(contents core.SubpassContents)
	CmdWaitEvents(events []iface.Event, srcStageMask common.PipelineStages, dstStageMask common.PipelineStages, memoryBarriers []options.MemoryBarrierOptions, bufferMemoryBarriers []options.BufferMemoryBarrierOptions, imageMemoryBarriers []options.ImageMemoryBarrierOptions) error
	CmdSetEvent(event iface.Event, stageMask common.PipelineStages)
	CmdClearColorImage(image iface.Image, imageLayout common.ImageLayout, color core.ClearColorValue, ranges []common.ImageSubresourceRange)
	CmdResetQueryPool(queryPool iface.QueryPool, startQuery, queryCount int)
	CmdBeginQuery(queryPool iface.QueryPool, query int, flags common.QueryControlFlags)
	CmdEndQuery(queryPool iface.QueryPool, query int)
	CmdCopyQueryPoolResults(queryPool iface.QueryPool, firstQuery, queryCount int, dstBuffer iface.Buffer, dstOffset, stride int, flags QueryResultFlags)
	CmdExecuteCommands(commandBuffers []iface.CommandBuffer)
	CmdClearAttachments(attachments []ClearAttachment, rects []ClearRect) error
	CmdClearDepthStencilImage(image iface.Image, imageLayout common.ImageLayout, depthStencil *core.ClearValueDepthStencil, ranges []common.ImageSubresourceRange)
	CmdCopyImageToBuffer(srcImage iface.Image, srcImageLayout common.ImageLayout, dstBuffer iface.Buffer, regions []BufferImageCopy) error
	CmdDispatch(groupCountX, groupCountY, groupCountZ int)
	CmdDispatchIndirect(buffer iface.Buffer, offset int)
	CmdDrawIndexedIndirect(buffer iface.Buffer, offset int, drawCount, stride int)
	CmdDrawIndirect(buffer iface.Buffer, offset int, drawCount, stride int)
	CmdFillBuffer(dstBuffer iface.Buffer, dstOffset int, size int, data uint32)
	CmdResetEvent(event iface.Event, stageMask common.PipelineStages)
	CmdResolveImage(srcImage iface.Image, srcImageLayout common.ImageLayout, dstImage iface.Image, dstImageLayout common.ImageLayout, regions []ImageResolve) error
	CmdSetBlendConstants(blendConstants [4]float32)
	CmdSetDepthBias(depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32)
	CmdSetDepthBounds(min, max float32)
	CmdSetLineWidth(lineWidth float32)
	CmdSetStencilCompareMask(faceMask common.StencilFaces, compareMask uint32)
	CmdSetStencilReference(faceMask common.StencilFaces, reference uint32)
	CmdSetStencilWriteMask(faceMask common.StencilFaces, writeMask uint32)
	CmdUpdateBuffer(dstBuffer iface.Buffer, dstOffset int, dataSize int, data []byte)
	CmdWriteTimestamp(pipelineStage common.PipelineStages, queryPool iface.QueryPool, query int)
}

type CommandPool interface {
	iface.CommandPool
	Reset(flags core.CommandPoolResetFlags) (common.VkResult, error)
}

type DescriptorPool interface {
	iface.DescriptorPool
	Reset(flags DescriptorPoolResetFlags) (common.VkResult, error)
}

type DescriptorSet interface {
	iface.DescriptorSet
	Free() (common.VkResult, error)
}

type DescriptorSetLayout interface {
	iface.DescriptorSetLayout
}

type DeviceMemory interface {
	iface.DeviceMemory
	MapMemory(offset int, size int, flags options.MemoryMapFlags) (unsafe.Pointer, common.VkResult, error)
	UnmapMemory()
	Free(callbacks *driver.AllocationCallbacks)
	Commitment() int
	Flush() (common.VkResult, error)
	Invalidate() (common.VkResult, error)
}

type Device interface {
	iface.Device
	WaitForIdle() (common.VkResult, error)
	WaitForFences(waitForAll bool, timeout time.Duration, fences []iface.Fence) (common.VkResult, error)
	ResetFences(fences []iface.Fence) (common.VkResult, error)
	UpdateDescriptorSets(writes []options.WriteDescriptorSetOptions, copies []options.CopyDescriptorSetOptions) error
	FlushMappedMemoryRanges(ranges []options.MappedMemoryRange) (common.VkResult, error)
	InvalidateMappedMemoryRanges(ranges []options.MappedMemoryRange) (common.VkResult, error)
}

type Event interface {
	iface.Event
	Set() (common.VkResult, error)
	Reset() (common.VkResult, error)
	Status() (common.VkResult, error)
}

type Fence interface {
	iface.Fence
	Wait(timeout time.Duration) (common.VkResult, error)
	Reset() (common.VkResult, error)
	Status() (common.VkResult, error)
}

type Framebuffer interface {
	iface.Framebuffer
}

type Image interface {
	iface.Image
	MemoryRequirements() *MemoryRequirements
	BindImageMemory(memory iface.DeviceMemory, offset int) (common.VkResult, error)
	SubresourceLayout(subresource *common.ImageSubresource) *common.SubresourceLayout
	SparseMemoryRequirements() []SparseImageMemoryRequirements
}

type ImageView interface {
	iface.ImageView
}

type Instance interface {
	iface.Instance
}

type PhysicalDevice interface {
	iface.PhysicalDevice
	QueueFamilyProperties() []*common.QueueFamily
	Properties() *PhysicalDeviceProperties
	Features() *options.PhysicalDeviceFeatures
	AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error)
	MemoryProperties() *PhysicalDeviceMemoryProperties
	FormatProperties(format common.DataFormat) *FormatProperties
	ImageFormatProperties(format common.DataFormat, imageType common.ImageType, tiling common.ImageTiling, usages common.ImageUsages, flags options.ImageFlags) (*ImageFormatProperties, common.VkResult, error)
}

type Pipeline interface {
	iface.Pipeline
}

type PipelineCache interface {
	iface.PipelineCache
	CacheData() ([]byte, common.VkResult, error)
	MergePipelineCaches(srcCaches []iface.PipelineCache) (common.VkResult, error)
}

type PipelineLayout interface {
	iface.PipelineLayout
}

type QueryPool interface {
	iface.QueryPool
	PopulateResults(firstQuery, queryCount int, resultSize, resultStride int, flags QueryResultFlags) ([]byte, common.VkResult, error)
}

type Queue interface {
	iface.Queue
	WaitForIdle() (common.VkResult, error)
	SubmitToQueue(fence iface.Fence, o []options.SubmitOptions) (common.VkResult, error)
	BindSparse(fence iface.Fence, bindInfos []options.BindSparseOptions) (common.VkResult, error)
}

type RenderPass interface {
	iface.RenderPass
	RenderAreaGranularity() common.Extent2D
}

type Semaphore interface {
	iface.Semaphore
}

type ShaderModule interface {
	iface.ShaderModule
}

type Sampler interface {
	iface.Sampler
}
