package mocks

import (
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/golang/mock/gomock"
)

func EasyMockBuffer(ctrl *gomock.Controller) *MockBuffer {
	buffer := NewMockBuffer(ctrl)
	buffer.EXPECT().Handle().Return(NewFakeBufferHandle()).AnyTimes()

	return buffer
}

func EasyMockBufferView(ctrl *gomock.Controller) *MockBufferView {
	bufferView := NewMockBufferView(ctrl)
	bufferView.EXPECT().Handle().Return(NewFakeBufferViewHandle()).AnyTimes()

	return bufferView
}

func EasyMockCommandBuffer(ctrl *gomock.Controller) *MockCommandBuffer {
	commandBuffer := NewMockCommandBuffer(ctrl)
	commandBuffer.EXPECT().Handle().Return(NewFakeCommandBufferHandle()).AnyTimes()

	return commandBuffer
}

func EasyMockCommandPool(ctrl *gomock.Controller, device core1_0.Device) *MockCommandPool {
	commandPool := NewMockCommandPool(ctrl)
	commandPool.EXPECT().Handle().Return(NewFakeCommandPoolHandle()).AnyTimes()
	commandPool.EXPECT().Driver().Return(device.Driver()).AnyTimes()
	commandPool.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	commandPool.EXPECT().APIVersion().Return(device.APIVersion()).AnyTimes()

	return commandPool
}

func EasyMockDescriptorPool(ctrl *gomock.Controller, device core1_0.Device) *MockDescriptorPool {
	descriptorPool := NewMockDescriptorPool(ctrl)
	descriptorPool.EXPECT().Handle().Return(NewFakeDescriptorPool()).AnyTimes()
	descriptorPool.EXPECT().Driver().Return(device.Driver()).AnyTimes()
	descriptorPool.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()
	descriptorPool.EXPECT().APIVersion().Return(device.APIVersion()).AnyTimes()

	return descriptorPool
}

func EasyMockDescriptorSet(ctrl *gomock.Controller) *MockDescriptorSet {
	set := NewMockDescriptorSet(ctrl)
	set.EXPECT().Handle().Return(NewFakeDescriptorSet()).AnyTimes()

	return set
}

func EasyMockDescriptorSetLayout(ctrl *gomock.Controller) *MockDescriptorSetLayout {
	layout := NewMockDescriptorSetLayout(ctrl)
	layout.EXPECT().Handle().Return(NewFakeDescriptorSetLayout()).AnyTimes()

	return layout
}

func EasyMockDevice(ctrl *gomock.Controller, driver driver.Driver) *MockDevice {
	device := NewMockDevice(ctrl)
	device.EXPECT().Handle().Return(NewFakeDeviceHandle()).AnyTimes()
	device.EXPECT().Driver().Return(driver).AnyTimes()

	if driver.Version() > 0 {
		device.EXPECT().APIVersion().Return(driver.Version()).AnyTimes()
	}

	return device
}

func EasyMockDeviceMemory(ctrl *gomock.Controller) *MockDeviceMemory {
	deviceMemory := NewMockDeviceMemory(ctrl)
	deviceMemory.EXPECT().Handle().Return(NewFakeDeviceMemoryHandle()).AnyTimes()

	return deviceMemory
}

func EasyMockEvent(ctrl *gomock.Controller) *MockEvent {
	event := NewMockEvent(ctrl)
	event.EXPECT().Handle().Return(NewFakeEventHandle()).AnyTimes()

	return event
}

func EasyMockFence(ctrl *gomock.Controller) *MockFence {
	fence := NewMockFence(ctrl)
	fence.EXPECT().Handle().Return(NewFakeFenceHandle()).AnyTimes()

	return fence
}

func EasyMockFramebuffer(ctrl *gomock.Controller) *MockFramebuffer {
	framebuffer := NewMockFramebuffer(ctrl)
	framebuffer.EXPECT().Handle().Return(NewFakeFramebufferHandle()).AnyTimes()

	return framebuffer
}

func EasyMockImage(ctrl *gomock.Controller) *MockImage {
	image := NewMockImage(ctrl)
	image.EXPECT().Handle().Return(NewFakeImageHandle()).AnyTimes()

	return image
}

