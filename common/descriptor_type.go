package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type DescriptorType int32

const (
	DescriptorSampler                  DescriptorType = C.VK_DESCRIPTOR_TYPE_SAMPLER
	DescriptorCombinedImageSampler     DescriptorType = C.VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
	DescriptorSampledImage             DescriptorType = C.VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE
	DescriptorStorageImage             DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_IMAGE
	DescriptorUniformTexelBuffer       DescriptorType = C.VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER
	DescriptorStorageTexelBuffer       DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER
	DescriptorUniformBuffer            DescriptorType = C.VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER
	DescriptorStorageBuffer            DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_BUFFER
	DescriptorUniformBufferDynamic     DescriptorType = C.VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC
	DescriptorStorageBufferDynamic     DescriptorType = C.VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC
	DescriptorInputAttachment          DescriptorType = C.VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT
	DescriptorInlineUniformBlockEXT    DescriptorType = C.VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK_EXT
	DescriptorAccelerationStructureKHR DescriptorType = C.VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR
	DescriptorAccelerationStructureNV  DescriptorType = C.VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV
	DescriptorMutableValve             DescriptorType = C.VK_DESCRIPTOR_TYPE_MUTABLE_VALVE
)

var descriptorTypeToString = map[DescriptorType]string{
	DescriptorSampler:                  "Sampler",
	DescriptorCombinedImageSampler:     "Combined Image Sampler",
	DescriptorSampledImage:             "Sampled Image",
	DescriptorStorageImage:             "Storage Image",
	DescriptorUniformTexelBuffer:       "Uniform Texel Buffer",
	DescriptorStorageTexelBuffer:       "Storage Texel Buffer",
	DescriptorUniformBuffer:            "Uniform Buffer",
	DescriptorStorageBuffer:            "Storage Buffer",
	DescriptorUniformBufferDynamic:     "Uniform Buffer Dynamic",
	DescriptorStorageBufferDynamic:     "Storage Buffer Dynamic",
	DescriptorInputAttachment:          "Input Attachment",
	DescriptorInlineUniformBlockEXT:    "Inline Uniform Block (Extension)",
	DescriptorAccelerationStructureKHR: "Acceleration Structure (Khronos Extension)",
	DescriptorAccelerationStructureNV:  "Acceleration Structure (Nvidia Extension)",
	DescriptorMutableValve:             "Mutable (Valve Extension)",
}

func (t DescriptorType) String() string {
	return descriptorTypeToString[t]
}
