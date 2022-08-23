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

// PipelineInputAssemblyStateCreateFlags reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateFlags.html
type PipelineInputAssemblyStateCreateFlags uint32

var pipelineInputAssemblyStateCreateFlagsMapping = common.NewFlagStringMapping[PipelineInputAssemblyStateCreateFlags]()

func (f PipelineInputAssemblyStateCreateFlags) Register(str string) {
	pipelineInputAssemblyStateCreateFlagsMapping.Register(f, str)
}

func (f PipelineInputAssemblyStateCreateFlags) String() string {
	return pipelineInputAssemblyStateCreateFlagsMapping.FlagsToString(f)
}

////

const (
	// PrimitiveTopologyPointList specifies a series of separate point primitives
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
	PrimitiveTopologyPointList PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_POINT_LIST
	// PrimitiveTopologyLineList specifies a series of separate line primitives
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
	PrimitiveTopologyLineList PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_LIST
	// PrimitiveTopologyLineStrip specifies a series of connected line primitives with consecutive
	// lines sharing a vertex
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
	PrimitiveTopologyLineStrip PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_STRIP
	// PrimitiveTopologyTriangleList specifies a series of separate triangle primitives
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
	PrimitiveTopologyTriangleList PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST
	// PrimitiveTopologyTriangleStrip specifies a series of connected triangle primitives with
	// consecutive triangles sharing an edge
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
	PrimitiveTopologyTriangleStrip PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP
	// PrimitiveTopologyTriangleFan specifies a series of connected triangle primitives with all triangles
	// sharing a common vertex
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
	PrimitiveTopologyTriangleFan PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN
	// PrimitiveTopologyLineListWithAdjacency specifies a series of separate line primitives with adjacency
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
	PrimitiveTopologyLineListWithAdjacency PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY
	// PrimitiveTopologyLineStripWithAdjacency specifies a series of connected line primitives
	// with adjacency, with consecutive primitives sharing three vertices
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
	PrimitiveTopologyLineStripWithAdjacency PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY
	// PrimitiveTopologyTriangleListWithAdjacency specifies a series of separate triangle primitives
	// with adjacency
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
	PrimitiveTopologyTriangleListWithAdjacency PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY
	// PrimitiveTopologyTriangleStripWithAdjacency specifies connected triangle primitives with
	// adjacency, with consecutive triangles sharing an edge
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
	PrimitiveTopologyTriangleStripWithAdjacency PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY
	// PrimitiveTopologyPatchList specifies separate patch primitives
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrimitiveTopology.html
	PrimitiveTopologyPatchList PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_PATCH_LIST
)

func init() {
	PrimitiveTopologyPointList.Register("Point List")
	PrimitiveTopologyLineList.Register("Line List")
	PrimitiveTopologyLineStrip.Register("Line Strip")
	PrimitiveTopologyTriangleList.Register("Triangle List")
	PrimitiveTopologyTriangleStrip.Register("Triangle Strip")
	PrimitiveTopologyTriangleFan.Register("Triangle Fan")
	PrimitiveTopologyLineListWithAdjacency.Register("Line List w/ Adjacency")
	PrimitiveTopologyLineStripWithAdjacency.Register("Line Strip w/ Adjacency")
	PrimitiveTopologyTriangleListWithAdjacency.Register("Triangle List w/ Adjacency")
	PrimitiveTopologyTriangleStripWithAdjacency.Register("Triangle Strip w/ Adjacency")
	PrimitiveTopologyPatchList.Register("Patch List")
}

// PipelineInputAssemblyStateCreateInfo specifies parameters of a newly-created Pipeline input
// assembly state
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateInfo.html
type PipelineInputAssemblyStateCreateInfo struct {
	// Flags is reserved for future use
	Flags PipelineInputAssemblyStateCreateFlags
	// Topology defines the primitive topology
	Topology PrimitiveTopology
	// PrimitiveRestartEnable controls whether a special vertex index value is treated as
	// restarting the assembly of primitives
	PrimitiveRestartEnable bool

	common.NextOptions
}

func (o PipelineInputAssemblyStateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineInputAssemblyStateCreateInfo)
	}
	createInfo := (*C.VkPipelineInputAssemblyStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO
	createInfo.flags = C.VkPipelineInputAssemblyStateCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.topology = C.VkPrimitiveTopology(o.Topology)
	createInfo.primitiveRestartEnable = C.VK_FALSE

	if o.PrimitiveRestartEnable {
		createInfo.primitiveRestartEnable = C.VK_TRUE
	}

	return preallocatedPointer, nil
}
