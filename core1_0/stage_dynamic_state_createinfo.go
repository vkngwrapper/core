package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// PipelineDynamicStateCreateFlags is reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateFlags.html
type PipelineDynamicStateCreateFlags uint32

var pipelineDynamicStateCreateFlagsMapping = common.NewFlagStringMapping[PipelineDynamicStateCreateFlags]()

func (f PipelineDynamicStateCreateFlags) Register(str string) {
	pipelineDynamicStateCreateFlagsMapping.Register(f, str)
}

func (f PipelineDynamicStateCreateFlags) String() string {
	return pipelineDynamicStateCreateFlagsMapping.FlagsToString(f)
}

////

const (
	// DynamicStateViewport specifies that PipelineViewportStateCreateInfo.Viewports will be ignored
	// and must be set dynamically with CommandBuffer.CmdSetViewport before any drawing commands
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html
	DynamicStateViewport DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT
	// DynamicStateScissor specifies that PipelineViewportStateCreateInfo.Scissors will be ignored
	// and must be set dynamically with CommandBuffer.CmdSetScissor before any drawing commands
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html
	DynamicStateScissor DynamicState = C.VK_DYNAMIC_STATE_SCISSOR
	// DynamicStateLineWidth specifies that PipelineRasterizationStateCreateInfo.LineWidth
	// will be ignored and must be set dynamically with CommandBuffer.CmdSetLineWidth before any
	// drawing commands
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html
	DynamicStateLineWidth DynamicState = C.VK_DYNAMIC_STATE_LINE_WIDTH
	// DynamicStateDepthBias specifies that PipelineRasterizationStateCreateInfo.DepthBiasConstantFactor,
	// PipelineRasterizationStateCreateInfo.DepthBiasClamp, and PipelineRasterizationStateCreateInfo.DepthBiasSlopeFactor
	// will be ignored and must be set dynamically with CommandBuffer.CmdSetDepthBias before any
	// draws are performed with PipelineRasterizationStateCreateInfo.DepthBiasEnabled set to true
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html
	DynamicStateDepthBias DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BIAS
	// DynamicStateBlendConstants specifies that PipelineColorBlendStateCreateInfo.BlendConstants
	// will be ignored and must be set dynamically with CommandBuffer.CmdSetBlendConstants before
	// any draws are performed with PipelineColorBlendAttachmentState.BlendEnabled set to true
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html
	DynamicStateBlendConstants DynamicState = C.VK_DYNAMIC_STATE_BLEND_CONSTANTS
	// DynamicStateDepthBounds specifies that PipelineDepthStencilStateCreateInfo.MinDepthBounds and
	// PipelineDepthStencilStateCreateInfo.MaxDepthBounds will be ignored and must be set dynamically
	// with CommandBuffer.CmdSetDepthBounds before any draws are performed with
	// PipelineDepthStencilStateCreateInfo.DepthBoundsTestEnable set to true
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html
	DynamicStateDepthBounds DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BOUNDS
	// DynamicStateStencilCompareMask specifies that PipelineDepthStencilStateCreateInfo.Front.CompareMask and
	// PipelineDepthStencilStateCreateInfo.Back.CompareMask will be ignored and must be set dynamically with
	// CommandBuffer.CmdSetStencilCompareMask before any draws are performed with
	// PipelineDepthStencilStateCreateInfo.StencilTestEnable set to true
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html
	DynamicStateStencilCompareMask DynamicState = C.VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK
	// DynamicStateStencilWriteMask specifies that
	// PipelineDepthStencilStateCreateInfo.Front.WriteMask and PipelineDepthStencilStateCreateInfo.Back.WriteMask
	// will be ignored and must be set dynamically with CommandBuffer.CmdSetStencilWriteMask before any draws
	// are performed with PipelineDepthStencilStateCreateInfo.StencilTestEnable
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html
	DynamicStateStencilWriteMask DynamicState = C.VK_DYNAMIC_STATE_STENCIL_WRITE_MASK
	// DynamicStateStencilReference specifies that PipelineDepthStencilStateCreateInfo.Front.Reference
	// and PipelineDepthStencilStateCreateInfo.Back.Reference must be set dynamically with
	// CommandBuffer.CmdSetStencilReference before any draws are performed with
	// PipelineDepthStencilStateCreateInfo.StencilTestEnable set to true
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html
	DynamicStateStencilReference DynamicState = C.VK_DYNAMIC_STATE_STENCIL_REFERENCE
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

// PipelineDynamicStateCreateInfo specifies parameters of a newly-created Pipeline dynamic state
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html
type PipelineDynamicStateCreateInfo struct {
	// Flags is reserved for future use
	Flags PipelineDynamicStateCreateFlags
	// DynamicStates is a slice of DynamicState values specifying which pieces of Pipeline state
	// will use the values from dynamic state commands rather than from Pipeline state creation
	// information
	DynamicStates []DynamicState

	common.NextOptions
}

func (o PipelineDynamicStateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineDynamicStateCreateInfo)
	}
	createInfo := (*C.VkPipelineDynamicStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO
	createInfo.flags = C.VkPipelineDynamicStateCreateFlags(o.Flags)
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
