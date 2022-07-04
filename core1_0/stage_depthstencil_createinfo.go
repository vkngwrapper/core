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
	StencilKeep              StencilOp = C.VK_STENCIL_OP_KEEP
	StencilZero              StencilOp = C.VK_STENCIL_OP_ZERO
	StencilReplace           StencilOp = C.VK_STENCIL_OP_REPLACE
	StencilIncrementAndClamp StencilOp = C.VK_STENCIL_OP_INCREMENT_AND_CLAMP
	StencilDecrementAndClamp StencilOp = C.VK_STENCIL_OP_DECREMENT_AND_CLAMP
	StencilInvert            StencilOp = C.VK_STENCIL_OP_INVERT
	StencilIncrementAndWrap  StencilOp = C.VK_STENCIL_OP_INCREMENT_AND_WRAP
	StencilDecrementAndWrap  StencilOp = C.VK_STENCIL_OP_DECREMENT_AND_WRAP
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

type StencilOpState struct {
	FailOp      StencilOp
	PassOp      StencilOp
	DepthFailOp StencilOp

	CompareOp   CompareOp
	CompareMask uint32
	WriteMask   uint32

	Reference uint32
}

type DepthStencilStateOptions struct {
	DepthTestEnable  bool
	DepthWriteEnable bool
	DepthCompareOp   CompareOp

	DepthBoundsTestEnable bool
	StencilTestEnable     bool

	FrontStencilState StencilOpState
	BackStencilState  StencilOpState

	MinDepthBounds float32
	MaxDepthBounds float32

	common.NextOptions
}

func (o DepthStencilStateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPipelineDepthStencilStateCreateInfo)
	}
	createInfo := (*C.VkPipelineDepthStencilStateCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO
	createInfo.flags = 0
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

	createInfo.front.failOp = C.VkStencilOp(o.FrontStencilState.FailOp)
	createInfo.front.passOp = C.VkStencilOp(o.FrontStencilState.PassOp)
	createInfo.front.depthFailOp = C.VkStencilOp(o.FrontStencilState.DepthFailOp)
	createInfo.front.compareOp = C.VkCompareOp(o.FrontStencilState.CompareOp)
	createInfo.front.compareMask = C.uint32_t(o.FrontStencilState.CompareMask)
	createInfo.front.writeMask = C.uint32_t(o.FrontStencilState.WriteMask)
	createInfo.front.reference = C.uint32_t(o.FrontStencilState.Reference)

	createInfo.back.failOp = C.VkStencilOp(o.BackStencilState.FailOp)
	createInfo.back.passOp = C.VkStencilOp(o.BackStencilState.PassOp)
	createInfo.back.depthFailOp = C.VkStencilOp(o.BackStencilState.DepthFailOp)
	createInfo.back.compareOp = C.VkCompareOp(o.BackStencilState.CompareOp)
	createInfo.back.compareMask = C.uint32_t(o.BackStencilState.CompareMask)
	createInfo.back.writeMask = C.uint32_t(o.BackStencilState.WriteMask)
	createInfo.back.reference = C.uint32_t(o.BackStencilState.Reference)

	return preallocatedPointer, nil
}
