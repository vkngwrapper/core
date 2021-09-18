package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type PrimitiveTopology int32

const (
	TopologyPointList                  PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_POINT_LIST
	TopologyLineList                   PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_LIST
	TopologyLineStrip                  PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_STRIP
	TopologyTriangleList               PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST
	TopologyTriangleStrip              PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP
	TopologyTriangleFan                PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN
	TopologyLineListWithAdjacency      PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY
	TopologyLineStripWithAdjacency     PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY
	TopologyTriangleListWithAdjacency  PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY
	TopologyTriangleStripWithAdjacency PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY
	TopologyPatchlist                  PrimitiveTopology = C.VK_PRIMITIVE_TOPOLOGY_PATCH_LIST
)

var primitiveTopologyToString = map[PrimitiveTopology]string{
	TopologyPointList:                  "Point List",
	TopologyLineList:                   "Line List",
	TopologyLineStrip:                  "Line Strip",
	TopologyTriangleList:               "Triangle List",
	TopologyTriangleStrip:              "Triangle Strip",
	TopologyTriangleFan:                "Triangle Fan",
	TopologyLineListWithAdjacency:      "Line List w/ Adjacency",
	TopologyLineStripWithAdjacency:     "Line Strip w/ Adjacency",
	TopologyTriangleListWithAdjacency:  "Triangle List w/ Adjacency",
	TopologyTriangleStripWithAdjacency: "Triangle Strip w/ Adjacency",
	TopologyPatchlist:                  "Patch List",
}

func (t PrimitiveTopology) String() string {
	return primitiveTopologyToString[t]
}
