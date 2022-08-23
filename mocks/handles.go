package mocks

import (
	"github.com/vkngwrapper/core/v2/driver"
	"math/rand"
	"unsafe"
)

func fakePointer() unsafe.Pointer {
	return unsafe.Pointer(uintptr(rand.Int()))
}

func NewFakeBufferHandle() driver.VkBuffer {
	return driver.VkBuffer(fakePointer())
}

func NewFakeBufferViewHandle() driver.VkBufferView {
	return driver.VkBufferView(fakePointer())
}

func NewFakeCommandBufferHandle() driver.VkCommandBuffer {
	return driver.VkCommandBuffer(fakePointer())
}

func NewFakeCommandPoolHandle() driver.VkCommandPool {
	return driver.VkCommandPool(fakePointer())
}

func NewFakeDescriptorPool() driver.VkDescriptorPool {
	return driver.VkDescriptorPool(fakePointer())
}

func NewFakeDescriptorSet() driver.VkDescriptorSet {
	return driver.VkDescriptorSet(fakePointer())
}

func NewFakeDescriptorSetLayout() driver.VkDescriptorSetLayout {
	return driver.VkDescriptorSetLayout(fakePointer())
}

func NewFakeDescriptorUpdateTemplate() driver.VkDescriptorUpdateTemplate {
	return driver.VkDescriptorUpdateTemplate(fakePointer())
}

func NewFakeDeviceHandle() driver.VkDevice {
	return driver.VkDevice(fakePointer())
}

func NewFakeDeviceMemoryHandle() driver.VkDeviceMemory {
	return driver.VkDeviceMemory(fakePointer())
}

func NewFakeEventHandle() driver.VkEvent {
	return driver.VkEvent(fakePointer())
}

func NewFakeFenceHandle() driver.VkFence {
	return driver.VkFence(fakePointer())
}

func NewFakeFramebufferHandle() driver.VkFramebuffer {
	return driver.VkFramebuffer(fakePointer())
}

func NewFakeImageHandle() driver.VkImage {
	return driver.VkImage(fakePointer())
}

func NewFakeImageViewHandle() driver.VkImageView {
	return driver.VkImageView(fakePointer())
}

func NewFakeInstanceHandle() driver.VkInstance {
	return driver.VkInstance(fakePointer())
}

func NewFakePhysicalDeviceHandle() driver.VkPhysicalDevice {
	return driver.VkPhysicalDevice(fakePointer())
}

func NewFakePipeline() driver.VkPipeline {
	return driver.VkPipeline(fakePointer())
}

func NewFakePipelineCache() driver.VkPipelineCache {
	return driver.VkPipelineCache(fakePointer())
}

func NewFakePipelineLayout() driver.VkPipelineLayout {
	return driver.VkPipelineLayout(fakePointer())
}

func NewFakeQueryPool() driver.VkQueryPool {
	return driver.VkQueryPool(fakePointer())
}

func NewFakeQueue() driver.VkQueue {
	return driver.VkQueue(fakePointer())
}

func NewFakeRenderPassHandle() driver.VkRenderPass {
	return driver.VkRenderPass(fakePointer())
}

func NewFakeSamplerHandle() driver.VkSampler {
	return driver.VkSampler(fakePointer())
}

func NewFakeSamplerYcbcrConversionHandle() driver.VkSamplerYcbcrConversion {
	return driver.VkSamplerYcbcrConversion(fakePointer())
}

func NewFakeSemaphore() driver.VkSemaphore {
	return driver.VkSemaphore(fakePointer())
}

func NewFakeShaderModule() driver.VkShaderModule {
	return driver.VkShaderModule(fakePointer())
}
