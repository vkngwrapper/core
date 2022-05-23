package core

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	internal1_0 "github.com/CannibalVox/VKng/core/internal/core1_0"
	internal1_1 "github.com/CannibalVox/VKng/core/internal/core1_1"
)

func CreateInstance(instanceDriver driver.Driver, handle driver.VkInstance, version common.APIVersion) core1_0.Instance {
	return internal1_0.CreateInstanceObject(instanceDriver, handle, version)
}

func CreatePhysicalDevice(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, instanceVersion, deviceVersion common.APIVersion) core1_0.PhysicalDevice {
	return internal1_0.CreatePhysicalDeviceObject(coreDriver, instance, handle, instanceVersion, deviceVersion)
}

func CreateDevice(deviceDriver driver.Driver, handle driver.VkDevice, version common.APIVersion) core1_0.Device {
	return internal1_0.CreateDeviceObject(deviceDriver, handle, version)
}

func CreateBuffer(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBuffer, version common.APIVersion) core1_0.Buffer {
	return internal1_0.CreateBufferObject(coreDriver, device, handle, version)
}

func CreateBufferView(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBufferView, version common.APIVersion) core1_0.BufferView {
	return internal1_0.CreateBufferViewObject(coreDriver, device, handle, version)
}

func CreateCommandBuffer(coreDriver driver.Driver, commandPool driver.VkCommandPool, device driver.VkDevice, handle driver.VkCommandBuffer, version common.APIVersion) core1_0.CommandBuffer {
	return internal1_0.CreateCommandBufferObject(coreDriver, commandPool, device, handle, version)
}

func CreateCommandPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkCommandPool, version common.APIVersion) core1_0.CommandPool {
	return internal1_0.CreateCommandPoolObject(coreDriver, device, handle, version)
}

func CreateDescriptorPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorPool, version common.APIVersion) core1_0.DescriptorPool {
	return internal1_0.CreateDescriptorPoolObject(coreDriver, device, handle, version)
}

func CreateDescriptorSet(coreDriver driver.Driver, device driver.VkDevice, descriptorPool driver.VkDescriptorPool, handle driver.VkDescriptorSet, version common.APIVersion) core1_0.DescriptorSet {
	return internal1_0.CreateDescriptorSetObject(coreDriver, device, descriptorPool, handle, version)
}

func CreateDescriptorSetLayout(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorSetLayout, version common.APIVersion) core1_0.DescriptorSetLayout {
	return internal1_0.CreateDescriptorSetLayoutObject(coreDriver, device, handle, version)
}

func CreateDescriptorUpdateTemplate(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorUpdateTemplate, version common.APIVersion) core1_1.DescriptorUpdateTemplate {
	return internal1_1.CreateDescriptorUpdateTemplate(coreDriver, device, handle, version)
}

func CreateDeviceMemory(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDeviceMemory, version common.APIVersion, size int) core1_0.DeviceMemory {
	return internal1_0.CreateDeviceMemoryObject(coreDriver, device, handle, version, size)
}

func CreateEvent(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkEvent, version common.APIVersion) core1_0.Event {
	return internal1_0.CreateEventObject(coreDriver, device, handle, version)
}

func CreateFence(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFence, version common.APIVersion) core1_0.Fence {
	return internal1_0.CreateFenceObject(coreDriver, device, handle, version)
}

func CreateFramebuffer(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFramebuffer, version common.APIVersion) core1_0.Framebuffer {
	return internal1_0.CreateFramebufferObject(coreDriver, device, handle, version)
}

func CreateImage(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImage, version common.APIVersion) core1_0.Image {
	return internal1_0.CreateImageObject(coreDriver, device, handle, version)
}

func CreateImageView(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImageView, version common.APIVersion) core1_0.ImageView {
	return internal1_0.CreateImageViewObject(coreDriver, device, handle, version)
}

func CreatePipeline(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipeline, version common.APIVersion) core1_0.Pipeline {
	return internal1_0.CreatePipelineObject(coreDriver, device, handle, version)
}

func CreatePipelineCache(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineCache, version common.APIVersion) core1_0.PipelineCache {
	return internal1_0.CreatePipelineCacheObject(coreDriver, device, handle, version)
}

func CreatePipelineLayout(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineLayout, version common.APIVersion) core1_0.PipelineLayout {
	return internal1_0.CreatePipelineLayoutObject(coreDriver, device, handle, version)
}

func CreateQueryPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueryPool, version common.APIVersion) core1_0.QueryPool {
	return internal1_0.CreateQueryPoolObject(coreDriver, device, handle, version)
}

func CreateQueue(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) core1_0.Queue {
	return internal1_0.CreateQueueObject(coreDriver, device, handle, version)
}

func CreateRenderPass(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkRenderPass, version common.APIVersion) core1_0.RenderPass {
	return internal1_0.CreateRenderPassObject(coreDriver, device, handle, version)
}

func CreateSampler(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSampler, version common.APIVersion) core1_0.Sampler {
	return internal1_0.CreateSamplerObject(coreDriver, device, handle, version)
}

func CreateSamplerYcbcrConversion(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSamplerYcbcrConversion, version common.APIVersion) core1_1.SamplerYcbcrConversion {
	return internal1_1.CreateSamplerYcbcrConversion(coreDriver, device, handle, version)
}

func CreateSemaphore(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSemaphore, version common.APIVersion) core1_0.Semaphore {
	return internal1_0.CreateSemaphoreObject(coreDriver, device, handle, version)
}

func CreateShaderModule(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkShaderModule, version common.APIVersion) core1_0.ShaderModule {
	return internal1_0.CreateShaderModuleObject(coreDriver, device, handle, version)
}
