package dummies

import (
	"github.com/vkngwrapper/core/v3/common/extensions"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/mocks"
)

func EasyDummyBuffer(driver driver.Driver, device core1_0.Device) core1_0.Buffer {
	handle := mocks.NewFakeBufferHandle()

	return extensions.CreateBufferObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyCommandPool(driver driver.Driver, device core1_0.Device) core1_0.CommandPool {
	handle := mocks.NewFakeCommandPoolHandle()

	return extensions.CreateCommandPoolObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyCommandBuffer(driver driver.Driver, device core1_0.Device, commandPool core1_0.CommandPool) core1_0.CommandBuffer {
	handle := mocks.NewFakeCommandBufferHandle()

	return extensions.CreateCommandBufferObject(driver, commandPool.Handle(), device.Handle(), handle, driver.Version())
}

func EasyDummyDescriptorPool(driver driver.Driver, device core1_0.Device) core1_0.DescriptorPool {
	handle := mocks.NewFakeDescriptorPool()

	return extensions.CreateDescriptorPoolObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyDescriptorSet(driver driver.Driver, pool core1_0.DescriptorPool) core1_0.DescriptorSet {
	handle := mocks.NewFakeDescriptorSet()

	return extensions.CreateDescriptorSetObject(driver, pool.DeviceHandle(), pool.Handle(), handle, driver.Version())
}

func EasyDummyDescriptorSetLayout(driver driver.Driver, device core1_0.Device) core1_0.DescriptorSetLayout {
	handle := mocks.NewFakeDescriptorSetLayout()

	return extensions.CreateDescriptorSetLayoutObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyDevice(driver driver.Driver) core1_0.Device {
	handle := mocks.NewFakeDeviceHandle()

	return extensions.CreateDeviceObject(driver, handle, driver.Version())
}

func EasyDummyDeviceMemory(driver driver.Driver, device core1_0.Device, size int) core1_0.DeviceMemory {
	handle := mocks.NewFakeDeviceMemoryHandle()

	return extensions.CreateDeviceMemoryObject(driver, device.Handle(), handle, driver.Version(), size)
}

func EasyDummyEvent(driver driver.Driver, device core1_0.Device) core1_0.Event {
	handle := mocks.NewFakeEventHandle()

	return extensions.CreateEventObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyFence(driver driver.Driver, device core1_0.Device) core1_0.Fence {
	handle := mocks.NewFakeFenceHandle()

	return extensions.CreateFenceObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyFramebuffer(driver driver.Driver, device core1_0.Device) core1_0.Framebuffer {
	handle := mocks.NewFakeFramebufferHandle()

	return extensions.CreateFramebufferObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyImage(driver driver.Driver, device core1_0.Device) core1_0.Image {
	handle := mocks.NewFakeImageHandle()

	return extensions.CreateImageObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyInstance(driver driver.Driver) core1_0.Instance {
	handle := mocks.NewFakeInstanceHandle()

	return extensions.CreateInstanceObject(driver, handle, driver.Version())
}

func EasyDummyPhysicalDevice(driver driver.Driver, instance core1_0.Instance) core1_0.PhysicalDevice {
	handle := mocks.NewFakePhysicalDeviceHandle()

	return extensions.CreatePhysicalDeviceObject(driver, instance.Handle(), handle, driver.Version(), driver.Version())
}

func EasyDummyPipeline(driver driver.Driver, device core1_0.Device) core1_0.Pipeline {
	handle := mocks.NewFakePipeline()

	return extensions.CreatePipelineObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyPipelineCache(driver driver.Driver, device core1_0.Device) core1_0.PipelineCache {
	handle := mocks.NewFakePipelineCache()

	return extensions.CreatePipelineCacheObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyQueryPool(driver driver.Driver, device core1_0.Device) core1_0.QueryPool {
	handle := mocks.NewFakeQueryPool()

	return extensions.CreateQueryPoolObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyQueue(driver driver.Driver, device core1_0.Device) core1_0.Queue {
	handle := mocks.NewFakeQueue()

	return extensions.CreateQueueObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyRenderPass(driver driver.Driver, device core1_0.Device) core1_0.RenderPass {
	handle := mocks.NewFakeRenderPassHandle()

	return extensions.CreateRenderPassObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummySemaphore(driver driver.Driver, device core1_0.Device) core1_0.Semaphore {
	handle := mocks.NewFakeSemaphore()

	return extensions.CreateSemaphoreObject(driver, device.Handle(), handle, driver.Version())
}
