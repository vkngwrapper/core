package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// FramebufferCreateInfo specifies parameters of a newly-created Framebuffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html
type FramebufferCreateInfo struct {
	// Attachments is a slice ImageView objects, each of which will be used as the corresponding
	// attachment in a RenderPass instance
	Attachments []ImageView
	// Flags is a bitmask of FramebufferCreateFlags
	Flags FramebufferCreateFlags

	// Width is the width of the Framebuffer
	Width int
	// Height is the height of the Framebuffer
	Height int
	// Layers is the depth of the Framebuffer
	Layers uint32

	// RenderPass is a RenderPass defining what render passes the Framebuffer will be compatible with
	RenderPass RenderPass

	common.NextOptions
}

func (o FramebufferCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkFramebufferCreateInfo)
	}
	createInfo := (*C.VkFramebufferCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO
	createInfo.flags = C.VkFramebufferCreateFlags(o.Flags)
	createInfo.pNext = next

	if o.RenderPass != nil {
		createInfo.renderPass = (C.VkRenderPass)(unsafe.Pointer(o.RenderPass.Handle()))
	}

	attachmentCount := len(o.Attachments)
	createInfo.attachmentCount = C.uint32_t(attachmentCount)
	if attachmentCount > 0 {
		attachmentsPtr := (*C.VkImageView)(allocator.Malloc(attachmentCount * int(unsafe.Sizeof([1]C.VkImageView{}))))
		attachmentsSlice := ([]C.VkImageView)(unsafe.Slice(attachmentsPtr, attachmentCount))

		for i := 0; i < attachmentCount; i++ {
			if o.Attachments[i] == nil {
				return nil, errors.Newf("core1_0.FrameBufferCreateInfo.Attachments cannot contain nil elements, but element %d is nil", i)
			}
			attachmentsSlice[i] = C.VkImageView(unsafe.Pointer(o.Attachments[i].Handle()))
		}

		createInfo.pAttachments = attachmentsPtr
	}

	createInfo.width = C.uint32_t(o.Width)
	createInfo.height = C.uint32_t(o.Height)
	createInfo.layers = C.uint32_t(o.Layers)

	return unsafe.Pointer(createInfo), nil
}
