package internal1_0

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

func CreateInstanceObject(instanceDriver driver.Driver, handle driver.VkInstance, version common.APIVersion) *VulkanInstance {
	return instanceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			instance := &VulkanInstance{
				InstanceDriver:           instanceDriver,
				InstanceHandle:           handle,
				MaximumVersion:           version,
				ActiveInstanceExtensions: make(map[string]struct{}),
			}

			return instance
		}).(*VulkanInstance)
}

func CreatePhysicalDeviceObject(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, instanceVersion, deviceVersion common.APIVersion) *VulkanPhysicalDevice {
	physicalDevice := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			device := &VulkanPhysicalDevice{
				InstanceDriver:       coreDriver,
				PhysicalDeviceHandle: handle,
				InstanceVersion:      instanceVersion,
				MaximumDeviceVersion: deviceVersion,
			}

			return device
		}).(*VulkanPhysicalDevice)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(instance), driver.VulkanHandle(handle))
	return physicalDevice
}

func createPhysicalDeviceCore1_0(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, instanceVersion, deviceVersion common.APIVersion) core1_0.PhysicalDevice {
	physicalDevice := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			device := &VulkanPhysicalDevice{
				InstanceDriver:       coreDriver,
				PhysicalDeviceHandle: handle,
				InstanceVersion:      instanceVersion,
				MaximumDeviceVersion: deviceVersion,
			}

			return device
		}).(*VulkanPhysicalDevice)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(instance), driver.VulkanHandle(handle))
	return physicalDevice
}

func CreateDeviceObject(deviceDriver driver.Driver, handle driver.VkDevice, version common.APIVersion) *VulkanDevice {
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			device := &VulkanDevice{
				DeviceDriver:           deviceDriver,
				DeviceHandle:           handle,
				MaximumAPIVersion:      version,
				ActiveDeviceExtensions: make(map[string]struct{}),
			}

			return device
		}).(*VulkanDevice)
}

func CreateBufferObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBuffer, version common.APIVersion) *VulkanBuffer {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			buffer := &VulkanBuffer{
				DeviceDriver:      coreDriver,
				Device:            device,
				BufferHandle:      handle,
				MaximumAPIVersion: version,
			}

			return buffer
		}).(*VulkanBuffer)
}

func CreateBufferViewObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBufferView, version common.APIVersion) *VulkanBufferView {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			bufferView := &VulkanBufferView{
				DeviceDriver:      coreDriver,
				Device:            device,
				BufferViewHandle:  handle,
				MaximumAPIVersion: version,
			}

			return bufferView
		}).(*VulkanBufferView)
}

func CreateCommandBufferObject(coreDriver driver.Driver, commandPool driver.VkCommandPool, device driver.VkDevice, handle driver.VkCommandBuffer, version common.APIVersion) *VulkanCommandBuffer {
	commandBuffer := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			var commandCount int
			var drawCount int
			var dispatchCount int

			commandBuffer := &VulkanCommandBuffer{
				DeviceDriver:        coreDriver,
				Device:              device,
				CommandBufferHandle: handle,
				CommandPool:         commandPool,
				MaximumAPIVersion:   version,

				CommandCount:  &commandCount,
				DrawCallCount: &drawCount,
				DispatchCount: &dispatchCount,
			}

			return commandBuffer
		}).(*VulkanCommandBuffer)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(commandPool), driver.VulkanHandle(handle))
	return commandBuffer
}

func CreateCommandPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkCommandPool, version common.APIVersion) *VulkanCommandPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			commandPool := &VulkanCommandPool{
				DeviceDriver:      coreDriver,
				Device:            device,
				CommandPoolHandle: handle,
				MaximumAPIVersion: version,
			}

			return commandPool
		}).(*VulkanCommandPool)
}

func CreateDescriptorPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorPool, version common.APIVersion) *VulkanDescriptorPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			descriptorPool := &VulkanDescriptorPool{
				DeviceDriver:         coreDriver,
				Device:               device,
				DescriptorPoolHandle: handle,
				MaximumAPIVersion:    version,
			}

			return descriptorPool
		}).(*VulkanDescriptorPool)
}

func CreateDescriptorSetObject(coreDriver driver.Driver, device driver.VkDevice, descriptorPool driver.VkDescriptorPool, handle driver.VkDescriptorSet, version common.APIVersion) *VulkanDescriptorSet {
	descriptorSet := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			descriptorSet := &VulkanDescriptorSet{
				DeviceDriver:        coreDriver,
				Device:              device,
				DescriptorSetHandle: handle,
				MaximumAPIVersion:   version,
				DescriptorPool:      descriptorPool,
			}

			return descriptorSet
		}).(*VulkanDescriptorSet)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(descriptorPool), driver.VulkanHandle(handle))
	return descriptorSet
}

func CreateDescriptorSetLayoutObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorSetLayout, version common.APIVersion) *VulkanDescriptorSetLayout {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			descriptorSetLayout := &VulkanDescriptorSetLayout{
				DeviceDriver:              coreDriver,
				Device:                    device,
				DescriptorSetLayoutHandle: handle,
				MaximumAPIVersion:         version,
			}

			return descriptorSetLayout
		}).(*VulkanDescriptorSetLayout)
}

func CreateDeviceMemoryObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDeviceMemory, version common.APIVersion, size int) *VulkanDeviceMemory {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			deviceMemory := &VulkanDeviceMemory{
				DeviceDriver:       coreDriver,
				Device:             device,
				DeviceMemoryHandle: handle,
				MaximumAPIVersion:  version,
				Size:               size,
			}

			return deviceMemory
		}).(*VulkanDeviceMemory)
}

func CreateEventObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkEvent, version common.APIVersion) *VulkanEvent {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			event := &VulkanEvent{
				DeviceDriver:      coreDriver,
				Device:            device,
				EventHandle:       handle,
				MaximumAPIVersion: version,
			}

			return event
		}).(*VulkanEvent)
}

func CreateFenceObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFence, version common.APIVersion) *VulkanFence {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			fence := &VulkanFence{
				DeviceDriver:      coreDriver,
				Device:            device,
				FenceHandle:       handle,
				MaximumAPIVersion: version,
			}

			return fence
		}).(*VulkanFence)
}

func CreateFramebufferObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFramebuffer, version common.APIVersion) *VulkanFramebuffer {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			framebuffer := &VulkanFramebuffer{
				DeviceDriver:      coreDriver,
				Device:            device,
				FramebufferHandle: handle,
				MaximumAPIVersion: version,
			}

			return framebuffer
		}).(*VulkanFramebuffer)
}

func CreateImageObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImage, version common.APIVersion) *VulkanImage {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			image := &VulkanImage{
				DeviceDriver:      coreDriver,
				Device:            device,
				ImageHandle:       handle,
				MaximumAPIVersion: version,
			}

			return image
		}).(*VulkanImage)
}

func CreateImageViewObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImageView, version common.APIVersion) *VulkanImageView {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			imageView := &VulkanImageView{
				DeviceDriver:      coreDriver,
				Device:            device,
				ImageViewHandle:   handle,
				MaximumAPIVersion: version,
			}

			return imageView
		}).(*VulkanImageView)
}

func CreatePipelineObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipeline, version common.APIVersion) *VulkanPipeline {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &VulkanPipeline{
				DeviceDriver:      coreDriver,
				Device:            device,
				PipelineHandle:    handle,
				MaximumAPIVersion: version,
			}

			return pipeline
		}).(*VulkanPipeline)
}

func CreatePipelineCacheObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineCache, version common.APIVersion) *VulkanPipelineCache {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipelineCache := &VulkanPipelineCache{
				DeviceDriver:        coreDriver,
				Device:              device,
				PipelineCacheHandle: handle,
				MaximumAPIVersion:   version,
			}

			return pipelineCache
		}).(*VulkanPipelineCache)
}

func CreatePipelineLayoutObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineLayout, version common.APIVersion) *VulkanPipelineLayout {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipelineLayout := &VulkanPipelineLayout{
				DeviceDriver:         coreDriver,
				Device:               device,
				PipelineLayoutHandle: handle,
				MaximumAPIVersion:    version,
			}

			return pipelineLayout
		}).(*VulkanPipelineLayout)
}

func CreateQueryPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueryPool, version common.APIVersion) *VulkanQueryPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &VulkanQueryPool{
				DeviceDriver:      coreDriver,
				Device:            device,
				QueryPoolHandle:   handle,
				MaximumAPIVersion: version,
			}

			return pipeline
		}).(*VulkanQueryPool)
}

func CreateQueueObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) *VulkanQueue {
	queue := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			queue := &VulkanQueue{
				DeviceDriver:      coreDriver,
				QueueHandle:       handle,
				MaximumAPIVersion: version,
			}

			return queue
		}).(*VulkanQueue)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(device), driver.VulkanHandle(handle))
	return queue
}

func CreateRenderPassObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkRenderPass, version common.APIVersion) *VulkanRenderPass {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			renderPass := &VulkanRenderPass{
				DeviceDriver:      coreDriver,
				Device:            device,
				RenderPassHandle:  handle,
				MaximumAPIVersion: version,
			}

			return renderPass
		}).(*VulkanRenderPass)
}

func CreateSamplerObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSampler, version common.APIVersion) *VulkanSampler {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &VulkanSampler{
				DeviceDriver:      coreDriver,
				Device:            device,
				SamplerHandle:     handle,
				MaximumAPIVersion: version,
			}

			return pipeline
		}).(*VulkanSampler)
}

func CreateSemaphoreObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSemaphore, version common.APIVersion) *VulkanSemaphore {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &VulkanSemaphore{
				DeviceDriver:      coreDriver,
				Device:            device,
				SemaphoreHandle:   handle,
				MaximumAPIVersion: version,
			}

			return pipeline
		}).(*VulkanSemaphore)
}

func CreateShaderModuleObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkShaderModule, version common.APIVersion) *VulkanShaderModule {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &VulkanShaderModule{
				DeviceDriver:       coreDriver,
				Device:             device,
				ShaderModuleHandle: handle,
				MaximumAPIVersion:  version,
			}

			return pipeline
		}).(*VulkanShaderModule)
}
