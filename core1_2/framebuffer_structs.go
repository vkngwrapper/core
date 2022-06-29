package core1_2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const (
	FramebufferCreateImageless core1_0.FramebufferCreateFlags = C.VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT
)

func init() {
	FramebufferCreateImageless.Register("Imageless")
}

////

type FramebufferAttachmentImageOptions struct {
	Flags      core1_0.ImageCreateFlags
	Usage      core1_0.ImageUsages
	Width      int
	Height     int
	LayerCount int

	ViewFormats []core1_0.DataFormat

	common.HaveNext
}

func (o FramebufferAttachmentImageOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkFramebufferAttachmentImageInfo{})))
	}

	info := (*C.VkFramebufferAttachmentImageInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO
	info.pNext = next
	info.flags = C.VkImageCreateFlags(o.Flags)
	info.usage = C.VkImageUsageFlags(o.Usage)
	info.width = C.uint32_t(o.Width)
	info.height = C.uint32_t(o.Height)
	info.layerCount = C.uint32_t(o.LayerCount)

	count := len(o.ViewFormats)
	info.viewFormatCount = C.uint32_t(count)
	info.pViewFormats = nil

	if count > 0 {
		info.pViewFormats = (*C.VkFormat)(allocator.Malloc(count * int(unsafe.Sizeof(C.VkFormat(0)))))
		viewFormatSlice := ([]C.VkFormat)(unsafe.Slice(info.pViewFormats, count))
		for i := 0; i < count; i++ {
			viewFormatSlice[i] = C.VkFormat(o.ViewFormats[i])
		}
	}

	return preallocatedPointer, nil
}

func (o FramebufferAttachmentImageOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkFramebufferAttachmentImageInfo)(cDataPointer)
	return info.pNext, nil
}

////

type FramebufferAttachmentsCreateOptions struct {
	AttachmentImageInfos []FramebufferAttachmentImageOptions

	common.HaveNext
}

func (o FramebufferAttachmentsCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkFramebufferAttachmentsCreateInfo{})))
	}

	info := (*C.VkFramebufferAttachmentsCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO
	info.pNext = next

	count := len(o.AttachmentImageInfos)
	info.attachmentImageInfoCount = C.uint32_t(count)
	info.pAttachmentImageInfos = nil

	if count > 0 {
		infosPtr, err := common.AllocOptionSlice[C.VkFramebufferAttachmentImageInfo, FramebufferAttachmentImageOptions](allocator, o.AttachmentImageInfos)
		if err != nil {
			return nil, err
		}

		info.pAttachmentImageInfos = infosPtr
	}

	return preallocatedPointer, nil
}

func (o FramebufferAttachmentsCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkFramebufferAttachmentsCreateInfo)(cDataPointer)
	return info.pNext, nil
}
