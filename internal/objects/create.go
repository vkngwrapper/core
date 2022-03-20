package objects

import (
	"github.com/CannibalVox/VKng/core/common"
	iface "github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/internal/core1_0"
	"github.com/CannibalVox/VKng/core/internal/core1_1"
)

func CreateInstance(instanceDriver driver.Driver, handle driver.VkInstance, version common.APIVersion) *core1_0.VulkanInstance {
	return instanceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			instance := &core1_0.VulkanInstance{
				InstanceDriver: instanceDriver,
				InstanceHandle: handle,
				MaximumVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				instance.Instance1_1 = &core1_1.VulkanInstance{
					InstanceDriver: instanceDriver,
					InstanceHandle: handle,
				}
			}

			return instance
		}).(*core1_0.VulkanInstance)
}

func CreatePhysicalDevice(instance iface.Instance, handle driver.VkPhysicalDevice, version common.APIVersion) *core1_0.VulkanPhysicalDevice {
	coreDriver := instance.Driver()
	physicalDevice := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			device := &core1_0.VulkanPhysicalDevice{
				InstanceDriver:       coreDriver,
				PhysicalDeviceHandle: handle,
				MaximumVersion:       version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				device.PhysicalDevice1_1 = &core1_1.VulkanPhysicalDevice{
					InstanceDriver:       coreDriver,
					PhysicalDeviceHandle: handle,
				}
			}

			return device
		}).(*core1_0.VulkanPhysicalDevice)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(instance.Handle()), driver.VulkanHandle(handle))
	return physicalDevice
}

func CreateDevice(deviceDriver driver.Driver, handle driver.VkDevice, version common.APIVersion) *core1_0.VulkanDevice {
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			device := &core1_0.VulkanDevice{
				DeviceDriver:      deviceDriver,
				DeviceHandle:      handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				device.Device1_1 = &core1_1.VulkanDevice{
					DeviceDriver: deviceDriver,
					DeviceHandle: handle,
				}
			}

			return device
		}).(*core1_0.VulkanDevice)
}

func CreateBuffer(device iface.Device, handle driver.VkBuffer) *core1_0.VulkanBuffer {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			buffer := &core1_0.VulkanBuffer{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				BufferHandle:      handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				buffer.Buffer1_1 = &core1_1.VulkanBuffer{
					Driver:       deviceDriver,
					Device:       deviceHandle,
					BufferHandle: handle,
				}
			}

			return buffer
		}).(*core1_0.VulkanBuffer)
}

func CreateBufferView(device iface.Device, handle driver.VkBufferView) *core1_0.VulkanBufferView {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			bufferView := &core1_0.VulkanBufferView{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				BufferViewHandle:  handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				bufferView.BufferView1_1 = &core1_1.VulkanBufferView{
					Driver:           deviceDriver,
					Device:           deviceHandle,
					BufferViewHandle: handle,
				}
			}

			return bufferView
		}).(*core1_0.VulkanBufferView)
}

func CreateCommandBuffer(commandPool iface.CommandPool, handle driver.VkCommandBuffer) *core1_0.VulkanCommandBuffer {
	deviceDriver := commandPool.Driver()
	deviceHandle := commandPool.Device()
	commandPoolHandle := commandPool.Handle()
	version := commandPool.APIVersion()
	commandBuffer := deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			commandBuffer := &core1_0.VulkanCommandBuffer{
				DeviceDriver:        deviceDriver,
				Device:              deviceHandle,
				CommandBufferHandle: handle,
				CommandPool:         commandPoolHandle,
				MaximumAPIVersion:   version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				commandBuffer.CommandBuffer1_1 = &core1_1.VulkanCommandBuffer{
					DeviceDriver:        deviceDriver,
					Device:              deviceHandle,
					CommandPool:         commandPoolHandle,
					CommandBufferHandle: handle,
				}
			}

			return commandBuffer
		}).(*core1_0.VulkanCommandBuffer)
	deviceDriver.ObjectStore().SetParent(driver.VulkanHandle(commandPoolHandle), driver.VulkanHandle(handle))
	return commandBuffer
}

func CreateCommandPool(device iface.Device, handle driver.VkCommandPool) *core1_0.VulkanCommandPool {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			commandPool := &core1_0.VulkanCommandPool{
				DeviceDriver:      deviceDriver,
				DeviceHandle:      deviceHandle,
				CommandPoolHandle: handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				commandPool.CommandPool1_1 = &core1_1.VulkanCommandPool{
					DeviceDriver:      deviceDriver,
					DeviceHandle:      deviceHandle,
					CommandPoolHandle: handle,
				}
			}

			return commandPool
		}).(*core1_0.VulkanCommandPool)
}

func CreateDescriptorPool(device iface.Device, handle driver.VkDescriptorPool) *core1_0.VulkanDescriptorPool {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			descriptorPool := &core1_0.VulkanDescriptorPool{
				DeviceDriver:         deviceDriver,
				Device:               deviceHandle,
				DescriptorPoolHandle: handle,
				MaximumAPIVersion:    version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				descriptorPool.DescriptorPool1_1 = &core1_1.VulkanDescriptorPool{
					DeviceDriver:         deviceDriver,
					Device:               deviceHandle,
					DescriptorPoolHandle: handle,
				}
			}

			return descriptorPool
		}).(*core1_0.VulkanDescriptorPool)
}

