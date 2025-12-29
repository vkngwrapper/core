package mocks

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
)

func NewDummyInstance(version common.APIVersion, instanceExtensions []string) core1_0.Instance {
	return core1_0.InternalInstance(NewFakeInstanceHandle(), version, instanceExtensions)
}

func NewDummyPhysicalDevice(instance core1_0.Instance, deviceVersion common.APIVersion) core1_0.PhysicalDevice {
	return core1_0.InternalPhysicalDevice(NewFakePhysicalDeviceHandle(), instance.APIVersion(), deviceVersion)
}

func NewDummyDevice(deviceVersion common.APIVersion, deviceExtensions []string) core1_0.Device {
	return core1_0.InternalDevice(NewFakeDeviceHandle(), deviceVersion, deviceExtensions)
}

func NewDummyBuffer(device core1_0.Device) core1_0.Buffer {
	return core1_0.InternalBuffer(device.Handle(), NewFakeBufferHandle(), device.APIVersion())
}

func NewDummyBufferView(device core1_0.Device) core1_0.BufferView {
	return core1_0.InternalBufferView(device.Handle(), NewFakeBufferViewHandle(), device.APIVersion())
}

func NewDummyCommandBuffer(commandPool core1_0.CommandPool, device core1_0.Device) core1_0.CommandBuffer {
	return core1_0.InternalCommandBuffer(device.Handle(), commandPool.Handle(), NewFakeCommandBufferHandle(), device.APIVersion())
}

func NewDummyCommandPool(device core1_0.Device) core1_0.CommandPool {
	return core1_0.InternalCommandPool(device.Handle(), NewFakeCommandPoolHandle(), device.APIVersion())
}

func NewDummyDescriptorPool(device core1_0.Device) core1_0.DescriptorPool {
	return core1_0.InternalDescriptorPool(device.Handle(), NewFakeDescriptorPool(), device.APIVersion())
}

func NewDummyDescriptorSet(descriptorPool core1_0.DescriptorPool, device core1_0.Device) core1_0.DescriptorSet {
	return core1_0.InternalDescriptorSet(device.Handle(), descriptorPool.Handle(), NewFakeDescriptorSet(), device.APIVersion())
}

func NewDummyDescriptorSetLayout(device core1_0.Device) core1_0.DescriptorSetLayout {
	return core1_0.InternalDescriptorSetLayout(device.Handle(), NewFakeDescriptorSetLayout(), device.APIVersion())
}

func NewDummyDeviceMemory(device core1_0.Device, size int) core1_0.DeviceMemory {
	return core1_0.InternalDeviceMemory(device.Handle(), NewFakeDeviceMemoryHandle(), device.APIVersion(), size)
}

func NewDummyEvent(device core1_0.Device) core1_0.Event {
	return core1_0.InternalEvent(device.Handle(), NewFakeEventHandle(), device.APIVersion())
}

func NewDummyFence(device core1_0.Device) core1_0.Fence {
	return core1_0.InternalFence(device.Handle(), NewFakeFenceHandle(), device.APIVersion())
}

func NewDummyFramebuffer(device core1_0.Device) core1_0.Framebuffer {
	return core1_0.InternalFramebuffer(device.Handle(), NewFakeFramebufferHandle(), device.APIVersion())
}

func NewDummyImage(device core1_0.Device) core1_0.Image {
	return core1_0.InternalImage(device.Handle(), NewFakeImageHandle(), device.APIVersion())
}

func NewDummyImageView(device core1_0.Device) core1_0.ImageView {
	return core1_0.InternalImageView(device.Handle(), NewFakeImageViewHandle(), device.APIVersion())
}

func NewDummyPipeline(device core1_0.Device) core1_0.Pipeline {
	return core1_0.InternalPipeline(device.Handle(), NewFakePipeline(), device.APIVersion())
}

func NewDummyPipelineCache(device core1_0.Device) core1_0.PipelineCache {
	return core1_0.InternalPipelineCache(device.Handle(), NewFakePipelineCache(), device.APIVersion())
}

func NewDummyPipelineLayout(device core1_0.Device) core1_0.PipelineLayout {
	return core1_0.InternalPipelineLayout(device.Handle(), NewFakePipelineLayout(), device.APIVersion())
}

func NewDummyQueryPool(device core1_0.Device) core1_0.QueryPool {
	return core1_0.InternalQueryPool(device.Handle(), NewFakeQueryPool(), device.APIVersion())
}

func NewDummyQueue(device core1_0.Device) core1_0.Queue {
	return core1_0.InternalQueue(device.Handle(), NewFakeQueue(), device.APIVersion())
}

func NewDummyRenderPass(device core1_0.Device) core1_0.RenderPass {
	return core1_0.InternalRenderPass(device.Handle(), NewFakeRenderPassHandle(), device.APIVersion())
}

func NewDummySampler(device core1_0.Device) core1_0.Sampler {
	return core1_0.InternalSampler(device.Handle(), NewFakeSamplerHandle(), device.APIVersion())
}

func NewDummySemaphore(device core1_0.Device) core1_0.Semaphore {
	return core1_0.InternalSemaphore(device.Handle(), NewFakeSemaphore(), device.APIVersion())
}

func NewDummyShaderModule(device core1_0.Device) core1_0.ShaderModule {
	return core1_0.InternalShaderModule(device.Handle(), NewFakeShaderModule(), device.APIVersion())
}

func NewDummySamplerYcbcrConversion(device core1_0.Device) core1_1.SamplerYcbcrConversion {
	return core1_1.InternalSamplerYcbcrConversion(device.Handle(), NewFakeSamplerYcbcrConversionHandle(), device.APIVersion())
}

func NewDummyDescriptorUpdateTemplate(device core1_0.Device) core1_1.DescriptorUpdateTemplate {
	return core1_1.InternalDescriptorUpdateTemplate(device.Handle(), NewFakeDescriptorUpdateTemplate(), device.APIVersion())
}
