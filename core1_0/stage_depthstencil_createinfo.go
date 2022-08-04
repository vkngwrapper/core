package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// PipelineDepthStencilStateCreateFlags are reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateFlags.html
type PipelineDepthStencilStateCreateFlags uint32

var pipelineDepthStencilStateCreateFlagsMapping = common.NewFlagStringMapping[PipelineDepthStencilStateCreateFlags]()

func (f PipelineDepthStencilStateCreateFlags) Register(str string) {
	pipelineDepthStencilStateCreateFlagsMapping.Register(f, str)
}

func (f PipelineDepthStencilStateCreateFlags) String() string {
	return pipelineDepthStencilStateCreateFlagsMapping.FlagsToString(f)
}

////

const (
	// StencilKeep keeps the current value
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html
	StencilKeep StencilOp = C.VK_STENCIL_OP_KEEP
	// StencilZero sets the value to 0
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html
	StencilZero StencilOp = C.VK_STENCIL_OP_ZERO
	// StencilReplace sets the value to Reference
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html
	StencilReplace StencilOp = C.VK_STENCIL_OP_REPLACE
	// StencilIncrementAndClamp increments the current value and clamps to the maximum
	// representable unsigned value
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html
	StencilIncrementAndClamp StencilOp = C.VK_STENCIL_OP_INCREMENT_AND_CLAMP
	// StencilDecrementAndClamp decrements the current value and clamps to 0
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html
	StencilDecrementAndClamp StencilOp = C.VK_STENCIL_OP_DECREMENT_AND_CLAMP
	// StencilInvert bitwise-inverts the current value
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html
	StencilInvert StencilOp = C.VK_STENCIL_OP_INVERT
	// StencilIncrementAndWrap increments the current value and wraps to 0 when the
	// maximum would have been exceeded
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html
	StencilIncrementAndWrap StencilOp = C.VK_STENCIL_OP_INCREMENT_AND_WRAP
	// StencilDecrementAndWrap decrements the current value and wraps to the maximum possible
	// value when the value would go below 0
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html
	StencilDecrementAndWrap StencilOp = C.VK_STENCIL_OP_DECREMENT_AND_WRAP
)

func init() {
	StencilKeep.Register("Keep")
	StencilZero.Register("Zero")
	StencilReplace.Register("Replace")
	StencilIncrementAndClamp.Register("Increment and Clamp")
	StencilDecrementAndClamp.Register("Decrement and Clamp")
	StencilInvert.Register("Invert")
	StencilIncrementAndWrap.Register("Increment and Wrap")
	StencilDecrementAndWrap.Register("Decrement and Wrap")
}

// StencilOpState specifies stencil operation state
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html
type StencilOpState struct {
	// FailOp specifies the action performed on samples that fail the stencil test
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html
	FailOp StencilOp
	// PassOp specifies the action performed on samples that pass both the depth and stencil tests
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html
	PassOp StencilOp
	// DepthFailOp specifies the action performed on samples that pass the stencil test and fail
	// the depth test
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html
	DepthFailOp StencilOp

	// CompareOp specifies the comparison operator used in the stencil test
	CompareOp CompareOp
	// CompareMask selects the bits of the unsigned integer stencil values participating in the
	// stencil test
	CompareMask uint32
	// WriteMask selects the bits of the unsigned integer stencil values updated by the stencil
	// test in the stencil Framebuffer attachment
	WriteMask uint32

	// Reference is an integer stencil reference value that is used in the unsigned stencil
	// comparison
	Reference uint32
}

// PipelineDepthStencilStateCreateInfo specifies parameters of a newly-created Pipeline depth stencil
// state
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html
type PipelineDepthStencilStateCreateInfo struct {
	// Flags specifies additional depth/stencil state information
	Flags PipelineDepthStencilStateCreateFlags

	// DepthTestEnable controls whether depth testing is enabled
	DepthTestEnable bool
	// DepthWriteEnable controls whether depth writes are enabled when DepthTestEnable is true
	DepthWriteEnable bool
	// DepthCompareOp specifies the comparison operator to use in the depth comparison step
	// of the depth test
	DepthCompareOp CompareOp

	// DepthBoundsTestEnable controls whether depth bounds testing is enabled
	DepthBoundsTestEnable bool
	// StencilTestEnable controls whether stencil testing is enabled
	StencilTestEnable bool

	// Front controls the parameters of the stencil test for front-facing triangles
	Front StencilOpState
	// Back controls the parameters of the stencil test for back-facing triangles
	Back StencilOpState

	// MinDepthBounds is the minimum depth bound used in the depth bounds test
	MinDepthBounds float32
	// MaxDepthBounds is the maximum depth bound used in the depth bounds test
	MaxDepthBounds float32

	common.NextOptions
}

func (o PipelineDepthStencilStateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineDepthStencilStateCreateInfo)
	}
	createInfo := (*C.VkPipelineDepthStencilStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO
	createInfo.flags = C.VkPipelineDepthStencilStateCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.depthTestEnable = C.VK_FALSE
	createInfo.depthWriteEnable = C.VK_FALSE
	createInfo.depthBoundsTestEnable = C.VK_FALSE
	createInfo.stencilTestEnable = C.VK_FALSE

	if o.DepthTestEnable {
		createInfo.depthTestEnable = C.VK_TRUE
	}
	if o.DepthWriteEnable {
		createInfo.depthWriteEnable = C.VK_TRUE
	}
	if o.DepthBoundsTestEnable {
		createInfo.depthBoundsTestEnable = C.VK_TRUE
	}
	if o.StencilTestEnable {
		createInfo.stencilTestEnable = C.VK_TRUE
	}

	createInfo.depthCompareOp = C.VkCompareOp(o.DepthCompareOp)
	createInfo.minDepthBounds = C.float(o.MinDepthBounds)
	createInfo.maxDepthBounds = C.float(o.MaxDepthBounds)

	createInfo.front.failOp = C.VkStencilOp(o.Front.FailOp)
	createInfo.front.passOp = C.VkStencilOp(o.Front.PassOp)
	createInfo.front.depthFailOp = C.VkStencilOp(o.Front.DepthFailOp)
	createInfo.front.compareOp = C.VkCompareOp(o.Front.CompareOp)
	createInfo.front.compareMask = C.uint32_t(o.Front.CompareMask)
	createInfo.front.writeMask = C.uint32_t(o.Front.WriteMask)
	createInfo.front.reference = C.uint32_t(o.Front.Reference)

	createInfo.back.failOp = C.VkStencilOp(o.Back.FailOp)
	createInfo.back.passOp = C.VkStencilOp(o.Back.PassOp)
	createInfo.back.depthFailOp = C.VkStencilOp(o.Back.DepthFailOp)
	createInfo.back.compareOp = C.VkCompareOp(o.Back.CompareOp)
	createInfo.back.compareMask = C.uint32_t(o.Back.CompareMask)
	createInfo.back.writeMask = C.uint32_t(o.Back.WriteMask)
	createInfo.back.reference = C.uint32_t(o.Back.Reference)

	return preallocatedPointer, nil
}