func CreateDescriptorSet(descriptorPool iface.DescriptorPool, handle driver.VkDescriptorSet) *core1_0.VulkanDescriptorSet {
	deviceDriver := descriptorPool.Driver()
	deviceHandle := descriptorPool.DeviceHandle()
	version := descriptorPool.APIVersion()
	descriptorPoolHandle := descriptorPool.Handle()
	descriptorSet := deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			descriptorSet := &core1_0.VulkanDescriptorSet{
				DeviceDriver:        deviceDriver,
				Device:              deviceHandle,
				DescriptorSetHandle: handle,
				MaximumAPIVersion:   version,
				DescriptorPool:      descriptorPoolHandle,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				descriptorSet.DescriptorSet1_1 = &core1_1.VulkanDescriptorSet{
					DeviceDriver:        deviceDriver,
					Device:              deviceHandle,
					DescriptorSetHandle: handle,
					DescriptorPool:      descriptorPoolHandle,
				}
			}

			return descriptorSet
		}).(*core1_0.VulkanDescriptorSet)
	deviceDriver.ObjectStore().SetParent(driver.VulkanHandle(descriptorPoolHandle), driver.VulkanHandle(handle))
	return descriptorSet
}

func CreateDescriptorSetLayout(device iface.Device, handle driver.VkDescriptorSetLayout) *core1_0.VulkanDescriptorSetLayout {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			descriptorSetLayout := &core1_0.VulkanDescriptorSetLayout{
				Driver:                    deviceDriver,
				Device:                    deviceHandle,
				DescriptorSetLayoutHandle: handle,
				MaximumAPIVersion:         version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				descriptorSetLayout.DescriptorSetLayout1_1 = &core1_1.VulkanDescriptorSetLayout{
					Driver:                    deviceDriver,
					Device:                    deviceHandle,
					DescriptorSetLayoutHandle: handle,
				}
			}

			return descriptorSetLayout
		}).(*core1_0.VulkanDescriptorSetLayout)
}

func CreateDeviceMemory(device iface.Device, handle driver.VkDeviceMemory, size int) *core1_0.VulkanDeviceMemory {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			deviceMemory := &core1_0.VulkanDeviceMemory{
				DeviceDriver:       deviceDriver,
				Device:             deviceHandle,
				DeviceMemoryHandle: handle,
				MaximumAPIVersion:  version,
				Size:               size,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				deviceMemory.DeviceMemory1_1 = &core1_1.VulkanDeviceMemory{
					DeviceDriver:       deviceDriver,
					Device:             deviceHandle,
					DeviceMemoryHandle: handle,
					Size:               size,
				}
			}

			return deviceMemory
		}).(*core1_0.VulkanDeviceMemory)
}

func CreateEvent(device iface.Device, handle driver.VkEvent) *core1_0.VulkanEvent {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			event := &core1_0.VulkanEvent{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				EventHandle:       handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				event.Event1_1 = &core1_1.VulkanEvent{
					Driver:      deviceDriver,
					Device:      deviceHandle,
					EventHandle: handle,
				}
			}

			return event
		}).(*core1_0.VulkanEvent)
}

func CreateFence(device iface.Device, handle driver.VkFence) *core1_0.VulkanFence {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			fence := &core1_0.VulkanFence{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				FenceHandle:       handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				fence.Fence1_1 = &core1_1.VulkanFence{
					Driver:      deviceDriver,
					Device:      deviceHandle,
					FenceHandle: handle,
				}
			}

			return fence
		}).(*core1_0.VulkanFence)
}

func CreateFramebuffer(device iface.Device, handle driver.VkFramebuffer) *core1_0.VulkanFramebuffer {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			framebuffer := &core1_0.VulkanFramebuffer{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				FramebufferHandle: handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				framebuffer.Framebuffer1_1 = &core1_1.VulkanFramebuffer{
					Driver:            deviceDriver,
					Device:            deviceHandle,
					FramebufferHandle: handle,
				}
			}

			return framebuffer
		}).(*core1_0.VulkanFramebuffer)
}

func CreateImage(device iface.Device, handle driver.VkImage) *core1_0.VulkanImage {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			image := &core1_0.VulkanImage{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				ImageHandle:       handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				image.Image1_1 = &core1_1.VulkanImage{
					Driver:      deviceDriver,
					Device:      deviceHandle,
					ImageHandle: handle,
				}
			}

			return image
		}).(*core1_0.VulkanImage)
}

func CreateImageView(device iface.Device, handle driver.VkImageView) *core1_0.VulkanImageView {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			imageView := &core1_0.VulkanImageView{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				ImageViewHandle:   handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				imageView.ImageView1_1 = &core1_1.VulkanImageView{
					Driver:          deviceDriver,
					Device:          deviceHandle,
					ImageViewHandle: handle,
				}
			}

			return imageView
		}).(*core1_0.VulkanImageView)
}

