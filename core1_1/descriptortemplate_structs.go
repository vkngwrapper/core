package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

// DescriptorUpdateTemplateType indicates the valid usage of the DescriptorUpdateTemplate
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateType.html
type DescriptorUpdateTemplateType int32

var descriptorTemplateTypeMapping = make(map[DescriptorUpdateTemplateType]string)

func (e DescriptorUpdateTemplateType) Register(str string) {
	descriptorTemplateTypeMapping[e] = str
}

func (e DescriptorUpdateTemplateType) String() string {
	return descriptorTemplateTypeMapping[e]
}

////

// DescriptorUpdateTemplateCreateFlags is reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateFlags.html
type DescriptorUpdateTemplateCreateFlags int32

var descriptorTemplateFlagsMapping = common.NewFlagStringMapping[DescriptorUpdateTemplateCreateFlags]()

func (f DescriptorUpdateTemplateCreateFlags) Register(str string) {
	descriptorTemplateFlagsMapping.Register(f, str)
}
func (f DescriptorUpdateTemplateCreateFlags) String() string {
	return descriptorTemplateFlagsMapping.FlagsToString(f)
}

////

const (
	// DescriptorUpdateTemplateTypeDescriptorSet indicates the valid usage of the DescriptorUpdateTemplate
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateType.html
	DescriptorUpdateTemplateTypeDescriptorSet DescriptorUpdateTemplateType = C.VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET
)

func init() {
	DescriptorUpdateTemplateTypeDescriptorSet.Register("Descriptor Set")
}

////

// DescriptorUpdateTemplateEntry describes a single descriptor update of the DescriptorUpdateTemplate
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateEntry.html
type DescriptorUpdateTemplateEntry struct {
	// DstBinding is the descriptor binding to update when using this DescriptorUpdateTemplate
	DstBinding int
	// DstArrayElement is the starting element in the array belonging to DstBinding
	DstArrayElement int
	// DescriptorCount is the number of descriptors to update
	DescriptorCount int

	// DescriptorType specifies the type of the descriptor
	DescriptorType core1_0.DescriptorType

	// Offset is the offset in bytes of the first binding in the raw data structure
	Offset int
	// Stride is the stride in bytes between two consecutive array elements of the
	// descriptor update informations in the raw data structure
	Stride int
}

func (e DescriptorUpdateTemplateEntry) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorUpdateTemplateEntry{})))
	}

	entry := (*C.VkDescriptorUpdateTemplateEntry)(preallocatedPointer)
	entry.dstBinding = C.uint32_t(e.DstBinding)
	entry.dstArrayElement = C.uint32_t(e.DstArrayElement)
	entry.descriptorCount = C.uint32_t(e.DescriptorCount)
	entry.descriptorType = C.VkDescriptorType(e.DescriptorType)
	entry.offset = C.size_t(e.Offset)
	entry.stride = C.size_t(e.Stride)

	return preallocatedPointer, nil
}

// DescriptorUpdateTemplateCreateInfo specifies parameters of a newly-created Descriptor Update
// Template
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html
type DescriptorUpdateTemplateCreateInfo struct {
	// Flags is reserved for future use
	Flags DescriptorUpdateTemplateCreateFlags
	// DescriptorUpdateEntries is a slice of DescriptorUpdateTemplateEntry structures describing
	// the descriptors to be updated by the DescriptorUpdateTEmplate
	DescriptorUpdateEntries []DescriptorUpdateTemplateEntry
	// TemplateType specifies the type of the DescriptorUpdateTemplate
	TemplateType DescriptorUpdateTemplateType

	// DescriptorSetLayout is the DescriptorSetLayout used to build the DescriptorUpdateTemplate
	DescriptorSetLayout core.DescriptorSetLayout

	// PipelineBindPoint indicates the type of the Pipeline that will use the descriptors
	PipelineBindPoint core1_0.PipelineBindPoint
	// PipelineLayout is a PipelineLayout object used to program the bindings
	PipelineLayout core.PipelineLayout
	// Set is the set number of the DescriptorSet in the PipelineLayout that will be updated
	Set int

	common.NextOptions
}

func (o DescriptorUpdateTemplateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorUpdateTemplateCreateInfo{})))
	}

	createInfo := (*C.VkDescriptorUpdateTemplateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkDescriptorUpdateTemplateCreateFlags(o.Flags)

	entryCount := len(o.DescriptorUpdateEntries)
	createInfo.descriptorUpdateEntryCount = C.uint32_t(entryCount)

	var err error
	createInfo.pDescriptorUpdateEntries, err = common.AllocSlice[C.VkDescriptorUpdateTemplateEntry, DescriptorUpdateTemplateEntry](allocator, o.DescriptorUpdateEntries)
	if err != nil {
		return nil, err
	}

	createInfo.templateType = C.VkDescriptorUpdateTemplateType(o.TemplateType)
	createInfo.descriptorSetLayout = nil
	createInfo.pipelineLayout = nil

	if o.DescriptorSetLayout.Initialized() {
		createInfo.descriptorSetLayout = C.VkDescriptorSetLayout(unsafe.Pointer(o.DescriptorSetLayout.Handle()))
	}

	if o.PipelineLayout.Initialized() {
		createInfo.pipelineLayout = C.VkPipelineLayout(unsafe.Pointer(o.PipelineLayout.Handle()))
	}

	createInfo.pipelineBindPoint = C.VkPipelineBindPoint(o.PipelineBindPoint)
	createInfo.set = C.uint32_t(o.Set)

	return preallocatedPointer, nil
}
