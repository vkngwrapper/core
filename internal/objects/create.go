package objects

import (
	"github.com/CannibalVox/VKng/core/common"
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

func CreatePhysicalDevice(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, version common.APIVersion) *core1_0.VulkanPhysicalDevice {
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
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(instance), driver.VulkanHandle(handle))
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

func CreateBuffer(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBuffer, version common.APIVersion) *core1_0.VulkanBuffer {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			buffer := &core1_0.VulkanBuffer{
				Driver:            coreDriver,
				Device:            device,
				BufferHandle:      handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				buffer.Buffer1_1 = &core1_1.VulkanBuffer{
					Driver:       coreDriver,
					Device:       device,
					BufferHandle: handle,
				}
			}

			return buffer
		}).(*core1_0.VulkanBuffer)
}

func CreateBufferView(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBufferView, version common.APIVersion) *core1_0.VulkanBufferView {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			bufferView := &core1_0.VulkanBufferView{
				Driver:            coreDriver,
				Device:            device,
				BufferViewHandle:  handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				bufferView.BufferView1_1 = &core1_1.VulkanBufferView{
					Driver:           coreDriver,
					Device:           device,
					BufferViewHandle: handle,
				}
			}

			return bufferView
		}).(*core1_0.VulkanBufferView)
}

func CreateCommandBuffer(coreDriver driver.Driver, commandPool driver.VkCommandPool, device driver.VkDevice, handle driver.VkCommandBuffer, version common.APIVersion) *core1_0.VulkanCommandBuffer {
	commandBuffer := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			commandBuffer := &core1_0.VulkanCommandBuffer{
				DeviceDriver:        coreDriver,
				Device:              device,
				CommandBufferHandle: handle,
				CommandPool:         commandPool,
				MaximumAPIVersion:   version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				commandBuffer.CommandBuffer1_1 = &core1_1.VulkanCommandBuffer{
					DeviceDriver:        coreDriver,
					Device:              device,
					CommandPool:         commandPool,
					CommandBufferHandle: handle,
				}
			}

			return commandBuffer
		}).(*core1_0.VulkanCommandBuffer)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(commandPool), driver.VulkanHandle(handle))
	return commandBuffer
}

func CreateCommandPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkCommandPool, version common.APIVersion) *core1_0.VulkanCommandPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			commandPool := &core1_0.VulkanCommandPool{
				DeviceDriver:      coreDriver,
				DeviceHandle:      device,
				CommandPoolHandle: handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				commandPool.CommandPool1_1 = &core1_1.VulkanCommandPool{
					DeviceDriver:      coreDriver,
					DeviceHandle:      device,
					CommandPoolHandle: handle,
				}
			}

			return commandPool
		}).(*core1_0.VulkanCommandPool)
}

func CreateDescriptorPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorPool, version common.APIVersion) *core1_0.VulkanDescriptorPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			descriptorPool := &core1_0.VulkanDescriptorPool{
				DeviceDriver:         coreDriver,
				Device:               device,
				DescriptorPoolHandle: handle,
				MaximumAPIVersion:    version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				descriptorPool.DescriptorPool1_1 = &core1_1.VulkanDescriptorPool{
					DeviceDriver:         coreDriver,
					Device:               device,
					DescriptorPoolHandle: handle,
				}
			}

			return descriptorPool
		}).(*core1_0.VulkanDescriptorPool)
}

func CreateDescriptorSet(coreDriver driver.Driver, device driver.VkDevice, descriptorPool driver.VkDescriptorPool, handle driver.VkDescriptorSet, version common.APIVersion) *core1_0.VulkanDescriptorSet {
	descriptorSet := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			descriptorSet := &core1_0.VulkanDescriptorSet{
				DeviceDriver:        coreDriver,
				Device:              device,
				DescriptorSetHandle: handle,
				MaximumAPIVersion:   version,
				DescriptorPool:      descriptorPool,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				descriptorSet.DescriptorSet1_1 = &core1_1.VulkanDescriptorSet{
					DeviceDriver:        coreDriver,
					Device:              device,
					DescriptorSetHandle: handle,
					DescriptorPool:      descriptorPool,
				}
			}

			return descriptorSet
		}).(*core1_0.VulkanDescriptorSet)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(descriptorPool), driver.VulkanHandle(handle))
	return descriptorSet
}

