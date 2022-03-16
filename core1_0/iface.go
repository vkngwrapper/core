package core1_0

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	"time"
	"unsafe"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/core1_0_mocks.go -package mocks

type Buffer interface {
	Handle() driver.VkBuffer

	Core1_1() core1_1.Buffer

	Destroy(callbacks *driver.AllocationCallbacks)
	MemoryRequirements() *MemoryRequirements
	BindBufferMemory(memory DeviceMemory, offset int) (common.VkResult, error)
}

type BufferView interface {
	Handle() driver.VkBufferView

	Core1_1() core1_1.BufferView

	Destroy(callbacks *driver.AllocationCallbacks)
}

type CommandBuffer interface {
	Handle() driver.VkCommandBuffer
	Driver() driver.Driver
	DeviceHandle() driver.VkDevice
	CommandPoolHandle() driver.VkCommandPool

	Core1_1() core1_1.CommandBuffer

	Free()
	Begin(o *BeginOptions) (common.VkResult, error)
	End() (common.VkResult, error)
	Reset(flags common.CommandBufferResetFlags) (common.VkResult, error)

	CmdBeginRenderPass(contents common.SubpassContents, o *RenderPassBeginOptions) error
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
	Device() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Core1_1() core1_1.CommandPool

	Destroy(callbacks *driver.AllocationCallbacks)
	Reset(flags common.CommandPoolResetFlags) (common.VkResult, error)
}

type DescriptorPool interface {
	Handle() driver.VkDescriptorPool
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Core1_1() core1_1.DescriptorPool

	Destroy(callbacks *driver.AllocationCallbacks)
	Reset(flags common.DescriptorPoolResetFlags) (common.VkResult, error)
}

type DescriptorSet interface {
	Handle() driver.VkDescriptorSet
	PoolHandle() driver.VkDescriptorPool
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver

	Core1_1() core1_1.DescriptorSet

	Free() (common.VkResult, error)
}

type DescriptorSetLayout interface {
	Handle() driver.VkDescriptorSetLayout

	Core1_1() core1_1.DescriptorSetLayout

	Destroy(callbacks *driver.AllocationCallbacks)
}

type DeviceMemory interface {
	Handle() driver.VkDeviceMemory
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver

	Core1_1() core1_1.DeviceMemory

	MapMemory(offset int, size int, flags MemoryMapFlags) (unsafe.Pointer, common.VkResult, error)
	UnmapMemory()
	Free(callbacks *driver.AllocationCallbacks)
	Commitment() int
	Flush() (common.VkResult, error)
	Invalidate() (common.VkResult, error)
}

type Device interface {
	Handle() driver.VkDevice
	Driver() driver.Driver
	Destroy(callbacks *driver.AllocationCallbacks)
	APIVersion() common.APIVersion

	Core1_1() core1_1.Device

	WaitForIdle() (common.VkResult, error)
	WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (common.VkResult, error)
	ResetFences(fences []Fence) (common.VkResult, error)
	UpdateDescriptorSets(writes []WriteDescriptorSetOptions, copies []CopyDescriptorSetOptions) error
	FlushMappedMemoryRanges(ranges []MappedMemoryRange) (common.VkResult, error)
	InvalidateMappedMemoryRanges(ranges []MappedMemoryRange) (common.VkResult, error)

	GetQueue(queueFamilyIndex int, queueIndex int) Queue
	AllocateMemory(allocationCallbacks *driver.AllocationCallbacks, o *DeviceMemoryOptions) (DeviceMemory, common.VkResult, error)
	FreeMemory(deviceMemory DeviceMemory, allocationCallbacks *driver.AllocationCallbacks)
}

type Event interface {
	Handle() driver.VkEvent

	Core1_1() core1_1.Event

	Destroy(callbacks *driver.AllocationCallbacks)
	Set() (common.VkResult, error)
	Reset() (common.VkResult, error)
	Status() (common.VkResult, error)
}

type Fence interface {
	Handle() driver.VkFence

	Core1_1() core1_1.Fence

	Destroy(callbacks *driver.AllocationCallbacks)
	Wait(timeout time.Duration) (common.VkResult, error)
	Reset() (common.VkResult, error)
	Status() (common.VkResult, error)
}

type Framebuffer interface {
	Handle() driver.VkFramebuffer

	Core1_1() core1_1.Framebuffer

	Destroy(callbacks *driver.AllocationCallbacks)
}

type Image interface {
	Handle() driver.VkImage

	Core1_1() core1_1.Image

	Destroy(callbacks *driver.AllocationCallbacks)
	MemoryRequirements() *MemoryRequirements
	BindImageMemory(memory DeviceMemory, offset int) (common.VkResult, error)
	SubresourceLayout(subresource *common.ImageSubresource) *common.SubresourceLayout
	SparseMemoryRequirements() []SparseImageMemoryRequirements
}

type ImageView interface {
	Handle() driver.VkImageView

	Core1_1() core1_1.ImageView

	Destroy(callbacks *driver.AllocationCallbacks)
}

type Instance interface {
	Handle() driver.VkInstance
	Driver() driver.Driver

	Core1_1() core1_1.Instance

	Destroy(callbacks *driver.AllocationCallbacks)
	PhysicalDevices() ([]PhysicalDevice, common.VkResult, error)
}

type PhysicalDevice interface {
	Handle() driver.VkPhysicalDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Core1_1() core1_1.PhysicalDevice

	QueueFamilyProperties() []*QueueFamily
	Properties() *PhysicalDeviceProperties
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

	Core1_1() core1_1.Pipeline

	Destroy(callbacks *driver.AllocationCallbacks)
}

type PipelineCache interface {
	Handle() driver.VkPipelineCache

	Core1_1() core1_1.PipelineCache

	Destroy(callbacks *driver.AllocationCallbacks)
	CacheData() ([]byte, common.VkResult, error)
	MergePipelineCaches(srcCaches []PipelineCache) (common.VkResult, error)
}

type PipelineLayout interface {
	Handle() driver.VkPipelineLayout

	Core1_1() core1_1.PipelineLayout

	Destroy(callbacks *driver.AllocationCallbacks)
}

type QueryPool interface {
	Handle() driver.VkQueryPool

	Core1_1() core1_1.QueryPool

	Destroy(callbacks *driver.AllocationCallbacks)
	PopulateResults(firstQuery, queryCount int, resultSize, resultStride int, flags common.QueryResultFlags) ([]byte, common.VkResult, error)
}

type Queue interface {
	Handle() driver.VkQueue
	Driver() driver.Driver

	Core1_1() core1_1.Queue

	WaitForIdle() (common.VkResult, error)
	SubmitToQueue(fence Fence, o []SubmitOptions) (common.VkResult, error)
	BindSparse(fence Fence, bindInfos []BindSparseOptions) (common.VkResult, error)
}

type RenderPass interface {
	Handle() driver.VkRenderPass

	Core1_1() core1_1.RenderPass

	Destroy(callbacks *driver.AllocationCallbacks)
	RenderAreaGranularity() common.Extent2D
}

type Sampler interface {
	Handle() driver.VkSampler

	Core1_1() core1_1.Sampler

	Destroy(callbacks *driver.AllocationCallbacks)
}

type Semaphore interface {
	Handle() driver.VkSemaphore

	Core1_1() core1_1.Semaphore

	Destroy(callbacks *driver.AllocationCallbacks)
}

type ShaderModule interface {
	Handle() driver.VkShaderModule

	Core1_1() core1_1.ShaderModule

	Destroy(callbacks *driver.AllocationCallbacks)
}
