package mocks

import (
	"github.com/CannibalVox/VKng/core"
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

func EasyMockDescriptorSet(ctrl *gomock.Controller) *MockDescriptorSet {
	set := NewMockDescriptorSet(ctrl)
	set.EXPECT().Handle().Return(NewFakeDescriptorSet()).AnyTimes()

	return set
}

func EasyMockDevice(ctrl *gomock.Controller, driver core.Driver) *MockDevice {
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

func EasyMockPhysicalDevice(ctrl *gomock.Controller, driver core.Driver) *MockPhysicalDevice {
	physicalDevice := NewMockPhysicalDevice(ctrl)
	physicalDevice.EXPECT().Handle().Return(NewFakePhysicalDeviceHandle()).AnyTimes()
	physicalDevice.EXPECT().Driver().Return(driver).AnyTimes()

	return physicalDevice
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

func EasyDummyBuffer(t *testing.T, loader core.Loader1_0, device core.Device) core.Buffer {
	handle := NewFakeBufferHandle()
	driver := device.Driver().(*MockDriver)
	driver.EXPECT().VkCreateBuffer(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkBufferCreateInfo, pAllocator *core.VkAllocationCallbacks, pBuffer *core.VkBuffer) (core.VkResult, error) {
			*pBuffer = handle
			return core.VKSuccess, nil
		})

	buffer, _, err := loader.CreateBuffer(device, &core.BufferOptions{})
	require.NoError(t, err)
	require.NotNil(t, buffer)

	return buffer
}

func EasyDummyCommandPool(t *testing.T, loader core.Loader1_0, device core.Device) core.CommandPool {
	handle := NewFakeCommandPoolHandle()
	driver := device.Driver().(*MockDriver)
	driver.EXPECT().VkCreateCommandPool(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, createInfo *core.VkCommandPoolCreateInfo, allocator *core.VkAllocationCallbacks, commandPool *core.VkCommandPool) (core.VkResult, error) {
			*commandPool = handle

			return core.VKSuccess, nil
		})

	graphicsFamily := 0
	pool, res, err := loader.CreateCommandPool(device, &core.CommandPoolOptions{
		Flags:               core.CommandPoolResetBuffer,
		GraphicsQueueFamily: &graphicsFamily,
	})
	require.NoError(t, err)
	require.Equal(t, core.VKSuccess, res)

	return pool
}

func EasyDummyCommandBuffer(t *testing.T, device core.Device, commandPool core.CommandPool) core.CommandBuffer {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkAllocateCommandBuffers(gomock.Any(), gomock.Any(), gomock.Any()).Do(
		func(device core.VkDevice, pAllocateInfo *core.VkCommandBufferAllocateInfo, pCommandBuffers *core.VkCommandBuffer) {
			*pCommandBuffers = NewFakeCommandBufferHandle()
		})

	buffers, _, err := commandPool.AllocateCommandBuffers(&core.CommandBufferOptions{
		BufferCount: 1,
	})
	require.NoError(t, err)

	return buffers[0]
}

func EasyDummyDescriptorPool(t *testing.T, loader core.Loader1_0, device core.Device) core.DescriptorPool {
	mockDriver := device.Driver().(*MockDriver)

	mockDriver.EXPECT().VkCreateDescriptorPool(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkDescriptorPoolCreateInfo, pAllocator *core.VkAllocationCallbacks, pDescriptorPool *core.VkDescriptorPool) (core.VkResult, error) {
			*pDescriptorPool = NewFakeDescriptorPool()
			return core.VKSuccess, nil
		})

	pool, _, err := loader.CreateDescriptorPool(device, &core.DescriptorPoolOptions{})
	require.NoError(t, err)

	return pool
}

func EasyDummyDescriptorSetLayout(t *testing.T, loader core.Loader1_0, device core.Device) core.DescriptorSetLayout {
	mockDriver := device.Driver().(*MockDriver)

	mockDriver.EXPECT().VkCreateDescriptorSetLayout(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkDescriptorSetLayoutCreateInfo, pAllocator *core.VkAllocationCallbacks, pDescriptorSetLayout *core.VkDescriptorSetLayout) (core.VkResult, error) {
			*pDescriptorSetLayout = NewFakeDescriptorSetLayout()
			return core.VKSuccess, nil
		})

	layout, _, err := loader.CreateDescriptorSetLayout(device, &core.DescriptorSetLayoutOptions{})
	require.NoError(t, err)
	return layout
}

func EasyDummyDevice(t *testing.T, ctrl *gomock.Controller, loader core.Loader1_0) core.Device {
	mockDriver := loader.Driver().(*MockDriver)

	mockDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(mockDriver, nil)
	mockDriver.EXPECT().VkCreateDevice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(physicalDevice core.VkPhysicalDevice, pCreateInfo *core.VkDeviceCreateInfo, pAllocator *core.VkAllocationCallbacks, pDevice *core.VkDevice) (core.VkResult, error) {
			*pDevice = NewFakeDeviceHandle()
			return core.VKSuccess, nil
		})

	device, _, err := loader.CreateDevice(EasyMockPhysicalDevice(ctrl, mockDriver), &core.DeviceOptions{
		QueueFamilies: []*core.QueueFamilyOptions{
			{
				QueuePriorities: []float32{1},
			},
		},
	})
	require.NoError(t, err)

	return device
}

