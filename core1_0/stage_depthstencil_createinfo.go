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

type PipelineDepthStencilStateCreateInfo struct {
	Flags PipelineDepthStencilStateCreateFlags

	DepthTestEnable  bool
	DepthWriteEnable bool
	DepthCompareOp   CompareOp

	DepthBoundsTestEnable bool
	StencilTestEnable     bool

	Front StencilOpState
	Back  StencilOpState

	MinDepthBounds float32
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
