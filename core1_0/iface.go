package core1_0

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"time"
	"unsafe"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/core1_0_mocks.go -package mocks

type Buffer interface {
	Handle() driver.VkBuffer
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	MemoryRequirements() *MemoryRequirements
	BindBufferMemory(memory DeviceMemory, offset int) (common.VkResult, error)
}

type BufferView interface {
	Handle() driver.VkBufferView
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type CommandBuffer interface {
	Handle() driver.VkCommandBuffer
	Driver() driver.Driver
	DeviceHandle() driver.VkDevice
	CommandPoolHandle() driver.VkCommandPool
	APIVersion() common.APIVersion

	Free()
	Begin(o BeginOptions) (common.VkResult, error)
	End() (common.VkResult, error)
	Reset(flags common.CommandBufferResetFlags) (common.VkResult, error)
	CommandsRecorded() int
	DrawsRecorded() int
	DispatchesRecorded() int

	CmdBeginRenderPass(contents common.SubpassContents, o RenderPassBeginOptions) error
	CmdEndRenderPass()
	CmdBindPipeline(bindPoint common.PipelineBindPoint, pipeline Pipeline)
	CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32)
	CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32)
	CmdBindVertexBuffers(buffers []Buffer, bufferOffsets []int)
	CmdBindIndexBuffer(buffer Buffer, offset int, indexType common.IndexType)
	CmdCopyBuffer(srcBuffer Buffer, dstBuffer Buffer, copyRegions []BufferCopy) error
	CmdBindDescriptorSets(bindPoint common.PipelineBindPoint, layout PipelineLayout, sets []DescriptorSet, dynamicOffsets []int)
	CmdPipelineBarrier(srcStageMask, dstStageMask common.PipelineStages, dependencies common.DependencyFlags, memoryBarriers []MemoryBarrierOptions, bufferMemoryBarriers []BufferMemoryBarrierOptions, imageMemoryBarriers []ImageMemoryBarrierOptions) error
	CmdCopyBufferToImage(buffer Buffer, image Image, layout common.ImageLayout, regions []BufferImageCopy) error
	CmdBlitImage(sourceImage Image, sourceImageLayout common.ImageLayout, destinationImage Image, destinationImageLayout common.ImageLayout, regions []ImageBlit, filter common.Filter) error
	CmdPushConstants(layout PipelineLayout, stageFlags common.ShaderStages, offset int, valueBytes []byte)
	CmdSetViewport(viewports []common.Viewport)
	CmdSetScissor(scissors []common.Rect2D)
	CmdCopyImage(srcImage Image, srcImageLayout common.ImageLayout, dstImage Image, dstImageLayout common.ImageLayout, regions []ImageCopy) error
	CmdNextSubpass(contents common.SubpassContents)
	CmdWaitEvents(events []Event, srcStageMask common.PipelineStages, dstStageMask common.PipelineStages, memoryBarriers []MemoryBarrierOptions, bufferMemoryBarriers []BufferMemoryBarrierOptions, imageMemoryBarriers []ImageMemoryBarrierOptions) error
	CmdSetEvent(event Event, stageMask common.PipelineStages)
	CmdClearColorImage(image Image, imageLayout common.ImageLayout, color common.ClearColorValue, ranges []common.ImageSubresourceRange)
	CmdResetQueryPool(queryPool QueryPool, startQuery, queryCount int)
	CmdBeginQuery(queryPool QueryPool, query int, flags common.QueryControlFlags)
	CmdEndQuery(queryPool QueryPool, query int)
	CmdCopyQueryPoolResults(queryPool QueryPool, firstQuery, queryCount int, dstBuffer Buffer, dstOffset, stride int, flags common.QueryResultFlags)
	CmdExecuteCommands(commandBuffers []CommandBuffer)
	CmdClearAttachments(attachments []ClearAttachment, rects []ClearRect) error
	CmdClearDepthStencilImage(image Image, imageLayout common.ImageLayout, depthStencil *common.ClearValueDepthStencil, ranges []common.ImageSubresourceRange)
	CmdCopyImageToBuffer(srcImage Image, srcImageLayout common.ImageLayout, dstBuffer Buffer, regions []BufferImageCopy) error
	CmdDispatch(groupCountX, groupCountY, groupCountZ int)
	CmdDispatchIndirect(buffer Buffer, offset int)
	CmdDrawIndexedIndirect(buffer Buffer, offset int, drawCount, stride int)
	CmdDrawIndirect(buffer Buffer, offset int, drawCount, stride int)
	CmdFillBuffer(dstBuffer Buffer, dstOffset int, size int, data uint32)
	CmdResetEvent(event Event, stageMask common.PipelineStages)
	CmdResolveImage(srcImage Image, srcImageLayout common.ImageLayout, dstImage Image, dstImageLayout common.ImageLayout, regions []ImageResolve) error
	CmdSetBlendConstants(blendConstants [4]float32)
	CmdSetDepthBias(depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32)
	CmdSetDepthBounds(min, max float32)
	CmdSetLineWidth(lineWidth float32)
	CmdSetStencilCompareMask(faceMask common.StencilFaces, compareMask uint32)
	CmdSetStencilReference(faceMask common.StencilFaces, reference uint32)
	CmdSetStencilWriteMask(faceMask common.StencilFaces, writeMask uint32)
	CmdUpdateBuffer(dstBuffer Buffer, dstOffset int, dataSize int, data []byte)
	CmdWriteTimestamp(pipelineStage common.PipelineStages, queryPool QueryPool, query int)
}

