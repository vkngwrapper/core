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

type InputRate int32

const (
	RateVertex   InputRate = C.VK_VERTEX_INPUT_RATE_VERTEX
	RateInstance InputRate = C.VK_VERTEX_INPUT_RATE_INSTANCE
)

var inputRateToString = map[InputRate]string{
	RateVertex:   "Vertex",
	RateInstance: "Instance",
}

func (r InputRate) String() string {
	return inputRateToString[r]
}

type VertexBindingDescription struct {
	InputRate InputRate
	Binding   int
	Stride    int
}

type VertexAttributeDescription struct {
	Location uint32
	Binding  int
	Format   common.DataFormat
	Offset   int
}

type VertexInputOptions struct {
	VertexBindingDescriptions   []VertexBindingDescription
	VertexAttributeDescriptions []VertexAttributeDescription

	common.HaveNext
}

func (o *VertexInputOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineVertexInputStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineVertexInputStateCreateInfo))

	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next

	bindingCount := len(o.VertexBindingDescriptions)
	attributeCount := len(o.VertexAttributeDescriptions)
	createInfo.vertexBindingDescriptionCount = C.uint(bindingCount)
	createInfo.pVertexBindingDescriptions = nil
	createInfo.vertexAttributeDescriptionCount = C.uint(attributeCount)
	createInfo.pVertexAttributeDescriptions = nil

	if bindingCount > 0 {
		bindingPtr := (*C.VkVertexInputBindingDescription)(allocator.Malloc(bindingCount * C.sizeof_struct_VkVertexInputBindingDescription))
		bindingSlice := ([]C.VkVertexInputBindingDescription)(unsafe.Slice(bindingPtr, bindingCount))

		for i := 0; i < bindingCount; i++ {
			bindingSlice[i].binding = C.uint32_t(o.VertexBindingDescriptions[i].Binding)
			bindingSlice[i].stride = C.uint32_t(o.VertexBindingDescriptions[i].Stride)
			bindingSlice[i].inputRate = C.VkVertexInputRate(o.VertexBindingDescriptions[i].InputRate)
		}
		createInfo.pVertexBindingDescriptions = bindingPtr
	}

	if attributeCount > 0 {
		attributePtr := (*C.VkVertexInputAttributeDescription)(allocator.Malloc(attributeCount * C.sizeof_struct_VkVertexInputAttributeDescription))
		attributeSlice := ([]C.VkVertexInputAttributeDescription)(unsafe.Slice(attributePtr, attributeCount))

		for i := 0; i < attributeCount; i++ {
			attributeSlice[i].location = C.uint32_t(o.VertexAttributeDescriptions[i].Location)
			attributeSlice[i].binding = C.uint32_t(o.VertexAttributeDescriptions[i].Binding)
			attributeSlice[i].format = C.VkFormat(o.VertexAttributeDescriptions[i].Format)
			attributeSlice[i].offset = C.uint32_t(o.VertexAttributeDescriptions[i].Offset)
		}
		createInfo.pVertexAttributeDescriptions = attributePtr
	}

	return unsafe.Pointer(createInfo), nil
}
