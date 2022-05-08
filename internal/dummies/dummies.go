package dummies

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func EasyDummyBuffer(t *testing.T, loader core.Loader, device core.Device) core.Buffer {
	handle := mocks.NewFakeBufferHandle()
	mockDriver := device.Driver().(*mock_driver.MockDriver)
	mockDriver.EXPECT().VkCreateBuffer(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkBufferCreateInfo, pAllocator *driver.VkAllocationCallbacks, pBuffer *driver.VkBuffer) (common.VkResult, error) {
			*pBuffer = handle
			return core1_0.VKSuccess, nil
		})

	buffer, _, err := loader.CreateBuffer(device, nil, core1_0.BufferCreateOptions{})
	require.NoError(t, err)
	require.NotNil(t, buffer)

	return buffer
}

func EasyDummyCommandPool(t *testing.T, loader core.Loader, device core.Device) core.CommandPool {
	handle := mocks.NewFakeCommandPoolHandle()
	mockDriver := device.Driver().(*mock_driver.MockDriver)
	mockDriver.EXPECT().VkCreateCommandPool(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, createInfo *driver.VkCommandPoolCreateInfo, allocator *driver.VkAllocationCallbacks, commandPool *driver.VkCommandPool) (common.VkResult, error) {
			*commandPool = handle

			return core1_0.VKSuccess, nil
		})

	graphicsFamily := 0
	pool, res, err := loader.CreateCommandPool(device, nil, core1_0.CommandPoolCreateOptions{
		Flags:               core1_0.CommandPoolCreateResetBuffer,
		GraphicsQueueFamily: &graphicsFamily,
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)

	return pool
}

func EasyDummyCommandBuffer(t *testing.T, loader core.Loader, device core.Device, commandPool core.CommandPool) core.CommandBuffer {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkAllocateCommandBuffers(gomock.Any(), gomock.Any(), gomock.Any()).Do(
		func(device driver.VkDevice, pAllocateInfo *driver.VkCommandBufferAllocateInfo, pCommandBuffers *driver.VkCommandBuffer) {
			*pCommandBuffers = mocks.NewFakeCommandBufferHandle()
		})

	buffers, _, err := loader.AllocateCommandBuffers(core1_0.CommandBufferAllocateOptions{
		CommandPool: commandPool,
		BufferCount: 1,
	})
	require.NoError(t, err)

	return buffers[0]
}

func EasyDummyDescriptorPool(t *testing.T, loader core.Loader, device core.Device) core.DescriptorPool {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateDescriptorPool(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorPoolCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDescriptorPool *driver.VkDescriptorPool) (common.VkResult, error) {
			*pDescriptorPool = mocks.NewFakeDescriptorPool()
			return core1_0.VKSuccess, nil
		})

	pool, _, err := loader.CreateDescriptorPool(device, nil, core1_0.DescriptorPoolCreateOptions{})
	require.NoError(t, err)

	return pool
}

func EasyDummyDescriptorSet(t *testing.T, loader core.Loader, pool core.DescriptorPool, layout core.DescriptorSetLayout) core1_0.DescriptorSet {
	mockDriver := pool.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkAllocateDescriptorSets(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetAllocateInfo, pDescriptorSets *driver.VkDescriptorSet) (common.VkResult, error) {
			*pDescriptorSets = mocks.NewFakeDescriptorSet()

			return core1_0.VKSuccess, nil
		})

	sets, _, err := loader.AllocateDescriptorSets(core1_0.DescriptorSetAllocateOptions{
		DescriptorPool:    pool,
		AllocationLayouts: []core1_0.DescriptorSetLayout{layout},
	})
	require.NoError(t, err)

	return sets[0]
}

func EasyDummyDescriptorSetLayout(t *testing.T, loader core.Loader, device core.Device) core.DescriptorSetLayout {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateDescriptorSetLayout(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDescriptorSetLayout *driver.VkDescriptorSetLayout) (common.VkResult, error) {
			*pDescriptorSetLayout = mocks.NewFakeDescriptorSetLayout()
			return core1_0.VKSuccess, nil
		})

	layout, _, err := loader.CreateDescriptorSetLayout(device, nil, core1_0.DescriptorSetLayoutCreateOptions{})
	require.NoError(t, err)
	return layout
}

func EasyDummyDevice(t *testing.T, ctrl *gomock.Controller, loader core.Loader) core.Device {
	mockDriver := loader.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(mockDriver, nil)
	mockDriver.EXPECT().VkCreateDevice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice, pCreateInfo *driver.VkDeviceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDevice *driver.VkDevice) (common.VkResult, error) {
			*pDevice = mocks.NewFakeDeviceHandle()
			return core1_0.VKSuccess, nil
		})

	device, _, err := loader.CreateDevice(mocks.EasyMockPhysicalDevice(ctrl, mockDriver), nil, core1_0.DeviceCreateOptions{
		QueueFamilies: []core1_0.DeviceQueueCreateOptions{
			{
				CreatedQueuePriorities: []float32{1},
			},
		},
	})
	require.NoError(t, err)

	return device
}

