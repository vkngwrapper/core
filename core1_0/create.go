package core1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

func createInstanceObject(instanceDriver driver.Driver, handle driver.VkInstance, version common.APIVersion) *VulkanInstance {
	return instanceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			instance := &VulkanInstance{
				instanceDriver:           instanceDriver,
				instanceHandle:           handle,
				maximumVersion:           version,
				ActiveInstanceExtensions: make(map[string]struct{}),
			}

			return instance
		}).(*VulkanInstance)
}

func createPhysicalDeviceObject(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, instanceVersion, deviceVersion common.APIVersion) *VulkanPhysicalDevice {
	physicalDevice := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			device := &VulkanPhysicalDevice{
				instanceDriver:       coreDriver,
				physicalDeviceHandle: handle,
				instanceVersion:      instanceVersion,
				maximumDeviceVersion: deviceVersion,
			}

			return device
		}).(*VulkanPhysicalDevice)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(instance), driver.VulkanHandle(handle))
	return physicalDevice
}

func createPhysicalDeviceCore1_0(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, instanceVersion, deviceVersion common.APIVersion) PhysicalDevice {
	physicalDevice := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			device := &VulkanPhysicalDevice{
				instanceDriver:       coreDriver,
				physicalDeviceHandle: handle,
				instanceVersion:      instanceVersion,
				maximumDeviceVersion: deviceVersion,
			}

			return device
		}).(*VulkanPhysicalDevice)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(instance), driver.VulkanHandle(handle))
	return physicalDevice
}

func createDeviceObject(deviceDriver driver.Driver, handle driver.VkDevice, version common.APIVersion) *VulkanDevice {
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			device := &VulkanDevice{
				deviceDriver:           deviceDriver,
				deviceHandle:           handle,
				maximumAPIVersion:      version,
				activeDeviceExtensions: make(map[string]struct{}),
			}

			return device
		}).(*VulkanDevice)
}

func createBufferObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBuffer, version common.APIVersion) *VulkanBuffer {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			buffer := &VulkanBuffer{
				deviceDriver:      coreDriver,
				device:            device,
				bufferHandle:      handle,
				maximumAPIVersion: version,
			}

			return buffer
		}).(*VulkanBuffer)
}

func createBufferViewObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBufferView, version common.APIVersion) *VulkanBufferView {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			bufferView := &VulkanBufferView{
				deviceDriver:      coreDriver,
				device:            device,
				bufferViewHandle:  handle,
				maximumAPIVersion: version,
			}

			return bufferView
		}).(*VulkanBufferView)
}

func createCommandBufferObject(coreDriver driver.Driver, commandPool driver.VkCommandPool, device driver.VkDevice, handle driver.VkCommandBuffer, version common.APIVersion) *VulkanCommandBuffer {
	commandBuffer := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			commandBuffer := &VulkanCommandBuffer{
				deviceDriver:        coreDriver,
				device:              device,
				commandBufferHandle: handle,
				commandPool:         commandPool,
				maximumAPIVersion:   version,

				commandCounter: &CommandCounter{},
			}

			return commandBuffer
		}).(*VulkanCommandBuffer)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(commandPool), driver.VulkanHandle(handle))
	return commandBuffer
}

func createCommandPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkCommandPool, version common.APIVersion) *VulkanCommandPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			commandPool := &VulkanCommandPool{
				deviceDriver:      coreDriver,
				device:            device,
				commandPoolHandle: handle,
				maximumAPIVersion: version,
			}

			return commandPool
		}).(*VulkanCommandPool)
}

func createDescriptorPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorPool, version common.APIVersion) *VulkanDescriptorPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			descriptorPool := &VulkanDescriptorPool{
				deviceDriver:         coreDriver,
				device:               device,
				descriptorPoolHandle: handle,
				maximumAPIVersion:    version,
			}

			return descriptorPool
		}).(*VulkanDescriptorPool)
}

func createDescriptorSetObject(coreDriver driver.Driver, device driver.VkDevice, descriptorPool driver.VkDescriptorPool, handle driver.VkDescriptorSet, version common.APIVersion) *VulkanDescriptorSet {
	descriptorSet := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			descriptorSet := &VulkanDescriptorSet{
				deviceDriver:        coreDriver,
				device:              device,
				descriptorSetHandle: handle,
				maximumAPIVersion:   version,
				descriptorPool:      descriptorPool,
			}

			return descriptorSet
		}).(*VulkanDescriptorSet)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(descriptorPool), driver.VulkanHandle(handle))
	return descriptorSet
}

func createDescriptorSetLayoutObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorSetLayout, version common.APIVersion) *VulkanDescriptorSetLayout {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			descriptorSetLayout := &VulkanDescriptorSetLayout{
				deviceDriver:              coreDriver,
				device:                    device,
				descriptorSetLayoutHandle: handle,
				maximumAPIVersion:         version,
			}

			return descriptorSetLayout
		}).(*VulkanDescriptorSetLayout)
}

func createDeviceMemoryObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDeviceMemory, version common.APIVersion, size int) *VulkanDeviceMemory {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			deviceMemory := &VulkanDeviceMemory{
				deviceDriver:       coreDriver,
				device:             device,
				deviceMemoryHandle: handle,
				maximumAPIVersion:  version,
				size:               size,
			}

			return deviceMemory
		}).(*VulkanDeviceMemory)
}

func createEventObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkEvent, version common.APIVersion) *VulkanEvent {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			event := &VulkanEvent{
				deviceDriver:      coreDriver,
				device:            device,
				eventHandle:       handle,
				maximumAPIVersion: version,
			}

			return event
		}).(*VulkanEvent)
}

func createFenceObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFence, version common.APIVersion) *VulkanFence {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			fence := &VulkanFence{
				deviceDriver:      coreDriver,
				device:            device,
				fenceHandle:       handle,
				maximumAPIVersion: version,
			}

			return fence
		}).(*VulkanFence)
}

func createFramebufferObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFramebuffer, version common.APIVersion) *VulkanFramebuffer {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			framebuffer := &VulkanFramebuffer{
				deviceDriver:      coreDriver,
				device:            device,
				framebufferHandle: handle,
				maximumAPIVersion: version,
			}

			return framebuffer
		}).(*VulkanFramebuffer)
}

func createImageObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImage, version common.APIVersion) *VulkanImage {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			image := &VulkanImage{
				deviceDriver:      coreDriver,
				device:            device,
				imageHandle:       handle,
				maximumAPIVersion: version,
			}

			return image
		}).(*VulkanImage)
}

func createImageViewObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImageView, version common.APIVersion) *VulkanImageView {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			imageView := &VulkanImageView{
				deviceDriver:      coreDriver,
				device:            device,
				imageViewHandle:   handle,
				maximumAPIVersion: version,
			}

			return imageView
		}).(*VulkanImageView)
}

func createPipelineObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipeline, version common.APIVersion) *VulkanPipeline {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &VulkanPipeline{
				deviceDriver:      coreDriver,
				device:            device,
				pipelineHandle:    handle,
				maximumAPIVersion: version,
			}

			return pipeline
		}).(*VulkanPipeline)
}

func createPipelineCacheObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineCache, version common.APIVersion) *VulkanPipelineCache {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipelineCache := &VulkanPipelineCache{
				deviceDriver:        coreDriver,
				device:              device,
				pipelineCacheHandle: handle,
				maximumAPIVersion:   version,
			}

			return pipelineCache
		}).(*VulkanPipelineCache)
}

func createPipelineLayoutObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineLayout, version common.APIVersion) *VulkanPipelineLayout {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipelineLayout := &VulkanPipelineLayout{
				deviceDriver:         coreDriver,
				device:               device,
				pipelineLayoutHandle: handle,
				maximumAPIVersion:    version,
			}

			return pipelineLayout
		}).(*VulkanPipelineLayout)
}

func createQueryPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueryPool, version common.APIVersion) *VulkanQueryPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &VulkanQueryPool{
				deviceDriver:      coreDriver,
				device:            device,
				queryPoolHandle:   handle,
				maximumAPIVersion: version,
			}

			return pipeline
		}).(*VulkanQueryPool)
}

func createQueueObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) *VulkanQueue {
	queue := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			queue := &VulkanQueue{
				deviceDriver:      coreDriver,
				queueHandle:       handle,
				device:            device,
				maximumAPIVersion: version,
			}

			return queue
		}).(*VulkanQueue)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(device), driver.VulkanHandle(handle))
	return queue
}

func createRenderPassObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkRenderPass, version common.APIVersion) *VulkanRenderPass {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			renderPass := &VulkanRenderPass{
				deviceDriver:      coreDriver,
				device:            device,
				renderPassHandle:  handle,
				maximumAPIVersion: version,
			}

			return renderPass
		}).(*VulkanRenderPass)
}

func createSamplerObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSampler, version common.APIVersion) *VulkanSampler {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &VulkanSampler{
				deviceDriver:      coreDriver,
				device:            device,
				samplerHandle:     handle,
				maximumAPIVersion: version,
			}

			return pipeline
		}).(*VulkanSampler)
}

func createSemaphoreObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSemaphore, version common.APIVersion) *VulkanSemaphore {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &VulkanSemaphore{
				deviceDriver:      coreDriver,
				device:            device,
				semaphoreHandle:   handle,
				maximumAPIVersion: version,
			}

			return pipeline
		}).(*VulkanSemaphore)
}

func createShaderModuleObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkShaderModule, version common.APIVersion) *VulkanShaderModule {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &VulkanShaderModule{
				deviceDriver:       coreDriver,
				device:             device,
				shaderModuleHandle: handle,
				maximumAPIVersion:  version,
			}

			return pipeline
		}).(*VulkanShaderModule)
}
