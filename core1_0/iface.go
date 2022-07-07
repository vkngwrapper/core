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
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	MemoryRequirements() *MemoryRequirements
	BindBufferMemory(memory DeviceMemory, offset int) (common.VkResult, error)
}

type BufferView interface {
	Handle() driver.VkBufferView
	DeviceHandle() driver.VkDevice
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
	Begin(o CommandBufferBeginInfo) (common.VkResult, error)
	End() (common.VkResult, error)
	Reset(flags CommandBufferResetFlags) (common.VkResult, error)
	CommandsRecorded() int
	DrawsRecorded() int
	DispatchesRecorded() int

	CmdBeginRenderPass(contents SubpassContents, o RenderPassBeginInfo) error
	CmdEndRenderPass()
	CmdBindPipeline(bindPoint PipelineBindPoint, pipeline Pipeline)
	CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32)
	CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32)
	CmdBindVertexBuffers(buffers []Buffer, bufferOffsets []int)
	CmdBindIndexBuffer(buffer Buffer, offset int, indexType IndexType)
	CmdCopyBuffer(srcBuffer Buffer, dstBuffer Buffer, copyRegions []BufferCopy) error
	CmdBindDescriptorSets(bindPoint PipelineBindPoint, layout PipelineLayout, sets []DescriptorSet, dynamicOffsets []int)
	CmdPipelineBarrier(srcStageMask, dstStageMask PipelineStageFlags, dependencies DependencyFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) error
	CmdCopyBufferToImage(buffer Buffer, image Image, layout ImageLayout, regions []BufferImageCopy) error
	CmdBlitImage(sourceImage Image, sourceImageLayout ImageLayout, destinationImage Image, destinationImageLayout ImageLayout, regions []ImageBlit, filter Filter) error
	CmdPushConstants(layout PipelineLayout, stageFlags ShaderStageFlags, offset int, valueBytes []byte)
	CmdSetViewport(viewports []Viewport)
	CmdSetScissor(scissors []Rect2D)
	CmdCopyImage(srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regions []ImageCopy) error
	CmdNextSubpass(contents SubpassContents)
	CmdWaitEvents(events []Event, srcStageMask PipelineStageFlags, dstStageMask PipelineStageFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) error
	CmdSetEvent(event Event, stageMask PipelineStageFlags)
	CmdClearColorImage(image Image, imageLayout ImageLayout, color ClearColorValue, ranges []ImageSubresourceRange)
	CmdResetQueryPool(queryPool QueryPool, startQuery, queryCount int)
	CmdBeginQuery(queryPool QueryPool, query int, flags QueryControlFlags)
	CmdEndQuery(queryPool QueryPool, query int)
	CmdCopyQueryPoolResults(queryPool QueryPool, firstQuery, queryCount int, dstBuffer Buffer, dstOffset, stride int, flags QueryResultFlags)
	CmdExecuteCommands(commandBuffers []CommandBuffer)
	CmdClearAttachments(attachments []ClearAttachment, rects []ClearRect) error
	CmdClearDepthStencilImage(image Image, imageLayout ImageLayout, depthStencil *ClearValueDepthStencil, ranges []ImageSubresourceRange)
	CmdCopyImageToBuffer(srcImage Image, srcImageLayout ImageLayout, dstBuffer Buffer, regions []BufferImageCopy) error
	CmdDispatch(groupCountX, groupCountY, groupCountZ int)
	CmdDispatchIndirect(buffer Buffer, offset int)
	CmdDrawIndexedIndirect(buffer Buffer, offset int, drawCount, stride int)
	CmdDrawIndirect(buffer Buffer, offset int, drawCount, stride int)
	CmdFillBuffer(dstBuffer Buffer, dstOffset int, size int, data uint32)
	CmdResetEvent(event Event, stageMask PipelineStageFlags)
	CmdResolveImage(srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regions []ImageResolve) error
	CmdSetBlendConstants(blendConstants [4]float32)
	CmdSetDepthBias(depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32)
	CmdSetDepthBounds(min, max float32)
	CmdSetLineWidth(lineWidth float32)
	CmdSetStencilCompareMask(faceMask StencilFaceFlags, compareMask uint32)
	CmdSetStencilReference(faceMask StencilFaceFlags, reference uint32)
	CmdSetStencilWriteMask(faceMask StencilFaceFlags, writeMask uint32)
	CmdUpdateBuffer(dstBuffer Buffer, dstOffset int, dataSize int, data []byte)
	CmdWriteTimestamp(pipelineStage PipelineStageFlags, queryPool QueryPool, query int)
}