type CommandPool interface {
	Handle() driver.VkCommandPool
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	Reset(flags common.CommandPoolResetFlags) (common.VkResult, error)
}

type DescriptorPool interface {
	Handle() driver.VkDescriptorPool
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	Reset(flags common.DescriptorPoolResetFlags) (common.VkResult, error)
}

type DescriptorSet interface {
	Handle() driver.VkDescriptorSet
	DescriptorPoolHandle() driver.VkDescriptorPool
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Free() (common.VkResult, error)
}

type DescriptorSetLayout interface {
	Handle() driver.VkDescriptorSetLayout
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type DeviceMemory interface {
	Handle() driver.VkDeviceMemory
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	MapMemory(offset int, size int, flags MemoryMapFlags) (unsafe.Pointer, common.VkResult, error)
	UnmapMemory()
	Free(callbacks *driver.AllocationCallbacks)
	Commitment() int
	FlushAll() (common.VkResult, error)
	InvalidateAll() (common.VkResult, error)
}

type Device interface {
	Handle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	WaitForIdle() (common.VkResult, error)
	WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (common.VkResult, error)
	ResetFences(fences []Fence) (common.VkResult, error)
	UpdateDescriptorSets(writes []WriteDescriptorSetOptions, copies []CopyDescriptorSetOptions) error
	FlushMappedMemoryRanges(ranges []MappedMemoryRangeOptions) (common.VkResult, error)
	InvalidateMappedMemoryRanges(ranges []MappedMemoryRangeOptions) (common.VkResult, error)
}

type Event interface {
	Handle() driver.VkEvent
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	Set() (common.VkResult, error)
	Reset() (common.VkResult, error)
	Status() (common.VkResult, error)
}

type Fence interface {
	Handle() driver.VkFence
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	Wait(timeout time.Duration) (common.VkResult, error)
	Reset() (common.VkResult, error)
	Status() (common.VkResult, error)
}

type Framebuffer interface {
	Handle() driver.VkFramebuffer
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type Image interface {
	Handle() driver.VkImage
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	MemoryRequirements() *MemoryRequirements
	BindImageMemory(memory DeviceMemory, offset int) (common.VkResult, error)
	SubresourceLayout(subresource *common.ImageSubresource) *common.SubresourceLayout
	SparseMemoryRequirements() []SparseImageMemoryRequirements
}

type ImageView interface {
	Handle() driver.VkImageView
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type Instance interface {
	Handle() driver.VkInstance
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type PhysicalDevice interface {
	Handle() driver.VkPhysicalDevice
	Driver() driver.Driver
	InstanceAPIVersion() common.APIVersion
	DeviceAPIVersion() common.APIVersion

	QueueFamilyProperties() []*QueueFamily
	Properties() (*PhysicalDeviceProperties, error)
	Features() *PhysicalDeviceFeatures
	AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error)
	MemoryProperties() *PhysicalDeviceMemoryProperties
	FormatProperties(format common.DataFormat) *FormatProperties
	ImageFormatProperties(format common.DataFormat, imageType common.ImageType, tiling common.ImageTiling, usages common.ImageUsages, flags common.ImageCreateFlags) (*ImageFormatProperties, common.VkResult, error)
	SparseImageFormatProperties(format common.DataFormat, imageType common.ImageType, samples common.SampleCounts, usages common.ImageUsages, tiling common.ImageTiling) []SparseImageFormatProperties
}

type Pipeline interface {
	Handle() driver.VkPipeline
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type PipelineCache interface {
	Handle() driver.VkPipelineCache
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	CacheData() ([]byte, common.VkResult, error)
	MergePipelineCaches(srcCaches []PipelineCache) (common.VkResult, error)
}

type PipelineLayout interface {
	Handle() driver.VkPipelineLayout
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type QueryPool interface {
	Handle() driver.VkQueryPool
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	PopulateResults(firstQuery, queryCount int, results []byte, resultStride int, flags common.QueryResultFlags) (common.VkResult, error)
}

type Queue interface {
	Handle() driver.VkQueue
	Driver() driver.Driver
	APIVersion() common.APIVersion

	WaitForIdle() (common.VkResult, error)
	SubmitToQueue(fence Fence, o []SubmitOptions) (common.VkResult, error)
	BindSparse(fence Fence, bindInfos []BindSparseOptions) (common.VkResult, error)
}

type RenderPass interface {
	Handle() driver.VkRenderPass
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	RenderAreaGranularity() common.Extent2D
}

type Sampler interface {
	Handle() driver.VkSampler
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type Semaphore interface {
	Handle() driver.VkSemaphore
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type ShaderModule interface {
	Handle() driver.VkShaderModule
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}
