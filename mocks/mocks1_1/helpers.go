package mocks1_1

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/mocks"
	gomock "go.uber.org/mock/gomock"
)

func EasyMockBuffer(ctrl *gomock.Controller) *MockBuffer {
	buffer := NewMockBuffer(ctrl)
	buffer.EXPECT().Handle().Return(mocks.NewFakeBufferHandle()).AnyTimes()

	return buffer
}

func EasyMockBufferView(ctrl *gomock.Controller) *MockBufferView {
	bufferView := NewMockBufferView(ctrl)
	bufferView.EXPECT().Handle().Return(mocks.NewFakeBufferViewHandle()).AnyTimes()

	return bufferView
}

func EasyMockCommandBuffer(ctrl *gomock.Controller) *MockCommandBuffer {
	commandBuffer := NewMockCommandBuffer(ctrl)
	commandBuffer.EXPECT().Handle().Return(mocks.NewFakeCommandBufferHandle()).AnyTimes()

	return commandBuffer
}

func EasyMockCommandPool(ctrl *gomock.Controller, device core1_0.Device) *MockCommandPool {
	commandPool := NewMockCommandPool(ctrl)
	commandPool.EXPECT().Handle().Return(mocks.NewFakeCommandPoolHandle()).AnyTimes()
	commandPool.EXPECT().Driver().Return(device.Driver()).AnyTimes()
	commandPool.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	commandPool.EXPECT().APIVersion().Return(device.APIVersion()).AnyTimes()

	return commandPool
}

func EasyMockDescriptorPool(ctrl *gomock.Controller, device core1_0.Device) *MockDescriptorPool {
	descriptorPool := NewMockDescriptorPool(ctrl)
	descriptorPool.EXPECT().Handle().Return(mocks.NewFakeDescriptorPool()).AnyTimes()
	descriptorPool.EXPECT().Driver().Return(device.Driver()).AnyTimes()
	descriptorPool.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	descriptorPool.EXPECT().APIVersion().Return(device.APIVersion()).AnyTimes()

	return descriptorPool
}

func EasyMockDescriptorSet(ctrl *gomock.Controller) *MockDescriptorSet {
	set := NewMockDescriptorSet(ctrl)
	set.EXPECT().Handle().Return(mocks.NewFakeDescriptorSet()).AnyTimes()

	return set
}

func EasyMockDescriptorSetLayout(ctrl *gomock.Controller) *MockDescriptorSetLayout {
	layout := NewMockDescriptorSetLayout(ctrl)
	layout.EXPECT().Handle().Return(mocks.NewFakeDescriptorSetLayout()).AnyTimes()

	return layout
}

func EasyMockDevice(ctrl *gomock.Controller, driver driver.Driver) *MockDevice {
	device := NewMockDevice(ctrl)
	device.EXPECT().Handle().Return(mocks.NewFakeDeviceHandle()).AnyTimes()
	device.EXPECT().Driver().Return(driver).AnyTimes()

	if driver.Version() > 0 {
		device.EXPECT().APIVersion().Return(driver.Version()).AnyTimes()
	}

	return device
}

func EasyMockDeviceMemory(ctrl *gomock.Controller) *MockDeviceMemory {
	deviceMemory := NewMockDeviceMemory(ctrl)
	deviceMemory.EXPECT().Handle().Return(mocks.NewFakeDeviceMemoryHandle()).AnyTimes()

	return deviceMemory
}

func EasyMockEvent(ctrl *gomock.Controller) *MockEvent {
	event := NewMockEvent(ctrl)
	event.EXPECT().Handle().Return(mocks.NewFakeEventHandle()).AnyTimes()

	return event
}

func EasyMockFence(ctrl *gomock.Controller) *MockFence {
	fence := NewMockFence(ctrl)
	fence.EXPECT().Handle().Return(mocks.NewFakeFenceHandle()).AnyTimes()

	return fence
}

func EasyMockFramebuffer(ctrl *gomock.Controller) *MockFramebuffer {
	framebuffer := NewMockFramebuffer(ctrl)
	framebuffer.EXPECT().Handle().Return(mocks.NewFakeFramebufferHandle()).AnyTimes()

	return framebuffer
}

func EasyMockImage(ctrl *gomock.Controller) *MockImage {
	image := NewMockImage(ctrl)
	image.EXPECT().Handle().Return(mocks.NewFakeImageHandle()).AnyTimes()

	return image
}

