package mocks

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"unsafe"
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

	return physicalDevice
}

func EasyMockQueryPool(ctrl *gomock.Controller) *MockQueryPool {
	queryPool := NewMockQueryPool(ctrl)
	queryPool.EXPECT().Handle().Return(NewFakeQueryPool()).AnyTimes()

	return queryPool
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

func EasyDummyBuffer(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.Buffer {
	handle := NewFakeBufferHandle()
	mockDriver := device.Driver().(*mock_driver.MockDriver)
	mockDriver.EXPECT().VkCreateBuffer(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkBufferCreateInfo, pAllocator *driver.VkAllocationCallbacks, pBuffer *driver.VkBuffer) (common.VkResult, error) {
			*pBuffer = handle
			return common.VKSuccess, nil
		})

	buffer, _, err := loader.CreateBuffer(device, nil, &core1_0.BufferOptions{})
	require.NoError(t, err)
	require.NotNil(t, buffer)

	return buffer
}

func EasyDummyCommandPool(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.CommandPool {
	handle := NewFakeCommandPoolHandle()
	mockDriver := device.Driver().(*mock_driver.MockDriver)
	mockDriver.EXPECT().VkCreateCommandPool(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, createInfo *driver.VkCommandPoolCreateInfo, allocator *driver.VkAllocationCallbacks, commandPool *driver.VkCommandPool) (common.VkResult, error) {
			*commandPool = handle

			return common.VKSuccess, nil
		})

	graphicsFamily := 0
	pool, res, err := loader.CreateCommandPool(device, nil, &core1_0.CommandPoolOptions{
		Flags:               core1_0.CommandPoolResetBuffer,
		GraphicsQueueFamily: &graphicsFamily,
	})
	require.NoError(t, err)
	require.Equal(t, common.VKSuccess, res)

	return pool
}

func EasyDummyCommandBuffer(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.CommandBuffer {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkAllocateCommandBuffers(gomock.Any(), gomock.Any(), gomock.Any()).Do(
		func(device driver.VkDevice, pAllocateInfo *driver.VkCommandBufferAllocateInfo, pCommandBuffers *driver.VkCommandBuffer) {
			*pCommandBuffers = NewFakeCommandBufferHandle()
		})

	buffers, _, err := loader.AllocateCommandBuffers(&core1_0.CommandBufferOptions{
		BufferCount: 1,
	})
	require.NoError(t, err)

	return buffers[0]
}

func EasyDummyDescriptorPool(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.DescriptorPool {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateDescriptorPool(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorPoolCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDescriptorPool *driver.VkDescriptorPool) (common.VkResult, error) {
			*pDescriptorPool = NewFakeDescriptorPool()
			return common.VKSuccess, nil
		})

	pool, _, err := loader.CreateDescriptorPool(device, nil, &core1_0.DescriptorPoolOptions{})
	require.NoError(t, err)

	return pool
}

func EasyDummyDescriptorSetLayout(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.DescriptorSetLayout {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateDescriptorSetLayout(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDescriptorSetLayout *driver.VkDescriptorSetLayout) (common.VkResult, error) {
			*pDescriptorSetLayout = NewFakeDescriptorSetLayout()
			return common.VKSuccess, nil
		})

	layout, _, err := loader.CreateDescriptorSetLayout(device, nil, &core1_0.DescriptorSetLayoutOptions{})
	require.NoError(t, err)
	return layout
}

func EasyDummyDevice(t *testing.T, ctrl *gomock.Controller, loader loader.Loader) core1_0.Device {
	mockDriver := loader.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(mockDriver, nil)
	mockDriver.EXPECT().VkCreateDevice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice, pCreateInfo *driver.VkDeviceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDevice *driver.VkDevice) (common.VkResult, error) {
			*pDevice = NewFakeDeviceHandle()
			return common.VKSuccess, nil
		})

	device, _, err := loader.CreateDevice(EasyMockPhysicalDevice(ctrl, mockDriver), nil, &core1_0.DeviceOptions{
		QueueFamilies: []*core1_0.QueueFamilyOptions{
			{
				QueuePriorities: []float32{1},
			},
		},
	})
	require.NoError(t, err)

	return device
}

