package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type ObjectType int32

const (
	UsageUnknown                       ObjectType = C.VK_OBJECT_TYPE_UNKNOWN
	UsageInstance                      ObjectType = C.VK_OBJECT_TYPE_INSTANCE
	UsagePhysicalDevice                ObjectType = C.VK_OBJECT_TYPE_PHYSICAL_DEVICE
	UsageDevice                        ObjectType = C.VK_OBJECT_TYPE_DEVICE
	UsageQueue                         ObjectType = C.VK_OBJECT_TYPE_QUEUE
	UsageSemaphore                     ObjectType = C.VK_OBJECT_TYPE_SEMAPHORE
	UsageCommandBuffer                 ObjectType = C.VK_OBJECT_TYPE_COMMAND_BUFFER
	UsageFence                         ObjectType = C.VK_OBJECT_TYPE_FENCE
	UsageDeviceMemory                  ObjectType = C.VK_OBJECT_TYPE_DEVICE_MEMORY
	UsageBuffer                        ObjectType = C.VK_OBJECT_TYPE_BUFFER
	UsageImage                         ObjectType = C.VK_OBJECT_TYPE_IMAGE
	UsageEvent                         ObjectType = C.VK_OBJECT_TYPE_EVENT
	UsageQueryPool                     ObjectType = C.VK_OBJECT_TYPE_QUERY_POOL
	UsageBufferView                    ObjectType = C.VK_OBJECT_TYPE_BUFFER_VIEW
	UsageImageView                     ObjectType = C.VK_OBJECT_TYPE_IMAGE_VIEW
	UsageShaderModule                  ObjectType = C.VK_OBJECT_TYPE_SHADER_MODULE
	UsagePipelineCache                 ObjectType = C.VK_OBJECT_TYPE_PIPELINE_CACHE
	UsagePipelineLayout                ObjectType = C.VK_OBJECT_TYPE_PIPELINE_LAYOUT
	UsageRenderPass                    ObjectType = C.VK_OBJECT_TYPE_RENDER_PASS
	UsagePipeline                      ObjectType = C.VK_OBJECT_TYPE_PIPELINE
	UsageDescriptorSetLayout           ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT
	UsageSampler                       ObjectType = C.VK_OBJECT_TYPE_SAMPLER
	UsageDescriptorPool                ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_POOL
	UsageDescriptorSet                 ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_SET
	UsageFramebuffer                   ObjectType = C.VK_OBJECT_TYPE_FRAMEBUFFER
	UsageCommandPool                   ObjectType = C.VK_OBJECT_TYPE_COMMAND_POOL
	UsageSamplerYCbCrConversion        ObjectType = C.VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION
	UsageDescriptorUpdateTemplate      ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE
	UsageSurface                       ObjectType = C.VK_OBJECT_TYPE_SURFACE_KHR
	UsageSwapchain                     ObjectType = C.VK_OBJECT_TYPE_SWAPCHAIN_KHR
	UsageDisplay                       ObjectType = C.VK_OBJECT_TYPE_DISPLAY_KHR
	UsageDisplayMode                   ObjectType = C.VK_OBJECT_TYPE_DISPLAY_MODE_KHR
	UsageDebugReportCallback           ObjectType = C.VK_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT
	UsageCUModuleNVX                   ObjectType = C.VK_OBJECT_TYPE_CU_MODULE_NVX
	UsageCUFunctionNVX                 ObjectType = C.VK_OBJECT_TYPE_CU_FUNCTION_NVX
	UsageDebugUtilsMessenger           ObjectType = C.VK_OBJECT_TYPE_DEBUG_UTILS_MESSENGER_EXT
	UsageAccelerationStructure         ObjectType = C.VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_KHR
	UsageValidationCache               ObjectType = C.VK_OBJECT_TYPE_VALIDATION_CACHE_EXT
	UsageAccelerationStructureNV       ObjectType = C.VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV
	UsagePerformanceConfigurationIntel ObjectType = C.VK_OBJECT_TYPE_PERFORMANCE_CONFIGURATION_INTEL
	UsageDeferredOperation             ObjectType = C.VK_OBJECT_TYPE_DEFERRED_OPERATION_KHR
	UsageIndirectCommandsLayoutNV      ObjectType = C.VK_OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NV
	UsagePrivateDataSlot               ObjectType = C.VK_OBJECT_TYPE_PRIVATE_DATA_SLOT_EXT
)

var vkObjectTypeToString = map[ObjectType]string{
	UsageUnknown:                       "Unknown",
	UsageInstance:                      "Instance",
	UsagePhysicalDevice:                "Physical Device",
	UsageDevice:                        "Device",
	UsageQueue:                         "Queue",
	UsageSemaphore:                     "Semaphore",
	UsageCommandBuffer:                 "Command Buffer",
	UsageFence:                         "Fence",
	UsageDeviceMemory:                  "Device Memory",
	UsageBuffer:                        "Buffer",
	UsageImage:                         "Image",
	UsageEvent:                         "Event",
	UsageQueryPool:                     "Query Pool",
	UsageBufferView:                    "Buffer View",
	UsageImageView:                     "Image View",
	UsageShaderModule:                  "Shader Module",
	UsagePipelineCache:                 "Pipeline Cache",
	UsagePipelineLayout:                "Pipeline Layout",
	UsageRenderPass:                    "Render Pass",
	UsagePipeline:                      "Pipeline",
	UsageDescriptorSetLayout:           "Descriptor Set Layout",
	UsageSampler:                       "Sampler",
	UsageDescriptorPool:                "Descriptor Pool",
	UsageDescriptorSet:                 "Descriptor Set",
	UsageFramebuffer:                   "Framebuffer",
	UsageCommandPool:                   "Command Pool",
	UsageSamplerYCbCrConversion:        "Sampler YCbCr Conversion",
	UsageDescriptorUpdateTemplate:      "Descriptor Update Template",
	UsageSurface:                       "Surface",
	UsageSwapchain:                     "Swapchain",
	UsageDisplay:                       "Display",
	UsageDisplayMode:                   "Display Mode",
	UsageDebugReportCallback:           "Debug Report Callback",
	UsageCUModuleNVX:                   "CU Module (Nvidia VX)",
	UsageCUFunctionNVX:                 "CU Function (Nvidia VX)",
	UsageDebugUtilsMessenger:           "Debug Utils Messenger",
	UsageAccelerationStructure:         "Acceleration Structure",
	UsageValidationCache:               "Validation Cache",
	UsageAccelerationStructureNV:       "Acceleration Structure (Nvidia)",
	UsagePerformanceConfigurationIntel: "Performance Configuration (Intel)",
	UsageDeferredOperation:             "Deferred Operation",
	UsageIndirectCommandsLayoutNV:      "Indirect Commands Layout (Nvidia)",
	UsagePrivateDataSlot:               "Private Data Slot",
}

func (t ObjectType) String() string {
	return vkObjectTypeToString[t]
}
