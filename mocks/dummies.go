package mocks

import (
	"github.com/vkngwrapper/core/v3/common"
	types "github.com/vkngwrapper/core/v3/types"
)

func NewDummyInstance(version common.APIVersion, instanceExtensions []string) types.Instance {
	return types.InternalInstance(NewFakeInstanceHandle(), version, instanceExtensions)
}

func NewDummyPhysicalDevice(instance types.Instance, deviceVersion common.APIVersion) types.PhysicalDevice {
	return types.InternalPhysicalDevice(NewFakePhysicalDeviceHandle(), instance.APIVersion(), deviceVersion)
}

func NewDummyDevice(deviceVersion common.APIVersion, deviceExtensions []string) types.Device {
	return types.InternalDevice(NewFakeDeviceHandle(), deviceVersion, deviceExtensions)
}

func NewDummyBuffer(device types.Device) types.Buffer {
	return types.InternalBuffer(device.Handle(), NewFakeBufferHandle(), device.APIVersion())
}

func NewDummyBufferView(device types.Device) types.BufferView {
	return types.InternalBufferView(device.Handle(), NewFakeBufferViewHandle(), device.APIVersion())
}

func NewDummyCommandBuffer(commandPool types.CommandPool, device types.Device) types.CommandBuffer {
	return types.InternalCommandBuffer(device.Handle(), commandPool.Handle(), NewFakeCommandBufferHandle(), device.APIVersion())
}

func NewDummyCommandPool(device types.Device) types.CommandPool {
	return types.InternalCommandPool(device.Handle(), NewFakeCommandPoolHandle(), device.APIVersion())
}

func NewDummyDescriptorPool(device types.Device) types.DescriptorPool {
	return types.InternalDescriptorPool(device.Handle(), NewFakeDescriptorPool(), device.APIVersion())
}

func NewDummyDescriptorSet(descriptorPool types.DescriptorPool, device types.Device) types.DescriptorSet {
	return types.InternalDescriptorSet(device.Handle(), descriptorPool.Handle(), NewFakeDescriptorSet(), device.APIVersion())
}

func NewDummyDescriptorSetLayout(device types.Device) types.DescriptorSetLayout {
	return types.InternalDescriptorSetLayout(device.Handle(), NewFakeDescriptorSetLayout(), device.APIVersion())
}

func NewDummyDeviceMemory(device types.Device, size int) types.DeviceMemory {
	return types.InternalDeviceMemory(device.Handle(), NewFakeDeviceMemoryHandle(), device.APIVersion(), size)
}

func NewDummyEvent(device types.Device) types.Event {
	return types.InternalEvent(device.Handle(), NewFakeEventHandle(), device.APIVersion())
}

func NewDummyFence(device types.Device) types.Fence {
	return types.InternalFence(device.Handle(), NewFakeFenceHandle(), device.APIVersion())
}

func NewDummyFramebuffer(device types.Device) types.Framebuffer {
	return types.InternalFramebuffer(device.Handle(), NewFakeFramebufferHandle(), device.APIVersion())
}

func NewDummyImage(device types.Device) types.Image {
	return types.InternalImage(device.Handle(), NewFakeImageHandle(), device.APIVersion())
}

func NewDummyImageView(device types.Device) types.ImageView {
	return types.InternalImageView(device.Handle(), NewFakeImageViewHandle(), device.APIVersion())
}

func NewDummyPipeline(device types.Device) types.Pipeline {
	return types.InternalPipeline(device.Handle(), NewFakePipeline(), device.APIVersion())
}

func NewDummyPipelineCache(device types.Device) types.PipelineCache {
	return types.InternalPipelineCache(device.Handle(), NewFakePipelineCache(), device.APIVersion())
}

func NewDummyPipelineLayout(device types.Device) types.PipelineLayout {
	return types.InternalPipelineLayout(device.Handle(), NewFakePipelineLayout(), device.APIVersion())
}

func NewDummyQueryPool(device types.Device) types.QueryPool {
	return types.InternalQueryPool(device.Handle(), NewFakeQueryPool(), device.APIVersion())
}

func NewDummyQueue(device types.Device) types.Queue {
	return types.InternalQueue(device.Handle(), NewFakeQueue(), device.APIVersion())
}

func NewDummyRenderPass(device types.Device) types.RenderPass {
	return types.InternalRenderPass(device.Handle(), NewFakeRenderPassHandle(), device.APIVersion())
}

func NewDummySampler(device types.Device) types.Sampler {
	return types.InternalSampler(device.Handle(), NewFakeSamplerHandle(), device.APIVersion())
}

func NewDummySemaphore(device types.Device) types.Semaphore {
	return types.InternalSemaphore(device.Handle(), NewFakeSemaphore(), device.APIVersion())
}

func NewDummyShaderModule(device types.Device) types.ShaderModule {
	return types.InternalShaderModule(device.Handle(), NewFakeShaderModule(), device.APIVersion())
}
