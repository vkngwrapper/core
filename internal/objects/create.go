package objects

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	internal1_0 "github.com/CannibalVox/VKng/core/internal/core1_0"
	"github.com/CannibalVox/VKng/core/internal/core1_1"
)

func CreateInstance(instanceDriver driver.Driver, handle driver.VkInstance, version common.APIVersion) *internal1_0.VulkanInstance {
	return instanceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			instance := &internal1_0.VulkanInstance{
				InstanceDriver: instanceDriver,
				InstanceHandle: handle,
				MaximumVersion: version,
			}

			return instance
		}).(*internal1_0.VulkanInstance)
}

func CreatePhysicalDevice(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, instanceVersion, deviceVersion common.APIVersion) *internal1_0.VulkanPhysicalDevice {
	physicalDevice := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			device := &internal1_0.VulkanPhysicalDevice{
				InstanceDriver:       coreDriver,
				PhysicalDeviceHandle: handle,
				InstanceVersion:      instanceVersion,
				MaximumDeviceVersion: deviceVersion,
			}

			return device
		}).(*internal1_0.VulkanPhysicalDevice)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(instance), driver.VulkanHandle(handle))
	return physicalDevice
}

func createPhysicalDeviceCore1_0(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, instanceVersion, deviceVersion common.APIVersion) core1_0.PhysicalDevice {
	physicalDevice := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			device := &internal1_0.VulkanPhysicalDevice{
				InstanceDriver:       coreDriver,
				PhysicalDeviceHandle: handle,
				InstanceVersion:      instanceVersion,
				MaximumDeviceVersion: deviceVersion,
			}

			return device
		}).(*internal1_0.VulkanPhysicalDevice)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(instance), driver.VulkanHandle(handle))
	return physicalDevice
}

func CreateDevice(deviceDriver driver.Driver, handle driver.VkDevice, version common.APIVersion) *internal1_0.VulkanDevice {
	return deviceDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			device := &internal1_0.VulkanDevice{
				DeviceDriver:      deviceDriver,
				DeviceHandle:      handle,
				MaximumAPIVersion: version,
			}

			return device
		}).(*internal1_0.VulkanDevice)
}

func CreateBuffer(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBuffer, version common.APIVersion) *internal1_0.VulkanBuffer {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			buffer := &internal1_0.VulkanBuffer{
				DeviceDriver:      coreDriver,
				Device:            device,
				BufferHandle:      handle,
				MaximumAPIVersion: version,
			}

			return buffer
		}).(*internal1_0.VulkanBuffer)
}

func CreateBufferView(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBufferView, version common.APIVersion) *internal1_0.VulkanBufferView {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			bufferView := &internal1_0.VulkanBufferView{
				DeviceDriver:      coreDriver,
				Device:            device,
				BufferViewHandle:  handle,
				MaximumAPIVersion: version,
			}

			return bufferView
		}).(*internal1_0.VulkanBufferView)
}

func CreateCommandBuffer(coreDriver driver.Driver, commandPool driver.VkCommandPool, device driver.VkDevice, handle driver.VkCommandBuffer, version common.APIVersion) *internal1_0.VulkanCommandBuffer {
	commandBuffer := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			var commandCount int
			var drawCount int
			var dispatchCount int

			commandBuffer := &internal1_0.VulkanCommandBuffer{
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
		}).(*internal1_0.VulkanCommandBuffer)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(commandPool), driver.VulkanHandle(handle))
	return commandBuffer
}

func CreateCommandPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkCommandPool, version common.APIVersion) *internal1_0.VulkanCommandPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			commandPool := &internal1_0.VulkanCommandPool{
				DeviceDriver:      coreDriver,
				Device:            device,
				CommandPoolHandle: handle,
				MaximumAPIVersion: version,
			}

			return commandPool
		}).(*internal1_0.VulkanCommandPool)
}

func CreateDescriptorPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorPool, version common.APIVersion) *internal1_0.VulkanDescriptorPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			descriptorPool := &internal1_0.VulkanDescriptorPool{
				DeviceDriver:         coreDriver,
				Device:               device,
				DescriptorPoolHandle: handle,
				MaximumAPIVersion:    version,
			}

			return descriptorPool
		}).(*internal1_0.VulkanDescriptorPool)
}

