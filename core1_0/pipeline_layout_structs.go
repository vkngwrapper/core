package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// PipelineLayoutCreateFlags represents PipelineLayout creation flag bits
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateFlagBits.html
type PipelineLayoutCreateFlags uint32

var pipelineLayoutCreateFlagsMapping = common.NewFlagStringMapping[PipelineLayoutCreateFlags]()

func (f PipelineLayoutCreateFlags) Register(str string) {
	pipelineLayoutCreateFlagsMapping.Register(f, str)
}

func (f PipelineLayoutCreateFlags) String() string {
	return pipelineLayoutCreateFlagsMapping.FlagsToString(f)
}

////

// PushConstantRange specifies a push constant range
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantRange.html
type PushConstantRange struct {
	// StageFlags describes the shader stages that will access a range of push constants
	StageFlags ShaderStageFlags
	// Offset is the start offset consumed by the range, in bytes. Must be a multiple of 4
	Offset int
	// Size is the size consumed by the range, in bytes. Must be a multiple of 4
	Size int
}

// PipelineLayoutCreateInfo creates a new PipelineLayout object
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html
type PipelineLayoutCreateInfo struct {
	// Flags specifies options for PipelineLayout creation
	Flags PipelineLayoutCreateFlags

	// SetLayouts is a slice of DescriptorSetLayout objects
	SetLayouts []DescriptorSetLayout
	// PushConstantRanges is a slice of PushConstantRange structures defining a set of push constant
	// ranges for use in a single PipelineLayout
	PushConstantRanges []PushConstantRange

	common.NextOptions
}

func (o PipelineLayoutCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineLayoutCreateInfo)
	}
	createInfo := (*C.VkPipelineLayoutCreateInfo)(preallocatedPointer)

	setLayoutCount := len(o.SetLayouts)
	constantRangesCount := len(o.PushConstantRanges)

	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkPipelineLayoutCreateFlags(o.Flags)
	createInfo.setLayoutCount = C.uint32_t(setLayoutCount)
	createInfo.pushConstantRangeCount = C.uint32_t(constantRangesCount)

	createInfo.pSetLayouts = nil
	if setLayoutCount > 0 {
		setLayoutPtr := (*C.VkDescriptorSetLayout)(allocator.Malloc(setLayoutCount * int(unsafe.Sizeof([1]C.VkDescriptorSetLayout{}))))
		setLayoutSlice := ([]C.VkDescriptorSetLayout)(unsafe.Slice(setLayoutPtr, setLayoutCount))

		for i := 0; i < setLayoutCount; i++ {
			setLayoutSlice[i] = (C.VkDescriptorSetLayout)(unsafe.Pointer(o.SetLayouts[i].Handle()))
		}
		createInfo.pSetLayouts = setLayoutPtr
	}

	createInfo.pPushConstantRanges = nil
	if constantRangesCount > 0 {
		constantRangesPtr := (*C.VkPushConstantRange)(allocator.Malloc(constantRangesCount * C.sizeof_struct_VkPushConstantRange))
		constantRangesSlice := ([]C.VkPushConstantRange)(unsafe.Slice(constantRangesPtr, constantRangesCount))

		for i := 0; i < constantRangesCount; i++ {
			constantRangesSlice[i].stageFlags = C.VkShaderStageFlags(o.PushConstantRanges[i].StageFlags)
			constantRangesSlice[i].offset = C.uint32_t(o.PushConstantRanges[i].Offset)
			constantRangesSlice[i].size = C.uint32_t(o.PushConstantRanges[i].Size)
		}
		createInfo.pPushConstantRanges = constantRangesPtr
	}

	return preallocatedPointer, nil
}
