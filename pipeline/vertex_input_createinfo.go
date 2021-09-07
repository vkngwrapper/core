package pipeline

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
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
	Stride    uintptr
}

type VertexAttributeDescription struct {
	Location uint32
	Binding  int
	Format   core.DataFormat
	Offset   uintptr
}

type VertexInputOptions struct {
	VertexBindingDescriptions   []VertexBindingDescription
	VertexAttributeDescriptions []VertexAttributeDescription

	Next core.Options
}

func (o *VertexInputOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineVertexInputStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineVertexInputStateCreateInfo))

	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO

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
