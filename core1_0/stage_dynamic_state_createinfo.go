package core1_0

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
	DynamicStateViewport                       DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT
	DynamicStateScissor                        DynamicState = C.VK_DYNAMIC_STATE_SCISSOR
	DynamicStateLineWidth                      DynamicState = C.VK_DYNAMIC_STATE_LINE_WIDTH
	DynamicStateDepthBias                      DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BIAS
	DynamicStateBlendConstants                 DynamicState = C.VK_DYNAMIC_STATE_BLEND_CONSTANTS
	DynamicStateDepthBounds                    DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BOUNDS
	DynamicStateStencilCompareMask             DynamicState = C.VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK
	DynamicStateStencilWriteMask               DynamicState = C.VK_DYNAMIC_STATE_STENCIL_WRITE_MASK
	DynamicStateStencilReference               DynamicState = C.VK_DYNAMIC_STATE_STENCIL_REFERENCE
	DynamicStateViewportWithScalingNV          DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV
	DynamicStateDiscardRectangleEXT            DynamicState = C.VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT
	DynamicStateSampleLocationsEXT             DynamicState = C.VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT
	DynamicStateRayTracingPipelineStackSizeKHR DynamicState = C.VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR
	DynamicStateViewportShadingRatePaletteNV   DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV
	DynamicStateViewportCoarseSampleOrderNV    DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV
	DynamicStateExclusiveScissorNV             DynamicState = C.VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV
	DynamicStateFragmentShadingRateKHR         DynamicState = C.VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR
	DynamicStateLineStippleEXT                 DynamicState = C.VK_DYNAMIC_STATE_LINE_STIPPLE_EXT
	DynamicStateCullModeEXT                    DynamicState = C.VK_DYNAMIC_STATE_CULL_MODE_EXT
	DynamicStateFrontFaceEXT                   DynamicState = C.VK_DYNAMIC_STATE_FRONT_FACE_EXT
	DynamicStatePrimitiveTopologyEXT           DynamicState = C.VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY_EXT
	DynamicStateViewportWithCountEXT           DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT_EXT
	DynamicStateScissorWithCountEXT            DynamicState = C.VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT_EXT
	DynamicStateVertexInputBindingStrideEXT    DynamicState = C.VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE_EXT
	DynamicStateDepthTestEnableEXT             DynamicState = C.VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE_EXT
	DynamicStateDepthWriteEnableEXT            DynamicState = C.VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE_EXT
	DynamicStateDepthCompareOpEXT              DynamicState = C.VK_DYNAMIC_STATE_DEPTH_COMPARE_OP_EXT
	DynamicStateDepthBoundsTestEnableEXT       DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE_EXT
	DynamicStateStencilTestEnableEXT           DynamicState = C.VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE_EXT
	DynamicStateStencilOpEXT                   DynamicState = C.VK_DYNAMIC_STATE_STENCIL_OP_EXT
	DynamicStateVertexInputEXT                 DynamicState = C.VK_DYNAMIC_STATE_VERTEX_INPUT_EXT
	DynamicStatePatchControlPointsEXT          DynamicState = C.VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT
	DynamicStateRasterizerDiscardEnableEXT     DynamicState = C.VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE_EXT
	DynamicStateDepthBiasEnableEXT             DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE_EXT
	DynamicStateLogicOpEXT                     DynamicState = C.VK_DYNAMIC_STATE_LOGIC_OP_EXT
	DynamicStatePrimitiveRestartEnableEXT      DynamicState = C.VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE_EXT
	DynamicStateColorWriteEnableEXT            DynamicState = C.VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT
)

var dynamicStateToString = map[DynamicState]string{
	DynamicStateViewport:                       "Viewport",
	DynamicStateScissor:                        "Scissor",
	DynamicStateLineWidth:                      "Line Width",
	DynamicStateDepthBias:                      "Depth Bias",
	DynamicStateBlendConstants:                 "Blend Constants",
	DynamicStateDepthBounds:                    "Depth Bounds",
	DynamicStateStencilCompareMask:             "Stencil Compare Mask",
	DynamicStateStencilWriteMask:               "Stencil Write Mask",
	DynamicStateStencilReference:               "Stencil Reference",
	DynamicStateViewportWithScalingNV:          "Viewport With Scaling (Nvidia Extension)",
	DynamicStateDiscardRectangleEXT:            "Discard Rectangle (Extension)",
	DynamicStateSampleLocationsEXT:             "Sample Locations (Extension)",
	DynamicStateRayTracingPipelineStackSizeKHR: "Ray Tracing Pipeline Stack Size (Khronos Extension)",
	DynamicStateViewportShadingRatePaletteNV:   "Viewport Shading Rate Palette (Nvidia Extension)",
	DynamicStateViewportCoarseSampleOrderNV:    "Viewport Coarse Sample Order (Nvidia Extension)",
	DynamicStateExclusiveScissorNV:             "Exclusive Scissor (Nvidia Extension)",
	DynamicStateFragmentShadingRateKHR:         "Fragment Shading Rate (Khronos Extension)",
	DynamicStateLineStippleEXT:                 "Line Stipple (Extension)",
	DynamicStateCullModeEXT:                    "Cull Mode (Extension)",
	DynamicStateFrontFaceEXT:                   "Front Face (Extension)",
	DynamicStatePrimitiveTopologyEXT:           "Primitive Topology (Extension)",
	DynamicStateViewportWithCountEXT:           "Viewport With Count (Extension)",
	DynamicStateScissorWithCountEXT:            "Scissor With Count (Extension)",
	DynamicStateVertexInputBindingStrideEXT:    "Vertex Input Binding State (Extension)",
	DynamicStateDepthTestEnableEXT:             "Depth Test Enable (Extension)",
	DynamicStateDepthWriteEnableEXT:            "Depth Write Enable (Extension)",
	DynamicStateDepthCompareOpEXT:              "Depth Compare Op (Extension)",
	DynamicStateDepthBoundsTestEnableEXT:       "Depth Bounds Test Enable (Extension)",
	DynamicStateStencilTestEnableEXT:           "Stencil Test Enable (Extension)",
	DynamicStateStencilOpEXT:                   "Stencil Op (Extension)",
	DynamicStateVertexInputEXT:                 "Vertex Input (Extension)",
	DynamicStatePatchControlPointsEXT:          "Patch Control Points (Extension)",
	DynamicStateRasterizerDiscardEnableEXT:     "Rasterizer Discard Enable (Extension)",
	DynamicStateDepthBiasEnableEXT:             "Depth Bias Enable (Extension)",
	DynamicStateLogicOpEXT:                     "Logic Op (Extension)",
	DynamicStatePrimitiveRestartEnableEXT:      "Primitive Restart Enable (Extension)",
	DynamicStateColorWriteEnableEXT:            "Color Write Enable (Extension)",
}

func (s DynamicState) String() string {
	return dynamicStateToString[s]
}

type DynamicStateOptions struct {
	DynamicStates []DynamicState

	core.HaveNext
}

func (o DynamicStateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineDynamicStateCreateInfo)
	}
	createInfo := (*C.VkPipelineDynamicStateCreateInfo)(preallocatedPointer)
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

	return preallocatedPointer, nil
}
