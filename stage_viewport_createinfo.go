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

type ViewportOptions struct {
	Viewports []common.Viewport
	Scissors  []common.Rect2D

	common.HaveNext
}

func (o *ViewportOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkPipelineViewportStateCreateInfo)(allocator.Malloc(C.sizeof_struct_VkPipelineViewportStateCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next

	viewportCount := len(o.Viewports)
	scissorsCount := len(o.Scissors)

	createInfo.viewportCount = C.uint(viewportCount)
	if viewportCount > 0 {
		viewportPtr := (*C.VkViewport)(allocator.Malloc(viewportCount * C.sizeof_struct_VkViewport))
		viewportSlice := ([]C.VkViewport)(unsafe.Slice(viewportPtr, viewportCount))
		for i := 0; i < viewportCount; i++ {
			viewportSlice[i].x = C.float(o.Viewports[i].X)
			viewportSlice[i].y = C.float(o.Viewports[i].Y)
			viewportSlice[i].width = C.float(o.Viewports[i].Width)
			viewportSlice[i].height = C.float(o.Viewports[i].Height)
			viewportSlice[i].minDepth = C.float(o.Viewports[i].MinDepth)
			viewportSlice[i].maxDepth = C.float(o.Viewports[i].MaxDepth)
		}
		createInfo.pViewports = viewportPtr
	}

	createInfo.scissorCount = C.uint(scissorsCount)
	if scissorsCount > 0 {
		scissorPtr := (*C.VkRect2D)(allocator.Malloc(scissorsCount * C.sizeof_struct_VkRect2D))
		scissorSlice := ([]C.VkRect2D)(unsafe.Slice(scissorPtr, scissorsCount))
		for i := 0; i < scissorsCount; i++ {
			scissorSlice[i].offset.x = C.int32_t(o.Scissors[i].Offset.X)
			scissorSlice[i].offset.y = C.int32_t(o.Scissors[i].Offset.Y)
			scissorSlice[i].extent.width = C.uint32_t(o.Scissors[i].Extent.Width)
			scissorSlice[i].extent.height = C.uint32_t(o.Scissors[i].Extent.Height)
		}
		createInfo.pScissors = scissorPtr
	}

	return unsafe.Pointer(createInfo), nil
}
