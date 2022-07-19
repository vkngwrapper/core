package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"

const (
	// ObjectTypeUnknown specifies an unknown or undefined handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeUnknown ObjectType = C.VK_OBJECT_TYPE_UNKNOWN
	// ObjectTypeInstance specifies an Instance handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeInstance ObjectType = C.VK_OBJECT_TYPE_INSTANCE
	// ObjectTypePhysicalDevice specifies a PhysicalDevice handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypePhysicalDevice ObjectType = C.VK_OBJECT_TYPE_PHYSICAL_DEVICE
	// ObjectTypeDevice specifies a Device handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeDevice ObjectType = C.VK_OBJECT_TYPE_DEVICE
	// ObjectTypeQueue specifies a Queue handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeQueue ObjectType = C.VK_OBJECT_TYPE_QUEUE
	// ObjectTypeSemaphore specifies a Semaphore handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeSemaphore ObjectType = C.VK_OBJECT_TYPE_SEMAPHORE
	// ObjectTypeCommandBuffer specifies a CommandBuffer handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeCommandBuffer ObjectType = C.VK_OBJECT_TYPE_COMMAND_BUFFER
	// ObjectTypeFence specifies a Fence handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeFence ObjectType = C.VK_OBJECT_TYPE_FENCE
	// ObjectTypeDeviceMemory specifies a DeviceMemory handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeDeviceMemory ObjectType = C.VK_OBJECT_TYPE_DEVICE_MEMORY
	// ObjectTypeBuffer specifies a Buffer handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeBuffer ObjectType = C.VK_OBJECT_TYPE_BUFFER
	// ObjectTypeImage specifies an Image handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeImage ObjectType = C.VK_OBJECT_TYPE_IMAGE
	// ObjectTypeEvent specifies an Event handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeEvent ObjectType = C.VK_OBJECT_TYPE_EVENT
	// ObjectTypeQueryPool specifies a QueryPool handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeQueryPool ObjectType = C.VK_OBJECT_TYPE_QUERY_POOL
	// ObjectTypeBufferView specifies a BufferView handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeBufferView ObjectType = C.VK_OBJECT_TYPE_BUFFER_VIEW
	// ObjectTypeImageView specifies an ImageView handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeImageView ObjectType = C.VK_OBJECT_TYPE_IMAGE_VIEW
	// ObjectTypeShaderModule specifies a ShaderModule handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeShaderModule ObjectType = C.VK_OBJECT_TYPE_SHADER_MODULE
	// ObjectTypePipelineCache specifies a PipelineCache handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypePipelineCache ObjectType = C.VK_OBJECT_TYPE_PIPELINE_CACHE
	// ObjectTypePipelineLayout specifies a PipelineLayout handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypePipelineLayout ObjectType = C.VK_OBJECT_TYPE_PIPELINE_LAYOUT
	// ObjectTypeRenderPass specifies a RenderPass handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeRenderPass ObjectType = C.VK_OBJECT_TYPE_RENDER_PASS
	// ObjectTypePipeline specifies a Pipeline handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypePipeline ObjectType = C.VK_OBJECT_TYPE_PIPELINE
	// ObjectTypeDescriptorSetLayout specifies a DescriptorSetLayout handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeDescriptorSetLayout ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT
	// ObjectTypeSampler specifies a Sampler handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeSampler ObjectType = C.VK_OBJECT_TYPE_SAMPLER
	// ObjectTypeDescriptorPool specifies a DescriptorPool handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeDescriptorPool ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_POOL
	// ObjectTypeDescriptorSet specifies a DescriptorSet handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeDescriptorSet ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_SET
	// ObjectTypeFramebuffer specifies a Framebuffer handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeFramebuffer ObjectType = C.VK_OBJECT_TYPE_FRAMEBUFFER
	// ObjectTypeCommandPool specifies a CommandPool handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeCommandPool ObjectType = C.VK_OBJECT_TYPE_COMMAND_POOL
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
