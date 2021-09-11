package pipeline

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type PolygonMode int32

const (
	ModeFill            PolygonMode = C.VK_POLYGON_MODE_FILL
	ModeLine            PolygonMode = C.VK_POLYGON_MODE_LINE
	ModePoint           PolygonMode = C.VK_POLYGON_MODE_POINT
	ModeFillRectangleNV PolygonMode = C.VK_POLYGON_MODE_FILL_RECTANGLE_NV
)

var polygonModeToString = map[PolygonMode]string{
	ModeFill:            "Fill",
	ModeLine:            "Line",
	ModePoint:           "Point",
	ModeFillRectangleNV: "Fill Rectangle (Nvidia)",
}

func (m PolygonMode) String() string {
	return polygonModeToString[m]
}

type RasterizationOptions struct {
	DepthClamp        bool
	RasterizerDiscard bool

	PolygonMode PolygonMode
	CullMode    core.CullModes
	FrontFace   core.FrontFace

	DepthBias               bool
	DepthBiasClamp          float32
	DepthBiasConstantFactor float32
	DepthBiasSlopeFactor    float32

	LineWidth float32

	Next core.Options
}

func (o *RasterizationOptions) AllocForC(allocator *cgoparam.Allocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineRasterizationStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineRasterizationStateCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.depthClampEnable = C.VK_FALSE
	createInfo.rasterizerDiscardEnable = C.VK_FALSE
	createInfo.depthBiasEnable = C.VK_FALSE

	if o.DepthClamp {
		createInfo.depthClampEnable = C.VK_TRUE
	}

	if o.RasterizerDiscard {
		createInfo.rasterizerDiscardEnable = C.VK_TRUE
	}

	if o.DepthBias {
		createInfo.depthBiasEnable = C.VK_TRUE
	}

	createInfo.polygonMode = C.VkPolygonMode(o.PolygonMode)
	createInfo.cullMode = C.VkCullModeFlags(o.CullMode)
	createInfo.frontFace = C.VkFrontFace(o.FrontFace)

	createInfo.depthBiasClamp = C.float(o.DepthBiasClamp)
	createInfo.depthBiasConstantFactor = C.float(o.DepthBiasConstantFactor)
	createInfo.depthBiasSlopeFactor = C.float(o.DepthBiasSlopeFactor)

	createInfo.lineWidth = C.float(o.LineWidth)

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
