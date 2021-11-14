package core

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

type DynamicState int32

const (
	StateViewport                       DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT
	StateScissor                        DynamicState = C.VK_DYNAMIC_STATE_SCISSOR
	StateLineWidth                      DynamicState = C.VK_DYNAMIC_STATE_LINE_WIDTH
	StateDepthBias                      DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BIAS
	StateBlendConstants                 DynamicState = C.VK_DYNAMIC_STATE_BLEND_CONSTANTS
	StateDepthBounds                    DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BOUNDS
	StateStencilCompareMask             DynamicState = C.VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK
	StateStencilWriteMask               DynamicState = C.VK_DYNAMIC_STATE_STENCIL_WRITE_MASK
	StateStencilReference               DynamicState = C.VK_DYNAMIC_STATE_STENCIL_REFERENCE
	StateViewportWithScalingNV          DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV
	StateDiscardRectangleEXT            DynamicState = C.VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT
	StateSampleLocationsEXT             DynamicState = C.VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT
	StateRayTracingPipelineStackSizeKHR DynamicState = C.VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR
	StateViewportShadingRatePaletteNV   DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV
	StateViewportCoarseSampleOrderNV    DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV
	StateExclusiveScissorNV             DynamicState = C.VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV
	StateFragmentShadingRateKHR         DynamicState = C.VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR
	StateLineStippleEXT                 DynamicState = C.VK_DYNAMIC_STATE_LINE_STIPPLE_EXT
	StateCullModeEXT                    DynamicState = C.VK_DYNAMIC_STATE_CULL_MODE_EXT
	StateFrontFaceEXT                   DynamicState = C.VK_DYNAMIC_STATE_FRONT_FACE_EXT
	StatePrimitiveTopologyEXT           DynamicState = C.VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY_EXT
	StateViewportWithCountEXT           DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT_EXT
	StateScissorWithCountEXT            DynamicState = C.VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT_EXT
	StateVertexInputBindingStrideEXT    DynamicState = C.VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE_EXT
	StateDepthTestEnableEXT             DynamicState = C.VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE_EXT
	StateDepthWriteEnableEXT            DynamicState = C.VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE_EXT
	StateDepthCompareOpEXT              DynamicState = C.VK_DYNAMIC_STATE_DEPTH_COMPARE_OP_EXT
	StateDepthBoundsTestEnableEXT       DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE_EXT
	StateStencilTestEnableEXT           DynamicState = C.VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE_EXT
	StateStencilOpEXT                   DynamicState = C.VK_DYNAMIC_STATE_STENCIL_OP_EXT
	StateVertexInputEXT                 DynamicState = C.VK_DYNAMIC_STATE_VERTEX_INPUT_EXT
	StatePatchControlPointsEXT          DynamicState = C.VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT
	StateRasterizerDiscardEnableEXT     DynamicState = C.VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE_EXT
	StateDepthBiasEnableEXT             DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE_EXT
	StateLogicOpEXT                     DynamicState = C.VK_DYNAMIC_STATE_LOGIC_OP_EXT
	StatePrimitiveRestartEnableEXT      DynamicState = C.VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE_EXT
	StateColorWriteEnableEXT            DynamicState = C.VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT
)

var dynamicStateToString = map[DynamicState]string{
	StateViewport:                       "Viewport",
	StateScissor:                        "Scissor",
	StateLineWidth:                      "Line Width",
	StateDepthBias:                      "Depth Bias",
	StateBlendConstants:                 "Blend Constants",
	StateDepthBounds:                    "Depth Bounds",
	StateStencilCompareMask:             "Stencil Compare Mask",
	StateStencilWriteMask:               "Stencil Write Mask",
	StateStencilReference:               "Stencil Reference",
	StateViewportWithScalingNV:          "Viewport With Scaling (Nvidia Extension)",
	StateDiscardRectangleEXT:            "Discard Rectangle (Extension)",
	StateSampleLocationsEXT:             "Sample Locations (Extension)",
	StateRayTracingPipelineStackSizeKHR: "Ray Tracing Pipeline Stack Size (Khronos Extension)",
	StateViewportShadingRatePaletteNV:   "Viewport Shading Rate Palette (Nvidia Extension)",
	StateViewportCoarseSampleOrderNV:    "Viewport Coarse Sample Order (Nvidia Extension)",
	StateExclusiveScissorNV:             "Exclusive Scissor (Nvidia Extension)",
	StateFragmentShadingRateKHR:         "Fragment Shading Rate (Khronos Extension)",
	StateLineStippleEXT:                 "Line Stipple (Extension)",
	StateCullModeEXT:                    "Cull Mode (Extension)",
	StateFrontFaceEXT:                   "Front Face (Extension)",
	StatePrimitiveTopologyEXT:           "Primitive Topology (Extension)",
	StateViewportWithCountEXT:           "Viewport With Count (Extension)",
	StateScissorWithCountEXT:            "Scissor With Count (Extension)",
	StateVertexInputBindingStrideEXT:    "Vertex Input Binding State (Extension)",
	StateDepthTestEnableEXT:             "Depth Test Enable (Extension)",
	StateDepthWriteEnableEXT:            "Depth Write Enable (Extension)",
	StateDepthCompareOpEXT:              "Depth Compare Op (Extension)",
	StateDepthBoundsTestEnableEXT:       "Depth Bounds Test Enable (Extension)",
	StateStencilTestEnableEXT:           "Stencil Test Enable (Extension)",
	StateStencilOpEXT:                   "Stencil Op (Extension)",
	StateVertexInputEXT:                 "Vertex Input (Extension)",
	StatePatchControlPointsEXT:          "Patch Control Points (Extension)",
	StateRasterizerDiscardEnableEXT:     "Rasterizer Discard Enable (Extension)",
	StateDepthBiasEnableEXT:             "Depth Bias Enable (Extension)",
	StateLogicOpEXT:                     "Logic Op (Extension)",
	StatePrimitiveRestartEnableEXT:      "Primitive Restart Enable (Extension)",
	StateColorWriteEnableEXT:            "Color Write Enable (Extension)",
}

func (s DynamicState) String() string {
	return dynamicStateToString[s]
}

type DynamicStateOptions struct {
	DynamicStates []DynamicState

	common.HaveNext
}

func (o *DynamicStateOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineDynamicStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineDynamicStateCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next

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

	return unsafe.Pointer(createInfo), nil
}
