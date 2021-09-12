package render_pass

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/resources"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type FramebufferOptions struct {
	Attachments []resources.ImageView
	RenderPass  RenderPass

	Width  uint32
	Height uint32
	Layers uint32

	core.HaveNext
}

func (o *FramebufferOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkFramebufferCreateInfo)(allocator.Malloc(C.sizeof_struct_VkFramebufferCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next

	createInfo.renderPass = (C.VkRenderPass)(unsafe.Pointer(o.RenderPass.Handle()))

	attachmentCount := len(o.Attachments)
	createInfo.attachmentCount = C.uint32_t(attachmentCount)
	if attachmentCount > 0 {
		attachmentsPtr := (*C.VkImageView)(allocator.Malloc(attachmentCount * int(unsafe.Sizeof([1]C.VkImageView{}))))
		attachmentsSlice := ([]C.VkImageView)(unsafe.Slice(attachmentsPtr, attachmentCount))

		for i := 0; i < attachmentCount; i++ {
			attachmentsSlice[i] = C.VkImageView(unsafe.Pointer(o.Attachments[i].Handle()))
		}

		createInfo.pAttachments = attachmentsPtr
	}

	createInfo.width = C.uint32_t(o.Width)
	createInfo.height = C.uint32_t(o.Height)
	createInfo.layers = C.uint32_t(o.Layers)

	return unsafe.Pointer(createInfo), nil
}
