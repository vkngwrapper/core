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

type SubpassContents int32

const (
	ContentsInline                  SubpassContents = C.VK_SUBPASS_CONTENTS_INLINE
	ContentsSecondaryCommandBuffers SubpassContents = C.VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
)

var subpassContentsToString = map[SubpassContents]string{
	ContentsInline:                  "Inline",
	ContentsSecondaryCommandBuffers: "Secondary Command Buffers",
}

func (c SubpassContents) String() string {
	return subpassContentsToString[c]
}

type ClearValue interface {
	populateUnion(v *C.VkClearValue)
}

type ClearValueInt32 [4]int32

func (v ClearValueInt32) populateUnion(c *C.VkClearValue) {
	colorInt32 := unsafe.Slice((*C.int32_t)(unsafe.Pointer(c)), 4)
	for i := 0; i < 4; i++ {
		colorInt32[i] = C.int32_t(v[i])
	}
}

type ClearValueUint32 [4]uint32

func (v ClearValueUint32) populateUnion(c *C.VkClearValue) {
	colorUint32 := unsafe.Slice((*C.uint32_t)(unsafe.Pointer(c)), 4)
	for i := 0; i < 4; i++ {
		colorUint32[i] = C.uint32_t(v[i])
	}
}

type ClearValueFloat [4]float32

func (v ClearValueFloat) populateUnion(c *C.VkClearValue) {
	colorFloat := unsafe.Slice((*C.float)(unsafe.Pointer(c)), 4)
	for i := 0; i < 4; i++ {
		colorFloat[i] = C.float(v[i])
	}
}

type ClearValueDepthStencil struct {
	Depth   float32
	Stencil uint32
}

func (s *ClearValueDepthStencil) populateUnion(c *C.VkClearValue) {
	depthStencil := (*C.VkClearDepthStencilValue)(unsafe.Pointer(c))
	depthStencil.depth = C.float(s.Depth)
	depthStencil.stencil = C.uint32_t(s.Stencil)
}

type RenderPassBeginOptions struct {
	RenderPass  RenderPass
	Framebuffer Framebuffer

	RenderArea  common.Rect2D
	ClearValues []ClearValue

	common.HaveNext
}

func (o *RenderPassBeginOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkRenderPassBeginInfo)(allocator.Malloc(C.sizeof_struct_VkRenderPassBeginInfo))
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
			o.ClearValues[i].populateUnion(&(valueSlice[i]))
		}

		createInfo.pClearValues = valuePtr
	}

	return unsafe.Pointer(createInfo), nil
}
