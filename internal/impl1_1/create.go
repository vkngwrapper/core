package impl1_1

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
)

func CreateInstanceObject(instanceDriver driver.Driver, handle driver.VkInstance, version common.APIVersion, instanceExtensions []string) core1_0.Instance {
	instance := &VulkanInstance{
		VulkanInstance: impl1_0.VulkanInstance{
			InstanceDriver:           instanceDriver,
			InstanceHandle:           handle,
			MaximumVersion:           version,
			ActiveInstanceExtensions: make(map[string]struct{}),
			InstanceObjectBuilder:    &InstanceObjectBuilderImpl{},
		},
	}

	for _, extension := range instanceExtensions {
		instance.ActiveInstanceExtensions[extension] = struct{}{}
	}

	return instance
}

type InstanceObjectBuilderImpl struct{}

func (b *InstanceObjectBuilderImpl) CreatePhysicalDeviceObject(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, instanceVersion, deviceVersion common.APIVersion) core1_0.PhysicalDevice {
	return &VulkanPhysicalDevice{
		VulkanPhysicalDevice: impl1_0.VulkanPhysicalDevice{
			InstanceDriver:        coreDriver,
			PhysicalDeviceHandle:  handle,
			InstanceVersion:       instanceVersion,
			MaximumDeviceVersion:  deviceVersion,
			InstanceObjectBuilder: b,
		},
	}
}

func CreateDeviceObject(deviceDriver driver.Driver, handle driver.VkDevice, version common.APIVersion, deviceExtensionNames []string) core1_0.Device {
	if version < common.Vulkan1_1 {
		return impl1_0.CreateDeviceObject(deviceDriver, handle, version, deviceExtensionNames)
	}

	device := &VulkanDevice{
		VulkanDevice: impl1_0.VulkanDevice{
			DeviceDriver:           deviceDriver,
			DeviceHandle:           handle,
			MaximumAPIVersion:      version,
			ActiveDeviceExtensions: make(map[string]struct{}),
			DeviceObjectBuilder:    &DeviceObjectBuilderImpl{},
		},
		DeviceObjectBuilder: &DeviceObjectBuilderImpl1_1{},
	}

	for _, extension := range deviceExtensionNames {
		device.ActiveDeviceExtensions[extension] = struct{}{}
	}

	return device
}

func (b *InstanceObjectBuilderImpl) CreateDeviceObject(deviceDriver driver.Driver, handle driver.VkDevice, version common.APIVersion, deviceExtensionNames []string) core1_0.Device {
	return CreateDeviceObject(deviceDriver, handle, version, deviceExtensionNames)
}

type DeviceObjectBuilderImpl struct{}