func CreateDescriptorSetLayout(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorSetLayout, version common.APIVersion) *core1_0.VulkanDescriptorSetLayout {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			descriptorSetLayout := &core1_0.VulkanDescriptorSetLayout{
				Driver:                    coreDriver,
				Device:                    device,
				DescriptorSetLayoutHandle: handle,
				MaximumAPIVersion:         version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				descriptorSetLayout.DescriptorSetLayout1_1 = &core1_1.VulkanDescriptorSetLayout{
					Driver:                    coreDriver,
					Device:                    device,
					DescriptorSetLayoutHandle: handle,
				}
			}

			return descriptorSetLayout
		}).(*core1_0.VulkanDescriptorSetLayout)
}

func CreateDeviceMemory(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDeviceMemory, version common.APIVersion, size int) *core1_0.VulkanDeviceMemory {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			deviceMemory := &core1_0.VulkanDeviceMemory{
				DeviceDriver:       coreDriver,
				Device:             device,
				DeviceMemoryHandle: handle,
				MaximumAPIVersion:  version,
				Size:               size,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				deviceMemory.DeviceMemory1_1 = &core1_1.VulkanDeviceMemory{
					DeviceDriver:       coreDriver,
					Device:             device,
					DeviceMemoryHandle: handle,
					Size:               size,
				}
			}

			return deviceMemory
		}).(*core1_0.VulkanDeviceMemory)
}

func CreateEvent(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkEvent, version common.APIVersion) *core1_0.VulkanEvent {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			event := &core1_0.VulkanEvent{
				Driver:            coreDriver,
				Device:            device,
				EventHandle:       handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				event.Event1_1 = &core1_1.VulkanEvent{
					Driver:      coreDriver,
					Device:      device,
					EventHandle: handle,
				}
			}

			return event
		}).(*core1_0.VulkanEvent)
}

func CreateFence(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFence, version common.APIVersion) *core1_0.VulkanFence {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			fence := &core1_0.VulkanFence{
				Driver:            coreDriver,
				Device:            device,
				FenceHandle:       handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				fence.Fence1_1 = &core1_1.VulkanFence{
					Driver:      coreDriver,
					Device:      device,
					FenceHandle: handle,
				}
			}

			return fence
		}).(*core1_0.VulkanFence)
}

func CreateFramebuffer(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFramebuffer, version common.APIVersion) *core1_0.VulkanFramebuffer {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			framebuffer := &core1_0.VulkanFramebuffer{
				Driver:            coreDriver,
				Device:            device,
				FramebufferHandle: handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				framebuffer.Framebuffer1_1 = &core1_1.VulkanFramebuffer{
					Driver:            coreDriver,
					Device:            device,
					FramebufferHandle: handle,
				}
			}

			return framebuffer
		}).(*core1_0.VulkanFramebuffer)
}

func CreateImage(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImage, version common.APIVersion) *core1_0.VulkanImage {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			image := &core1_0.VulkanImage{
				Driver:            coreDriver,
				Device:            device,
				ImageHandle:       handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				image.Image1_1 = &core1_1.VulkanImage{
					Driver:      coreDriver,
					Device:      device,
					ImageHandle: handle,
				}
			}

			return image
		}).(*core1_0.VulkanImage)
}

func CreateImageView(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImageView, version common.APIVersion) *core1_0.VulkanImageView {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			imageView := &core1_0.VulkanImageView{
				Driver:            coreDriver,
				Device:            device,
				ImageViewHandle:   handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				imageView.ImageView1_1 = &core1_1.VulkanImageView{
					Driver:          coreDriver,
					Device:          device,
					ImageViewHandle: handle,
				}
			}

			return imageView
		}).(*core1_0.VulkanImageView)
}

func CreatePipeline(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipeline, version common.APIVersion) *core1_0.VulkanPipeline {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipeline := &core1_0.VulkanPipeline{
				Driver:            coreDriver,
				Device:            device,
				PipelineHandle:    handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipeline.Pipeline1_1 = &core1_1.VulkanPipeline{
					Driver:         coreDriver,
					Device:         device,
					PipelineHandle: handle,
				}
			}

			return pipeline
		}).(*core1_0.VulkanPipeline)
}

