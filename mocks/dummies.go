package mocks

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
)

func NewDummyInstance(version common.APIVersion, instanceExtensions []string) core.Instance {
	return core.InternalInstance(NewFakeInstanceHandle(), version, instanceExtensions)
}

func NewDummyPhysicalDevice(instance core.Instance, deviceVersion common.APIVersion) core.PhysicalDevice {
	return core.InternalPhysicalDevice(NewFakePhysicalDeviceHandle(), instance.APIVersion(), deviceVersion)
}

func NewDummyDevice(deviceVersion common.APIVersion, deviceExtensions []string) core.Device {
	return core.InternalDevice(NewFakeDeviceHandle(), deviceVersion, deviceExtensions)
}

func NewDummyBuffer(device core.Device) core.Buffer {
	return core.InternalBuffer(device.Handle(), NewFakeBufferHandle(), device.APIVersion())
}

func NewDummyBufferView(device core.Device) core.BufferView {
	return core.InternalBufferView(device.Handle(), NewFakeBufferViewHandle(), device.APIVersion())
}

func NewDummyCommandBuffer(commandPool core.CommandPool, device core.Device) core.CommandBuffer {
	return core.InternalCommandBuffer(device.Handle(), commandPool.Handle(), NewFakeCommandBufferHandle(), device.APIVersion())
}

func NewDummyCommandPool(device core.Device) core.CommandPool {
	return core.InternalCommandPool(device.Handle(), NewFakeCommandPoolHandle(), device.APIVersion())
}

func NewDummyDescriptorPool(device core.Device) core.DescriptorPool {
	return core.InternalDescriptorPool(device.Handle(), NewFakeDescriptorPool(), device.APIVersion())
}

func NewDummyDescriptorSet(descriptorPool core.DescriptorPool, device core.Device) core.DescriptorSet {
	return core.InternalDescriptorSet(device.Handle(), descriptorPool.Handle(), NewFakeDescriptorSet(), device.APIVersion())
}

func NewDummyDescriptorSetLayout(device core.Device) core.DescriptorSetLayout {
	return core.InternalDescriptorSetLayout(device.Handle(), NewFakeDescriptorSetLayout(), device.APIVersion())
}

func NewDummyDeviceMemory(device core.Device, size int) core.DeviceMemory {
	return core.InternalDeviceMemory(device.Handle(), NewFakeDeviceMemoryHandle(), device.APIVersion(), size)
}

func NewDummyEvent(device core.Device) core.Event {
	return core.InternalEvent(device.Handle(), NewFakeEventHandle(), device.APIVersion())
}

func NewDummyFence(device core.Device) core.Fence {
	return core.InternalFence(device.Handle(), NewFakeFenceHandle(), device.APIVersion())
}

func NewDummyFramebuffer(device core.Device) core.Framebuffer {
	return core.InternalFramebuffer(device.Handle(), NewFakeFramebufferHandle(), device.APIVersion())
}

func NewDummyImage(device core.Device) core.Image {
	return core.InternalImage(device.Handle(), NewFakeImageHandle(), device.APIVersion())
}

func NewDummyImageView(device core.Device) core.ImageView {
	return core.InternalImageView(device.Handle(), NewFakeImageViewHandle(), device.APIVersion())
}

func NewDummyPipeline(device core.Device) core.Pipeline {
	return core.InternalPipeline(device.Handle(), NewFakePipeline(), device.APIVersion())
}

func NewDummyPipelineCache(device core.Device) core.PipelineCache {
	return core.InternalPipelineCache(device.Handle(), NewFakePipelineCache(), device.APIVersion())
}

func NewDummyPipelineLayout(device core.Device) core.PipelineLayout {
	return core.InternalPipelineLayout(device.Handle(), NewFakePipelineLayout(), device.APIVersion())
}

func NewDummyQueryPool(device core.Device) core.QueryPool {
	return core.InternalQueryPool(device.Handle(), NewFakeQueryPool(), device.APIVersion())
}

func NewDummyQueue(device core.Device) core.Queue {
	return core.InternalQueue(device.Handle(), NewFakeQueue(), device.APIVersion())
}

func NewDummyRenderPass(device core.Device) core.RenderPass {
	return core.InternalRenderPass(device.Handle(), NewFakeRenderPassHandle(), device.APIVersion())
}

func NewDummySampler(device core.Device) core.Sampler {
	return core.InternalSampler(device.Handle(), NewFakeSamplerHandle(), device.APIVersion())
}

func NewDummySemaphore(device core.Device) core.Semaphore {
	return core.InternalSemaphore(device.Handle(), NewFakeSemaphore(), device.APIVersion())
}

func NewDummyShaderModule(device core.Device) core.ShaderModule {
	return core.InternalShaderModule(device.Handle(), NewFakeShaderModule(), device.APIVersion())
}

func NewDummySamplerYcbcrConversion(device core.Device) core.SamplerYcbcrConversion {
	return core.InternalSamplerYcbcrConversion(device.Handle(), NewFakeSamplerYcbcrConversionHandle(), device.APIVersion())
}

func NewDummyDescriptorUpdateTemplate(device core.Device) core.DescriptorUpdateTemplate {
	return core.InternalDescriptorUpdateTemplate(device.Handle(), NewFakeDescriptorUpdateTemplate(), device.APIVersion())
}