func EasyDummyDeviceMemory(t *testing.T, loader core.Loader, device core.Device, size int) core.DeviceMemory {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkAllocateMemory(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkMemoryAllocateInfo, pAllocator *driver.VkAllocationCallbacks, pDeviceMemory *driver.VkDeviceMemory) (common.VkResult, error) {
			*pDeviceMemory = mocks.NewFakeDeviceMemoryHandle()
			return core1_0.VKSuccess, nil
		})

	memory, _, err := loader.AllocateMemory(device, nil, core1_0.MemoryAllocateOptions{
		AllocationSize: size,
	})
	require.NoError(t, err)

	return memory
}

func EasyDummyEvent(t *testing.T, loader core.Loader, device core.Device) core.Event {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateEvent(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkEventCreateInfo, pAllocator *driver.VkAllocationCallbacks, pEvent *driver.VkEvent) (common.VkResult, error) {
			*pEvent = mocks.NewFakeEventHandle()
			return core1_0.VKSuccess, nil
		})

	event, _, err := loader.CreateEvent(device, nil, core1_0.EventCreateOptions{})
	require.NoError(t, err)

	return event
}

func EasyDummyFence(t *testing.T, loader core.Loader, device core.Device) core.Fence {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateFence(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkFenceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pFence *driver.VkFence) (common.VkResult, error) {
			*pFence = mocks.NewFakeFenceHandle()
			return core1_0.VKSuccess, nil
		})

	fence, _, err := loader.CreateFence(device, nil, core1_0.FenceCreateOptions{})
	require.NoError(t, err)

	return fence
}

func EasyDummyFramebuffer(t *testing.T, loader core.Loader, device core.Device) core.Framebuffer {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateFramebuffer(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkFramebufferCreateInfo, pAllocator *driver.VkAllocationCallbacks, pFramebuffer *driver.VkFramebuffer) (common.VkResult, error) {
			*pFramebuffer = mocks.NewFakeFramebufferHandle()
			return core1_0.VKSuccess, nil
		})

	framebuffer, _, err := loader.CreateFrameBuffer(device, nil, core1_0.FramebufferCreateOptions{})
	require.NoError(t, err)

	return framebuffer
}

func EasyDummyGraphicsPipeline(t *testing.T, loader core.Loader, device core.Device) core.Pipeline {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateGraphicsPipelines(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pPipelines *driver.VkPipeline) (common.VkResult, error) {
			*pPipelines = mocks.NewFakePipeline()
			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineCreateOptions{{}})
	require.NoError(t, err)

	return pipelines[0]
}

func EasyDummyImage(t *testing.T, loader core.Loader, device core.Device) core.Image {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateImage(device.Handle(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkImageCreateInfo, pAllocator *driver.VkAllocationCallbacks, pImage *driver.VkImage) (common.VkResult, error) {
			*pImage = mocks.NewFakeImageHandle()

			return core1_0.VKSuccess, nil
		})

	image, _, err := loader.CreateImage(device, nil, core1_0.ImageCreateOptions{})
	require.NoError(t, err)

	return image
}

func EasyDummyInstance(t *testing.T, loader core.Loader) core.Instance {
	mockDriver := loader.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().CreateInstanceDriver(gomock.Any()).Return(mockDriver, nil).AnyTimes()
	mockDriver.EXPECT().VkCreateInstance(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(pCreateInfo *driver.VkInstanceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pInstance *driver.VkInstance) (common.VkResult, error) {
			*pInstance = mocks.NewFakeInstanceHandle()
			return core1_0.VKSuccess, nil
		})

	instance, _, err := loader.CreateInstance(nil, core1_0.InstanceCreateOptions{
		VulkanVersion: loader.Version(),
	})
	require.NoError(t, err)

	return instance
}

