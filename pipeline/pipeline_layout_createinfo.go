package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type PipelineLayoutOptions struct {
	SetLayouts         []*DescriptorSetLayout
	PushConstantRanges []*core.PushConstantRange

	Next core.Options
}

func (o *PipelineLayoutOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineLayoutCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineLayoutCreateInfo))

	setLayoutCount := len(o.SetLayouts)
	constantRangesCount := len(o.PushConstantRanges)

	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
	createInfo.flags = 0
	createInfo.setLayoutCount = C.uint32_t(setLayoutCount)
	createInfo.pushConstantRangeCount = C.uint32_t(constantRangesCount)

	createInfo.pSetLayouts = nil
	if setLayoutCount > 0 {
		setLayoutPtr := (*C.VkDescriptorSetLayout)(allocator.Malloc(setLayoutCount * int(unsafe.Sizeof([1]C.VkDescriptorSetLayout{}))))
		setLayoutSlice := ([]C.VkDescriptorSetLayout)(unsafe.Slice(setLayoutPtr, setLayoutCount))

		for i := 0; i < setLayoutCount; i++ {
			setLayoutSlice[i] = (C.VkDescriptorSetLayout)(unsafe.Pointer(o.SetLayouts[i].handle))
		}
		createInfo.pSetLayouts = setLayoutPtr
	}

	createInfo.pPushConstantRanges = nil
	if constantRangesCount > 0 {
		constantRangesPtr := (*C.VkPushConstantRange)(allocator.Malloc(constantRangesCount * C.sizeof_struct_VkPushConstantRange))
		constantRangesSlice := ([]C.VkPushConstantRange)(unsafe.Slice(constantRangesPtr, constantRangesCount))

		for i := 0; i < constantRangesCount; i++ {
			constantRangesSlice[i].stageFlags = C.VkShaderStageFlags(o.PushConstantRanges[i].Stages)
			constantRangesSlice[i].offset = C.uint32_t(o.PushConstantRanges[i].Offset)
			constantRangesSlice[i].size = C.uint32_t(o.PushConstantRanges[i].Size)
		}
	}

	var err error
	var next unsafe.Pointer
	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}

	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
