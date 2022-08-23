package extensions

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
	_ "unsafe"
)

//go:linkname CreateInstanceObject github.com/vkngwrapper/core/v2/core1_0.createInstanceObject
func CreateInstanceObject(instanceDriver driver.Driver, handle driver.VkInstance, version common.APIVersion) *core1_0.VulkanInstance

//go:linkname CreatePhysicalDeviceObject github.com/vkngwrapper/core/v2/core1_0.createPhysicalDeviceObject
func CreatePhysicalDeviceObject(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, instanceVersion, deviceVersion common.APIVersion) *core1_0.VulkanPhysicalDevice

//go:linkname CreateDeviceObject github.com/vkngwrapper/core/v2/core1_0.createDeviceObject
func CreateDeviceObject(deviceDriver driver.Driver, handle driver.VkDevice, version common.APIVersion) *core1_0.VulkanDevice

//go:linkname CreateBufferObject github.com/vkngwrapper/core/v2/core1_0.createBufferObject
func CreateBufferObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBuffer, version common.APIVersion) *core1_0.VulkanBuffer

//go:linkname CreateBufferViewObject github.com/vkngwrapper/core/v2/core1_0.createBufferViewObject
func CreateBufferViewObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkBufferView, version common.APIVersion) *core1_0.VulkanBufferView

//go:linkname CreateCommandBufferObject github.com/vkngwrapper/core/v2/core1_0.createCommandBufferObject
func CreateCommandBufferObject(coreDriver driver.Driver, commandPool driver.VkCommandPool, device driver.VkDevice, handle driver.VkCommandBuffer, version common.APIVersion) *core1_0.VulkanCommandBuffer

//go:linkname CreateCommandPoolObject github.com/vkngwrapper/core/v2/core1_0.createCommandPoolObject
func CreateCommandPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkCommandPool, version common.APIVersion) *core1_0.VulkanCommandPool

//go:linkname CreateDescriptorPoolObject github.com/vkngwrapper/core/v2/core1_0.createDescriptorPoolObject
func CreateDescriptorPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorPool, version common.APIVersion) *core1_0.VulkanDescriptorPool

//go:linkname CreateDescriptorSetObject github.com/vkngwrapper/core/v2/core1_0.createDescriptorSetObject
func CreateDescriptorSetObject(coreDriver driver.Driver, device driver.VkDevice, descriptorPool driver.VkDescriptorPool, handle driver.VkDescriptorSet, version common.APIVersion) *core1_0.VulkanDescriptorSet

//go:linkname CreateDescriptorSetLayoutObject github.com/vkngwrapper/core/v2/core1_0.createDescriptorSetLayoutObject
func CreateDescriptorSetLayoutObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorSetLayout, version common.APIVersion) *core1_0.VulkanDescriptorSetLayout

//go:linkname CreateDescriptorUpdateTemplateObject github.com/vkngwrapper/core/v2/core1_0.createDescriptorUpdateTemplateObject
func CreateDescriptorUpdateTemplateObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDescriptorUpdateTemplate, version common.APIVersion) core1_1.DescriptorUpdateTemplate

//go:linkname CreateDeviceMemoryObject github.com/vkngwrapper/core/v2/core1_0.createDeviceMemoryObject
func CreateDeviceMemoryObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkDeviceMemory, version common.APIVersion, size int) *core1_0.VulkanDeviceMemory

//go:linkname CreateEventObject github.com/vkngwrapper/core/v2/core1_0.createEventObject
func CreateEventObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkEvent, version common.APIVersion) *core1_0.VulkanEvent

//go:linkname CreateFenceObject github.com/vkngwrapper/core/v2/core1_0.createFenceObject
func CreateFenceObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFence, version common.APIVersion) *core1_0.VulkanFence

//go:linkname CreateFramebufferObject github.com/vkngwrapper/core/v2/core1_0.createFramebufferObject
func CreateFramebufferObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkFramebuffer, version common.APIVersion) *core1_0.VulkanFramebuffer

//go:linkname CreateImageObject github.com/vkngwrapper/core/v2/core1_0.createImageObject
func CreateImageObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImage, version common.APIVersion) *core1_0.VulkanImage

//go:linkname CreateImageViewObject github.com/vkngwrapper/core/v2/core1_0.createImageViewObject
func CreateImageViewObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkImageView, version common.APIVersion) *core1_0.VulkanImageView

//go:linkname CreatePipelineObject github.com/vkngwrapper/core/v2/core1_0.createPipelineObject
func CreatePipelineObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipeline, version common.APIVersion) *core1_0.VulkanPipeline

//go:linkname CreatePipelineCacheObject github.com/vkngwrapper/core/v2/core1_0.createPipelineCacheObject
func CreatePipelineCacheObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineCache, version common.APIVersion) *core1_0.VulkanPipelineCache

//go:linkname CreatePipelineLayoutObject github.com/vkngwrapper/core/v2/core1_0.createPipelineLayoutObject
func CreatePipelineLayoutObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkPipelineLayout, version common.APIVersion) *core1_0.VulkanPipelineLayout

//go:linkname CreateQueryPoolObject github.com/vkngwrapper/core/v2/core1_0.createQueryPoolObject
func CreateQueryPoolObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueryPool, version common.APIVersion) *core1_0.VulkanQueryPool

//go:linkname CreateQueueObject github.com/vkngwrapper/core/v2/core1_0.createQueueObject
func CreateQueueObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkQueue, version common.APIVersion) *core1_0.VulkanQueue

//go:linkname CreateRenderPassObject github.com/vkngwrapper/core/v2/core1_0.createRenderPassObject
func CreateRenderPassObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkRenderPass, version common.APIVersion) *core1_0.VulkanRenderPass

//go:linkname CreateSamplerObject github.com/vkngwrapper/core/v2/core1_0.createSamplerObject
func CreateSamplerObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSampler, version common.APIVersion) *core1_0.VulkanSampler

//go:linkname CreateSamplerYcbcrConversionObject github.com/vkngwrapper/core/v2/core1_0.createSamplerYcbcrConversionObject
func CreateSamplerYcbcrConversionObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSamplerYcbcrConversion, version common.APIVersion) core1_1.SamplerYcbcrConversion

//go:linkname CreateSemaphoreObject github.com/vkngwrapper/core/v2/core1_0.createSemaphoreObject
func CreateSemaphoreObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkSemaphore, version common.APIVersion) *core1_0.VulkanSemaphore

//go:linkname CreateShaderModuleObject github.com/vkngwrapper/core/v2/core1_0.createShaderModuleObject
func CreateShaderModuleObject(coreDriver driver.Driver, device driver.VkDevice, handle driver.VkShaderModule, version common.APIVersion) *core1_0.VulkanShaderModule
