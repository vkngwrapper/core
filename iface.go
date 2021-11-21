package core

import (
	"github.com/CannibalVox/VKng/core/common"
	"time"
	"unsafe"
)

//go:generate mockgen -source iface.go -destination ./mocks/core.go -package=mocks

type Buffer interface {
	Handle() VkBuffer
	Destroy() error
	MemoryRequirements() (*common.MemoryRequirements, error)
	BindBufferMemory(memory DeviceMemory, offset int) (VkResult, error)
}

type BufferView interface {
	Handle() VkBufferView
}

type CommandBuffer interface {
	Handle() VkCommandBuffer

	Begin(o *BeginOptions) (VkResult, error)
	End() (VkResult, error)

	CmdBeginRenderPass(contents SubpassContents, o *RenderPassBeginOptions) error
	CmdEndRenderPass() error
	CmdBindPipeline(bindPoint common.PipelineBindPoint, pipeline Pipeline) error
	CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) error
	CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) error
	CmdBindVertexBuffers(firstBinding uint32, buffers []Buffer, bufferOffsets []int) error
	CmdBindIndexBuffer(buffer Buffer, offset int, indexType common.IndexType) error
	CmdCopyBuffer(srcBuffer Buffer, dstBuffer Buffer, copyRegions []BufferCopy) error
	CmdBindDescriptorSets(bindPoint common.PipelineBindPoint, layout PipelineLayout, firstSet int, sets []DescriptorSet, dynamicOffsets []int) error
	CmdPipelineBarrier(srcStageMask, dstStageMask common.PipelineStages, dependencies common.DependencyFlags, memoryBarriers []*MemoryBarrierOptions, bufferMemoryBarriers []*BufferMemoryBarrierOptions, imageMemoryBarriers []*ImageMemoryBarrierOptions) error
	CmdCopyBufferToImage(buffer Buffer, image Image, layout common.ImageLayout, regions []*BufferImageCopy) error
	CmdBlitImage(sourceImage Image, sourceImageLayout common.ImageLayout, destinationImage Image, destinationImageLayout common.ImageLayout, regions []*ImageBlit, filter common.Filter) error
}

type CommandPool interface {
	Handle() VkCommandPool
	Destroy() error
	AllocateCommandBuffers(o *CommandBufferOptions) ([]CommandBuffer, VkResult, error)
	FreeCommandBuffers(buffers []CommandBuffer) error
}

type DescriptorPool interface {
	Handle() VkDescriptorPool
	Destroy() error
	AllocateDescriptorSets(o *DescriptorSetOptions) ([]DescriptorSet, VkResult, error)
	FreeDescriptorSets(sets []DescriptorSet) (VkResult, error)
}

type DescriptorSet interface {
	Handle() VkDescriptorSet
}

type DescriptorSetLayout interface {
	Handle() VkDescriptorSetLayout
	Destroy() error
}

type DeviceMemory interface {
	Handle() VkDeviceMemory
	MapMemory(offset int, size int, flags MemoryMapFlags) (unsafe.Pointer, VkResult, error)
	UnmapMemory() error
}

type Device interface {
	Handle() VkDevice
	Destroy() error
	Driver() Driver
	GetQueue(queueFamilyIndex int, queueIndex int) (Queue, error)
	WaitForIdle() (VkResult, error)
	WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (VkResult, error)
	ResetFences(fences []Fence) (VkResult, error)
	AllocateMemory(o *DeviceMemoryOptions) (DeviceMemory, VkResult, error)
	FreeMemory(memory DeviceMemory) error
	UpdateDescriptorSets(writes []WriteDescriptorSetOptions, copies []CopyDescriptorSetOptions) error
}

type Fence interface {
	Handle() VkFence
	Destroy() error
}

type Framebuffer interface {
	Handle() VkFramebuffer
	Destroy() error
}

type Image interface {
	Handle() VkImage
	Destroy() error
	MemoryRequirements() (*common.MemoryRequirements, error)
	BindImageMemory(memory DeviceMemory, offset int) (VkResult, error)
}

type ImageView interface {
	Handle() VkImageView
	Destroy() error
}

type Instance interface {
	Handle() VkInstance
	Destroy() error
	Driver() Driver
	PhysicalDevices() ([]PhysicalDevice, VkResult, error)
}

type Loader1_0 interface {
	Version() common.APIVersion
	Driver() Driver

	AvailableExtensions() (map[string]*common.ExtensionProperties, VkResult, error)
	AvailableLayers() (map[string]*common.LayerProperties, VkResult, error)

	CreateBuffer(device Device, o *BufferOptions) (Buffer, VkResult, error)
	CreateCommandPool(device Device, o *CommandPoolOptions) (CommandPool, VkResult, error)
	CreateDescriptorPool(device Device, o *DescriptorPoolOptions) (DescriptorPool, VkResult, error)
	CreateDescriptorSetLayout(device Device, o *DescriptorSetLayoutOptions) (DescriptorSetLayout, VkResult, error)
	CreateDevice(physicalDevice PhysicalDevice, options *DeviceOptions) (Device, VkResult, error)
	CreateFence(device Device, o *FenceOptions) (Fence, VkResult, error)
	CreateFrameBuffer(device Device, o *FramebufferOptions) (Framebuffer, VkResult, error)
	CreateGraphicsPipelines(device Device, o []*GraphicsPipelineOptions) ([]Pipeline, VkResult, error)
	CreateInstance(options *InstanceOptions) (Instance, VkResult, error)
	CreateImage(device Device, options *ImageOptions) (Image, VkResult, error)
	CreateImageView(device Device, o *ImageViewOptions) (ImageView, VkResult, error)
	CreatePipelineLayout(device Device, o *PipelineLayoutOptions) (PipelineLayout, VkResult, error)
	CreateRenderPass(device Device, o *RenderPassOptions) (RenderPass, VkResult, error)
	CreateSampler(device Device, o *SamplerOptions) (Sampler, VkResult, error)
	CreateSemaphore(device Device, o *SemaphoreOptions) (Semaphore, VkResult, error)
	CreateShaderModule(device Device, o *ShaderModuleOptions) (ShaderModule, VkResult, error)
}

type PhysicalDevice interface {
	Handle() VkPhysicalDevice
	Driver() Driver
	QueueFamilyProperties() ([]*common.QueueFamily, error)
	Properties() (*common.PhysicalDeviceProperties, error)
	Features() (*common.PhysicalDeviceFeatures, error)
	AvailableExtensions() (map[string]*common.ExtensionProperties, VkResult, error)
	MemoryProperties() (*PhysicalDeviceMemoryProperties, error)
	FormatProperties(format common.DataFormat) (*common.FormatProperties, error)
}

type Pipeline interface {
	Handle() VkPipeline
	Destroy() error
}

type PipelineLayout interface {
	Handle() VkPipelineLayout
	Destroy() error
}

type Queue interface {
	Handle() VkQueue
	Driver() Driver
	WaitForIdle() (VkResult, error)
	SubmitToQueue(fence Fence, o []*SubmitOptions) (VkResult, error)
}

type RenderPass interface {
	Handle() VkRenderPass
	Destroy() error
}

type Semaphore interface {
	Handle() VkSemaphore
	Destroy() error
}

type ShaderModule interface {
	Handle() VkShaderModule
	Destroy() error
}

type Sampler interface {
	Handle() VkSampler
	Destroy() error
}
