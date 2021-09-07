package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type DepthStencilOptions struct {
	DepthTestEnable  bool
	DepthWriteEnable bool
	DepthCompareOp   core.CompareOp

	DepthBoundsTestEnable bool
	StencilTestEnable     bool

	FrontStencilState StencilOpState
	BackStencilState  StencilOpState

	MinDepthBounds float32
	MaxDepthBounds float32

	Next core.Options
}

func (o *DepthStencilOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineDepthStencilStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineDepthStencilStateCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO
	createInfo.flags = 0
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
