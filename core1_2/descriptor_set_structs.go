package core1_2

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// DescriptorBindingFlags specifies DescriptorSetLayout binding properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html
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
	// DescriptorBindingPartiallyBound indicates that descriptors in this binding that are
	// not dynamically used need not contain valid descriptors at the time the descriptors
	// are consumed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html
	DescriptorBindingPartiallyBound DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT
	// DescriptorBindingUpdateAfterBind indicates that if descriptors in this binding are updated
	// between when the DescriptorSet is bound in a CommandBuffer and when that CommandBuffer is
	// submitted to a Queue, then the submission will use the most recently-set descriptors
	// for this binding and the updates do not invalidate the CommandBuffer
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html
	DescriptorBindingUpdateAfterBind DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT
	// DescriptorBindingUpdateUnusedWhilePending indicates that descriptors in this binding can be
	// updated after a CommandBuffer has bound this DescriptorSet, or while a CommandBuffer that
	// uses this DescriptorSet is pending execution, as long as the descriptors that are updated
	// are not used by those CommandBuffer objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html
	DescriptorBindingUpdateUnusedWhilePending DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT
	// DescriptorBindingVariableDescriptorCount indicates that this is a variable-sized descriptor
	// binding whose size will be specified when a DescriptorSet is allocated using this layout
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html
	DescriptorBindingVariableDescriptorCount DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT

	// DescriptorSetLayoutCreateUpdateAfterBindPool specifies that DescriptorSet objects using this
	// layout must be allocated from a DescriptorPool created with DescriptorPoolCreateUpdateAfterBind
	// set
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlagBits.html
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

// DescriptorSetVariableDescriptorCountAllocateInfo specifies additional allocation parameters
// for DescriptorSet objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfo.html
type DescriptorSetVariableDescriptorCountAllocateInfo struct {
	// DescriptorCounts is a slice of descriptor counts, with each member specifying the number
	// of descriptors in a variable-sized descriptor binding in the corresponding DescriptorSet
	// being allocated
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

// DescriptorSetLayoutBindingFlagsCreateInfo specifies parameters of a newly-created
// DescriptorSetLayout
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfo.html
type DescriptorSetLayoutBindingFlagsCreateInfo struct {
	// BindingFlags is a slice of DescriptorBindingFlags, one for each DescriptorSetLayout binding
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

// DescriptorSetVariableDescriptorCountLayoutSupport returns information about whether a
// DescriptorSetLayout can be supported
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetVariableDescriptorCountLayoutSupport.html
type DescriptorSetVariableDescriptorCountLayoutSupport struct {
	// MaxVariableDescriptorCount indicates the maximum number of descriptors supported in the
	// highest numbered binding of the layout, if that binding is variable-sized
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
