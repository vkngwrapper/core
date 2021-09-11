package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DynamicState int32

const (
	StateViewport                     DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT
	StateScissor                      DynamicState = C.VK_DYNAMIC_STATE_SCISSOR
	StateLineWidth                    DynamicState = C.VK_DYNAMIC_STATE_LINE_WIDTH
	StateDepthBias                    DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BIAS
	StateBlendConstants               DynamicState = C.VK_DYNAMIC_STATE_BLEND_CONSTANTS
	StateDepthBounds                  DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BOUNDS
	StateStencilCompareMask           DynamicState = C.VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK
	StateStencilWriteMask             DynamicState = C.VK_DYNAMIC_STATE_STENCIL_WRITE_MASK
	StateStencilReference             DynamicState = C.VK_DYNAMIC_STATE_STENCIL_REFERENCE
	StateViewportWithScalingNV        DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV
	StateDiscardRectangle             DynamicState = C.VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT
	StateSampleLocations              DynamicState = C.VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT
	StateRayTracingPipelineStackSize  DynamicState = C.VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR
	StateViewportShadingRatePaletteNV DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV
	StateViewportCoarseSampleOrderNV  DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV
	StateExclusiveScissorNV           DynamicState = C.VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV
	StateFragmentShadingRate          DynamicState = C.VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR
	StateLineStipple                  DynamicState = C.VK_DYNAMIC_STATE_LINE_STIPPLE_EXT
	StateCullMode                     DynamicState = C.VK_DYNAMIC_STATE_CULL_MODE_EXT
	StateFrontFace                    DynamicState = C.VK_DYNAMIC_STATE_FRONT_FACE_EXT
	StatePrimitiveTopology            DynamicState = C.VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY_EXT
	StateViewportWithCount            DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT_EXT
	StateScissorWithCount             DynamicState = C.VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT_EXT
	StateVertexInputBindingStride     DynamicState = C.VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE_EXT
	StateDepthTestEnable              DynamicState = C.VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE_EXT
	StateDepthWriteEnable             DynamicState = C.VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE_EXT
	StateDepthCompareOp               DynamicState = C.VK_DYNAMIC_STATE_DEPTH_COMPARE_OP_EXT
	StateDepthBoundsTestEnable        DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE_EXT
	StateStencilTestEnable            DynamicState = C.VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE_EXT
	StateStencilOp                    DynamicState = C.VK_DYNAMIC_STATE_STENCIL_OP_EXT
	StateVertexInput                  DynamicState = C.VK_DYNAMIC_STATE_VERTEX_INPUT_EXT
	StatePatchControlPoints           DynamicState = C.VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT
	StateRasterizerDiscardEnable      DynamicState = C.VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE_EXT
	StateDepthBiasEnable              DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE_EXT
	StateLogicOp                      DynamicState = C.VK_DYNAMIC_STATE_LOGIC_OP_EXT
	StatePrimitiveRestartEnable       DynamicState = C.VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE_EXT
	StateColorWriteEnable             DynamicState = C.VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT
)

var dynamicStateToString = map[DynamicState]string{
	StateViewport:                     "Viewport",
	StateScissor:                      "Scissor",
	StateLineWidth:                    "Line Width",
	StateDepthBias:                    "Depth Bias",
	StateBlendConstants:               "Blend Constants",
	StateDepthBounds:                  "Depth Bounds",
	StateStencilCompareMask:           "Stencil Compare Mask",
	StateStencilWriteMask:             "Stencil Write Mask",
	StateStencilReference:             "Stencil Reference",
	StateViewportWithScalingNV:        "Viewport With Scaling (Nvidia)",
	StateDiscardRectangle:             "Discard Rectangle",
	StateSampleLocations:              "Sample Locations",
	StateRayTracingPipelineStackSize:  "Ray Tracing Pipeline Stack Size",
	StateViewportShadingRatePaletteNV: "Viewport Shading Rate Palette (Nvidia)",
	StateViewportCoarseSampleOrderNV:  "Viewport Coarse Sample Order (Nvidia)",
	StateExclusiveScissorNV:           "Exclusive Scissor (Nvidia)",
	StateFragmentShadingRate:          "Fragment Shading Rate",
	StateLineStipple:                  "Line Stipple",
	StateCullMode:                     "Cull Mode",
	StateFrontFace:                    "Front Face",
	StatePrimitiveTopology:            "Primitive Topology",
	StateViewportWithCount:            "Viewport With Count",
	StateScissorWithCount:             "Scissor With Count",
	StateVertexInputBindingStride:     "Vertex Input Binding State",
	StateDepthTestEnable:              "Depth Test Enable",
	StateDepthWriteEnable:             "Depth Write Enable",
	StateDepthCompareOp:               "Depth Compare Op",
	StateDepthBoundsTestEnable:        "Depth Bounds Test Enable",
	StateStencilTestEnable:            "Stencil Test Enable",
	StateStencilOp:                    "Stencil Op",
	StateVertexInput:                  "Vertex Input",
	StatePatchControlPoints:           "Patch Control Points",
	StateRasterizerDiscardEnable:      "Rasterizer Discard Enable",
	StateDepthBiasEnable:              "Depth Bias Enable",
	StateLogicOp:                      "Logic Op",
	StatePrimitiveRestartEnable:       "Primitive Restart Enable",
	StateColorWriteEnable:             "Color Write Enable",
}

func (s DynamicState) String() string {
	return dynamicStateToString[s]
}

type DynamicStateOptions struct {
	DynamicStates []DynamicState

	Next core.Options
}

func (o *DynamicStateOptions) AllocForC(allocator *cgoparam.Allocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineDynamicStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineDynamicStateCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO
	createInfo.flags = 0

	stateCount := len(o.DynamicStates)
	createInfo.dynamicStateCount = C.uint32_t(stateCount)

	if stateCount > 0 {
		statesPtr := (*C.VkDynamicState)(allocator.Malloc(stateCount * int(unsafe.Sizeof(C.VkDynamicState(0)))))
		stateSlice := ([]C.VkDynamicState)(unsafe.Slice(statesPtr, stateCount))

		for i := 0; i < stateCount; i++ {
			stateSlice[i] = C.VkDynamicState(o.DynamicStates[i])
		}

		createInfo.pDynamicStates = statesPtr
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
