package impl1_0

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) DestroyBufferView(bufferView core.BufferView, callbacks *loader.AllocationCallbacks) {
	if bufferView.Handle() == 0 {
		panic("bufferView cannot be uninitialized")
	}

	v.LoaderObj.VkDestroyBufferView(bufferView.DeviceHandle(), bufferView.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyDescriptorSetLayout(layout core.DescriptorSetLayout, callbacks *loader.AllocationCallbacks) {
	if layout.Handle() == 0 {
		panic("layout was uninitialiazed")
	}
	v.LoaderObj.VkDestroyDescriptorSetLayout(layout.DeviceHandle(), layout.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyFramebuffer(framebuffer core.Framebuffer, callbacks *loader.AllocationCallbacks) {
	if framebuffer.Handle() == 0 {
		panic("framebuffer was uninitialized")
	}
	v.LoaderObj.VkDestroyFramebuffer(framebuffer.DeviceHandle(), framebuffer.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyImageView(imageView core.ImageView, callbacks *loader.AllocationCallbacks) {
	if imageView.Handle() == 0 {
		panic("imageView was uninitialized")
	}
	v.LoaderObj.VkDestroyImageView(imageView.DeviceHandle(), imageView.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyPipeline(pipeline core.Pipeline, callbacks *loader.AllocationCallbacks) {
	if pipeline.Handle() == 0 {
		panic("pipeline was uninitialized")
	}
	v.LoaderObj.VkDestroyPipeline(pipeline.DeviceHandle(), pipeline.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyPipelineLayout(pipelineLayout core.PipelineLayout, callbacks *loader.AllocationCallbacks) {
	if pipelineLayout.Handle() == 0 {
		panic("pipelineLayout was uninitialized")
	}
	v.LoaderObj.VkDestroyPipelineLayout(pipelineLayout.DeviceHandle(), pipelineLayout.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroySampler(sampler core.Sampler, callbacks *loader.AllocationCallbacks) {
	if sampler.Handle() == 0 {
		panic("sampler was uninitialized")
	}
	v.LoaderObj.VkDestroySampler(sampler.DeviceHandle(), sampler.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroySemaphore(semaphore core.Semaphore, callbacks *loader.AllocationCallbacks) {
	if semaphore.Handle() == 0 {
		panic("semaphore was uninitialized")
	}
	v.LoaderObj.VkDestroySemaphore(semaphore.DeviceHandle(), semaphore.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) DestroyShaderModule(shaderModule core.ShaderModule, callbacks *loader.AllocationCallbacks) {
	if shaderModule.Handle() == 0 {
		panic("shaderModule is uninitialized")
	}
	v.LoaderObj.VkDestroyShaderModule(shaderModule.DeviceHandle(), shaderModule.Handle(), callbacks.Handle())
}
