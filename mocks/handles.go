package mocks

import (
	"math/rand"
	"unsafe"

	"github.com/vkngwrapper/core/v3/loader"
)

func fakePointer() unsafe.Pointer {
	return unsafe.Pointer(uintptr(rand.Int()))
}

func NewFakeBufferHandle() loader.VkBuffer {
	return loader.VkBuffer(fakePointer())
}

func NewFakeBufferViewHandle() loader.VkBufferView {
	return loader.VkBufferView(fakePointer())
}

func NewFakeCommandBufferHandle() loader.VkCommandBuffer {
	return loader.VkCommandBuffer(fakePointer())
}

func NewFakeCommandPoolHandle() loader.VkCommandPool {
	return loader.VkCommandPool(fakePointer())
}

func NewFakeDescriptorPool() loader.VkDescriptorPool {
	return loader.VkDescriptorPool(fakePointer())
}

func NewFakeDescriptorSet() loader.VkDescriptorSet {
	return loader.VkDescriptorSet(fakePointer())
}

func NewFakeDescriptorSetLayout() loader.VkDescriptorSetLayout {
	return loader.VkDescriptorSetLayout(fakePointer())
}

func NewFakeDescriptorUpdateTemplate() loader.VkDescriptorUpdateTemplate {
	return loader.VkDescriptorUpdateTemplate(fakePointer())
}

func NewFakeDeviceHandle() loader.VkDevice {
	return loader.VkDevice(fakePointer())
}

func NewFakeDeviceMemoryHandle() loader.VkDeviceMemory {
	return loader.VkDeviceMemory(fakePointer())
}

func NewFakeEventHandle() loader.VkEvent {
	return loader.VkEvent(fakePointer())
}

func NewFakeFenceHandle() loader.VkFence {
	return loader.VkFence(fakePointer())
}

func NewFakeFramebufferHandle() loader.VkFramebuffer {
	return loader.VkFramebuffer(fakePointer())
}

func NewFakeImageHandle() loader.VkImage {
	return loader.VkImage(fakePointer())
}

func NewFakeImageViewHandle() loader.VkImageView {
	return loader.VkImageView(fakePointer())
}

func NewFakeInstanceHandle() loader.VkInstance {
	return loader.VkInstance(fakePointer())
}

func NewFakePhysicalDeviceHandle() loader.VkPhysicalDevice {
	return loader.VkPhysicalDevice(fakePointer())
}

func NewFakePipeline() loader.VkPipeline {
	return loader.VkPipeline(fakePointer())
}

func NewFakePipelineCache() loader.VkPipelineCache {
	return loader.VkPipelineCache(fakePointer())
}

func NewFakePipelineLayout() loader.VkPipelineLayout {
	return loader.VkPipelineLayout(fakePointer())
}

func NewFakeQueryPool() loader.VkQueryPool {
	return loader.VkQueryPool(fakePointer())
}

func NewFakeQueue() loader.VkQueue {
	return loader.VkQueue(fakePointer())
}

func NewFakeRenderPassHandle() loader.VkRenderPass {
	return loader.VkRenderPass(fakePointer())
}

func NewFakeSamplerHandle() loader.VkSampler {
	return loader.VkSampler(fakePointer())
}

func NewFakeSamplerYcbcrConversionHandle() loader.VkSamplerYcbcrConversion {
	return loader.VkSamplerYcbcrConversion(fakePointer())
}

func NewFakeSemaphore() loader.VkSemaphore {
	return loader.VkSemaphore(fakePointer())
}

func NewFakeShaderModule() loader.VkShaderModule {
	return loader.VkShaderModule(fakePointer())
}
