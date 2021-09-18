package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"strings"
	"unsafe"
)

type DescriptorSetLayoutFlags int32

const (
	DescriptorSetLayoutUpdateAfterBindPool DescriptorSetLayoutFlags = C.VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT
	DescriptorSetLayoutPushDescriptor      DescriptorSetLayoutFlags = C.VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR
	DescriptorSetLayoutHostOnlyPoolValve   DescriptorSetLayoutFlags = C.VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_VALVE
)

var descriptorSetLayoutFlagsToString = map[DescriptorSetLayoutFlags]string{
	DescriptorSetLayoutUpdateAfterBindPool: "Update After Bind Pool",
	DescriptorSetLayoutPushDescriptor:      "Push Descriptor",
	DescriptorSetLayoutHostOnlyPoolValve:   "Host-Only Pool (Valve)",
}

func (f DescriptorSetLayoutFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := DescriptorSetLayoutFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := descriptorSetLayoutFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type DescriptorLayoutBinding struct {
	Binding      int
	Type         common.DescriptorType
	Count        int
	ShaderStages common.ShaderStages

	ImmutableSamplers []*Sampler
}

type DescriptorSetLayoutOptions struct {
	Flags    DescriptorSetLayoutFlags
	Bindings []*DescriptorLayoutBinding

	common.HaveNext
}

func (o *DescriptorSetLayoutOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkDescriptorSetLayoutCreateInfo)(allocator.Malloc(C.sizeof_struct_VkDescriptorSetLayoutCreateInfo))
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
			bindingsSlice[i].binding = C.uint32_t(o.Bindings[i].Binding)
			bindingsSlice[i].descriptorType = C.VkDescriptorType(o.Bindings[i].Type)
			bindingsSlice[i].descriptorCount = C.uint32_t(o.Bindings[i].Count)
			bindingsSlice[i].stageFlags = C.VkShaderStageFlags(o.Bindings[i].ShaderStages)

			//TODO: pImmutablesamplers
		}

		createInfo.pBindings = bindingsPtr
	}

	return unsafe.Pointer(createInfo), nil
}

type vulkanDescriptorSetLayout struct {
	driver Driver
	device VkDevice
	handle VkDescriptorSetLayout
}

func (h *vulkanDescriptorSetLayout) Handle() VkDescriptorSetLayout {
	return h.handle
}

func (h *vulkanDescriptorSetLayout) Destroy() error {
	return h.driver.VkDestroyDescriptorSetLayout(h.device, h.handle, nil)
}