func EasyDummyDeviceMemory(t *testing.T, device core.Device) core.DeviceMemory {
	mockDriver := device.Driver().(*MockDriver)

	mockDriver.EXPECT().VkAllocateMemory(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkMemoryAllocateInfo, pAllocator *core.VkAllocationCallbacks, pDeviceMemory *core.VkDeviceMemory) (core.VkResult, error) {
			*pDeviceMemory = NewFakeDeviceMemoryHandle()
			return core.VKSuccess, nil
		})

	memory, _, err := device.AllocateMemory(&core.DeviceMemoryOptions{})
	require.NoError(t, err)

	return memory
}

func EasyDummyFence(t *testing.T, loader core.Loader1_0, device core.Device) core.Fence {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkCreateFence(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkFenceCreateInfo, pAllocator *core.VkAllocationCallbacks, pFence *core.VkFence) (core.VkResult, error) {
			*pFence = NewFakeFenceHandle()
			return core.VKSuccess, nil
		})

	fence, _, err := loader.CreateFence(device, &core.FenceOptions{})
	require.NoError(t, err)

	return fence
}

func EasyDummyFramebuffer(t *testing.T, loader core.Loader1_0, device core.Device) core.Framebuffer {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkCreateFramebuffer(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkFramebufferCreateInfo, pAllocator *core.VkAllocationCallbacks, pFramebuffer *core.VkFramebuffer) (core.VkResult, error) {
			*pFramebuffer = NewFakeFramebufferHandle()
			return core.VKSuccess, nil
		})

	framebuffer, _, err := loader.CreateFrameBuffer(device, &core.FramebufferOptions{})
	require.NoError(t, err)

	return framebuffer
}

func EasyDummyGraphicsPipeline(t *testing.T, loader core.Loader1_0, device core.Device) core.Pipeline {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkCreateGraphicsPipelines(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, cache core.VkPipelineCache, createInfoCount core.Uint32, pCreateInfos *core.VkGraphicsPipelineCreateInfo, pAllocator *core.VkAllocationCallbacks, pPipelines *core.VkPipeline) (core.VkResult, error) {
			*pPipelines = NewFakePipeline()
			return core.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, []*core.GraphicsPipelineOptions{{}})
	require.NoError(t, err)

	return pipelines[0]
}

func EasyDummyInstance(t *testing.T, loader core.Loader1_0) core.Instance {
	driver := loader.Driver().(*MockDriver)

	driver.EXPECT().CreateInstanceDriver(gomock.Any()).Return(driver, nil).AnyTimes()
	driver.EXPECT().VkCreateInstance(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(pCreateInfo *core.VkInstanceCreateInfo, pAllocator *core.VkAllocationCallbacks, pInstance *core.VkInstance) (core.VkResult, error) {
			*pInstance = NewFakeInstanceHandle()
			return core.VKSuccess, nil
		})

	instance, _, err := loader.CreateInstance(&core.InstanceOptions{})
	require.NoError(t, err)

	return instance
}

func EasyDummyPhysicalDevice(t *testing.T, loader core.Loader1_0) core.PhysicalDevice {
	driver := loader.Driver().(*MockDriver)
	instance := EasyDummyInstance(t, loader)

	driver.EXPECT().VkEnumeratePhysicalDevices(instance.Handle(), gomock.Not(nil), nil).DoAndReturn(
		func(instance core.VkInstance, pPhysicalDeviceCount *core.Uint32, pPhysicalDevices *core.VkPhysicalDevice) (core.VkResult, error) {
			*pPhysicalDeviceCount = core.Uint32(1)
			return core.VKSuccess, nil
		})
	driver.EXPECT().VkEnumeratePhysicalDevices(instance.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(instance core.VkInstance, pPhysicalDeviceCount *core.Uint32, pPhysicalDevices *core.VkPhysicalDevice) (core.VkResult, error) {
			*pPhysicalDeviceCount = core.Uint32(1)
			devices := ([]core.VkPhysicalDevice)(unsafe.Slice(pPhysicalDevices, 1))
			devices[0] = NewFakePhysicalDeviceHandle()

			return core.VKSuccess, nil
		})
	devices, _, err := instance.PhysicalDevices()
	require.NoError(t, err)

	return devices[0]
}

func EasyDummyQueue(t *testing.T, device core.Device) core.Queue {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkGetDeviceQueue(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, queueFamilyIndex, queueIndex core.Uint32, pQueue *core.VkQueue) error {
			*pQueue = NewFakeQueue()
			return nil
		})

	queue, err := device.GetQueue(0, 0)
	require.NoError(t, err)

	return queue
}

func EasyDummyRenderPass(t *testing.T, loader core.Loader1_0, device core.Device) core.RenderPass {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkCreateRenderPass(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkRenderPassCreateInfo, pAllocator *core.VkAllocationCallbacks, pRenderPass *core.VkRenderPass) (core.VkResult, error) {
			*pRenderPass = NewFakeRenderPassHandle()
			return core.VKSuccess, nil
		})

	renderPass, _, err := loader.CreateRenderPass(device, &core.RenderPassOptions{})
	require.NoError(t, err)

	return renderPass
}

func EasyDummySemaphore(t *testing.T, loader core.Loader1_0, device core.Device) core.Semaphore {
	driver := device.Driver().(*MockDriver)

	driver.EXPECT().VkCreateSemaphore(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkSemaphoreCreateInfo, pAllocator *core.VkAllocationCallbacks, pSemaphore *core.VkSemaphore) (core.VkResult, error) {
			*pSemaphore = NewFakeSemaphore()
			return core.VKSuccess, nil
		})

	semaphore, _, err := loader.CreateSemaphore(device, &core.SemaphoreOptions{})
	require.NoError(t, err)

	return semaphore
}