func EasyMockImageView(ctrl *gomock.Controller) *MockImageView {
	imageView := NewMockImageView(ctrl)
	imageView.EXPECT().Handle().Return(NewFakeImageViewHandle()).AnyTimes()

	return imageView
}

func EasyMockInstance(ctrl *gomock.Controller, driver driver.Driver) *MockInstance {
	instance := NewMockInstance(ctrl)
	instance.EXPECT().Handle().Return(NewFakeInstanceHandle()).AnyTimes()
	instance.EXPECT().Driver().Return(driver).AnyTimes()
	instance.EXPECT().APIVersion().Return(driver.Version()).AnyTimes()

	return instance
}

func EasyMockPipeline(ctrl *gomock.Controller) *MockPipeline {
	pipeline := NewMockPipeline(ctrl)
	pipeline.EXPECT().Handle().Return(NewFakePipeline()).AnyTimes()

	return pipeline
}

func EasyMockPipelineCache(ctrl *gomock.Controller) *MockPipelineCache {
	pipelineCache := NewMockPipelineCache(ctrl)
	pipelineCache.EXPECT().Handle().Return(NewFakePipelineCache()).AnyTimes()

	return pipelineCache
}

func EasyMockPipelineLayout(ctrl *gomock.Controller) *MockPipelineLayout {
	pipelineLayout := NewMockPipelineLayout(ctrl)
	pipelineLayout.EXPECT().Handle().Return(NewFakePipelineLayout()).AnyTimes()

	return pipelineLayout
}

func EasyMockPhysicalDevice(ctrl *gomock.Controller, driver driver.Driver) *MockPhysicalDevice {
	physicalDevice := NewMockPhysicalDevice(ctrl)
	physicalDevice.EXPECT().Handle().Return(NewFakePhysicalDeviceHandle()).AnyTimes()
	physicalDevice.EXPECT().Driver().Return(driver).AnyTimes()

	if driver.Version() > 0 {
		physicalDevice.EXPECT().InstanceAPIVersion().Return(driver.Version()).AnyTimes()
		physicalDevice.EXPECT().DeviceAPIVersion().Return(driver.Version()).AnyTimes()
	}

	return physicalDevice
}

func EasyMockQueryPool(ctrl *gomock.Controller) *MockQueryPool {
	queryPool := NewMockQueryPool(ctrl)
	queryPool.EXPECT().Handle().Return(NewFakeQueryPool()).AnyTimes()

	return queryPool
}

func EasyMockQueue(ctrl *gomock.Controller) *MockQueue {
	queue := NewMockQueue(ctrl)
	queue.EXPECT().Handle().Return(NewFakeQueue()).AnyTimes()

	return queue
}

func EasyMockRenderPass(ctrl *gomock.Controller) *MockRenderPass {
	renderPass := NewMockRenderPass(ctrl)
	renderPass.EXPECT().Handle().Return(NewFakeRenderPassHandle()).AnyTimes()

	return renderPass
}

func EasyMockSampler(ctrl *gomock.Controller) *MockSampler {
	sampler := NewMockSampler(ctrl)
	sampler.EXPECT().Handle().Return(NewFakeSamplerHandle()).AnyTimes()

	return sampler
}

func EasyMockSamplerYcbcrConversion(ctrl *gomock.Controller) *MockSamplerYcbcrConversion {
	ycbcr := NewMockSamplerYcbcrConversion(ctrl)
	ycbcr.EXPECT().Handle().Return(NewFakeSamplerYcbcrConversionHandle()).AnyTimes()

	return ycbcr
}

func EasyMockSemaphore(ctrl *gomock.Controller) *MockSemaphore {
	semaphore := NewMockSemaphore(ctrl)
	semaphore.EXPECT().Handle().Return(NewFakeSemaphore()).AnyTimes()

	return semaphore
}

func EasyMockShaderModule(ctrl *gomock.Controller) *MockShaderModule {
	shader := NewMockShaderModule(ctrl)
	shader.EXPECT().Handle().Return(NewFakeShaderModule()).AnyTimes()

	return shader
}
