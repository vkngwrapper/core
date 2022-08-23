package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

const (
	// VertexInputRateVertex specifies that vertex attribute addressing is a function of the vertex index
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputRate.html
	VertexInputRateVertex VertexInputRate = C.VK_VERTEX_INPUT_RATE_VERTEX
	// VertexInputRateInstance specifies that vertex attribute addressing is a function of the
	// instance index
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputRate.html
	VertexInputRateInstance VertexInputRate = C.VK_VERTEX_INPUT_RATE_INSTANCE
)

func init() {
	VertexInputRateVertex.Register("Vertex")
	VertexInputRateInstance.Register("Instance")
}

// VertexInputBindingDescription specifies a vertex input binding description
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription.html
type VertexInputBindingDescription struct {
	// InputRate specifies whether vertex attribute addressing is a function of the vertex index
	// or the instance index
	InputRate VertexInputRate
	// Binding isthe bidning number that this structure describes
	Binding int
	// Stride is the byte stride between consecutive elements within the buffer
	Stride int
}

// VertexInputAttributeDescription specifies a vertex input attribute description
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription.html
type VertexInputAttributeDescription struct {
	// Location is the shader input location number for this attribute
	Location uint32
	// Binding is the binding number which this attribute takes its data from
	Binding int
	// Format is the size and type of the vertex attribute data
	Format Format
	// Offset is a byte offset of this attribute relative to the start of an element in the vertex
	// input binding
	Offset int
}

// PipelineVertexInputStateCreateInfo specifies parameters of a newly-created Pipeline vertex input state
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html
type PipelineVertexInputStateCreateInfo struct {
	// VertexBindingDescriptions is a slice of VertexInputBindingDescription structures
	VertexBindingDescriptions []VertexInputBindingDescription
	// VertexAttributeDescriptions is a slice of VertexInputAttributeDescription structures
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
