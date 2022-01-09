package mocks

import (
	"github.com/CannibalVox/VKng/core"
	"math/rand"
	"unsafe"
)

func fakePointer() unsafe.Pointer {
	return unsafe.Pointer(uintptr(rand.Int()))
}

func NewFakeBufferHandle() core.VkBuffer {
	return core.VkBuffer(fakePointer())
}

func NewFakeBufferViewHandle() core.VkBufferView {
	return core.VkBufferView(fakePointer())
}

func NewFakeCommandBufferHandle() core.VkCommandBuffer {
	return core.VkCommandBuffer(fakePointer())
}

func NewFakeCommandPoolHandle() core.VkCommandPool {
	return core.VkCommandPool(fakePointer())
}

func NewFakeDescriptorPool() core.VkDescriptorPool {
	return core.VkDescriptorPool(fakePointer())
}

func NewFakeDescriptorSet() core.VkDescriptorSet {
	return core.VkDescriptorSet(fakePointer())
}

func NewFakeDescriptorSetLayout() core.VkDescriptorSetLayout {
	return core.VkDescriptorSetLayout(fakePointer())
}

func NewFakeDeviceHandle() core.VkDevice {
	return core.VkDevice(fakePointer())
}

func NewFakeDeviceMemoryHandle() core.VkDeviceMemory {
	return core.VkDeviceMemory(fakePointer())
}

func NewFakeEventHandle() core.VkEvent {
	return core.VkEvent(fakePointer())
}

func NewFakeFenceHandle() core.VkFence {
	return core.VkFence(fakePointer())
}

func NewFakeFramebufferHandle() core.VkFramebuffer {
	return core.VkFramebuffer(fakePointer())
}

func NewFakeImageHandle() core.VkImage {
	return core.VkImage(fakePointer())
}

func NewFakeImageViewHandle() core.VkImageView {
	return core.VkImageView(fakePointer())
}

func NewFakeInstanceHandle() core.VkInstance {
	return core.VkInstance(fakePointer())
}

func NewFakePhysicalDeviceHandle() core.VkPhysicalDevice {
	return core.VkPhysicalDevice(fakePointer())
}

func NewFakePipeline() core.VkPipeline {
	return core.VkPipeline(fakePointer())
}

func NewFakePipelineCache() core.VkPipelineCache {
	return core.VkPipelineCache(fakePointer())
}

func NewFakePipelineLayout() core.VkPipelineLayout {
	return core.VkPipelineLayout(fakePointer())
}

func NewFakeQueryPool() core.VkQueryPool {
	return core.VkQueryPool(fakePointer())
}

func NewFakeQueue() core.VkQueue {
	return core.VkQueue(fakePointer())
}

func NewFakeRenderPassHandle() core.VkRenderPass {
	return core.VkRenderPass(fakePointer())
}

func NewFakeSamplerHandle() core.VkSampler {
	return core.VkSampler(fakePointer())
}

func NewFakeSemaphore() core.VkSemaphore {
	return core.VkSemaphore(fakePointer())
}

func NewFakeShaderModule() core.VkShaderModule {
	return core.VkShaderModule(fakePointer())
}
