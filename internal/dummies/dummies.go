package dummies

import (
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	internal1_0 "github.com/CannibalVox/VKng/core/internal/core1_0"
	"github.com/CannibalVox/VKng/core/mocks"
)

func EasyDummyBuffer(driver driver.Driver, device core1_0.Device) core1_0.Buffer {
	handle := mocks.NewFakeBufferHandle()

	return internal1_0.CreateBufferObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyCommandPool(driver driver.Driver, device core1_0.Device) core1_0.CommandPool {
	handle := mocks.NewFakeCommandPoolHandle()

	return internal1_0.CreateCommandPoolObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyCommandBuffer(driver driver.Driver, device core1_0.Device, commandPool core1_0.CommandPool) core1_0.CommandBuffer {
	handle := mocks.NewFakeCommandBufferHandle()

	return internal1_0.CreateCommandBufferObject(driver, commandPool.Handle(), device.Handle(), handle, driver.Version())
}

func EasyDummyDescriptorPool(driver driver.Driver, device core1_0.Device) core1_0.DescriptorPool {
	handle := mocks.NewFakeDescriptorPool()

	return internal1_0.CreateDescriptorPoolObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyDescriptorSet(driver driver.Driver, pool core1_0.DescriptorPool) core1_0.DescriptorSet {
	handle := mocks.NewFakeDescriptorSet()

	return internal1_0.CreateDescriptorSetObject(driver, pool.DeviceHandle(), pool.Handle(), handle, driver.Version())
}

func EasyDummyDescriptorSetLayout(driver driver.Driver, device core1_0.Device) core1_0.DescriptorSetLayout {
	handle := mocks.NewFakeDescriptorSetLayout()

	return internal1_0.CreateDescriptorSetLayoutObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyDevice(driver driver.Driver) core1_0.Device {
	handle := mocks.NewFakeDeviceHandle()

	return internal1_0.CreateDeviceObject(driver, handle, driver.Version())
}

func EasyDummyDeviceMemory(driver driver.Driver, device core1_0.Device, size int) core1_0.DeviceMemory {
	handle := mocks.NewFakeDeviceMemoryHandle()

	return internal1_0.CreateDeviceMemoryObject(driver, device.Handle(), handle, driver.Version(), size)
}

func EasyDummyEvent(driver driver.Driver, device core1_0.Device) core1_0.Event {
	handle := mocks.NewFakeEventHandle()

	return internal1_0.CreateEventObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyFence(driver driver.Driver, device core1_0.Device) core1_0.Fence {
	handle := mocks.NewFakeFenceHandle()

	return internal1_0.CreateFenceObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyFramebuffer(driver driver.Driver, device core1_0.Device) core1_0.Framebuffer {
	handle := mocks.NewFakeFramebufferHandle()

	return internal1_0.CreateFramebufferObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyImage(driver driver.Driver, device core1_0.Device) core1_0.Image {
	handle := mocks.NewFakeImageHandle()

	return internal1_0.CreateImageObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyInstance(driver driver.Driver) core1_0.Instance {
	handle := mocks.NewFakeInstanceHandle()

	return internal1_0.CreateInstanceObject(driver, handle, driver.Version())
}

func EasyDummyPhysicalDevice(driver driver.Driver, instance core1_0.Instance) core1_0.PhysicalDevice {
	handle := mocks.NewFakePhysicalDeviceHandle()

	return internal1_0.CreatePhysicalDeviceObject(driver, instance.Handle(), handle, driver.Version(), driver.Version())
}

func EasyDummyPipeline(driver driver.Driver, device core1_0.Device) core1_0.Pipeline {
	handle := mocks.NewFakePipeline()

	return internal1_0.CreatePipelineObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyPipelineCache(driver driver.Driver, device core1_0.Device) core1_0.PipelineCache {
	handle := mocks.NewFakePipelineCache()

	return internal1_0.CreatePipelineCacheObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyQueryPool(driver driver.Driver, device core1_0.Device) core1_0.QueryPool {
	handle := mocks.NewFakeQueryPool()

	return internal1_0.CreateQueryPoolObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyQueue(driver driver.Driver, device core1_0.Device) core1_0.Queue {
	handle := mocks.NewFakeQueue()

	return internal1_0.CreateQueueObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummyRenderPass(driver driver.Driver, device core1_0.Device) core1_0.RenderPass {
	handle := mocks.NewFakeRenderPassHandle()

	return internal1_0.CreateRenderPassObject(driver, device.Handle(), handle, driver.Version())
}

func EasyDummySemaphore(driver driver.Driver, device core1_0.Device) core1_0.Semaphore {
	handle := mocks.NewFakeSemaphore()

	return internal1_0.CreateSemaphoreObject(driver, device.Handle(), handle, driver.Version())
}
