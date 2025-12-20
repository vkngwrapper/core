package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

func CreateInstanceObject(instanceDriver driver.Driver, handle driver.VkInstance, version common.APIVersion, instanceExtensions []string) core1_0.Instance {
	instance := &VulkanInstance{
		InstanceDriver:           instanceDriver,
		InstanceHandle:           handle,
		MaximumVersion:           version,
		ActiveInstanceExtensions: make(map[string]struct{}),
		InstanceObjectBuilder:    &InstanceObjectBuilderImpl{},
	}

	for _, extension := range instanceExtensions {
		instance.ActiveInstanceExtensions[extension] = struct{}{}
	}

	return instance
}

type InstanceObjectBuilderImpl struct {
}

func (b *InstanceObjectBuilderImpl) CreatePhysicalDeviceObject(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, instanceVersion, deviceVersion common.APIVersion) core1_0.PhysicalDevice {
	return &VulkanPhysicalDevice{
		InstanceDriver:        coreDriver,
		PhysicalDeviceHandle:  handle,
		InstanceVersion:       instanceVersion,
		MaximumDeviceVersion:  deviceVersion,
		InstanceObjectBuilder: b,
	}
}

func CreateDeviceObject(deviceDriver driver.Driver, handle driver.VkDevice, version common.APIVersion, deviceExtensionNames []string) core1_0.Device {
	device := &VulkanDevice{
		DeviceDriver:           deviceDriver,
		DeviceHandle:           handle,
		MaximumAPIVersion:      version,
		ActiveDeviceExtensions: make(map[string]struct{}),
		DeviceObjectBuilder:    &DeviceObjectBuilderImpl{},
	}

	for _, extension := range deviceExtensionNames {
		device.ActiveDeviceExtensions[extension] = struct{}{}
	}

	return device
}

func (b *InstanceObjectBuilderImpl) CreateDeviceObject(deviceDriver driver.Driver, handle driver.VkDevice, version common.APIVersion, deviceExtensionNames []string) core1_0.Device {
	return CreateDeviceObject(deviceDriver, handle, version, deviceExtensionNames)
}

type DeviceObjectBuilderImpl struct {
}

