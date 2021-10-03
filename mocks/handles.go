package mocks

import (
	"github.com/CannibalVox/VKng/core"
	"unsafe"
)

func NewFakeBufferHandle() core.VkBuffer {
	val := 0
	return core.VkBuffer(unsafe.Pointer(&val))
}

func NewFakeCommandBufferHandle() core.VkCommandBuffer {
	val := 0
	return core.VkCommandBuffer(unsafe.Pointer(&val))
}

func NewFakeCommandPoolHandle() core.VkCommandPool {
	val := 0
	return core.VkCommandPool(unsafe.Pointer(&val))
}

func NewFakeDescriptorPool() core.VkDescriptorPool {
	val := 0
	return core.VkDescriptorPool(unsafe.Pointer(&val))
}

func NewFakeDescriptorSet() core.VkDescriptorSet {
	val := 0
	return core.VkDescriptorSet(unsafe.Pointer(&val))
}

func NewFakeDescriptorSetLayout() core.VkDescriptorSetLayout {
	val := 0
	return core.VkDescriptorSetLayout(unsafe.Pointer(&val))
}

func NewFakeDeviceHandle() core.VkDevice {
	val := 0
	return core.VkDevice(unsafe.Pointer(&val))
}

func NewFakeFence() core.VkFence {
	val := 0
	return core.VkFence(unsafe.Pointer(&val))
}

func NewFakeFramebufferHandle() core.VkFramebuffer {
	val := 0
	return core.VkFramebuffer(unsafe.Pointer(&val))
}

func NewFakePhysicalDeviceHandle() core.VkPhysicalDevice {
	val := 0
	return core.VkPhysicalDevice(unsafe.Pointer(&val))
}

func NewFakePipeline() core.VkPipeline {
	val := 0
	return core.VkPipeline(unsafe.Pointer(&val))
}

func NewFakePipelineLayout() core.VkPipelineLayout {
	val := 0
	return core.VkPipelineLayout(unsafe.Pointer(&val))
}

func NewFakeQueue() core.VkQueue {
	val := 0
	return core.VkQueue(unsafe.Pointer(&val))
}

func NewFakeRenderPassHandle() core.VkRenderPass {
	val := 0
	return core.VkRenderPass(unsafe.Pointer(&val))
}

func NewFakeSampler() core.VkSampler {
	val := 0
	return core.VkSampler(unsafe.Pointer(&val))
}

func NewFakeSemaphore() core.VkSemaphore {
	val := 0
	return core.VkSemaphore(unsafe.Pointer(&val))
}
