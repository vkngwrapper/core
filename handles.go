package core

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/internal/objects"
)

func CreateInstance(instanceDriver driver.Driver, handle driver.VkInstance, version common.APIVersion) core1_0.Instance {
	return objects.CreateInstance(instanceDriver, handle, version)
}

func CreatePhysicalDevice(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, version common.APIVersion) core1_0.PhysicalDevice {
	return objects.CreatePhysicalDevice(coreDriver, instance, handle, version)
}

func CreateDevice(deviceDriver driver.Driver, handle driver.VkDevice, version common.APIVersion) core1_0.Device {
	return objects.CreateDevice(deviceDriver, handle, version)
}

func CreateBuffer(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBuffer, version common.APIVersion) core1_0.Buffer {
	return objects.CreateBuffer(coreDriver, device, handle, version)
}

func CreateBufferView(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBufferView, version common.APIVersion) core1_0.BufferView {
	return objects.CreateBufferView(coreDriver, device, handle, version)
}

func CreateCommandBuffer(coreDriver driver.Driver, commandPool driver.VkCommandPool, device driver.VkDevice, handle driver.VkCommandBuffer, version common.APIVersion) core1_0.CommandBuffer {
	return objects.CreateCommandBuffer(coreDriver, commandPool, device, handle, version)
}

func CreateCommandPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkCommandPool, version common.APIVersion) core1_0.CommandPool {
	return objects.CreateCommandPool(coreDriver, device, handle, version)
}

func CreateDescriptorPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorPool, version common.APIVersion) core1_0.DescriptorPool {
	return objects.CreateDescriptorPool(coreDriver, device, handle, version)
}

func CreateDescriptorSet(coreDriver driver.Driver, device driver.VkDevice, descriptorPool driver.VkDescriptorPool, handle driver.VkDescriptorSet, version common.APIVersion) core1_0.DescriptorSet {
	return objects.CreateDescriptorSet(coreDriver, device, descriptorPool, handle, version)
}

func CreateDescriptorSetLayout(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorSetLayout, version common.APIVersion) core1_0.DescriptorSetLayout {
	return objects.CreateDescriptorSetLayout(coreDriver, device, handle, version)
}

func CreateDeviceMemory(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDeviceMemory, version common.APIVersion, size int) core1_0.DeviceMemory {
	return objects.CreateDeviceMemory(coreDriver, device, handle, version, size)
}

func CreateEvent(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkEvent, version common.APIVersion) core1_0.Event {
	return objects.CreateEvent(coreDriver, device, handle, version)
}

func CreateFence(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFence, version common.APIVersion) core1_0.Fence {
	return objects.CreateFence(coreDriver, device, handle, version)
}

func CreateFramebuffer(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFramebuffer, version common.APIVersion) core1_0.Framebuffer {
	return objects.CreateFramebuffer(coreDriver, device, handle, version)
}

func CreateImage(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImage, version common.APIVersion) core1_0.Image {
	return objects.CreateImage(coreDriver, device, handle, version)
}

func CreateImageView(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImageView, version common.APIVersion) core1_0.ImageView {
	return objects.CreateImageView(coreDriver, device, handle, version)
}

func CreatePipeline(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipeline, version common.APIVersion) core1_0.Pipeline {
	return objects.CreatePipeline(coreDriver, device, handle, version)
}

func CreatePipelineCache(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineCache, version common.APIVersion) core1_0.PipelineCache {
	return objects.CreatePipelineCache(coreDriver, device, handle, version)
}

func CreatePipelineLayout(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineLayout, version common.APIVersion) core1_0.PipelineLayout {
	return objects.CreatePipelineLayout(coreDriver, device, handle, version)
}

func CreateQueryPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueryPool, version common.APIVersion) core1_0.QueryPool {
	return objects.CreateQueryPool(coreDriver, device, handle, version)
}

func CreateQueue(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) core1_0.Queue {
	return objects.CreateQueue(coreDriver, device, handle, version)
}

func CreateRenderPass(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkRenderPass, version common.APIVersion) core1_0.RenderPass {
	return objects.CreateRenderPass(coreDriver, device, handle, version)
}

func CreateSampler(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSampler, version common.APIVersion) core1_0.Sampler {
	return objects.CreateSampler(coreDriver, device, handle, version)
}

func CreateSemaphore(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSemaphore, version common.APIVersion) core1_0.Semaphore {
	return objects.CreateSemaphore(coreDriver, device, handle, version)
}

func CreateShaderModule(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkShaderModule, version common.APIVersion) core1_0.ShaderModule {
	return objects.CreateShaderModule(coreDriver, device, handle, version)
}
