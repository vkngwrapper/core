package core1_0

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	ObjectTypeUnknown             ObjectType = C.VK_OBJECT_TYPE_UNKNOWN
	ObjectTypeInstance            ObjectType = C.VK_OBJECT_TYPE_INSTANCE
	ObjectTypePhysicalDevice      ObjectType = C.VK_OBJECT_TYPE_PHYSICAL_DEVICE
	ObjectTypeDevice              ObjectType = C.VK_OBJECT_TYPE_DEVICE
	ObjectTypeQueue               ObjectType = C.VK_OBJECT_TYPE_QUEUE
	ObjectTypeSemaphore           ObjectType = C.VK_OBJECT_TYPE_SEMAPHORE
	ObjectTypeCommandBuffer       ObjectType = C.VK_OBJECT_TYPE_COMMAND_BUFFER
	ObjectTypeFence               ObjectType = C.VK_OBJECT_TYPE_FENCE
	ObjectTypeDeviceMemory        ObjectType = C.VK_OBJECT_TYPE_DEVICE_MEMORY
	ObjectTypeBuffer              ObjectType = C.VK_OBJECT_TYPE_BUFFER
	ObjectTypeImage               ObjectType = C.VK_OBJECT_TYPE_IMAGE
	ObjectTypeEvent               ObjectType = C.VK_OBJECT_TYPE_EVENT
	ObjectTypeQueryPool           ObjectType = C.VK_OBJECT_TYPE_QUERY_POOL
	ObjectTypeBufferView          ObjectType = C.VK_OBJECT_TYPE_BUFFER_VIEW
	ObjectTypeImageView           ObjectType = C.VK_OBJECT_TYPE_IMAGE_VIEW
	ObjectTypeShaderModule        ObjectType = C.VK_OBJECT_TYPE_SHADER_MODULE
	ObjectTypePipelineCache       ObjectType = C.VK_OBJECT_TYPE_PIPELINE_CACHE
	ObjectTypePipelineLayout      ObjectType = C.VK_OBJECT_TYPE_PIPELINE_LAYOUT
	ObjectTypeRenderPass          ObjectType = C.VK_OBJECT_TYPE_RENDER_PASS
	ObjectTypePipeline            ObjectType = C.VK_OBJECT_TYPE_PIPELINE
	ObjectTypeDescriptorSetLayout ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT
	ObjectTypeSampler             ObjectType = C.VK_OBJECT_TYPE_SAMPLER
	ObjectTypeDescriptorPool      ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_POOL
	ObjectTypeDescriptorSet       ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_SET
	ObjectTypeFramebuffer         ObjectType = C.VK_OBJECT_TYPE_FRAMEBUFFER
	ObjectTypeCommandPool         ObjectType = C.VK_OBJECT_TYPE_COMMAND_POOL
)

func init() {
	ObjectTypeUnknown.Register("Unknown")
	ObjectTypeInstance.Register("Instance")
	ObjectTypePhysicalDevice.Register("Physical Device")
	ObjectTypeDevice.Register("Device")
	ObjectTypeQueue.Register("Queue")
	ObjectTypeSemaphore.Register("Semaphore")
	ObjectTypeCommandBuffer.Register("Command Buffer")
	ObjectTypeFence.Register("Fence")
	ObjectTypeDeviceMemory.Register("Device Memory")
	ObjectTypeBuffer.Register("Buffer")
	ObjectTypeImage.Register("Image")
	ObjectTypeEvent.Register("Event")
	ObjectTypeQueryPool.Register("Query Pool")
	ObjectTypeBufferView.Register("Buffer View")
	ObjectTypeImageView.Register("Image View")
	ObjectTypeShaderModule.Register("Shader Module")
	ObjectTypePipelineCache.Register("Pipeline Cache")
	ObjectTypePipelineLayout.Register("Pipeline Layout")
	ObjectTypeRenderPass.Register("Render Pass")
	ObjectTypePipeline.Register("Pipeline")
	ObjectTypeDescriptorSetLayout.Register("Descriptor Set Layout")
	ObjectTypeSampler.Register("Sampler")
	ObjectTypeDescriptorPool.Register("Descriptor Pool")
	ObjectTypeDescriptorSet.Register("Descriptor Set")
	ObjectTypeFramebuffer.Register("Framebuffer")
	ObjectTypeCommandPool.Register("Command Pool")
}