func (b *DeviceObjectBuilderImpl) CreateBufferObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBuffer, version common.APIVersion) core1_0.Buffer {
	return &VulkanBuffer{
		VulkanBuffer: impl1_0.VulkanBuffer{
			DeviceDriver:      coreDriver,
			Device:            device,
			BufferHandle:      handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateBufferViewObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBufferView, version common.APIVersion) core1_0.BufferView {
	return &VulkanBufferView{
		VulkanBufferView: impl1_0.VulkanBufferView{
			DeviceDriver:      coreDriver,
			Device:            device,
			BufferViewHandle:  handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateCommandBufferObject(coreDriver driver.Driver, commandPool driver.VkCommandPool, device driver.VkDevice, handle driver.VkCommandBuffer, version common.APIVersion) core1_0.CommandBuffer {
	return &VulkanCommandBuffer{
		VulkanCommandBuffer: impl1_0.VulkanCommandBuffer{
			DeviceDriver:        coreDriver,
			Device:              device,
			CommandBufferHandle: handle,
			CommandPool:         commandPool,
			MaximumAPIVersion:   version,

			Counter: &core1_0.CommandCounter{},
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateCommandPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkCommandPool, version common.APIVersion) core1_0.CommandPool {
	return &VulkanCommandPool{
		VulkanCommandPool: impl1_0.VulkanCommandPool{
			DeviceDriver:      coreDriver,
			Device:            device,
			CommandPoolHandle: handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateDescriptorPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorPool, version common.APIVersion) core1_0.DescriptorPool {
	return &VulkanDescriptorPool{
		VulkanDescriptorPool: impl1_0.VulkanDescriptorPool{
			DeviceDriver:         coreDriver,
			Device:               device,
			DescriptorPoolHandle: handle,
			MaximumAPIVersion:    version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateDescriptorSetObject(coreDriver driver.Driver, device driver.VkDevice, descriptorPool driver.VkDescriptorPool, handle driver.VkDescriptorSet, version common.APIVersion) core1_0.DescriptorSet {
	return &VulkanDescriptorSet{
		VulkanDescriptorSet: impl1_0.VulkanDescriptorSet{
			DeviceDriver:        coreDriver,
			Device:              device,
			DescriptorSetHandle: handle,
			MaximumAPIVersion:   version,
			DescriptorPool:      descriptorPool,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateDescriptorSetLayoutObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorSetLayout, version common.APIVersion) core1_0.DescriptorSetLayout {
	return &VulkanDescriptorSetLayout{
		VulkanDescriptorSetLayout: impl1_0.VulkanDescriptorSetLayout{
			DeviceDriver:              coreDriver,
			Device:                    device,
			DescriptorSetLayoutHandle: handle,
			MaximumAPIVersion:         version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateDeviceMemoryObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDeviceMemory, version common.APIVersion, size int) core1_0.DeviceMemory {
	return &VulkanDeviceMemory{
		VulkanDeviceMemory: impl1_0.VulkanDeviceMemory{
			DeviceDriver:       coreDriver,
			Device:             device,
			DeviceMemoryHandle: handle,
			MaximumAPIVersion:  version,
			Size:               size,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateEventObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkEvent, version common.APIVersion) core1_0.Event {
	return &VulkanEvent{
		VulkanEvent: impl1_0.VulkanEvent{
			DeviceDriver:      coreDriver,
			Device:            device,
			EventHandle:       handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateFenceObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFence, version common.APIVersion) core1_0.Fence {
	return &VulkanFence{
		VulkanFence: impl1_0.VulkanFence{
			DeviceDriver:      coreDriver,
			Device:            device,
			FenceHandle:       handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateFramebufferObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFramebuffer, version common.APIVersion) core1_0.Framebuffer {
	return &VulkanFramebuffer{
		VulkanFramebuffer: impl1_0.VulkanFramebuffer{
			DeviceDriver:      coreDriver,
			Device:            device,
			FramebufferHandle: handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateImageObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImage, version common.APIVersion) core1_0.Image {
	return &VulkanImage{
		VulkanImage: impl1_0.VulkanImage{
			DeviceDriver:      coreDriver,
			Device:            device,
			ImageHandle:       handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateImageViewObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImageView, version common.APIVersion) core1_0.ImageView {
	return &VulkanImageView{
		VulkanImageView: impl1_0.VulkanImageView{
			DeviceDriver:      coreDriver,
			Device:            device,
			ImageViewHandle:   handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreatePipelineObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipeline, version common.APIVersion) core1_0.Pipeline {
	return &VulkanPipeline{
		VulkanPipeline: impl1_0.VulkanPipeline{
			DeviceDriver:      coreDriver,
			Device:            device,
			PipelineHandle:    handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreatePipelineCacheObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineCache, version common.APIVersion) core1_0.PipelineCache {
	return &VulkanPipelineCache{
		VulkanPipelineCache: impl1_0.VulkanPipelineCache{
			DeviceDriver:        coreDriver,
			Device:              device,
			PipelineCacheHandle: handle,
			MaximumAPIVersion:   version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreatePipelineLayoutObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineLayout, version common.APIVersion) core1_0.PipelineLayout {
	return &VulkanPipelineLayout{
		VulkanPipelineLayout: impl1_0.VulkanPipelineLayout{
			DeviceDriver:         coreDriver,
			Device:               device,
			PipelineLayoutHandle: handle,
			MaximumAPIVersion:    version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateQueryPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueryPool, version common.APIVersion) core1_0.QueryPool {
	return &VulkanQueryPool{
		VulkanQueryPool: impl1_0.VulkanQueryPool{
			DeviceDriver:      coreDriver,
			Device:            device,
			QueryPoolHandle:   handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateQueueObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) core1_0.Queue {
	return &VulkanQueue{
		VulkanQueue: impl1_0.VulkanQueue{
			DeviceDriver:      coreDriver,
			QueueHandle:       handle,
			Device:            device,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateRenderPassObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkRenderPass, version common.APIVersion) core1_0.RenderPass {
	return &VulkanRenderPass{
		VulkanRenderPass: impl1_0.VulkanRenderPass{
			DeviceDriver:      coreDriver,
			Device:            device,
			RenderPassHandle:  handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateSamplerObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSampler, version common.APIVersion) core1_0.Sampler {
	return &VulkanSampler{
		VulkanSampler: impl1_0.VulkanSampler{
			DeviceDriver:      coreDriver,
			Device:            device,
			SamplerHandle:     handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateSemaphoreObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSemaphore, version common.APIVersion) core1_0.Semaphore {
	return &VulkanSemaphore{
		VulkanSemaphore: impl1_0.VulkanSemaphore{
			DeviceDriver:      coreDriver,
			Device:            device,
			SemaphoreHandle:   handle,
			MaximumAPIVersion: version,
		},
	}
}

func (b *DeviceObjectBuilderImpl) CreateShaderModuleObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkShaderModule, version common.APIVersion) core1_0.ShaderModule {
	return &VulkanShaderModule{
		VulkanShaderModule: impl1_0.VulkanShaderModule{
			DeviceDriver:       coreDriver,
			Device:             device,
			ShaderModuleHandle: handle,
			MaximumAPIVersion:  version,
		},
	}
}

type DeviceObjectBuilderImpl1_1 struct{}

func (b *DeviceObjectBuilderImpl1_1) CreateDescriptorUpdateTemplate(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorUpdateTemplate, version common.APIVersion) core1_1.DescriptorUpdateTemplate {
	return &VulkanDescriptorUpdateTemplate{
		DeviceDriver:             coreDriver,
		Device:                   device,
		DescriptorTemplateHandle: handle,
		MaximumAPIVersion:        version,
	}
}

func (b *DeviceObjectBuilderImpl1_1) CreateSamplerYcbcrConversion(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSamplerYcbcrConversion, version common.APIVersion) core1_1.SamplerYcbcrConversion {
	return &VulkanSamplerYcbcrConversion{
		DeviceDriver:      coreDriver,
		Device:            device,
		YcbcrHandle:       handle,
		MaximumAPIVersion: version,
	}
}
