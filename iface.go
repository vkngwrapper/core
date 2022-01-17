package core

import (
	"github.com/CannibalVox/VKng/core/common"
	driver2 "github.com/CannibalVox/VKng/core/driver"
	"time"
	"unsafe"
)

//go:generate mockgen -source iface.go -destination ./mocks/core.go -package=mocks

type Buffer interface {
	Handle() driver2.VkBuffer
	Destroy(callbacks *AllocationCallbacks)
	MemoryRequirements() *common.MemoryRequirements
	BindBufferMemory(memory DeviceMemory, offset int) (common.VkResult, error)
}

type BufferView interface {
	Handle() driver2.VkBufferView
	Destroy(callbacks *AllocationCallbacks)
}

type CommandBuffer interface {
	Handle() driver2.VkCommandBuffer

	Begin(o *BeginOptions) (common.VkResult, error)
	End() (common.VkResult, error)
	Reset(flags CommandBufferResetFlags) (common.VkResult, error)

	CmdBeginRenderPass(contents SubpassContents, o *RenderPassBeginOptions) error
	CmdEndRenderPass()
	CmdBindPipeline(bindPoint common.PipelineBindPoint, pipeline Pipeline)
	CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32)
	CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32)
	CmdBindVertexBuffers(buffers []Buffer, bufferOffsets []int)
	CmdBindIndexBuffer(buffer Buffer, offset int, indexType common.IndexType)
	CmdCopyBuffer(srcBuffer Buffer, dstBuffer Buffer, copyRegions []BufferCopy) error
	CmdBindDescriptorSets(bindPoint common.PipelineBindPoint, layout PipelineLayout, sets []DescriptorSet, dynamicOffsets []int)
	CmdPipelineBarrier(srcStageMask, dstStageMask common.PipelineStages, dependencies common.DependencyFlags, memoryBarriers []*MemoryBarrierOptions, bufferMemoryBarriers []*BufferMemoryBarrierOptions, imageMemoryBarriers []*ImageMemoryBarrierOptions) error
	CmdCopyBufferToImage(buffer Buffer, image Image, layout common.ImageLayout, regions []*BufferImageCopy) error
	CmdBlitImage(sourceImage Image, sourceImageLayout common.ImageLayout, destinationImage Image, destinationImageLayout common.ImageLayout, regions []*ImageBlit, filter common.Filter) error
	CmdPushConstants(layout PipelineLayout, stageFlags common.ShaderStages, offset int, valueBytes []byte)
	CmdSetViewport(viewports []common.Viewport)
	CmdSetScissor(scissors []common.Rect2D)
	CmdCopyImage(srcImage Image, srcImageLayout common.ImageLayout, dstImage Image, dstImageLayout common.ImageLayout, regions []ImageCopy) error
	CmdNextSubpass(contents SubpassContents)
	CmdWaitEvents(events []Event, srcStageMask common.PipelineStages, dstStageMask common.PipelineStages, memoryBarriers []*MemoryBarrierOptions, bufferMemoryBarriers []*BufferMemoryBarrierOptions, imageMemoryBarriers []*ImageMemoryBarrierOptions) error
	CmdSetEvent(event Event, stageMask common.PipelineStages)
	CmdClearColorImage(image Image, imageLayout common.ImageLayout, color ClearColorValue, ranges []common.ImageSubresourceRange)
	CmdResetQueryPool(queryPool QueryPool, startQuery, queryCount int)
	CmdBeginQuery(queryPool QueryPool, query int, flags common.QueryControlFlags)
	CmdEndQuery(queryPool QueryPool, query int)
	CmdCopyQueryPoolResults(queryPool QueryPool, firstQuery, queryCount int, dstBuffer Buffer, dstOffset, stride int, flags common.QueryResultFlags)
	CmdExecuteCommands(commandBuffers []CommandBuffer)
	CmdClearAttachments(attachments []ClearAttachment, rects []ClearRect)
	CmdClearDepthStencilImage(image Image, imageLayout common.ImageLayout, depthStencil *ClearValueDepthStencil, ranges []common.ImageSubresourceRange)
	CmdCopyImageToBuffer(srcImage Image, srcImageLayout common.ImageLayout, dstBuffer Buffer, regions []BufferImageCopy) error
	CmdDispatch(groupCountX, groupCountY, groupCountZ int)
	CmdDispatchIndirect(buffer Buffer, offset int)
	CmdDrawIndexedIndirect(buffer Buffer, offset int, drawCount, stride int)
	CmdDrawIndirect(buffer Buffer, offset int, drawCount, stride int)
	CmdFillBuffer(dstBuffer Buffer, dstOffset int, size int, data uint32)
	CmdResetEvent(event Event, stageMask common.PipelineStages)
	CmdResolveImage(srcImage Image, srcImageLayout common.ImageLayout, dstImage Image, dstImageLayout common.ImageLayout, regions []ImageResolve)
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
	Handle() driver2.VkCommandPool
	Destroy(callbacks *AllocationCallbacks)
	AllocateCommandBuffers(o *CommandBufferOptions) ([]CommandBuffer, common.VkResult, error)
	FreeCommandBuffers(buffers []CommandBuffer)
	Reset(flags CommandPoolResetFlags) (common.VkResult, error)
}

