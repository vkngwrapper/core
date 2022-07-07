package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DescriptorUpdateTemplateType int32

var descriptorTemplateTypeMapping = make(map[DescriptorUpdateTemplateType]string)

func (e DescriptorUpdateTemplateType) Register(str string) {
	descriptorTemplateTypeMapping[e] = str
}

func (e DescriptorUpdateTemplateType) String() string {
	return descriptorTemplateTypeMapping[e]
}

////

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
	DescriptorUpdateTemplateTypeDescriptorSet DescriptorUpdateTemplateType = C.VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET
)

func init() {
	DescriptorUpdateTemplateTypeDescriptorSet.Register("Descriptor Set")
}

////

type DescriptorUpdateTemplateEntry struct {
	DstBinding      int
	DstArrayElement int
	DescriptorCount int

	DescriptorType core1_0.DescriptorType

	Offset int
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

type DescriptorUpdateTemplateCreateInfo struct {
	Flags                   DescriptorUpdateTemplateCreateFlags
	DescriptorUpdateEntries []DescriptorUpdateTemplateEntry
	TemplateType            DescriptorUpdateTemplateType

	DescriptorSetLayout core1_0.DescriptorSetLayout

	PipelineBindPoint core1_0.PipelineBindPoint
	PipelineLayout    core1_0.PipelineLayout
	Set               int

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

	if o.DescriptorSetLayout != nil {
		createInfo.descriptorSetLayout = C.VkDescriptorSetLayout(unsafe.Pointer(o.DescriptorSetLayout.Handle()))
	}

	if o.PipelineLayout != nil {
		createInfo.pipelineLayout = C.VkPipelineLayout(unsafe.Pointer(o.PipelineLayout.Handle()))
	}

	createInfo.pipelineBindPoint = C.VkPipelineBindPoint(o.PipelineBindPoint)
	createInfo.set = C.uint32_t(o.Set)

	return preallocatedPointer, nil
}
