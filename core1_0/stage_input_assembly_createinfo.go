package core1_0

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

const (
	TopologyPointList                  common.PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_POINT_LIST
	TopologyLineList                   common.PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_LIST
	TopologyLineStrip                  common.PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_STRIP
	TopologyTriangleList               common.PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST
	TopologyTriangleStrip              common.PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP
	TopologyTriangleFan                common.PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN
	TopologyLineListWithAdjacency      common.PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY
	TopologyLineStripWithAdjacency     common.PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY
	TopologyTriangleListWithAdjacency  common.PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY
	TopologyTriangleStripWithAdjacency common.PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY
	TopologyPatchlist                  common.PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_PATCH_LIST
)

func init() {
	TopologyPointList.Register("Point List")
	TopologyLineList.Register("Line List")
	TopologyLineStrip.Register("Line Strip")
	TopologyTriangleList.Register("Triangle List")
	TopologyTriangleStrip.Register("Triangle Strip")
	TopologyTriangleFan.Register("Triangle Fan")
	TopologyLineListWithAdjacency.Register("Line List w/ Adjacency")
	TopologyLineStripWithAdjacency.Register("Line Strip w/ Adjacency")
	TopologyTriangleListWithAdjacency.Register("Triangle List w/ Adjacency")
	TopologyTriangleStripWithAdjacency.Register("Triangle Strip w/ Adjacency")
	TopologyPatchlist.Register("Patch List")
}

type InputAssemblyStateOptions struct {
	Topology               common.PrimitiveTopology
	EnablePrimitiveRestart bool

	common.HaveNext
}

func (o InputAssemblyStateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineInputAssemblyStateCreateInfo)
	}
	createInfo := (*C.VkPipelineInputAssemblyStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
	createInfo.topology = C.VkPrimitiveTopology(o.Topology)
	createInfo.primitiveRestartEnable = C.VK_FALSE

	if o.EnablePrimitiveRestart {
		createInfo.primitiveRestartEnable = C.VK_TRUE
	}

	return preallocatedPointer, nil
}

func (o InputAssemblyStateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPipelineInputAssemblyStateCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
