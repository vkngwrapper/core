package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

const (
	DescriptorSampler              common.DescriptorType = C.VK_DESCRIPTOR_TYPE_SAMPLER
	DescriptorCombinedImageSampler common.DescriptorType = C.VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
	DescriptorSampledImage         common.DescriptorType = C.VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE
	DescriptorStorageImage         common.DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_IMAGE
	DescriptorUniformTexelBuffer   common.DescriptorType = C.VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER
	DescriptorStorageTexelBuffer   common.DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER
	DescriptorUniformBuffer        common.DescriptorType = C.VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
	DescriptorStorageBuffer        common.DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_BUFFER
	DescriptorUniformBufferDynamic common.DescriptorType = C.VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC
	DescriptorStorageBufferDynamic common.DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC
	DescriptorInputAttachment      common.DescriptorType = C.VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT
)

func init() {
	DescriptorSampler.Register("Sampler")
	DescriptorCombinedImageSampler.Register("Combined Image Sampler")
	DescriptorSampledImage.Register("Sampled Image")
	DescriptorStorageImage.Register("Storage Image")
	DescriptorUniformTexelBuffer.Register("Uniform Texel Buffer")
	DescriptorStorageTexelBuffer.Register("Storage Texel Buffer")
	DescriptorUniformBuffer.Register("Uniform Buffer")
	DescriptorStorageBuffer.Register("Storage Buffer")
	DescriptorUniformBufferDynamic.Register("Uniform Buffer Dynamic")
	DescriptorStorageBufferDynamic.Register("Storage Buffer Dynamic")
	DescriptorInputAttachment.Register("Input Attachment")
}

type DescriptorLayoutBinding struct {
	Binding         int
	DescriptorType  common.DescriptorType
	DescriptorCount int
	StageFlags      common.ShaderStages

	ImmutableSamplers []Sampler
}

type DescriptorSetLayoutOptions struct {
	Flags    common.DescriptorSetLayoutCreateFlags
	Bindings []DescriptorLayoutBinding

	common.HaveNext
}

func (o *DescriptorSetLayoutOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o DescriptorSetLayoutOptions) PopulateOutData(cDataPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkDescriptorSetLayoutCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
