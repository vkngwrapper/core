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
	PrimitiveTopologyPointList                  PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_POINT_LIST
	PrimitiveTopologyLineList                   PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_LIST
	PrimitiveTopologyLineStrip                  PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_STRIP
	PrimitiveTopologyTriangleList               PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST
	PrimitiveTopologyTriangleStrip              PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP
	PrimitiveTopologyTriangleFan                PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN
	PrimitiveTopologyLineListWithAdjacency      PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY
	PrimitiveTopologyLineStripWithAdjacency     PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY
	PrimitiveTopologyTriangleListWithAdjacency  PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY
	PrimitiveTopologyTriangleStripWithAdjacency PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY
	PrimitiveTopologyPatchlist                  PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_PATCH_LIST
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
	PrimitiveTopologyPatchlist.Register("Patch List")
}

type PipelineInputAssemblyStateCreateInfo struct {
	Flags                  PipelineInputAssemblyStateCreateFlags
	Topology               PrimitiveTopology
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
