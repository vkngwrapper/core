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
	DynamicStateViewport           DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT
	DynamicStateScissor            DynamicState = C.VK_DYNAMIC_STATE_SCISSOR
	DynamicStateLineWidth          DynamicState = C.VK_DYNAMIC_STATE_LINE_WIDTH
	DynamicStateDepthBias          DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BIAS
	DynamicStateBlendConstants     DynamicState = C.VK_DYNAMIC_STATE_BLEND_CONSTANTS
	DynamicStateDepthBounds        DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BOUNDS
	DynamicStateStencilCompareMask DynamicState = C.VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK
	DynamicStateStencilWriteMask   DynamicState = C.VK_DYNAMIC_STATE_STENCIL_WRITE_MASK
	DynamicStateStencilReference   DynamicState = C.VK_DYNAMIC_STATE_STENCIL_REFERENCE
)

func init() {
	DynamicStateViewport.Register("Viewport")
	DynamicStateScissor.Register("Scissor")
	DynamicStateLineWidth.Register("Line Width")
	DynamicStateDepthBias.Register("Depth Bias")
	DynamicStateBlendConstants.Register("Blend Constants")
	DynamicStateDepthBounds.Register("Depth Bounds")
	DynamicStateStencilCompareMask.Register("Stencil Compare Mask")
	DynamicStateStencilWriteMask.Register("Stencil Write Mask")
	DynamicStateStencilReference.Register("Stencil Reference")
}

type DynamicStateOptions struct {
	DynamicStates []DynamicState

	common.HaveNext
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

func (o DynamicStateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPipelineDynamicStateCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
