package core1_2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DescriptorBindingFlags int32

var descriptorBindingFlagsMapping = common.NewFlagStringMapping[DescriptorBindingFlags]()

func (f DescriptorBindingFlags) Register(str string) {
	descriptorBindingFlagsMapping.Register(f, str)
}
func (f DescriptorBindingFlags) String() string {
	return descriptorBindingFlagsMapping.FlagsToString(f)
}

////

const (
	DescriptorBindingPartiallyBound           DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT
	DescriptorBindingUpdateAfterBind          DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT
	DescriptorBindingUpdateUnusedWhilePending DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT
	DescriptorBindingVariableDescriptorCount  DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT

	DescriptorSetLayoutCreateUpdateAfterBindPool core1_0.DescriptorSetLayoutCreateFlags = C.VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT
)

func init() {
	DescriptorBindingPartiallyBound.Register("Partially-Bound")
	DescriptorBindingUpdateAfterBind.Register("Update After Bind")
	DescriptorBindingUpdateUnusedWhilePending.Register("Update Unused While Pending")
	DescriptorBindingVariableDescriptorCount.Register("Variable Descriptor Count")

	DescriptorSetLayoutCreateUpdateAfterBindPool.Register("Update After Bind Pool")
}

////

type DescriptorSetVariableDescriptorCountAllocateInfo struct {
	DescriptorCounts []int

	common.NextOptions
}

func (o DescriptorSetVariableDescriptorCountAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorSetVariableDescriptorCountAllocateInfo{})))
	}

	info := (*C.VkDescriptorSetVariableDescriptorCountAllocateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO
	info.pNext = next

	count := len(o.DescriptorCounts)
	info.descriptorSetCount = C.uint32_t(count)
	info.pDescriptorCounts = nil

	if count > 0 {
		info.pDescriptorCounts = (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		descriptorCountSlice := unsafe.Slice(info.pDescriptorCounts, count)
		for i := 0; i < count; i++ {
			descriptorCountSlice[i] = C.uint32_t(o.DescriptorCounts[i])
		}
	}

	return preallocatedPointer, nil
}

////

type DescriptorSetLayoutBindingFlagsCreateInfo struct {
	BindingFlags []DescriptorBindingFlags

	common.NextOptions
}

func (o DescriptorSetLayoutBindingFlagsCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorSetLayoutBindingFlagsCreateInfo{})))
	}

	info := (*C.VkDescriptorSetLayoutBindingFlagsCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO
	info.pNext = next

	count := len(o.BindingFlags)
	info.bindingCount = C.uint32_t(count)
	info.pBindingFlags = nil

	if count > 0 {
		info.pBindingFlags = (*C.VkDescriptorBindingFlags)(allocator.Malloc(count * int(unsafe.Sizeof(C.VkDescriptorBindingFlags(0)))))
		flagSlice := unsafe.Slice(info.pBindingFlags, count)

		for i := 0; i < count; i++ {
			flagSlice[i] = C.VkDescriptorBindingFlags(o.BindingFlags[i])
		}
	}

	return preallocatedPointer, nil
}

////

type DescriptorSetVariableDescriptorCountLayoutSupport struct {
	MaxVariableDescriptorCount int

	common.NextOutData
}

func (o *DescriptorSetVariableDescriptorCountLayoutSupport) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorSetVariableDescriptorCountLayoutSupport{})))
	}

	info := (*C.VkDescriptorSetVariableDescriptorCountLayoutSupport)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *DescriptorSetVariableDescriptorCountLayoutSupport) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDescriptorSetVariableDescriptorCountLayoutSupport)(cDataPointer)

	o.MaxVariableDescriptorCount = int(info.maxVariableDescriptorCount)

	return info.pNext, nil
}
