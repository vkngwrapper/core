package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type ObjectType int32

const (
	Unknown                       ObjectType = C.VK_OBJECT_TYPE_UNKNOWN
	Instance                      ObjectType = C.VK_OBJECT_TYPE_INSTANCE
	PhysicalDevice                ObjectType = C.VK_OBJECT_TYPE_PHYSICAL_DEVICE
	Device                        ObjectType = C.VK_OBJECT_TYPE_DEVICE
	Queue                         ObjectType = C.VK_OBJECT_TYPE_QUEUE
	Semaphore                     ObjectType = C.VK_OBJECT_TYPE_SEMAPHORE
	CommandBuffer                 ObjectType = C.VK_OBJECT_TYPE_COMMAND_BUFFER
	Fence                         ObjectType = C.VK_OBJECT_TYPE_FENCE
	DeviceMemory                  ObjectType = C.VK_OBJECT_TYPE_DEVICE_MEMORY
	Buffer                        ObjectType = C.VK_OBJECT_TYPE_BUFFER
	Image                         ObjectType = C.VK_OBJECT_TYPE_IMAGE
	Event                         ObjectType = C.VK_OBJECT_TYPE_EVENT
	QueryPool                     ObjectType = C.VK_OBJECT_TYPE_QUERY_POOL
	BufferView                    ObjectType = C.VK_OBJECT_TYPE_BUFFER_VIEW
	ImageView                     ObjectType = C.VK_OBJECT_TYPE_IMAGE_VIEW
	ShaderModule                  ObjectType = C.VK_OBJECT_TYPE_SHADER_MODULE
	PipelineCache                 ObjectType = C.VK_OBJECT_TYPE_PIPELINE_CACHE
	PipelineLayout                ObjectType = C.VK_OBJECT_TYPE_PIPELINE_LAYOUT
	RenderPass                    ObjectType = C.VK_OBJECT_TYPE_RENDER_PASS
	Pipeline                      ObjectType = C.VK_OBJECT_TYPE_PIPELINE
	DescriptorSetLayout           ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT
	Sampler                       ObjectType = C.VK_OBJECT_TYPE_SAMPLER
	DescriptorPool                ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_POOL
	DescriptorSet                 ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_SET
	Framebuffer                   ObjectType = C.VK_OBJECT_TYPE_FRAMEBUFFER
	CommandPool                   ObjectType = C.VK_OBJECT_TYPE_COMMAND_POOL
	SamplerYCbCrConversion        ObjectType = C.VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION
	DescriptorUpdateTemplate      ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE
	Surface                       ObjectType = C.VK_OBJECT_TYPE_SURFACE_KHR
	Swapchain                     ObjectType = C.VK_OBJECT_TYPE_SWAPCHAIN_KHR
	Display                       ObjectType = C.VK_OBJECT_TYPE_DISPLAY_KHR
	DisplayMode                   ObjectType = C.VK_OBJECT_TYPE_DISPLAY_MODE_KHR
	DebugReportCallback           ObjectType = C.VK_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT
	CUModuleNVX                   ObjectType = C.VK_OBJECT_TYPE_CU_MODULE_NVX
	CUFunctionNVX                 ObjectType = C.VK_OBJECT_TYPE_CU_FUNCTION_NVX
	DebugUtilsMessenger           ObjectType = C.VK_OBJECT_TYPE_DEBUG_UTILS_MESSENGER_EXT
	AccelerationStructure         ObjectType = C.VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_KHR
	ValidationCache               ObjectType = C.VK_OBJECT_TYPE_VALIDATION_CACHE_EXT
	AccelerationStructureNV       ObjectType = C.VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV
	PerformanceConfigurationIntel ObjectType = C.VK_OBJECT_TYPE_PERFORMANCE_CONFIGURATION_INTEL
	DeferredOperation             ObjectType = C.VK_OBJECT_TYPE_DEFERRED_OPERATION_KHR
	IndirectCommandsLayoutNV      ObjectType = C.VK_OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NV
	PrivateDataSlot               ObjectType = C.VK_OBJECT_TYPE_PRIVATE_DATA_SLOT_EXT
)

var vkObjectTypeToString = map[ObjectType]string{
	Unknown:                       "Unknown",
	Instance:                      "Instance",
	PhysicalDevice:                "Physical Device",
	Device:                        "Device",
	Queue:                         "Queue",
	Semaphore:                     "Semaphore",
	CommandBuffer:                 "Command Buffer",
	Fence:                         "Fence",
	DeviceMemory:                  "Device Memory",
	Buffer:                        "Buffer",
	Image:                         "Image",
	Event:                         "Event",
	QueryPool:                     "Query Pool",
	BufferView:                    "Buffer View",
	ImageView:                     "Image View",
	ShaderModule:                  "Shader Module",
	PipelineCache:                 "Pipeline Cache",
	PipelineLayout:                "Pipeline Layout",
	RenderPass:                    "Render Pass",
	Pipeline:                      "Pipeline",
	DescriptorSetLayout:           "Descriptor Set Layout",
	Sampler:                       "Sampler",
	DescriptorPool:                "Descriptor Pool",
	DescriptorSet:                 "Descriptor Set",
	Framebuffer:                   "Framebuffer",
	CommandPool:                   "Command Pool",
	SamplerYCbCrConversion:        "Sampler YCbCr Conversion",
	DescriptorUpdateTemplate:      "Descriptor Update Template",
	Surface:                       "Surface",
	Swapchain:                     "Swapchain",
	Display:                       "Display",
	DisplayMode:                   "Display Mode",
	DebugReportCallback:           "Debug Report Callback",
	CUModuleNVX:                   "CU Module (Nvidia VX)",
	CUFunctionNVX:                 "CU Function (Nvidia VX)",
	DebugUtilsMessenger:           "Debug Utils Messenger",
	AccelerationStructure:         "Acceleration Structure",
	ValidationCache:               "Validation Cache",
	AccelerationStructureNV:       "Acceleration Structure (Nvidia)",
	PerformanceConfigurationIntel: "Performance Configuration (Intel)",
	DeferredOperation:             "Deferred Operation",
	IndirectCommandsLayoutNV:      "Indirect Commands Layout (Nvidia)",
	PrivateDataSlot:               "Private Data Slot",
}

func (t ObjectType) String() string {
	return vkObjectTypeToString[t]
}
