package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

const (
	// DescriptorTypeSampler specifies a Sampler descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
	DescriptorTypeSampler DescriptorType = C.VK_DESCRIPTOR_TYPE_SAMPLER
	// DescriptorTypeCombinedImageSampler specifies a combined Image Sampler descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
	DescriptorTypeCombinedImageSampler DescriptorType = C.VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
	// DescriptorTypeSampledImage specifies a sampled Image descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
	DescriptorTypeSampledImage DescriptorType = C.VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE
	// DescriptorTypeStorageImage specifies a storage Image descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
	DescriptorTypeStorageImage DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_IMAGE
	// DescriptorTypeUniformTexelBuffer specifies a uniform texel Buffer descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
	DescriptorTypeUniformTexelBuffer DescriptorType = C.VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER
	// DescriptorTypeStorageTexelBuffer specifies a storage texel Buffer descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
	DescriptorTypeStorageTexelBuffer DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER
	// DescriptorTypeUniformBuffer specifies a uniform Buffer descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
	DescriptorTypeUniformBuffer DescriptorType = C.VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
	// DescriptorTypeStorageBuffer specifies a storage Buffer descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
	DescriptorTypeStorageBuffer DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_BUFFER
	// DescriptorTypeUniformBufferDynamic specifies a dynamic uniform Buffer descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
	DescriptorTypeUniformBufferDynamic DescriptorType = C.VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC
	// DescriptorTypeStorageBufferDynamic specifies a dynamic storage Buffer descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
	DescriptorTypeStorageBufferDynamic DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC
	// DescriptorTypeInputAttachment specifies an input attachment descriptor
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html
	DescriptorTypeInputAttachment DescriptorType = C.VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT
)

func init() {
	DescriptorTypeSampler.Register("Sampler")
	DescriptorTypeCombinedImageSampler.Register("Combined Image Sampler")
	DescriptorTypeSampledImage.Register("Sampled Image")
	DescriptorTypeStorageImage.Register("Storage Image")
	DescriptorTypeUniformTexelBuffer.Register("Uniform Texel Buffer")
	DescriptorTypeStorageTexelBuffer.Register("Storage Texel Buffer")
	DescriptorTypeUniformBuffer.Register("Uniform Buffer")
	DescriptorTypeStorageBuffer.Register("Storage Buffer")
	DescriptorTypeUniformBufferDynamic.Register("Uniform Buffer Dynamic")
	DescriptorTypeStorageBufferDynamic.Register("Storage Buffer Dynamic")
	DescriptorTypeInputAttachment.Register("Input Attachment")
}

// DescriptorSetLayoutBinding specifies a DescriptorSetLayout binding
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBinding.html
type DescriptorSetLayoutBinding struct {
	// Binding is the binding number of this entry and corresponds to a resource of the same
	// binding number in the shader stages
	Binding int
	// DescriptorType specifies which type of resource descriptors are used for this binding
	DescriptorType DescriptorType
	// DescriptorCount is the number of descriptors contained in the binding
	DescriptorCount int
	// StageFlags specifies which pipeline shader stages can access a resource for this binding
	StageFlags ShaderStageFlags

	// ImmutableSamplers is a slice of Sampler objects that will be copied into the set layout
	// and used for the corresponding binding.
	ImmutableSamplers []Sampler
}

// DescriptorSetLayoutCreateInfo specifies parameters of a newly-created DescriptorSetLayout
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html
type DescriptorSetLayoutCreateInfo struct {
	// Flags specifies options for DescriptorSetLayout creation
	Flags DescriptorSetLayoutCreateFlags
	// Bindings is a slice of DescriptorSetLayoutBinding structures
	Bindings []DescriptorSetLayoutBinding

	common.NextOptions
}

func (o DescriptorSetLayoutCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDescriptorSetLayoutCreateInfo)
	}
	createInfo := (*C.VkDescriptorSetLayoutCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
	createInfo.flags = C.VkDescriptorSetLayoutCreateFlags(o.Flags)
	createInfo.pNext = next

	bindingCount := len(o.Bindings)
	createInfo.bindingCount = C.uint32_t(bindingCount)
	createInfo.pBindings = nil

	if bindingCount > 0 {
		bindingsPtr := (*C.VkDescriptorSetLayoutBinding)(allocator.Malloc(bindingCount * C.sizeof_struct_VkDescriptorSetLayoutBinding))
		bindingsSlice := ([]C.VkDescriptorSetLayoutBinding)(unsafe.Slice(bindingsPtr, bindingCount))

		for i := 0; i < bindingCount; i++ {
			samplerCount := len(o.Bindings[i].ImmutableSamplers)
			if samplerCount != 0 && samplerCount != o.Bindings[i].DescriptorCount {
				return nil, errors.Newf("allocate descriptor set layout bindings: binding %d has %d descriptors, but %d immutable samplers. if immutable samplers are provided, they must match the descriptor count", i, o.Bindings[i].DescriptorCount, len(o.Bindings[i].ImmutableSamplers))
			}

			bindingsSlice[i].binding = C.uint32_t(o.Bindings[i].Binding)
			bindingsSlice[i].descriptorType = C.VkDescriptorType(o.Bindings[i].DescriptorType)
			bindingsSlice[i].descriptorCount = C.uint32_t(o.Bindings[i].DescriptorCount)
			bindingsSlice[i].stageFlags = C.VkShaderStageFlags(o.Bindings[i].StageFlags)

			bindingsSlice[i].pImmutableSamplers = nil
			if samplerCount > 0 {
				immutableSamplerPtr := (*C.VkSampler)(allocator.Malloc(samplerCount * int(unsafe.Sizeof([1]C.VkSampler{}))))
				immutableSamplerSlice := ([]C.VkSampler)(unsafe.Slice(immutableSamplerPtr, samplerCount))

				for samplerIndex := 0; samplerIndex < samplerCount; samplerIndex++ {
					immutableSamplerSlice[samplerIndex] = C.VkSampler(unsafe.Pointer(o.Bindings[i].ImmutableSamplers[samplerIndex].Handle()))
				}

				bindingsSlice[i].pImmutableSamplers = immutableSamplerPtr
			}
		}

		createInfo.pBindings = bindingsPtr
	}

	return preallocatedPointer, nil
}