func CreateDescriptorSet(coreDriver driver.Driver, device driver.VkDevice, descriptorPool driver.VkDescriptorPool, handle driver.VkDescriptorSet, version common.APIVersion) *internal1_0.VulkanDescriptorSet {
	descriptorSet := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			descriptorSet := &internal1_0.VulkanDescriptorSet{
				DeviceDriver:        coreDriver,
				Device:              device,
				DescriptorSetHandle: handle,
				MaximumAPIVersion:   version,
				DescriptorPool:      descriptorPool,
			}

			return descriptorSet
		}).(*internal1_0.VulkanDescriptorSet)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(descriptorPool), driver.VulkanHandle(handle))
	return descriptorSet
}

func CreateDescriptorSetLayout(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorSetLayout, version common.APIVersion) *internal1_0.VulkanDescriptorSetLayout {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			descriptorSetLayout := &internal1_0.VulkanDescriptorSetLayout{
				DeviceDriver:              coreDriver,
				Device:                    device,
				DescriptorSetLayoutHandle: handle,
				MaximumAPIVersion:         version,
			}

			return descriptorSetLayout
		}).(*internal1_0.VulkanDescriptorSetLayout)
}

func CreateDescriptorUpdateTemplate(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorUpdateTemplate, version common.APIVersion) *internal1_1.VulkanDescriptorUpdateTemplate {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_1,
		func() any {
			template := &internal1_1.VulkanDescriptorUpdateTemplate{
				DeviceDriver:             coreDriver,
				Device:                   device,
				DescriptorTemplateHandle: handle,
				MaximumAPIVersion:        version,
			}

			return template
		}).(*internal1_1.VulkanDescriptorUpdateTemplate)
}

func CreateDeviceMemory(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDeviceMemory, version common.APIVersion, size int) *internal1_0.VulkanDeviceMemory {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			deviceMemory := &internal1_0.VulkanDeviceMemory{
				DeviceDriver:       coreDriver,
				Device:             device,
				DeviceMemoryHandle: handle,
				MaximumAPIVersion:  version,
				Size:               size,
			}

			return deviceMemory
		}).(*internal1_0.VulkanDeviceMemory)
}

func CreateEvent(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkEvent, version common.APIVersion) *internal1_0.VulkanEvent {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			event := &internal1_0.VulkanEvent{
				DeviceDriver:      coreDriver,
				Device:            device,
				EventHandle:       handle,
				MaximumAPIVersion: version,
			}

			return event
		}).(*internal1_0.VulkanEvent)
}

func CreateFence(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFence, version common.APIVersion) *internal1_0.VulkanFence {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			fence := &internal1_0.VulkanFence{
				DeviceDriver:      coreDriver,
				Device:            device,
				FenceHandle:       handle,
				MaximumAPIVersion: version,
			}

			return fence
		}).(*internal1_0.VulkanFence)
}

func CreateFramebuffer(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFramebuffer, version common.APIVersion) *internal1_0.VulkanFramebuffer {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			framebuffer := &internal1_0.VulkanFramebuffer{
				DeviceDriver:      coreDriver,
				Device:            device,
				FramebufferHandle: handle,
				MaximumAPIVersion: version,
			}

			return framebuffer
		}).(*internal1_0.VulkanFramebuffer)
}

func CreateImage(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImage, version common.APIVersion) *internal1_0.VulkanImage {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			image := &internal1_0.VulkanImage{
				DeviceDriver:      coreDriver,
				Device:            device,
				ImageHandle:       handle,
				MaximumAPIVersion: version,
			}

			return image
		}).(*internal1_0.VulkanImage)
}

func CreateImageView(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImageView, version common.APIVersion) *internal1_0.VulkanImageView {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			imageView := &internal1_0.VulkanImageView{
				DeviceDriver:      coreDriver,
				Device:            device,
				ImageViewHandle:   handle,
				MaximumAPIVersion: version,
			}

			return imageView
		}).(*internal1_0.VulkanImageView)
}

