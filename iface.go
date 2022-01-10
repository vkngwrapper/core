package core

import (
	"github.com/CannibalVox/VKng/core/common"
	"time"
	"unsafe"
)

//go:generate mockgen -source iface.go -destination ./mocks/core.go -package=mocks

type Buffer interface {
	Handle() VkBuffer
	Destroy()
	MemoryRequirements() *common.MemoryRequirements
	BindBufferMemory(memory DeviceMemory, offset int) (VkResult, error)
}

type BufferView interface {
	Handle() VkBufferView
	Destroy()
}

type CommandBuffer interface {
	Handle() VkCommandBuffer

	Begin(o *BeginOptions) (VkResult, error)
	End() (VkResult, error)
	Reset(flags CommandBufferResetFlags) (VkResult, error)

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
	Handle() VkCommandPool
	Destroy()
	AllocateCommandBuffers(o *CommandBufferOptions) ([]CommandBuffer, VkResult, error)
	FreeCommandBuffers(buffers []CommandBuffer)
	Reset(flags CommandPoolResetFlags) (VkResult, error)
}

type DescriptorPool interface {
	Handle() VkDescriptorPool
	Destroy()
	AllocateDescriptorSets(o *DescriptorSetOptions) ([]DescriptorSet, VkResult, error)
	FreeDescriptorSets(sets []DescriptorSet) (VkResult, error)
	Reset(flags DescriptorPoolResetFlags) (VkResult, error)
}

type DescriptorSet interface {
	Handle() VkDescriptorSet
}

type DescriptorSetLayout interface {
	Handle() VkDescriptorSetLayout
	Destroy()
}

type DeviceMemory interface {
	Handle() VkDeviceMemory
	MapMemory(offset int, size int, flags MemoryMapFlags) (unsafe.Pointer, VkResult, error)
	UnmapMemory()
	Commitment() int
	Flush() (VkResult, error)
	Invalidate() (VkResult, error)
}

type Device interface {
	Handle() VkDevice
	Destroy()
	Driver() Driver
	GetQueue(queueFamilyIndex int, queueIndex int) Queue
	WaitForIdle() (VkResult, error)
	WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (VkResult, error)
	ResetFences(fences []Fence) (VkResult, error)
	AllocateMemory(allocationCallbacks *AllocationCallbacks, o *DeviceMemoryOptions) (DeviceMemory, VkResult, error)
	FreeMemory(memory DeviceMemory)
	UpdateDescriptorSets(writes []WriteDescriptorSetOptions, copies []CopyDescriptorSetOptions) error
	FlushMappedMemoryRanges(ranges []*MappedMemoryRange) (VkResult, error)
	InvalidateMappedMemoryRanges(ranges []*MappedMemoryRange) (VkResult, error)
}

type Event interface {
	Handle() VkEvent
	Destroy()
	Set() (VkResult, error)
	Reset() (VkResult, error)
	Status() (VkResult, error)
}

type Fence interface {
	Handle() VkFence
	Destroy()
	Wait(timeout time.Duration) (VkResult, error)
	Reset() (VkResult, error)
	Status() (VkResult, error)
}

type Framebuffer interface {
	Handle() VkFramebuffer
	Destroy()
}

type Image interface {
	Handle() VkImage
	Destroy()
	MemoryRequirements() *common.MemoryRequirements
	BindImageMemory(memory DeviceMemory, offset int) (VkResult, error)
	SubresourceLayout(subresource *common.ImageSubresource) *common.SubresourceLayout
	SparseMemoryRequirements() []SparseImageMemoryRequirements
}

type ImageView interface {
	Handle() VkImageView
	Destroy()
}

type Instance interface {
	Handle() VkInstance
	Destroy()
	Driver() Driver
	PhysicalDevices() ([]PhysicalDevice, VkResult, error)
}