type DescriptorPool interface {
	Handle() driver2.VkDescriptorPool
	Destroy(callbacks *AllocationCallbacks)
	AllocateDescriptorSets(o *DescriptorSetOptions) ([]DescriptorSet, common.VkResult, error)
	FreeDescriptorSets(sets []DescriptorSet) (common.VkResult, error)
	Reset(flags DescriptorPoolResetFlags) (common.VkResult, error)
}

type DescriptorSet interface {
	Handle() driver2.VkDescriptorSet
}

type DescriptorSetLayout interface {
	Handle() driver2.VkDescriptorSetLayout
	Destroy(callbacks *AllocationCallbacks)
}

type DeviceMemory interface {
	Handle() driver2.VkDeviceMemory
	MapMemory(offset int, size int, flags MemoryMapFlags) (unsafe.Pointer, common.VkResult, error)
	UnmapMemory()
	Free(callbacks *AllocationCallbacks)
	Commitment() int
	Flush() (common.VkResult, error)
	Invalidate() (common.VkResult, error)
}

type Device interface {
	Handle() driver2.VkDevice
	Destroy(callbacks *AllocationCallbacks)
	Driver() driver2.Driver
	GetQueue(queueFamilyIndex int, queueIndex int) Queue
	WaitForIdle() (common.VkResult, error)
	WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (common.VkResult, error)
	ResetFences(fences []Fence) (common.VkResult, error)
	AllocateMemory(allocationCallbacks *AllocationCallbacks, o *DeviceMemoryOptions) (DeviceMemory, common.VkResult, error)
	UpdateDescriptorSets(writes []WriteDescriptorSetOptions, copies []CopyDescriptorSetOptions) error
	FlushMappedMemoryRanges(ranges []*MappedMemoryRange) (common.VkResult, error)
	InvalidateMappedMemoryRanges(ranges []*MappedMemoryRange) (common.VkResult, error)
}

type Event interface {
	Handle() driver2.VkEvent
	Destroy(callbacks *AllocationCallbacks)
	Set() (common.VkResult, error)
	Reset() (common.VkResult, error)
	Status() (common.VkResult, error)
}

type Fence interface {
	Handle() driver2.VkFence
	Destroy(callbacks *AllocationCallbacks)
	Wait(timeout time.Duration) (common.VkResult, error)
	Reset() (common.VkResult, error)
	Status() (common.VkResult, error)
}

type Framebuffer interface {
	Handle() driver2.VkFramebuffer
	Destroy(callbacks *AllocationCallbacks)
}

type Image interface {
	Handle() driver2.VkImage
	Destroy(callbacks *AllocationCallbacks)
	MemoryRequirements() *common.MemoryRequirements
	BindImageMemory(memory DeviceMemory, offset int) (common.VkResult, error)
	SubresourceLayout(subresource *common.ImageSubresource) *common.SubresourceLayout
	SparseMemoryRequirements() []SparseImageMemoryRequirements
}

type ImageView interface {
	Handle() driver2.VkImageView
	Destroy(callbacks *AllocationCallbacks)
}

type Instance interface {
	Handle() driver2.VkInstance
	Destroy(callbacks *AllocationCallbacks)
	Driver() driver2.Driver
	PhysicalDevices() ([]PhysicalDevice, common.VkResult, error)
}