func CreatePipeline(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipeline, version common.APIVersion) *internal1_0.VulkanPipeline {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &internal1_0.VulkanPipeline{
				DeviceDriver:      coreDriver,
				Device:            device,
				PipelineHandle:    handle,
				MaximumAPIVersion: version,
			}

			return pipeline
		}).(*internal1_0.VulkanPipeline)
}

func CreatePipelineCache(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineCache, version common.APIVersion) *internal1_0.VulkanPipelineCache {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipelineCache := &internal1_0.VulkanPipelineCache{
				DeviceDriver:        coreDriver,
				Device:              device,
				PipelineCacheHandle: handle,
				MaximumAPIVersion:   version,
			}

			return pipelineCache
		}).(*internal1_0.VulkanPipelineCache)
}

func CreatePipelineLayout(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineLayout, version common.APIVersion) *internal1_0.VulkanPipelineLayout {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipelineLayout := &internal1_0.VulkanPipelineLayout{
				DeviceDriver:         coreDriver,
				Device:               device,
				PipelineLayoutHandle: handle,
				MaximumAPIVersion:    version,
			}

			return pipelineLayout
		}).(*internal1_0.VulkanPipelineLayout)
}

func CreateQueryPool(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueryPool, version common.APIVersion) *internal1_0.VulkanQueryPool {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &internal1_0.VulkanQueryPool{
				DeviceDriver:      coreDriver,
				Device:            device,
				QueryPoolHandle:   handle,
				MaximumAPIVersion: version,
			}

			return pipeline
		}).(*internal1_0.VulkanQueryPool)
}

func CreateQueue(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) *internal1_0.VulkanQueue {
	queue := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			queue := &internal1_0.VulkanQueue{
				DeviceDriver:      coreDriver,
				QueueHandle:       handle,
				MaximumAPIVersion: version,
			}

			return queue
		}).(*internal1_0.VulkanQueue)
	coreDriver.ObjectStore().SetParent(driver.VulkanHandle(device), driver.VulkanHandle(handle))
	return queue
}

func CreateRenderPass(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkRenderPass, version common.APIVersion) *internal1_0.VulkanRenderPass {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			renderPass := &internal1_0.VulkanRenderPass{
				DeviceDriver:      coreDriver,
				Device:            device,
				RenderPassHandle:  handle,
				MaximumAPIVersion: version,
			}

			return renderPass
		}).(*internal1_0.VulkanRenderPass)
}

func CreateSampler(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSampler, version common.APIVersion) *internal1_0.VulkanSampler {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &internal1_0.VulkanSampler{
				DeviceDriver:      coreDriver,
				Device:            device,
				SamplerHandle:     handle,
				MaximumAPIVersion: version,
			}

			return pipeline
		}).(*internal1_0.VulkanSampler)
}

func CreateSamplerYcbcrConversion(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSamplerYcbcrConversion, version common.APIVersion) *internal1_1.VulkanSamplerYcbcrConversion {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_1,
		func() any {
			return &internal1_1.VulkanSamplerYcbcrConversion{
				DeviceDriver:      coreDriver,
				Device:            device,
				YcbcrHandle:       handle,
				MaximumAPIVersion: version,
			}
		}).(*internal1_1.VulkanSamplerYcbcrConversion)
}

func CreateSemaphore(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSemaphore, version common.APIVersion) *internal1_0.VulkanSemaphore {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &internal1_0.VulkanSemaphore{
				DeviceDriver:      coreDriver,
				Device:            device,
				SemaphoreHandle:   handle,
				MaximumAPIVersion: version,
			}

			return pipeline
		}).(*internal1_0.VulkanSemaphore)
}

func CreateShaderModule(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkShaderModule, version common.APIVersion) *internal1_0.VulkanShaderModule {
	return coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(handle), driver.Core1_0,
		func() any {
			pipeline := &internal1_0.VulkanShaderModule{
				DeviceDriver:       coreDriver,
				Device:             device,
				ShaderModuleHandle: handle,
				MaximumAPIVersion:  version,
			}

			return pipeline
		}).(*internal1_0.VulkanShaderModule)
}
