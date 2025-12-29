package impl1_0

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) DestroyBufferView(bufferView core1_0.BufferView, callbacks *loader.AllocationCallbacks) {
	if !bufferView.Initialized() {
		panic("bufferView cannot be uninitialized")
	}

	v.LoaderObj.VkDestroyBufferView(bufferView.DeviceHandle(), bufferView.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyDescriptorSetLayout(layout core1_0.DescriptorSetLayout, callbacks *loader.AllocationCallbacks) {
	if !layout.Initialized() {
		panic("layout was uninitialiazed")
	}
	v.LoaderObj.VkDestroyDescriptorSetLayout(layout.DeviceHandle(), layout.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyFramebuffer(framebuffer core1_0.Framebuffer, callbacks *loader.AllocationCallbacks) {
	if !framebuffer.Initialized() {
		panic("framebuffer was uninitialized")
	}
	v.LoaderObj.VkDestroyFramebuffer(framebuffer.DeviceHandle(), framebuffer.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyImageView(imageView core1_0.ImageView, callbacks *loader.AllocationCallbacks) {
	if !imageView.Initialized() {
		panic("imageView was uninitialized")
	}
	v.LoaderObj.VkDestroyImageView(imageView.DeviceHandle(), imageView.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyPipeline(pipeline core1_0.Pipeline, callbacks *loader.AllocationCallbacks) {
	if !pipeline.Initialized() {
		panic("pipeline was uninitialized")
	}
	v.LoaderObj.VkDestroyPipeline(pipeline.DeviceHandle(), pipeline.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyPipelineLayout(pipelineLayout core1_0.PipelineLayout, callbacks *loader.AllocationCallbacks) {
	if !pipelineLayout.Initialized() {
		panic("pipelineLayout was uninitialized")
	}
	v.LoaderObj.VkDestroyPipelineLayout(pipelineLayout.DeviceHandle(), pipelineLayout.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroySampler(sampler core1_0.Sampler, callbacks *loader.AllocationCallbacks) {
	if !sampler.Initialized() {
		panic("sampler was uninitialized")
	}
	v.LoaderObj.VkDestroySampler(sampler.DeviceHandle(), sampler.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroySemaphore(semaphore core1_0.Semaphore, callbacks *loader.AllocationCallbacks) {
	if !semaphore.Initialized() {
		panic("semaphore was uninitialized")
	}
	v.LoaderObj.VkDestroySemaphore(semaphore.DeviceHandle(), semaphore.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyShaderModule(shaderModule core1_0.ShaderModule, callbacks *loader.AllocationCallbacks) {
	if !shaderModule.Initialized() {
		panic("shaderModule is uninitialized")
	}
	v.LoaderObj.VkDestroyShaderModule(shaderModule.DeviceHandle(), shaderModule.Handle(), callbacks.Handle())
}
