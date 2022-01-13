package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type PipelineLayoutOptions struct {
	SetLayouts         []DescriptorSetLayout
	PushConstantRanges []common.PushConstantRange

	common.HaveNext
}

func (o *PipelineLayoutOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineLayoutCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineLayoutCreateInfo))

	setLayoutCount := len(o.SetLayouts)
	constantRangesCount := len(o.PushConstantRanges)

	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
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
			constantRangesSlice[i].stageFlags = C.VkShaderStageFlags(o.PushConstantRanges[i].Stages)
			constantRangesSlice[i].offset = C.uint32_t(o.PushConstantRanges[i].Offset)
			constantRangesSlice[i].size = C.uint32_t(o.PushConstantRanges[i].Size)
		}
		createInfo.pPushConstantRanges = constantRangesPtr
	}

	return unsafe.Pointer(createInfo), nil
}

type vulkanPipelineLayout struct {
	driver Driver
	device VkDevice
	handle VkPipelineLayout
}

func (l *vulkanPipelineLayout) Handle() VkPipelineLayout {
	return l.handle
}

func (l *vulkanPipelineLayout) Destroy(callbacks *AllocationCallbacks) {
	l.driver.VkDestroyPipelineLayout(l.device, l.handle, nil)
}