func EasyDummyPhysicalDevice(t *testing.T, loader core.Loader) core.PhysicalDevice {
	mockDriver := loader.Driver().(*mock_driver.MockDriver)
	instance := EasyDummyInstance(t, loader)

	handle := mocks.NewFakePhysicalDeviceHandle()

	mockDriver.EXPECT().VkEnumeratePhysicalDevices(instance.Handle(), gomock.Not(nil), nil).DoAndReturn(
		func(instance driver.VkInstance, pPhysicalDeviceCount *driver.Uint32, pPhysicalDevices *driver.VkPhysicalDevice) (common.VkResult, error) {
			*pPhysicalDeviceCount = driver.Uint32(1)
			return core1_0.VKSuccess, nil
		})
	mockDriver.EXPECT().VkEnumeratePhysicalDevices(instance.Handle(), gomock.Not(nil), gomock.Not(nil)).DoAndReturn(
		func(instance driver.VkInstance, pPhysicalDeviceCount *driver.Uint32, pPhysicalDevices *driver.VkPhysicalDevice) (common.VkResult, error) {
			*pPhysicalDeviceCount = driver.Uint32(1)
			devices := ([]driver.VkPhysicalDevice)(unsafe.Slice(pPhysicalDevices, 1))
			devices[0] = handle

			return core1_0.VKSuccess, nil
		})
	mockDriver.EXPECT().VkGetPhysicalDeviceProperties(handle, gomock.Not(nil)).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			val := reflect.ValueOf(pProperties).Elem()

			*(*uint32)(unsafe.Pointer(val.FieldByName("apiVersion").UnsafeAddr())) = uint32(loader.Version())
		})

	devices, _, err := loader.PhysicalDevices(instance)
	require.NoError(t, err)

	return devices[0]
}

func EasyDummyPipeline(t *testing.T, device core.Device, loader core.Loader) core.Pipeline {
	mockDriver := loader.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateGraphicsPipelines(device.Handle(), driver.VkPipelineCache(driver.NullHandle), driver.Uint32(1), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, cache driver.VkPipelineCache, createInfoCount driver.Uint32, pCreateInfos *driver.VkGraphicsPipelineCreateInfo, pAllocator *driver.VkAllocationCallbacks, pPipelines *driver.VkPipeline) (common.VkResult, error) {
			pipelines := ([]driver.VkPipeline)(unsafe.Slice(pPipelines, 1))
			pipelines[0] = mocks.NewFakePipeline()

			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := loader.CreateGraphicsPipelines(device, nil, nil, []core1_0.GraphicsPipelineCreateOptions{
		{},
	})
	require.NoError(t, err)
	return pipelines[0]
}

func EasyDummyPipelineCache(t *testing.T, device core.Device, loader core.Loader) core.PipelineCache {
	mockDriver := loader.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreatePipelineCache(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkPipelineCacheCreateInfo, pAllocator *driver.VkAllocationCallbacks, pPipelineCache *driver.VkPipelineCache) (common.VkResult, error) {
			*pPipelineCache = mocks.NewFakePipelineCache()
			return core1_0.VKSuccess, nil
		})

	pipelineCache, _, err := loader.CreatePipelineCache(device, nil, core1_0.PipelineCacheCreateOptions{})
	require.NoError(t, err)
	return pipelineCache
}

func EasyDummyQueryPool(t *testing.T, loader core.Loader, device core.Device) core.QueryPool {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateQueryPool(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkQueryPoolCreateInfo, pAllocator *driver.VkAllocationCallbacks, pQueryPool *driver.VkQueryPool) (common.VkResult, error) {
			*pQueryPool = mocks.NewFakeQueryPool()
			return core1_0.VKSuccess, nil
		})

	queryPool, _, err := loader.CreateQueryPool(device, nil, core1_0.QueryPoolCreateOptions{})
	require.NoError(t, err)
	return queryPool
}

func EasyDummyQueue(loader core.Loader, device core.Device) core.Queue {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkGetDeviceQueue(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, queueFamilyIndex, queueIndex driver.Uint32, pQueue *driver.VkQueue) error {
			*pQueue = mocks.NewFakeQueue()
			return nil
		})

	queue := loader.GetQueue(device, 0, 0)

	return queue
}

func EasyDummyRenderPass(t *testing.T, loader core.Loader, device core.Device) core.RenderPass {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateRenderPass(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkRenderPassCreateInfo, pAllocator *driver.VkAllocationCallbacks, pRenderPass *driver.VkRenderPass) (common.VkResult, error) {
			*pRenderPass = mocks.NewFakeRenderPassHandle()
			return core1_0.VKSuccess, nil
		})

	renderPass, _, err := loader.CreateRenderPass(device, nil, core1_0.RenderPassCreateOptions{})
	require.NoError(t, err)

	return renderPass
}

func EasyDummySemaphore(t *testing.T, loader core.Loader, device core.Device) core.Semaphore {
	mockDriver := device.Driver().(*mock_driver.MockDriver)

	mockDriver.EXPECT().VkCreateSemaphore(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkSemaphoreCreateInfo, pAllocator *driver.VkAllocationCallbacks, pSemaphore *driver.VkSemaphore) (common.VkResult, error) {
			*pSemaphore = mocks.NewFakeSemaphore()
			return core1_0.VKSuccess, nil
		})

	semaphore, _, err := loader.CreateSemaphore(device, nil, core1_0.SemaphoreCreateOptions{})
	require.NoError(t, err)

	return semaphore
}
