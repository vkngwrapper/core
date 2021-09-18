package core

import (
	"github.com/CannibalVox/VKng/core/common"
	"time"
	"unsafe"
)

//go:generate mockgen -source iface.go -destination ../mocks/core.go -package=mocks

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
	Destroy() error

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
}

type CommandPool interface {
	Handle() VkCommandPool
	Destroy() error
	DestroyBuffers(buffers []CommandBuffer) error
}

type DescriptorPool interface {
	Handle() VkDescriptorPool
	Destroy() error
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
	Free() error
	MapMemory(offset int, size int) (unsafe.Pointer, VkResult, error)
	UnmapMemory() error
	WriteData(offset int, data interface{}) (VkResult, error)
}

type Device interface {
	Handle() VkDevice
	Destroy() error
	Driver() Driver
	GetQueue(queueFamilyIndex int, queueIndex int) (Queue, error)
	CreateShaderModule(o *ShaderModuleOptions) (ShaderModule, VkResult, error)
	CreateImageView(o *ImageViewOptions) (ImageView, VkResult, error)
	CreateSemaphore(o *SemaphoreOptions) (Semaphore, VkResult, error)
	CreateFence(o *FenceOptions) (Fence, VkResult, error)
	WaitForIdle() (VkResult, error)
	WaitForFences(waitForAll bool, timeout time.Duration, fences []Fence) (VkResult, error)
	ResetFences(fences []Fence) (VkResult, error)
	CreateBuffer(o *BufferOptions) (Buffer, VkResult, error)
	AllocateMemory(o *DeviceMemoryOptions) (DeviceMemory, VkResult, error)
	CreateDescriptorSetLayout(o *DescriptorSetLayoutOptions) (DescriptorSetLayout, VkResult, error)
	CreateDescriptorPool(o *DescriptorPoolOptions) (DescriptorPool, VkResult, error)
	AllocateDescriptorSet(o *DescriptorSetOptions) ([]DescriptorSet, VkResult, error)
	UpdateDescriptorSets(writes []*WriteDescriptorSetOptions, copies []*CopyDescriptorSetOptions) error
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

type PhysicalDevice interface {
	Handle() VkPhysicalDevice
	QueueFamilyProperties() ([]*common.QueueFamily, error)
	CreateDevice(options *DeviceOptions) (Device, VkResult, error)
	Properties() (*common.PhysicalDeviceProperties, error)
	Features() (*common.PhysicalDeviceFeatures, error)
	AvailableExtensions() (map[string]*common.ExtensionProperties, VkResult, error)
	MemoryProperties() *PhysicalDeviceMemoryProperties
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
}