type Loader1_0 interface {
	Version() common.APIVersion
	Driver() driver2.Driver

	AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error)

	CreateAllocationCallbacks(o *AllocationCallbackOptions) *AllocationCallbacks
	CreateBuffer(device Device, allocationCallbacks *AllocationCallbacks, o *BufferOptions) (Buffer, common.VkResult, error)
	CreateBufferView(device Device, allocationCallbacks *AllocationCallbacks, o *BufferViewOptions) (BufferView, common.VkResult, error)
	CreateCommandPool(device Device, allocationCallbacks *AllocationCallbacks, o *CommandPoolOptions) (CommandPool, common.VkResult, error)
	CreateDescriptorPool(device Device, allocationCallbacks *AllocationCallbacks, o *DescriptorPoolOptions) (DescriptorPool, common.VkResult, error)
	CreateDescriptorSetLayout(device Device, allocationCallbacks *AllocationCallbacks, o *DescriptorSetLayoutOptions) (DescriptorSetLayout, common.VkResult, error)
	CreateDevice(physicalDevice PhysicalDevice, allocationCallbacks *AllocationCallbacks, options *DeviceOptions) (Device, common.VkResult, error)
	CreateEvent(device Device, allocationCallbacks *AllocationCallbacks, options *EventOptions) (Event, common.VkResult, error)
	CreateFence(device Device, allocationCallbacks *AllocationCallbacks, o *FenceOptions) (Fence, common.VkResult, error)
	CreateFrameBuffer(device Device, allocationCallbacks *AllocationCallbacks, o *FramebufferOptions) (Framebuffer, common.VkResult, error)
	CreateGraphicsPipelines(device Device, pipelineCache PipelineCache, allocationCallbacks *AllocationCallbacks, o []*GraphicsPipelineOptions) ([]Pipeline, common.VkResult, error)
	CreateComputePipelines(device Device, pipelineCache PipelineCache, allocationCallbacks *AllocationCallbacks, o []*ComputePipelineOptions) ([]Pipeline, common.VkResult, error)
	CreateInstance(allocationCallbacks *AllocationCallbacks, options *InstanceOptions) (Instance, common.VkResult, error)
	CreateImage(device Device, allocationCallbacks *AllocationCallbacks, options *ImageOptions) (Image, common.VkResult, error)
	CreateImageView(device Device, allocationCallbacks *AllocationCallbacks, o *ImageViewOptions) (ImageView, common.VkResult, error)
	CreatePipelineCache(device Device, allocationCallbacks *AllocationCallbacks, o *PipelineCacheOptions) (PipelineCache, common.VkResult, error)
	CreatePipelineLayout(device Device, allocationCallbacks *AllocationCallbacks, o *PipelineLayoutOptions) (PipelineLayout, common.VkResult, error)
	CreateQueryPool(device Device, allocationCallbacks *AllocationCallbacks, o *QueryPoolOptions) (QueryPool, common.VkResult, error)
	CreateRenderPass(device Device, allocationCallbacks *AllocationCallbacks, o *RenderPassOptions) (RenderPass, common.VkResult, error)
	CreateSampler(device Device, allocationCallbacks *AllocationCallbacks, o *SamplerOptions) (Sampler, common.VkResult, error)
	CreateSemaphore(device Device, allocationCallbacks *AllocationCallbacks, o *SemaphoreOptions) (Semaphore, common.VkResult, error)
	CreateShaderModule(device Device, allocationCallbacks *AllocationCallbacks, o *ShaderModuleOptions) (ShaderModule, common.VkResult, error)
}

type PhysicalDevice interface {
	Handle() driver2.VkPhysicalDevice
	Driver() driver2.Driver
	QueueFamilyProperties() []*common.QueueFamily
	Properties() *common.PhysicalDeviceProperties
	Features() *common.PhysicalDeviceFeatures
	AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error)
	MemoryProperties() *PhysicalDeviceMemoryProperties
	FormatProperties(format common.DataFormat) *common.FormatProperties
	ImageFormatProperties(format common.DataFormat, imageType common.ImageType, tiling common.ImageTiling, usages common.ImageUsages, flags ImageFlags) (*common.ImageFormatProperties, common.VkResult, error)
}

type Pipeline interface {
	Handle() driver2.VkPipeline
	Destroy(callbacks *AllocationCallbacks)
}

type PipelineCache interface {
	Handle() driver2.VkPipelineCache
	Destroy(callbacks *AllocationCallbacks)
	CacheData() ([]byte, common.VkResult, error)
	MergePipelineCaches(srcCaches []PipelineCache) (common.VkResult, error)
}

type PipelineLayout interface {
	Handle() driver2.VkPipelineLayout
	Destroy(callbacks *AllocationCallbacks)
}

type QueryPool interface {
	Handle() driver2.VkQueryPool
	Destroy(callbacks *AllocationCallbacks)
	PopulateResults(firstQuery, queryCount int, resultSize, resultStride int, flags common.QueryResultFlags) ([]byte, common.VkResult, error)
}

type Queue interface {
	Handle() driver2.VkQueue
	Driver() driver2.Driver
	WaitForIdle() (common.VkResult, error)
	SubmitToQueue(fence Fence, o []*SubmitOptions) (common.VkResult, error)
	BindSparse(fence Fence, bindInfos []*BindSparseOptions) (common.VkResult, error)
}

type RenderPass interface {
	Handle() driver2.VkRenderPass
	Destroy(callbacks *AllocationCallbacks)
	RenderAreaGranularity() common.Extent2D
}

type Semaphore interface {
	Handle() driver2.VkSemaphore
	Destroy(callbacks *AllocationCallbacks)
}

type ShaderModule interface {
	Handle() driver2.VkShaderModule
	Destroy(callbacks *AllocationCallbacks)
}

type Sampler interface {
	Handle() driver2.VkSampler
	Destroy(callbacks *AllocationCallbacks)
}
