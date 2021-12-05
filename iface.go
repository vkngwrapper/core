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
	CmdBindVertexBuffers(firstBinding uint32, buffers []Buffer, bufferOffsets []int)
	CmdBindIndexBuffer(buffer Buffer, offset int, indexType common.IndexType)
	CmdCopyBuffer(srcBuffer Buffer, dstBuffer Buffer, copyRegions []BufferCopy) error
	CmdBindDescriptorSets(bindPoint common.PipelineBindPoint, layout PipelineLayout, firstSet int, sets []DescriptorSet, dynamicOffsets []int)
	CmdPipelineBarrier(srcStageMask, dstStageMask common.PipelineStages, dependencies common.DependencyFlags, memoryBarriers []*MemoryBarrierOptions, bufferMemoryBarriers []*BufferMemoryBarrierOptions, imageMemoryBarriers []*ImageMemoryBarrierOptions) error
	CmdCopyBufferToImage(buffer Buffer, image Image, layout common.ImageLayout, regions []*BufferImageCopy) error
	CmdBlitImage(sourceImage Image, sourceImageLayout common.ImageLayout, destinationImage Image, destinationImageLayout common.ImageLayout, regions []*ImageBlit, filter common.Filter) error
	CmdPushConstants(layout PipelineLayout, stageFlags common.ShaderStages, offset int, values interface{}) error
	CmdSetViewport(firstViewport int, viewports []common.Viewport)
	CmdSetScissor(firstScissor int, scissors []common.Rect2D)
	CmdCopyImage(srcImage Image, srcImageLayout common.ImageLayout, dstImage Image, dstImageLayout common.ImageLayout, regions []ImageCopy) error
	CmdNextSubpass(contents SubpassContents)
	CmdWaitEvents(events []Event, srcStageMask common.PipelineStages, dstStageMask common.PipelineStages, memoryBarriers []*MemoryBarrierOptions, bufferMemoryBarriers []*BufferMemoryBarrierOptions, imageMemoryBarriers []*ImageMemoryBarrierOptions) error
	CmdSetEvent(event Event, stageMask common.PipelineStages)
}

type CommandPool interface {
	Handle() VkCommandPool
	Destroy()
	AllocateCommandBuffers(o *CommandBufferOptions) ([]CommandBuffer, VkResult, error)
	FreeCommandBuffers(buffers []CommandBuffer)
}

type DescriptorPool interface {
	Handle() VkDescriptorPool
	Destroy()
	AllocateDescriptorSets(o *DescriptorSetOptions) ([]DescriptorSet, VkResult, error)
	FreeDescriptorSets(sets []DescriptorSet) (VkResult, error)
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
}

type Device interface {
	Handle() VkDevice
	Destroy()
	Driver() Driver
	GetQueue(queueFamilyIndex int, queueIndex int) Queue
	WaitForIdle() (VkResult, error)
	WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (VkResult, error)
	ResetFences(fences []Fence) (VkResult, error)
	AllocateMemory(o *DeviceMemoryOptions) (DeviceMemory, VkResult, error)
	FreeMemory(memory DeviceMemory)
	UpdateDescriptorSets(writes []WriteDescriptorSetOptions, copies []CopyDescriptorSetOptions) error
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

	CreateBuffer(device Device, o *BufferOptions) (Buffer, VkResult, error)
	CreateCommandPool(device Device, o *CommandPoolOptions) (CommandPool, VkResult, error)
	CreateDescriptorPool(device Device, o *DescriptorPoolOptions) (DescriptorPool, VkResult, error)
	CreateDescriptorSetLayout(device Device, o *DescriptorSetLayoutOptions) (DescriptorSetLayout, VkResult, error)
	CreateDevice(physicalDevice PhysicalDevice, options *DeviceOptions) (Device, VkResult, error)
	CreateEvent(device Device, options *EventOptions) (Event, VkResult, error)
	CreateFence(device Device, o *FenceOptions) (Fence, VkResult, error)
	CreateFrameBuffer(device Device, o *FramebufferOptions) (Framebuffer, VkResult, error)
	CreateGraphicsPipelines(device Device, pipelineCache PipelineCache, o []*GraphicsPipelineOptions) ([]Pipeline, VkResult, error)
	CreateInstance(options *InstanceOptions) (Instance, VkResult, error)
	CreateImage(device Device, options *ImageOptions) (Image, VkResult, error)
	CreateImageView(device Device, o *ImageViewOptions) (ImageView, VkResult, error)
	CreatePipelineCache(device Device, o *PipelineCacheOptions) (PipelineCache, VkResult, error)
	CreatePipelineLayout(device Device, o *PipelineLayoutOptions) (PipelineLayout, VkResult, error)
	CreateRenderPass(device Device, o *RenderPassOptions) (RenderPass, VkResult, error)
	CreateSampler(device Device, o *SamplerOptions) (Sampler, VkResult, error)
	CreateSemaphore(device Device, o *SemaphoreOptions) (Semaphore, VkResult, error)
	CreateShaderModule(device Device, o *ShaderModuleOptions) (ShaderModule, VkResult, error)
}

type PhysicalDevice interface {
	Handle() VkPhysicalDevice
	Driver() Driver
	QueueFamilyProperties() []*common.QueueFamily
	Properties() *common.PhysicalDeviceProperties
	Features() *common.PhysicalDeviceFeatures
	AvailableExtensions() (map[string]*common.ExtensionProperties, VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, VkResult, error)
	MemoryProperties() *PhysicalDeviceMemoryProperties
	FormatProperties(format common.DataFormat) *common.FormatProperties
}

type Pipeline interface {
	Handle() VkPipeline
	Destroy()
}

type PipelineCache interface {
	Handle() VkPipelineCache
	Destroy()
}

type PipelineLayout interface {
	Handle() VkPipelineLayout
	Destroy()
}

type Queue interface {
	Handle() VkQueue
	Driver() Driver
	WaitForIdle() (VkResult, error)
	SubmitToQueue(fence Fence, o []*SubmitOptions) (VkResult, error)
}

type RenderPass interface {
	Handle() VkRenderPass
	Destroy()
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
