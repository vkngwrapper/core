package mocks1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/internal/impl1_0"
	"github.com/vkngwrapper/core/v3/mocks"
)

func NewDummyInstance(coreDriver driver.Driver, version common.APIVersion, instanceExtensions []string) core1_0.Instance {
	return impl1_0.CreateInstanceObject(coreDriver, mocks.NewFakeInstanceHandle(), version, instanceExtensions)
}

func NewDummyPhysicalDevice(coreDriver driver.Driver, instance core1_0.Instance, deviceVersion common.APIVersion) core1_0.PhysicalDevice {
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	return builder.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), instance.APIVersion(), deviceVersion)
}

func NewDummyDevice(coreDriver driver.Driver, deviceVersion common.APIVersion, deviceExtensions []string) core1_0.Device {
	builder := &impl1_0.InstanceObjectBuilderImpl{}
	return builder.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), deviceVersion, deviceExtensions)
}

func NewDummyBuffer(coreDriver driver.Driver, device core1_0.Device) core1_0.Buffer {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateBufferObject(coreDriver, device.Handle(), mocks.NewFakeBufferHandle(), device.APIVersion())
}

func NewDummyBufferView(coreDriver driver.Driver, device core1_0.Device) core1_0.BufferView {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateBufferViewObject(coreDriver, device.Handle(), mocks.NewFakeBufferViewHandle(), device.APIVersion())
}

func NewDummyCommandBuffer(coreDriver driver.Driver, commandPool core1_0.CommandPool, device core1_0.Device) core1_0.CommandBuffer {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateCommandBufferObject(coreDriver, commandPool.Handle(), device.Handle(), mocks.NewFakeCommandBufferHandle(), device.APIVersion())
}

func NewDummyCommandPool(coreDriver driver.Driver, device core1_0.Device) core1_0.CommandPool {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateCommandPoolObject(coreDriver, device.Handle(), mocks.NewFakeCommandPoolHandle(), device.APIVersion())
}

func NewDummyDescriptorPool(coreDriver driver.Driver, device core1_0.Device) core1_0.DescriptorPool {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateDescriptorPoolObject(coreDriver, device.Handle(), mocks.NewFakeDescriptorPool(), device.APIVersion())
}

func NewDummyDescriptorSet(coreDriver driver.Driver, descriptorPool core1_0.DescriptorPool, device core1_0.Device) core1_0.DescriptorSet {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateDescriptorSetObject(coreDriver, device.Handle(), descriptorPool.Handle(), mocks.NewFakeDescriptorSet(), device.APIVersion())
}

func NewDummyDescriptorSetLayout(coreDriver driver.Driver, device core1_0.Device) core1_0.DescriptorSetLayout {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateDescriptorSetLayoutObject(coreDriver, device.Handle(), mocks.NewFakeDescriptorSetLayout(), device.APIVersion())
}

func NewDummyDeviceMemory(coreDriver driver.Driver, device core1_0.Device, size int) core1_0.DeviceMemory {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateDeviceMemoryObject(coreDriver, device.Handle(), mocks.NewFakeDeviceMemoryHandle(), device.APIVersion(), size)
}

func NewDummyEvent(coreDriver driver.Driver, device core1_0.Device) core1_0.Event {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateEventObject(coreDriver, device.Handle(), mocks.NewFakeEventHandle(), device.APIVersion())
}

func NewDummyFence(coreDriver driver.Driver, device core1_0.Device) core1_0.Fence {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateFenceObject(coreDriver, device.Handle(), mocks.NewFakeFenceHandle(), device.APIVersion())
}

func NewDummyFramebuffer(coreDriver driver.Driver, device core1_0.Device) core1_0.Framebuffer {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateFramebufferObject(coreDriver, device.Handle(), mocks.NewFakeFramebufferHandle(), device.APIVersion())
}

func NewDummyImage(coreDriver driver.Driver, device core1_0.Device) core1_0.Image {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateImageObject(coreDriver, device.Handle(), mocks.NewFakeImageHandle(), device.APIVersion())
}

func NewDummyImageView(coreDriver driver.Driver, device core1_0.Device) core1_0.ImageView {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateImageViewObject(coreDriver, device.Handle(), mocks.NewFakeImageViewHandle(), device.APIVersion())
}

func NewDummyPipeline(coreDriver driver.Driver, device core1_0.Device) core1_0.Pipeline {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreatePipelineObject(coreDriver, device.Handle(), mocks.NewFakePipeline(), device.APIVersion())
}

func NewDummyPipelineCache(coreDriver driver.Driver, device core1_0.Device) core1_0.PipelineCache {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreatePipelineCacheObject(coreDriver, device.Handle(), mocks.NewFakePipelineCache(), device.APIVersion())
}

func NewDummyPipelineLayout(coreDriver driver.Driver, device core1_0.Device) core1_0.PipelineLayout {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreatePipelineLayoutObject(coreDriver, device.Handle(), mocks.NewFakePipelineLayout(), device.APIVersion())
}

func NewDummyQueryPool(coreDriver driver.Driver, device core1_0.Device) core1_0.QueryPool {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateQueryPoolObject(coreDriver, device.Handle(), mocks.NewFakeQueryPool(), device.APIVersion())
}

func NewDummyQueue(coreDriver driver.Driver, device core1_0.Device) core1_0.Queue {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateQueueObject(coreDriver, device.Handle(), mocks.NewFakeQueue(), device.APIVersion())
}

func NewDummyRenderPass(coreDriver driver.Driver, device core1_0.Device) core1_0.RenderPass {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateRenderPassObject(coreDriver, device.Handle(), mocks.NewFakeRenderPassHandle(), device.APIVersion())
}

func NewDummySampler(coreDriver driver.Driver, device core1_0.Device) core1_0.Sampler {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateSamplerObject(coreDriver, device.Handle(), mocks.NewFakeSamplerHandle(), device.APIVersion())
}

func NewDummySemaphore(coreDriver driver.Driver, device core1_0.Device) core1_0.Semaphore {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateSemaphoreObject(coreDriver, device.Handle(), mocks.NewFakeSemaphore(), device.APIVersion())
}

func NewDummyShaderModule(coreDriver driver.Driver, device core1_0.Device) core1_0.ShaderModule {
	builder := &impl1_0.DeviceObjectBuilderImpl{}
	return builder.CreateShaderModuleObject(coreDriver, device.Handle(), mocks.NewFakeShaderModule(), device.APIVersion())
}