type Loader1_0 interface {
	Version() common.APIVersion
	Driver() Driver

	AvailableExtensions() (map[string]*common.ExtensionProperties, VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, VkResult, error)
	AvailableLayers() (map[string]*common.LayerProperties, VkResult, error)

	CreateAllocationCallbacks(allocation AllocationFunction, reallocation ReallocationFunction, free FreeFunction, internalAllocation InternalAllocationNotification, internalFree InternalFreeNotification, userData interface{}) *AllocationCallbacks
	CreateBuffer(device Device, allocationCallbacks *AllocationCallbacks, o *BufferOptions) (Buffer, VkResult, error)
	CreateBufferView(device Device, allocationCallbacks *AllocationCallbacks, o *BufferViewOptions) (BufferView, VkResult, error)
	CreateCommandPool(device Device, allocationCallbacks *AllocationCallbacks, o *CommandPoolOptions) (CommandPool, VkResult, error)
	CreateDescriptorPool(device Device, allocationCallbacks *AllocationCallbacks, o *DescriptorPoolOptions) (DescriptorPool, VkResult, error)
	CreateDescriptorSetLayout(device Device, allocationCallbacks *AllocationCallbacks, o *DescriptorSetLayoutOptions) (DescriptorSetLayout, VkResult, error)
	CreateDevice(physicalDevice PhysicalDevice, allocationCallbacks *AllocationCallbacks, options *DeviceOptions) (Device, VkResult, error)
	CreateEvent(device Device, allocationCallbacks *AllocationCallbacks, options *EventOptions) (Event, VkResult, error)
	CreateFence(device Device, allocationCallbacks *AllocationCallbacks, o *FenceOptions) (Fence, VkResult, error)
	CreateFrameBuffer(device Device, allocationCallbacks *AllocationCallbacks, o *FramebufferOptions) (Framebuffer, VkResult, error)
	CreateGraphicsPipelines(device Device, pipelineCache PipelineCache, allocationCallbacks *AllocationCallbacks, o []*GraphicsPipelineOptions) ([]Pipeline, VkResult, error)
	CreateComputePipelines(device Device, pipelineCache PipelineCache, allocationCallbacks *AllocationCallbacks, o []*ComputePipelineOptions) ([]Pipeline, VkResult, error)
	CreateInstance(allocationCallbacks *AllocationCallbacks, options *InstanceOptions) (Instance, VkResult, error)
	CreateImage(device Device, allocationCallbacks *AllocationCallbacks, options *ImageOptions) (Image, VkResult, error)
	CreateImageView(device Device, allocationCallbacks *AllocationCallbacks, o *ImageViewOptions) (ImageView, VkResult, error)
	CreatePipelineCache(device Device, allocationCallbacks *AllocationCallbacks, o *PipelineCacheOptions) (PipelineCache, VkResult, error)
	CreatePipelineLayout(device Device, allocationCallbacks *AllocationCallbacks, o *PipelineLayoutOptions) (PipelineLayout, VkResult, error)
	CreateQueryPool(device Device, allocationCallbacks *AllocationCallbacks, o *QueryPoolOptions) (QueryPool, VkResult, error)
	CreateRenderPass(device Device, allocationCallbacks *AllocationCallbacks, o *RenderPassOptions) (RenderPass, VkResult, error)
	CreateSampler(device Device, allocationCallbacks *AllocationCallbacks, o *SamplerOptions) (Sampler, VkResult, error)
	CreateSemaphore(device Device, allocationCallbacks *AllocationCallbacks, o *SemaphoreOptions) (Semaphore, VkResult, error)
	CreateShaderModule(device Device, allocationCallbacks *AllocationCallbacks, o *ShaderModuleOptions) (ShaderModule, VkResult, error)
}

type PhysicalDevice interface {
	Handle() VkPhysicalDevice
	Driver() Driver
	QueueFamilyProperties() []*common.QueueFamily
	Properties() *common.PhysicalDeviceProperties
	Features() *common.PhysicalDeviceFeatures
	AvailableExtensions() (map[string]*common.ExtensionProperties, VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, VkResult, error)
	AvailableLayers() (map[string]*common.LayerProperties, VkResult, error)
	MemoryProperties() *PhysicalDeviceMemoryProperties
	FormatProperties(format common.DataFormat) *common.FormatProperties
	ImageFormatProperties(format common.DataFormat, imageType common.ImageType, tiling common.ImageTiling, usages common.ImageUsages, flags ImageFlags) (*common.ImageFormatProperties, VkResult, error)
}

type Pipeline interface {
	Handle() VkPipeline
	Destroy()
}

type PipelineCache interface {
	Handle() VkPipelineCache
	Destroy()
	CacheData() ([]byte, VkResult, error)
	MergePipelineCaches(srcCaches []PipelineCache) (VkResult, error)
}

type PipelineLayout interface {
	Handle() VkPipelineLayout
	Destroy()
}

type QueryPool interface {
	Handle() VkQueryPool
	Destroy()
	PopulateResults(firstQuery, queryCount int, resultSize, resultStride int, flags common.QueryResultFlags) ([]byte, VkResult, error)
}

type Queue interface {
	Handle() VkQueue
	Driver() Driver
	WaitForIdle() (VkResult, error)
	SubmitToQueue(fence Fence, o []*SubmitOptions) (VkResult, error)
	BindSparse(fence Fence, bindInfos []*BindSparseOptions) (VkResult, error)
}

type RenderPass interface {
	Handle() VkRenderPass
	Destroy()
	RenderAreaGranularity() common.Extent2D
}

type Semaphore interface {
	Handle() VkSemaphore
	Destroy()
}

type ShaderModule interface {
	Handle() VkShaderModule
	Destroy()
}

type Sampler interface {
	Handle() VkSampler
	Destroy()
}