type CommandPool interface {
	Handle() driver.VkCommandPool
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	Reset(flags CommandPoolResetFlags) (common.VkResult, error)
}

type DescriptorPool interface {
	Handle() driver.VkDescriptorPool
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	Reset(flags DescriptorPoolResetFlags) (common.VkResult, error)
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
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type DeviceMemory interface {
	Handle() driver.VkDeviceMemory
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Map(offset int, size int, flags MemoryMapFlags) (unsafe.Pointer, common.VkResult, error)
	Unmap()
	Free(callbacks *driver.AllocationCallbacks)
	Commitment() int
	FlushAll() (common.VkResult, error)
	InvalidateAll() (common.VkResult, error)
}

type Device interface {
	Handle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	IsDeviceExtensionActive(extensionName string) bool

	CreateBuffer(allocationCallbacks *driver.AllocationCallbacks, o BufferCreateInfo) (Buffer, common.VkResult, error)
	CreateBufferView(allocationCallbacks *driver.AllocationCallbacks, o BufferViewCreateInfo) (BufferView, common.VkResult, error)
	CreateCommandPool(allocationCallbacks *driver.AllocationCallbacks, o CommandPoolCreateInfo) (CommandPool, common.VkResult, error)
	CreateDescriptorPool(allocationCallbacks *driver.AllocationCallbacks, o DescriptorPoolCreateInfo) (DescriptorPool, common.VkResult, error)
	CreateDescriptorSetLayout(allocationCallbacks *driver.AllocationCallbacks, o DescriptorSetLayoutCreateInfo) (DescriptorSetLayout, common.VkResult, error)
	CreateEvent(allocationCallbacks *driver.AllocationCallbacks, options EventCreateInfo) (Event, common.VkResult, error)
	CreateFence(allocationCallbacks *driver.AllocationCallbacks, o FenceCreateInfo) (Fence, common.VkResult, error)
	CreateFramebuffer(allocationCallbacks *driver.AllocationCallbacks, o FramebufferCreateInfo) (Framebuffer, common.VkResult, error)
	CreateGraphicsPipelines(pipelineCache PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []GraphicsPipelineCreateInfo) ([]Pipeline, common.VkResult, error)
	CreateComputePipelines(pipelineCache PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []ComputePipelineCreateInfo) ([]Pipeline, common.VkResult, error)
	CreateImage(allocationCallbacks *driver.AllocationCallbacks, options ImageCreateOptions) (Image, common.VkResult, error)
	CreateImageView(allocationCallbacks *driver.AllocationCallbacks, o ImageViewCreateInfo) (ImageView, common.VkResult, error)
	CreatePipelineCache(allocationCallbacks *driver.AllocationCallbacks, o PipelineCacheCreateInfo) (PipelineCache, common.VkResult, error)
	CreatePipelineLayout(allocationCallbacks *driver.AllocationCallbacks, o PipelineLayoutCreateInfo) (PipelineLayout, common.VkResult, error)
	CreateQueryPool(allocationCallbacks *driver.AllocationCallbacks, o QueryPoolCreateInfo) (QueryPool, common.VkResult, error)
	CreateRenderPass(allocationCallbacks *driver.AllocationCallbacks, o RenderPassCreateInfo) (RenderPass, common.VkResult, error)
	CreateSampler(allocationCallbacks *driver.AllocationCallbacks, o SamplerCreateInfo) (Sampler, common.VkResult, error)
	CreateSemaphore(allocationCallbacks *driver.AllocationCallbacks, o SemaphoreCreateInfo) (Semaphore, common.VkResult, error)
	CreateShaderModule(allocationCallbacks *driver.AllocationCallbacks, o ShaderModuleCreateInfo) (ShaderModule, common.VkResult, error)

	GetQueue(queueFamilyIndex int, queueIndex int) Queue
	AllocateMemory(allocationCallbacks *driver.AllocationCallbacks, o MemoryAllocateInfo) (DeviceMemory, common.VkResult, error)
	FreeMemory(deviceMemory DeviceMemory, allocationCallbacks *driver.AllocationCallbacks)

	AllocateCommandBuffers(o CommandBufferAllocateInfo) ([]CommandBuffer, common.VkResult, error)
	FreeCommandBuffers(buffers []CommandBuffer)
	AllocateDescriptorSets(o DescriptorSetAllocateInfo) ([]DescriptorSet, common.VkResult, error)
	FreeDescriptorSets(sets []DescriptorSet) (common.VkResult, error)

	Destroy(callbacks *driver.AllocationCallbacks)
	WaitIdle() (common.VkResult, error)
	WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (common.VkResult, error)
	ResetFences(fences []Fence) (common.VkResult, error)
	UpdateDescriptorSets(writes []WriteDescriptorSet, copies []CopyDescriptorSet) error
	FlushMappedMemoryRanges(ranges []MappedMemoryRange) (common.VkResult, error)
	InvalidateMappedMemoryRanges(ranges []MappedMemoryRange) (common.VkResult, error)
}

type Event interface {
	Handle() driver.VkEvent
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	Set() (common.VkResult, error)
	Reset() (common.VkResult, error)
	Status() (common.VkResult, error)
}

type Fence interface {
	Handle() driver.VkFence
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	Wait(timeout time.Duration) (common.VkResult, error)
	Reset() (common.VkResult, error)
	Status() (common.VkResult, error)
}

type Framebuffer interface {
	Handle() driver.VkFramebuffer
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type Image interface {
	Handle() driver.VkImage
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	MemoryRequirements() *MemoryRequirements
	BindImageMemory(memory DeviceMemory, offset int) (common.VkResult, error)
	SubresourceLayout(subresource *ImageSubresource) *SubresourceLayout
	SparseMemoryRequirements() []SparseImageMemoryRequirements
}

type ImageView interface {
	Handle() driver.VkImageView
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type Instance interface {
	Handle() driver.VkInstance
	Driver() driver.Driver
	APIVersion() common.APIVersion

	IsInstanceExtensionActive(extensionName string) bool
	EnumeratePhysicalDevices() ([]PhysicalDevice, common.VkResult, error)

	Destroy(callbacks *driver.AllocationCallbacks)
}

type PhysicalDevice interface {
	Handle() driver.VkPhysicalDevice
	Driver() driver.Driver
	InstanceAPIVersion() common.APIVersion
	DeviceAPIVersion() common.APIVersion

	CreateDevice(allocationCallbacks *driver.AllocationCallbacks, options DeviceCreateInfo) (Device, common.VkResult, error)

	QueueFamilyProperties() []*QueueFamily
	Properties() (*PhysicalDeviceProperties, error)
	Features() *PhysicalDeviceFeatures
	EnumerateDeviceExtensionProperties() (map[string]*ExtensionProperties, common.VkResult, error)
	EnumerateDeviceExtensionPropertiesForLayer(layerName string) (map[string]*ExtensionProperties, common.VkResult, error)
	EnumerateDeviceLayerProperties() (map[string]*LayerProperties, common.VkResult, error)
	MemoryProperties() *PhysicalDeviceMemoryProperties
	FormatProperties(format Format) *FormatProperties
	ImageFormatProperties(format Format, imageType ImageType, tiling ImageTiling, usages ImageUsageFlags, flags ImageCreateFlags) (*ImageFormatProperties, common.VkResult, error)
	SparseImageFormatProperties(format Format, imageType ImageType, samples SampleCountFlags, usages ImageUsageFlags, tiling ImageTiling) []SparseImageFormatProperties
}

type Pipeline interface {
	Handle() driver.VkPipeline
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type PipelineCache interface {
	Handle() driver.VkPipelineCache
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	CacheData() ([]byte, common.VkResult, error)
	MergePipelineCaches(srcCaches []PipelineCache) (common.VkResult, error)
}

type PipelineLayout interface {
	Handle() driver.VkPipelineLayout
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type QueryPool interface {
	Handle() driver.VkQueryPool
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	PopulateResults(firstQuery, queryCount int, results []byte, resultStride int, flags QueryResultFlags) (common.VkResult, error)
}

type Queue interface {
	Handle() driver.VkQueue
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	WaitIdle() (common.VkResult, error)
	Submit(fence Fence, o []SubmitInfo) (common.VkResult, error)
	BindSparse(fence Fence, bindInfos []BindSparseOptions) (common.VkResult, error)
}

type RenderPass interface {
	Handle() driver.VkRenderPass
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
	RenderAreaGranularity() Extent2D
}

type Sampler interface {
	Handle() driver.VkSampler
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type Semaphore interface {
	Handle() driver.VkSemaphore
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}

type ShaderModule interface {
	Handle() driver.VkShaderModule
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	APIVersion() common.APIVersion

	Destroy(callbacks *driver.AllocationCallbacks)
}