func EasyDummyDeviceMemory(t *testing.T, loader loader.Loader, device core1_0.Device, size int) core1_0.DeviceMemory {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkAllocateMemory(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkMemoryAllocateInfo, pAllocator *driver.VkAllocationCallbacks, pDeviceMemory *driver.VkDeviceMemory) (common.VkResult, error) {
			*pDeviceMemory = NewFakeDeviceMemoryHandle()
			return common.VKSuccess, nil
		})

	memory, _, err := device.AllocateMemory(nil, &core1_0.DeviceMemoryOptions{
		AllocationSize: size,
	})
	require.NoError(t, err)

	return memory
}

func EasyDummyEvent(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.Event {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateEvent(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkEventCreateInfo, pAllocator *driver.VkAllocationCallbacks, pEvent *driver.VkEvent) (common.VkResult, error) {
			*pEvent = NewFakeEventHandle()
			return common.VKSuccess, nil
		})

	event, _, err := loader.CreateEvent(device, nil, &core1_0.EventOptions{})
	require.NoError(t, err)

	return event
}

func EasyDummyFence(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.Fence {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateFence(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkFenceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pFence *driver.VkFence) (common.VkResult, error) {
			*pFence = NewFakeFenceHandle()
			return common.VKSuccess, nil
		})

	fence, _, err := loader.CreateFence(device, nil, &core1_0.FenceOptions{})
	require.NoError(t, err)

	return fence
}

func EasyDummyFramebuffer(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.Framebuffer {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateFramebuffer(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkFramebufferCreateInfo, pAllocator *driver.VkAllocationCallbacks, pFramebuffer *driver.VkFramebuffer) (common.VkResult, error) {
			*pFramebuffer = NewFakeFramebufferHandle()
			return common.VKSuccess, nil
		})

	framebuffer, _, err := loader.CreateFrameBuffer(device, nil, &core1_0.FramebufferOptions{})
	require.NoError(t, err)

	return framebuffer
}

func EasyDummyGraphicsPipeline(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.Pipeline {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateGraphicsPipelines(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pPipelines *driver.VkPipeline) (common.VkResult, error) {
			*pPipelines = NewFakePipeline()
			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{{}})
	require.NoError(t, err)

	return pipelines[0]
}

func EasyDummyImage(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.Image {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateImage(device.Handle(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkImageCreateInfo, pAllocator *driver.VkAllocationCallbacks, pImage *driver.VkImage) (common.VkResult, error) {
			*pImage = NewFakeImageHandle()

			return common.VKSuccess, nil
		})

	image, _, err := loader.CreateImage(device, nil, &core1_0.ImageOptions{})
	require.NoError(t, err)

	return image
}

func EasyDummyInstance(t *testing.T, loader loader.Loader) core1_0.Instance {
	mockDriver := loader.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().CreateInstanceDriver(gomock.Any()).Return(mockDriver, nil).AnyTimes()
	mockDriver.EXPECT().VkCreateInstance(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(pCreateInfo *driver.VkInstanceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pInstance *driver.VkInstance) (common.VkResult, error) {
			*pInstance = NewFakeInstanceHandle()
			return common.VKSuccess, nil
		})

	instance, _, err := loader.CreateInstance(nil, &core1_0.InstanceOptions{})
	require.NoError(t, err)

	return instance
}

func EasyDummyPhysicalDevice(t *testing.T, loader loader.Loader) core1_0.PhysicalDevice {
	mockDriver := loader.Driver().(*mock_driver.MockDriver)
	instance := EasyDummyInstance(t, loader)

	mockDriver.EXPECT().VkEnumeratePhysicalDevices(instance.Handle(), gomock.Not(nil), nil).DoAndReturn(
		func(instance driver.VkInstance, pPhysicalDeviceCount *driver.Uint32, pPhysicalDevices *driver.VkPhysicalDevice) (common.VkResult, error) {
			*pPhysicalDeviceCount = driver.Uint32(1)
			return common.VKSuccess, nil
		})
	mockDriver.EXPECT().VkEnumeratePhysicalDevices(instance.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(instance driver.VkInstance, pPhysicalDeviceCount *driver.Uint32, pPhysicalDevices *driver.VkPhysicalDevice) (common.VkResult, error) {
			*pPhysicalDeviceCount = driver.Uint32(1)
			devices := ([]driver.VkPhysicalDevice)(unsafe.Slice(pPhysicalDevices, 1))
			devices[0] = NewFakePhysicalDeviceHandle()

			return common.VKSuccess, nil
		})
	devices, _, err := instance.PhysicalDevices()
	require.NoError(t, err)

	return devices[0]
}

func EasyDummyPipeline(t *testing.T, device core1_0.Device, loader loader.Loader) core1_0.Pipeline {
	mockDriver := loader.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), nil, driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pPipelines *driver.VkPipeline) (common.VkResult, error) {
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pPipelines, 1))
			pipelines[0] = NewFakePipeline()

			return common.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineOptions{
		{},
	})
	require.NoError(t, err)
	return pipelines[0]
}

