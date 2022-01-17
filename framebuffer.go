package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	driver3 "github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type FramebufferFlags int32

const (
	FramebufferCreateImageless FramebufferFlags = C.VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT
)

var framebufferFlagsToString = map[FramebufferFlags]string{
	FramebufferCreateImageless: "Create Imageless",
}

func (f FramebufferFlags) String() string {
	return common.FlagsToString(f, framebufferFlagsToString)
}

type FramebufferOptions struct {
	Attachments []ImageView
	Flags       FramebufferFlags

	Width  int
	Height int
	Layers uint32

	RenderPass RenderPass

	common.HaveNext
}

func (o *FramebufferOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkFramebufferCreateInfo)(allocator.Malloc(C.sizeof_struct_VkFramebufferCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO
	createInfo.flags = 0
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
			attachmentsSlice[i] = C.VkImageView(unsafe.Pointer(o.Attachments[i].Handle()))
		}

		createInfo.pAttachments = attachmentsPtr
	}

	createInfo.width = C.uint32_t(o.Width)
	createInfo.height = C.uint32_t(o.Height)
	createInfo.layers = C.uint32_t(o.Layers)
	createInfo.flags = C.VkFramebufferCreateFlags(o.Flags)

	return unsafe.Pointer(createInfo), nil
}

type vulkanFramebuffer struct {
	driver driver3.Driver
	device driver3.VkDevice
	handle driver3.VkFramebuffer
}

func (b *vulkanFramebuffer) Handle() driver3.VkFramebuffer {
	return b.handle
}

func (b *vulkanFramebuffer) Destroy(callbacks *AllocationCallbacks) {
	b.driver.VkDestroyFramebuffer(b.device, b.handle, callbacks.Handle())
}
