package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

const (
	// SubpassContentsInline specifies that the contents of the subpass will be recorded inline in
	// the primary CommandBuffer and secondary CommandBuffer objects must not be executed within
	// the subpass.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html
	SubpassContentsInline SubpassContents = C.VK_SUBPASS_CONTENTS_INLINE
	// SubpassContentsSecondaryCommandBuffers specifies that the contents are recorded in
	// secondary CommandBuffer objects that will be called from the primary CommandBuffer
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html
	SubpassContentsSecondaryCommandBuffers SubpassContents = C.VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
)

func init() {
	SubpassContentsInline.Register("Inline")
	SubpassContentsSecondaryCommandBuffers.Register("Secondary Command Buffers")
}

// RenderPassBeginInfo specifies RenderPass begin information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html
type RenderPassBeginInfo struct {
	// RenderPass is the RenderPass to begin an instance of
	RenderPass RenderPass
	// Framebuffer is the Framebuffer containing the attachments that are used with the
	// RenderPass
	Framebuffer Framebuffer

	// RenderArea is the render area that is affected by this RenderPass instance
	RenderArea Rect2D
	// ClearValues is a slice of ClearValue structures containing clear values for each attachment
	// if the attachment uses an AttachmentLoadOp value of AttachmentLoadOpClear. Elements of the slice
	// corresponding to attachments that do not use AttachmentLoadOpClear are ignored
	ClearValues []ClearValue

	common.NextOptions
}

func (o RenderPassBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkRenderPassBeginInfo)
	}

	createInfo := (*C.VkRenderPassBeginInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
	createInfo.pNext = next
	createInfo.renderPass = nil
	createInfo.framebuffer = nil

	if o.RenderPass.Initialized() {
		createInfo.renderPass = (C.VkRenderPass)(unsafe.Pointer(o.RenderPass.Handle()))
	}

	if o.Framebuffer.Initialized() {
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