func CreatePipelineCache(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineCache, version common.APIVersion) *core1_0.VulkanPipelineCache {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipelineCache := &core1_0.VulkanPipelineCache{
				Driver:              coreDriver,
				Device:              device,
				PipelineCacheHandle: handle,
				MaximumAPIVersion:   version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipelineCache.PipelineCache1_1 = &core1_1.VulkanPipelineCache{
					Driver:              coreDriver,
					Device:              device,
					PipelineCacheHandle: handle,
				}
			}

			return pipelineCache
		}).(*core1_0.VulkanPipelineCache)
}

func CreatePipelineLayout(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineLayout, version common.APIVersion) *core1_0.VulkanPipelineLayout {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipelineLayout := &core1_0.VulkanPipelineLayout{
				Driver:               coreDriver,
				Device:               device,
				PipelineLayoutHandle: handle,
				MaximumAPIVersion:    version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipelineLayout.PipelineLayout1_1 = &core1_1.VulkanPipelineLayout{
					Driver:               coreDriver,
					Device:               device,
					PipelineLayoutHandle: handle,
				}
			}

			return pipelineLayout
		}).(*core1_0.VulkanPipelineLayout)
}

func CreateQueryPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueryPool, version common.APIVersion) *core1_0.VulkanQueryPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipeline := &core1_0.VulkanQueryPool{
				Driver:            coreDriver,
				Device:            device,
				QueryPoolHandle:   handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipeline.QueryPool1_1 = &core1_1.VulkanQueryPool{
					Driver:          coreDriver,
					Device:          device,
					QueryPoolHandle: handle,
				}
			}

			return pipeline
		}).(*core1_0.VulkanQueryPool)
}

func CreateQueue(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) *core1_0.VulkanQueue {
	queue := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			queue := &core1_0.VulkanQueue{
				DeviceDriver:      coreDriver,
				QueueHandle:       handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				queue.Queue1_1 = &core1_1.VulkanQueue{
					DeviceDriver: coreDriver,
					QueueHandle:  handle,
				}
			}

			return queue
		}).(*core1_0.VulkanQueue)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(device), driver.VulkanHandle(handle))
	return queue
}

func CreateRenderPass(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkRenderPass, version common.APIVersion) *core1_0.VulkanRenderPass {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			renderPass := &core1_0.VulkanRenderPass{
				Driver:            coreDriver,
				Device:            device,
				RenderPassHandle:  handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				renderPass.RenderPass1_1 = &core1_1.VulkanRenderPass{
					Driver:           coreDriver,
					Device:           device,
					RenderPassHandle: handle,
				}
			}

			return renderPass
		}).(*core1_0.VulkanRenderPass)
}

func CreateSampler(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSampler, version common.APIVersion) *core1_0.VulkanSampler {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipeline := &core1_0.VulkanSampler{
				Driver:            coreDriver,
				Device:            device,
				SamplerHandle:     handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipeline.Sampler1_1 = &core1_1.VulkanSampler{
					Driver:        coreDriver,
					Device:        device,
					SamplerHandle: handle,
				}
			}

			return pipeline
		}).(*core1_0.VulkanSampler)
}

func CreateSemaphore(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSemaphore, version common.APIVersion) *core1_0.VulkanSemaphore {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipeline := &core1_0.VulkanSemaphore{
				Driver:            coreDriver,
				Device:            device,
				SemaphoreHandle:   handle,
				MaximumAPIVersion: version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipeline.Semaphore1_1 = &core1_1.VulkanSemaphore{
					Driver:          coreDriver,
					Device:          device,
					SemaphoreHandle: handle,
				}
			}

			return pipeline
		}).(*core1_0.VulkanSemaphore)
}

func CreateShaderModule(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkShaderModule, version common.APIVersion) *core1_0.VulkanShaderModule {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle),
		func() interface{} {
			pipeline := &core1_0.VulkanShaderModule{
				Driver:             coreDriver,
				Device:             device,
				ShaderModuleHandle: handle,
				MaximumAPIVersion:  version,
			}

			if version.IsAtLeast(common.Vulkan1_1) {
				pipeline.ShaderModule1_1 = &core1_1.VulkanShaderModule{
					Driver:             coreDriver,
					Device:             device,
					ShaderModuleHandle: handle,
				}
			}

			return pipeline
		}).(*core1_0.VulkanShaderModule)
}
