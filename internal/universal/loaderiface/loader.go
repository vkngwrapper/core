package loaderiface

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0/options"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/VKng/core/internal/universal"
)

type Loader1_0[Buffer iface.Buffer, BufferView iface.BufferView, CommandPool iface.CommandPool,
	DescriptorPool iface.DescriptorPool, DescriptorSet iface.DescriptorSet, Device iface.Device, Event iface.Event,
	Fence iface.Fence, Framebuffer iface.Framebuffer, Instance iface.Instance, Image iface.Image,
	ImageView iface.ImageView, PipelineCache iface.PipelineCache, PipelineLayout iface.PipelineLayout,
	QueryPool iface.QueryPool, RenderPass iface.RenderPass, Sampler iface.Sampler, Semaphore iface.Semaphore,
	ShaderModule iface.ShaderModule, CommandBuffer iface.CommandBuffer, DeviceMemory iface.DeviceMemory,
	DescriptorSetLayout iface.DescriptorSetLayout, Pipeline iface.Pipeline, PhysicalDevice iface.PhysicalDevice,
	Queue iface.Queue] interface {
	iface.Loader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event,
		Fence, Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass,
		Sampler, Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout, Pipeline, PhysicalDevice,
		Queue]

	AvailableExtensions() (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*common.ExtensionProperties, common.VkResult, error)
	AvailableLayers() (map[string]*common.LayerProperties, common.VkResult, error)

	CreateBuffer(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.BufferOptions) (Buffer, common.VkResult, error)
	CreateBufferView(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.BufferViewOptions) (BufferView, common.VkResult, error)
	CreateCommandPool(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.CommandPoolOptions) (CommandPool, common.VkResult, error)
	CreateDescriptorPool(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.DescriptorPoolOptions) (DescriptorPool, common.VkResult, error)
	CreateDescriptorSetLayout(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.DescriptorSetLayoutOptions) (DescriptorSetLayout, common.VkResult, error)
	CreateDevice(physicalDevice iface.PhysicalDevice, allocationCallbacks *driver.AllocationCallbacks, options *options.DeviceOptions) (Device, common.VkResult, error)
	CreateEvent(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, options *options.EventOptions) (Event, common.VkResult, error)
	CreateFence(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.FenceOptions) (Fence, common.VkResult, error)
	CreateFrameBuffer(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.FramebufferOptions) (Framebuffer, common.VkResult, error)
	CreateGraphicsPipelines(device iface.Device, pipelineCache iface.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []options.GraphicsPipelineOptions) ([]Pipeline, common.VkResult, error)
	CreateComputePipelines(device iface.Device, pipelineCache iface.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []options.ComputePipelineOptions) ([]Pipeline, common.VkResult, error)
	CreateInstance(allocationCallbacks *driver.AllocationCallbacks, options *options.InstanceOptions) (Instance, common.VkResult, error)
	CreateImage(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, options *options.ImageOptions) (Image, common.VkResult, error)
	CreateImageView(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.ImageViewOptions) (ImageView, common.VkResult, error)
	CreatePipelineCache(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.PipelineCacheOptions) (PipelineCache, common.VkResult, error)
	CreatePipelineLayout(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.PipelineLayoutOptions) (PipelineLayout, common.VkResult, error)
	CreateQueryPool(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.QueryPoolOptions) (QueryPool, common.VkResult, error)
	CreateRenderPass(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.RenderPassOptions) (RenderPass, common.VkResult, error)
	CreateSampler(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.SamplerOptions) (Sampler, common.VkResult, error)
	CreateSemaphore(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.SemaphoreOptions) (Semaphore, common.VkResult, error)
	CreateShaderModule(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.ShaderModuleOptions) (ShaderModule, common.VkResult, error)

	PhysicalDevices(instance iface.Instance) ([]PhysicalDevice, common.VkResult, error)
	DeviceQueue(device iface.Device, queueFamilyIndex int, queueIndex int) Queue

	AllocateCommandBuffers(o *options.CommandBufferOptions) ([]CommandBuffer, common.VkResult, error)
	FreeCommandBuffers(buffers []iface.CommandBuffer)
	AllocateDescriptorSets(o *options.DescriptorSetOptions) ([]DescriptorSet, common.VkResult, error)
	FreeDescriptorSets(sets []iface.DescriptorSet) (common.VkResult, error)
	AllocateMemory(device iface.Device, allocationCallbacks *driver.AllocationCallbacks, o *options.DeviceMemoryOptions) (DeviceMemory, common.VkResult, error)
	FreeDeviceMemory(deviceMemory iface.DeviceMemory, allocationCallbacks *driver.AllocationCallbacks)
}

func CreateLoaderFromDriver[Buffer iface.Buffer, BufferView iface.BufferView, CommandPool iface.CommandPool,
	DescriptorPool iface.DescriptorPool, DescriptorSet iface.DescriptorSet, Device iface.Device, Event iface.Event,
	Fence iface.Fence, Framebuffer iface.Framebuffer, Instance iface.Instance, Image iface.Image,
	ImageView iface.ImageView, PipelineCache iface.PipelineCache, PipelineLayout iface.PipelineLayout,
	QueryPool iface.QueryPool, RenderPass iface.RenderPass, Sampler iface.Sampler, Semaphore iface.Semaphore,
	ShaderModule iface.ShaderModule, CommandBuffer iface.CommandBuffer, DeviceMemory iface.DeviceMemory,
	DescriptorSetLayout iface.DescriptorSetLayout, Pipeline iface.Pipeline, PhysicalDevice iface.PhysicalDevice,
	Queue iface.Queue](driver driver.Driver) *universal.VulkanLoader[Buffer,
	BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event, Fence, Framebuffer, Instance, Image,
	ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass, Sampler, Semaphore, ShaderModule,
	CommandBuffer, DeviceMemory, DescriptorSetLayout, Pipeline, PhysicalDevice, Queue] {

	return universal.NewLoader[Buffer, BufferView, CommandPool, DescriptorPool, DescriptorSet, Device, Event,
		Fence, Framebuffer, Instance, Image, ImageView, PipelineCache, PipelineLayout, QueryPool, RenderPass,
		Sampler, Semaphore, ShaderModule, CommandBuffer, DeviceMemory, DescriptorSetLayout, Pipeline, PhysicalDevice,
		Queue](driver)
}
