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

const (
	RateVertex   InputRate = C.VK_VERTEX_INPUT_RATE_VERTEX
	RateInstance InputRate = C.VK_VERTEX_INPUT_RATE_INSTANCE
)

func init() {
	RateVertex.Register("Vertex")
	RateInstance.Register("Instance")
}

type VertexInputBindingDescription struct {
	InputRate InputRate
	Binding   int
	Stride    int
}

type VertexInputAttributeDescription struct {
	Location uint32
	Binding  int
	Format   Format
	Offset   int
}

type PipelineVertexInputStateCreateInfo struct {
	VertexBindingDescriptions   []VertexInputBindingDescription
	VertexAttributeDescriptions []VertexInputAttributeDescription

	common.NextOptions
}

func (o PipelineVertexInputStateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineVertexInputStateCreateInfo)
	}
	createInfo := (*C.VkPipelineVertexInputStateCreateInfo)(preallocatedPointer)

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