func CreatePipeline(device iface.Device, handle driver.VkPipeline) *core1_0.VulkanPipeline {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipeline := &core1_0.VulkanPipeline{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				PipelineHandle:    handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipeline.Pipeline1_1 = &core1_1.VulkanPipeline{
					Driver:         deviceDriver,
					Device:         deviceHandle,
					PipelineHandle: handle,
				}
			}

			return pipeline
		}).(*core1_0.VulkanPipeline)
}

func CreatePipelineCache(device iface.Device, handle driver.VkPipelineCache) *core1_0.VulkanPipelineCache {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipelineCache := &core1_0.VulkanPipelineCache{
				Driver:              deviceDriver,
				Device:              deviceHandle,
				PipelineCacheHandle: handle,
				MaximumAPIVersion:   version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipelineCache.PipelineCache1_1 = &core1_1.VulkanPipelineCache{
					Driver:              deviceDriver,
					Device:              deviceHandle,
					PipelineCacheHandle: handle,
				}
			}

			return pipelineCache
		}).(*core1_0.VulkanPipelineCache)
}

func CreatePipelineLayout(device iface.Device, handle driver.VkPipelineLayout) *core1_0.VulkanPipelineLayout {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipelineLayout := &core1_0.VulkanPipelineLayout{
				Driver:               deviceDriver,
				Device:               deviceHandle,
				PipelineLayoutHandle: handle,
				MaximumAPIVersion:    version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipelineLayout.PipelineLayout1_1 = &core1_1.VulkanPipelineLayout{
					Driver:               deviceDriver,
					Device:               deviceHandle,
					PipelineLayoutHandle: handle,
				}
			}

			return pipelineLayout
		}).(*core1_0.VulkanPipelineLayout)
}

func CreateQueryPool(device iface.Device, handle driver.VkQueryPool) *core1_0.VulkanQueryPool {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipeline := &core1_0.VulkanQueryPool{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				QueryPoolHandle:   handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipeline.QueryPool1_1 = &core1_1.VulkanQueryPool{
					Driver:          deviceDriver,
					Device:          deviceHandle,
					QueryPoolHandle: handle,
				}
			}

			return pipeline
		}).(*core1_0.VulkanQueryPool)
}

func CreateQueue(device iface.Device, handle driver.VkQueue) *core1_0.VulkanQueue {
	deviceDriver := device.Driver()
	version := device.APIVersion()
	queue := deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			queue := &core1_0.VulkanQueue{
				DeviceDriver:      deviceDriver,
				QueueHandle:       handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				queue.Queue1_1 = &core1_1.VulkanQueue{
					DeviceDriver: deviceDriver,
					QueueHandle:  handle,
				}
			}

			return queue
		}).(*core1_0.VulkanQueue)
	deviceDriver.ObjectStore().SetParent(driver.VulkanHandle(device.Handle()), driver.VulkanHandle(handle))
	return queue
}

func CreateRenderPass(device iface.Device, handle driver.VkRenderPass) *core1_0.VulkanRenderPass {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			renderPass := &core1_0.VulkanRenderPass{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				RenderPassHandle:  handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				renderPass.RenderPass1_1 = &core1_1.VulkanRenderPass{
					Driver:           deviceDriver,
					Device:           deviceHandle,
					RenderPassHandle: handle,
				}
			}

			return renderPass
		}).(*core1_0.VulkanRenderPass)
}

func CreateSampler(device iface.Device, handle driver.VkSampler) *core1_0.VulkanSampler {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipeline := &core1_0.VulkanSampler{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				SamplerHandle:     handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipeline.Sampler1_1 = &core1_1.VulkanSampler{
					Driver:        deviceDriver,
					Device:        deviceHandle,
					SamplerHandle: handle,
				}
			}

			return pipeline
		}).(*core1_0.VulkanSampler)
}

func CreateSemaphore(device iface.Device, handle driver.VkSemaphore) *core1_0.VulkanSemaphore {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipeline := &core1_0.VulkanSemaphore{
				Driver:            deviceDriver,
				Device:            deviceHandle,
				SemaphoreHandle:   handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipeline.Semaphore1_1 = &core1_1.VulkanSemaphore{
					Driver:          deviceDriver,
					Device:          deviceHandle,
					SemaphoreHandle: handle,
				}
			}

			return pipeline
		}).(*core1_0.VulkanSemaphore)
}

func CreateShaderModule(device iface.Device, handle driver.VkShaderModule) *core1_0.VulkanShaderModule {
	deviceDriver := device.Driver()
	deviceHandle := device.Handle()
	version := device.APIVersion()
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipeline := &core1_0.VulkanShaderModule{
				Driver:             deviceDriver,
				Device:             deviceHandle,
				ShaderModuleHandle: handle,
				MaximumAPIVersion:  version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipeline.ShaderModule1_1 = &core1_1.VulkanShaderModule{
					Driver:             deviceDriver,
					Device:             deviceHandle,
					ShaderModuleHandle: handle,
				}
			}

			return pipeline
		}).(*core1_0.VulkanShaderModule)
}
