package options

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type DescriptorSetLayoutFlags int32

const (
	DescriptorSetLayoutUpdateAfterBindPool DescriptorSetLayoutFlags = C.VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT
	DescriptorSetLayoutPushDescriptorKHR   DescriptorSetLayoutFlags = C.VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR
	DescriptorSetLayoutHostOnlyPoolValve   DescriptorSetLayoutFlags = C.VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_VALVE
)

var descriptorSetLayoutFlagsToString = map[DescriptorSetLayoutFlags]string{
	DescriptorSetLayoutUpdateAfterBindPool: "Update After Bind Pool",
	DescriptorSetLayoutPushDescriptorKHR:   "Push Descriptor (Khronos Extension)",
	DescriptorSetLayoutHostOnlyPoolValve:   "Host-Only Pool (Valve Extension)",
}

func (f DescriptorSetLayoutFlags) String() string {
	return common.FlagsToString(f, descriptorSetLayoutFlagsToString)
}

type DescriptorLayoutBinding struct {
	Binding         int
	DescriptorType  common.DescriptorType
	DescriptorCount int
	StageFlags      common.ShaderStages

	ImmutableSamplers []iface.Sampler
}

type DescriptorSetLayoutOptions struct {
	Flags    DescriptorSetLayoutFlags
	Bindings []*DescriptorLayoutBinding

	core.HaveNext
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
