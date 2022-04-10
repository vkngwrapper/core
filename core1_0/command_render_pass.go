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
	SubpassContentsInline                  common.SubpassContents = C.VK_SUBPASS_CONTENTS_INLINE
	SubpassContentsSecondaryCommandBuffers common.SubpassContents = C.VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
)

func init() {
	SubpassContentsInline.Register("Inline")
	SubpassContentsSecondaryCommandBuffers.Register("Secondary Command Buffers")
}

type RenderPassBeginOptions struct {
	RenderPass  RenderPass
	Framebuffer Framebuffer

	RenderArea  common.Rect2D
	ClearValues []common.ClearValue

	common.HaveNext
}

func (o RenderPassBeginOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkRenderPassBeginInfo)
	}

	createInfo := (*C.VkRenderPassBeginInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
	createInfo.pNext = next
	createInfo.renderPass = nil
	createInfo.framebuffer = nil

	if o.RenderPass != nil {
		createInfo.renderPass = (C.VkRenderPass)(unsafe.Pointer(o.RenderPass.Handle()))
	}

	if o.Framebuffer != nil {
		createInfo.framebuffer = (C.VkFramebuffer)(unsafe.Pointer(o.Framebuffer.Handle()))
	}

	createInfo.renderArea.offset.x = C.int32_t(o.RenderArea.Offset.X)
	createInfo.renderArea.offset.y = C.int32_t(o.RenderArea.Offset.Y)
	createInfo.renderArea.extent.width = C.uint32_t(o.RenderArea.Extent.Width)
	createInfo.renderArea.extent.height = C.uint32_t(o.RenderArea.Extent.Height)

	clearValueCount := len(o.ClearValues)
	createInfo.clearValueCount = C.uint32_t(clearValueCount)
	createInfo.pClearValues = nil

	if clearValueCount > 0 {
		valuePtr := (*C.VkClearValue)(allocator.Malloc(clearValueCount * C.sizeof_union_VkClearValue))
		valueSlice := ([]C.VkClearValue)(unsafe.Slice(valuePtr, clearValueCount))

		for i := 0; i < clearValueCount; i++ {
			o.ClearValues[i].PopulateValueUnion(unsafe.Pointer(&(valueSlice[i])))
		}

		createInfo.pClearValues = valuePtr
	}

	return preallocatedPointer, nil
}

func (o RenderPassBeginOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkRenderPassBeginInfo)(cDataPointer)
	return createInfo.pNext, nil
}
