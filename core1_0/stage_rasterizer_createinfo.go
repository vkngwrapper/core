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
	PolygonModeFill  PolygonMode = C.VK_POLYGON_MODE_FILL
	PolygonModeLine  PolygonMode = C.VK_POLYGON_MODE_LINE
	PolygonModePoint PolygonMode = C.VK_POLYGON_MODE_POINT

	CullFront CullModes = C.VK_CULL_MODE_FRONT_BIT
	CullBack  CullModes = C.VK_CULL_MODE_BACK_BIT

	FrontFaceCounterClockwise FrontFace = C.VK_FRONT_FACE_COUNTER_CLOCKWISE
	FrontFaceClockwise        FrontFace = C.VK_FRONT_FACE_CLOCKWISE
)

func init() {
	PolygonModeFill.Register("Fill")
	PolygonModeLine.Register("Line")
	PolygonModePoint.Register("Point")

	CullFront.Register("Front")
	CullBack.Register("Back")

	FrontFaceCounterClockwise.Register("Counter-Clockwise")
	FrontFaceClockwise.Register("Clockwise")
}

type RasterizationStateOptions struct {
	DepthClamp        bool
	RasterizerDiscard bool

	PolygonMode PolygonMode
	CullMode    CullModes
	FrontFace   FrontFace

	DepthBias               bool
	DepthBiasClamp          float32
	DepthBiasConstantFactor float32
	DepthBiasSlopeFactor    float32

	LineWidth float32

	common.NextOptions
}

func (o RasterizationStateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineRasterizationStateCreateInfo)
	}
	createInfo := (*C.VkPipelineRasterizationStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next
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

	return preallocatedPointer, nil
}
