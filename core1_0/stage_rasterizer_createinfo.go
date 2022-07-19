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

// PipelineRasterizationStateCreateFlags is reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateFlags.html
type PipelineRasterizationStateCreateFlags uint32

var pipelineRasterizationStateCreateFlagsMapping = common.NewFlagStringMapping[PipelineRasterizationStateCreateFlags]()

func (f PipelineRasterizationStateCreateFlags) Register(str string) {
	pipelineRasterizationStateCreateFlagsMapping.Register(f, str)
}

func (f PipelineRasterizationStateCreateFlags) String() string {
	return pipelineRasterizationStateCreateFlagsMapping.FlagsToString(f)
}

////

const (
	// PolygonModeFill specifies that polygons are rendered using the polygon rasterization rules
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPolygonMode.html
	PolygonModeFill PolygonMode = C.VK_POLYGON_MODE_FILL
	// PolygonModeLine specifies that polygon edges are drawn as line segments
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPolygonMode.html
	PolygonModeLine PolygonMode = C.VK_POLYGON_MODE_LINE
	// PolygonModePoint specifies that polygon vertices are drawn as points
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPolygonMode.html
	PolygonModePoint PolygonMode = C.VK_POLYGON_MODE_POINT

	// CullModeFront specifies that front-facing triangles are discarded
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCullModeFlagBits.html
	CullModeFront CullModeFlags = C.VK_CULL_MODE_FRONT_BIT
	// CullModeBack specifies that back-facing triangles are discarded
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCullModeFlagBits.html
	CullModeBack CullModeFlags = C.VK_CULL_MODE_BACK_BIT

	// FrontFaceCounterClockwise specifies that a triangle with positive area is considered
	// front-facing
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrontFace.html
	FrontFaceCounterClockwise FrontFace = C.VK_FRONT_FACE_COUNTER_CLOCKWISE
	// FrontFaceClockwise specifies that a triangle with negative area is considered front-facing
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrontFace.html
	FrontFaceClockwise FrontFace = C.VK_FRONT_FACE_CLOCKWISE
)

func init() {
	PolygonModeFill.Register("Fill")
	PolygonModeLine.Register("Line")
	PolygonModePoint.Register("Point")

	CullModeFront.Register("Front")
	CullModeBack.Register("Back")

	FrontFaceCounterClockwise.Register("Counter-Clockwise")
	FrontFaceClockwise.Register("Clockwise")
}

// PipelineRasterizationStateCreateInfo specifies parameters of a newly-created Pipeline
// rasterization state
type PipelineRasterizationStateCreateInfo struct {
	// Flags is reserved for future use
	Flags PipelineRasterizationStateCreateFlags
	// DepthClampEnable controls whether to clamp the fragment's depth values
	DepthClampEnable bool
	// RasterizerDiscardEnable controls whether primitives are discarded immediately before the
	// rasterization stage
	RasterizerDiscardEnable bool

	// PolygonMode is the triangle rendering mode
	PolygonMode PolygonMode
	// CullMode is the triangle facing direction used for primitive culling
	CullMode CullModeFlags
	// FrontFace specifies the front-facing triangle orientation to be used for culling
	FrontFace FrontFace

	// DepthBiasEnable controls whether to bias fragment depth values
	DepthBiasEnable bool
	// DepthBiasClamp is the maximum (or minimum) depth bias of a fragment
	DepthBiasClamp float32
	// DepthBiasConstantFactor is a scalar factor controlling the constant depth value added
	// to each fragment
	DepthBiasConstantFactor float32
	// DepthBiasSlopeFactor is a scalar factor applied to a fragment's slope in depth bias
	// calculations
	DepthBiasSlopeFactor float32

	// LineWidth is the width of rasterized line segments
	LineWidth float32

	common.NextOptions
}

func (o PipelineRasterizationStateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineRasterizationStateCreateInfo)
	}
	createInfo := (*C.VkPipelineRasterizationStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO
	createInfo.flags = C.VkPipelineRasterizationStateCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.depthClampEnable = C.VK_FALSE
	createInfo.rasterizerDiscardEnable = C.VK_FALSE
	createInfo.depthBiasEnable = C.VK_FALSE

	if o.DepthClampEnable {
		createInfo.depthClampEnable = C.VK_TRUE
	}

	if o.RasterizerDiscardEnable {
		createInfo.rasterizerDiscardEnable = C.VK_TRUE
	}

	if o.DepthBiasEnable {
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