func EasyDummyPipelineCache(t *testing.T, device core1_0.Device, loader loader.Loader) core1_0.PipelineCache {
	mockDriver := loader.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreatePipelineCache(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkPipelineCacheCreateInfo, pAllocator *driver.VkAllocationCallbacks, pPipelineCache *driver.VkPipelineCache) (common.VkResult, error) {
			*pPipelineCache = NewFakePipelineCache()
			return common.VKSuccess, nil
		})

	pipelineCache, _, err := loader.CreatePipelineCache(device, nil, &core1_0.PipelineCacheOptions{})
	require.NoError(t, err)
	return pipelineCache
}

func EasyDummyQueryPool(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.QueryPool {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateQueryPool(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkQueryPoolCreateInfo, pAllocator *driver.VkAllocationCallbacks, pQueryPool *driver.VkQueryPool) (common.VkResult, error) {
			*pQueryPool = NewFakeQueryPool()
			return common.VKSuccess, nil
		})

	queryPool, _, err := loader.CreateQueryPool(device, nil, &core1_0.QueryPoolOptions{})
	require.NoError(t, err)
	return queryPool
}

func EasyDummyQueue(device core1_0.Device) core1_0.Queue {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkGetDeviceQueue(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, queueFamilyIndex, queueIndex driver.Uint32, pQueue *driver.VkQueue) error {
			*pQueue = NewFakeQueue()
			return nil
		})

	queue := device.GetQueue(0, 0)

	return queue
}

func EasyDummyRenderPass(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.RenderPass {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateRenderPass(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkRenderPassCreateInfo, pAllocator *driver.VkAllocationCallbacks, pRenderPass *driver.VkRenderPass) (common.VkResult, error) {
			*pRenderPass = NewFakeRenderPassHandle()
			return common.VKSuccess, nil
		})

	renderPass, _, err := loader.CreateRenderPass(device, nil, &core1_0.RenderPassOptions{})
	require.NoError(t, err)

	return renderPass
}

func EasyDummySemaphore(t *testing.T, loader loader.Loader, device core1_0.Device) core1_0.Semaphore {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateSemaphore(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkSemaphoreCreateInfo, pAllocator *driver.VkAllocationCallbacks, pSemaphore *driver.VkSemaphore) (common.VkResult, error) {
			*pSemaphore = NewFakeSemaphore()
			return common.VKSuccess, nil
		})

	semaphore, _, err := loader.CreateSemaphore(device, nil, &core1_0.SemaphoreOptions{})
	require.NoError(t, err)

	return semaphore
}