func (b *DeviceObjectBuilderImpl) CreateBufferObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBuffer, version common.APIVersion) core1_0.Buffer {
	return &VulkanBuffer{
		DeviceDriver:      coreDriver,
		Device:            device,
		BufferHandle:      handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateBufferViewObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBufferView, version common.APIVersion) core1_0.BufferView {
	return &VulkanBufferView{
		DeviceDriver:      coreDriver,
		Device:            device,
		BufferViewHandle:  handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateCommandBufferObject(coreDriver driver.Driver, commandPool driver.VkCommandPool, device driver.VkDevice, handle driver.VkCommandBuffer, version common.APIVersion) core1_0.CommandBuffer {
	return &VulkanCommandBuffer{
		DeviceDriver:        coreDriver,
		Device:              device,
		CommandBufferHandle: handle,
		CommandPool:         commandPool,
		MaximumAPIVersion:   version,

		Counter: &core1_0.CommandCounter{},
	}
}

func (b *DeviceObjectBuilderImpl) CreateCommandPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkCommandPool, version common.APIVersion) core1_0.CommandPool {
	return &VulkanCommandPool{
		DeviceDriver:      coreDriver,
		Device:            device,
		CommandPoolHandle: handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateDescriptorPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorPool, version common.APIVersion) core1_0.DescriptorPool {
	return &VulkanDescriptorPool{
		DeviceDriver:         coreDriver,
		Device:               device,
		DescriptorPoolHandle: handle,
		MaximumAPIVersion:    version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateDescriptorSetObject(coreDriver driver.Driver, device driver.VkDevice, descriptorPool driver.VkDescriptorPool, handle driver.VkDescriptorSet, version common.APIVersion) core1_0.DescriptorSet {
	return &VulkanDescriptorSet{
		DeviceDriver:        coreDriver,
		Device:              device,
		DescriptorSetHandle: handle,
		MaximumAPIVersion:   version,
		DescriptorPool:      descriptorPool,
	}
}

func (b *DeviceObjectBuilderImpl) CreateDescriptorSetLayoutObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorSetLayout, version common.APIVersion) core1_0.DescriptorSetLayout {
	return &VulkanDescriptorSetLayout{
		DeviceDriver:              coreDriver,
		Device:                    device,
		DescriptorSetLayoutHandle: handle,
		MaximumAPIVersion:         version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateDeviceMemoryObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDeviceMemory, version common.APIVersion, size int) core1_0.DeviceMemory {
	return &VulkanDeviceMemory{
		DeviceDriver:       coreDriver,
		Device:             device,
		DeviceMemoryHandle: handle,
		MaximumAPIVersion:  version,
		Size:               size,
	}
}

func (b *DeviceObjectBuilderImpl) CreateEventObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkEvent, version common.APIVersion) core1_0.Event {
	return &VulkanEvent{
		DeviceDriver:      coreDriver,
		Device:            device,
		EventHandle:       handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateFenceObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFence, version common.APIVersion) core1_0.Fence {
	return &VulkanFence{
		DeviceDriver:      coreDriver,
		Device:            device,
		FenceHandle:       handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateFramebufferObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFramebuffer, version common.APIVersion) core1_0.Framebuffer {
	return &VulkanFramebuffer{
		DeviceDriver:      coreDriver,
		Device:            device,
		FramebufferHandle: handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateImageObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImage, version common.APIVersion) core1_0.Image {
	return &VulkanImage{
		DeviceDriver:      coreDriver,
		Device:            device,
		ImageHandle:       handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateImageViewObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImageView, version common.APIVersion) core1_0.ImageView {
	return &VulkanImageView{
		DeviceDriver:      coreDriver,
		Device:            device,
		ImageViewHandle:   handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreatePipelineObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipeline, version common.APIVersion) core1_0.Pipeline {
	return &VulkanPipeline{
		DeviceDriver:      coreDriver,
		Device:            device,
		PipelineHandle:    handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreatePipelineCacheObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineCache, version common.APIVersion) core1_0.PipelineCache {
	return &VulkanPipelineCache{
		DeviceDriver:        coreDriver,
		Device:              device,
		PipelineCacheHandle: handle,
		MaximumAPIVersion:   version,
	}
}

func (b *DeviceObjectBuilderImpl) CreatePipelineLayoutObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineLayout, version common.APIVersion) core1_0.PipelineLayout {
	return &VulkanPipelineLayout{
		DeviceDriver:         coreDriver,
		Device:               device,
		PipelineLayoutHandle: handle,
		MaximumAPIVersion:    version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateQueryPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueryPool, version common.APIVersion) core1_0.QueryPool {
	return &VulkanQueryPool{
		DeviceDriver:      coreDriver,
		Device:            device,
		QueryPoolHandle:   handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateQueueObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) core1_0.Queue {
	return &VulkanQueue{
		DeviceDriver:      coreDriver,
		QueueHandle:       handle,
		Device:            device,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateRenderPassObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkRenderPass, version common.APIVersion) core1_0.RenderPass {
	return &VulkanRenderPass{
		DeviceDriver:      coreDriver,
		Device:            device,
		RenderPassHandle:  handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateSamplerObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSampler, version common.APIVersion) core1_0.Sampler {
	return &VulkanSampler{
		DeviceDriver:      coreDriver,
		Device:            device,
		SamplerHandle:     handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateSemaphoreObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSemaphore, version common.APIVersion) core1_0.Semaphore {
	return &VulkanSemaphore{
		DeviceDriver:      coreDriver,
		Device:            device,
		SemaphoreHandle:   handle,
		MaximumAPIVersion: version,
	}
}

func (b *DeviceObjectBuilderImpl) CreateShaderModuleObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkShaderModule, version common.APIVersion) core1_0.ShaderModule {
	return &VulkanShaderModule{
		DeviceDriver:       coreDriver,
		Device:             device,
		ShaderModuleHandle: handle,
		MaximumAPIVersion:  version,
	}
}