func EasyMockImageView(ctrl *gomock.Controller) *MockImageView {
	imageView := NewMockImageView(ctrl)
	imageView.EXPECT().Handle().Return(mocks.NewFakeImageViewHandle()).AnyTimes()

	return imageView
}

func EasyMockInstance(ctrl *gomock.Controller, driver driver.Driver) *MockInstance {
	instance := NewMockInstance(ctrl)
	instance.EXPECT().Handle().Return(mocks.NewFakeInstanceHandle()).AnyTimes()
	instance.EXPECT().Driver().Return(driver).AnyTimes()
	instance.EXPECT().APIVersion().Return(driver.Version()).AnyTimes()

	return instance
}

func EasyMockPipeline(ctrl *gomock.Controller) *MockPipeline {
	pipeline := NewMockPipeline(ctrl)
	pipeline.EXPECT().Handle().Return(mocks.NewFakePipeline()).AnyTimes()

	return pipeline
}

func EasyMockPipelineCache(ctrl *gomock.Controller) *MockPipelineCache {
	pipelineCache := NewMockPipelineCache(ctrl)
	pipelineCache.EXPECT().Handle().Return(mocks.NewFakePipelineCache()).AnyTimes()

	return pipelineCache
}

func EasyMockPipelineLayout(ctrl *gomock.Controller) *MockPipelineLayout {
	pipelineLayout := NewMockPipelineLayout(ctrl)
	pipelineLayout.EXPECT().Handle().Return(mocks.NewFakePipelineLayout()).AnyTimes()

	return pipelineLayout
}

func EasyMockPhysicalDevice(ctrl *gomock.Controller, driver driver.Driver) *MockPhysicalDevice {
	physicalDevice := NewMockPhysicalDevice(ctrl)
	physicalDevice.EXPECT().Handle().Return(mocks.NewFakePhysicalDeviceHandle()).AnyTimes()
	physicalDevice.EXPECT().Driver().Return(driver).AnyTimes()

	if driver.Version() > 0 {
		physicalDevice.EXPECT().InstanceAPIVersion().Return(driver.Version()).AnyTimes()
		physicalDevice.EXPECT().DeviceAPIVersion().Return(driver.Version()).AnyTimes()
	}

	return physicalDevice
}

func EasyMockQueryPool(ctrl *gomock.Controller) *MockQueryPool {
	queryPool := NewMockQueryPool(ctrl)
	queryPool.EXPECT().Handle().Return(mocks.NewFakeQueryPool()).AnyTimes()

	return queryPool
}

func EasyMockQueue(ctrl *gomock.Controller) *MockQueue {
	queue := NewMockQueue(ctrl)
	queue.EXPECT().Handle().Return(mocks.NewFakeQueue()).AnyTimes()

	return queue
}

func EasyMockRenderPass(ctrl *gomock.Controller) *MockRenderPass {
	renderPass := NewMockRenderPass(ctrl)
	renderPass.EXPECT().Handle().Return(mocks.NewFakeRenderPassHandle()).AnyTimes()

	return renderPass
}

func EasyMockSampler(ctrl *gomock.Controller) *MockSampler {
	sampler := NewMockSampler(ctrl)
	sampler.EXPECT().Handle().Return(mocks.NewFakeSamplerHandle()).AnyTimes()

	return sampler
}

func EasyMockSemaphore(ctrl *gomock.Controller) *MockSemaphore {
	semaphore := NewMockSemaphore(ctrl)
	semaphore.EXPECT().Handle().Return(mocks.NewFakeSemaphore()).AnyTimes()

	return semaphore
}

func EasyMockShaderModule(ctrl *gomock.Controller) *MockShaderModule {
	shader := NewMockShaderModule(ctrl)
	shader.EXPECT().Handle().Return(mocks.NewFakeShaderModule()).AnyTimes()

	return shader
}

func EasyMockSamplerYcbcrConversion(ctrl *gomock.Controller) *MockSamplerYcbcrConversion {
	conversion := NewMockSamplerYcbcrConversion(ctrl)
	conversion.EXPECT().Handle().Return(mocks.NewFakeSamplerYcbcrConversionHandle()).AnyTimes()

	return conversion
}

func EasyMockDescriptorUpdateTemplate(ctrl *gomock.Controller) *MockDescriptorUpdateTemplate {
	template := NewMockDescriptorUpdateTemplate(ctrl)
	template.EXPECT().Handle().Return(mocks.NewFakeDescriptorUpdateTemplate())

	return template
}
