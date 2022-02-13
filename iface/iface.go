package iface

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type Buffer interface {
	Handle() driver.VkBuffer
	Destroy(callbacks *driver.AllocationCallbacks)
}

type BufferView interface {
	Handle() driver.VkBufferView
	Destroy(callbacks *driver.AllocationCallbacks)
}

type CommandBuffer interface {
	Handle() driver.VkCommandBuffer
	Driver() driver.Driver
	DeviceHandle() driver.VkDevice
	CommandPoolHandle() driver.VkCommandPool
}

type CommandPool interface {
	Handle() driver.VkCommandPool
	Device() driver.VkDevice
	Driver() driver.Driver
	Destroy(callbacks *driver.AllocationCallbacks)
}

type DescriptorPool interface {
	Handle() driver.VkDescriptorPool
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
	Destroy(callbacks *driver.AllocationCallbacks)
}

type DescriptorSet interface {
	Handle() driver.VkDescriptorSet
	PoolHandle() driver.VkDescriptorPool
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
}

type DescriptorSetLayout interface {
	Handle() driver.VkDescriptorSetLayout
	Destroy(callbacks *driver.AllocationCallbacks)
}

type DeviceMemory interface {
	Handle() driver.VkDeviceMemory
	DeviceHandle() driver.VkDevice
	Driver() driver.Driver
}

type Device interface {
	Handle() driver.VkDevice
	Driver() driver.Driver
	Destroy(callbacks *driver.AllocationCallbacks)
}

type Event interface {
	Handle() driver.VkEvent
	Destroy(callbacks *driver.AllocationCallbacks)
}

type Fence interface {
	Handle() driver.VkFence
	Destroy(callbacks *driver.AllocationCallbacks)
}

type Framebuffer interface {
	Handle() driver.VkFramebuffer
	Destroy(callbacks *driver.AllocationCallbacks)
}

type Image interface {
	Handle() driver.VkImage
	Destroy(callbacks *driver.AllocationCallbacks)
}

type ImageView interface {
	Handle() driver.VkImageView
	Destroy(callbacks *driver.AllocationCallbacks)
}

type Instance interface {
	Handle() driver.VkInstance
	Driver() driver.Driver
	Destroy(callbacks *driver.AllocationCallbacks)
}

type Loader[B Buffer, BV BufferView, CP CommandPool,
	DP DescriptorPool, DS DescriptorSet, D Device, E Event,
	F Fence, FB Framebuffer, I Instance, IM Image,
	IMV ImageView, PC PipelineCache, PL PipelineLayout,
	QP QueryPool, RP RenderPass, SA Sampler, S Semaphore,
	SM ShaderModule, CB CommandBuffer, DM DeviceMemory,
	DSL DescriptorSetLayout, P Pipeline, PD PhysicalDevice,
	Q Queue] interface {
	Driver() driver.Driver
	Version() common.APIVersion
}

type PhysicalDevice interface {
	Handle() driver.VkPhysicalDevice
	Driver() driver.Driver
}

type Pipeline interface {
	Handle() driver.VkPipeline
	Destroy(callbacks *driver.AllocationCallbacks)
}

type PipelineCache interface {
	Handle() driver.VkPipelineCache
	Destroy(callbacks *driver.AllocationCallbacks)
}

type PipelineLayout interface {
	Handle() driver.VkPipelineLayout
	Destroy(callbacks *driver.AllocationCallbacks)
}

type QueryPool interface {
	Handle() driver.VkQueryPool
	Destroy(callbacks *driver.AllocationCallbacks)
}

type Queue interface {
	Handle() driver.VkQueue
	Driver() driver.Driver
}

type RenderPass interface {
	Handle() driver.VkRenderPass
	Destroy(callbacks *driver.AllocationCallbacks)
}

type Semaphore interface {
	Handle() driver.VkSemaphore
	Destroy(callbacks *driver.AllocationCallbacks)
}

type ShaderModule interface {
	Handle() driver.VkShaderModule
	Destroy(callbacks *driver.AllocationCallbacks)
}

type Sampler interface {
	Handle() driver.VkSampler
	Destroy(callbacks *driver.AllocationCallbacks)
}
