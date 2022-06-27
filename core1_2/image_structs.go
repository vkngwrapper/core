package core1_2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const (
	ImageLayoutDepthAttachmentOptimal   common.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
	ImageLayoutDepthReadOnlyOptimal     common.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL
	ImageLayoutStencilAttachmentOptimal common.ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL
	ImageLayoutStencilReadOnlyOptimal   common.ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL
)

func init() {
	ImageLayoutDepthAttachmentOptimal.Register("Depth Attachment Optimal")
	ImageLayoutDepthReadOnlyOptimal.Register("Depth Read-Only Optimal")
	ImageLayoutStencilAttachmentOptimal.Register("Stencil Attachment Optimal")
	ImageLayoutStencilReadOnlyOptimal.Register("Stencil Read-Only Optimal")
}

////

type ImageStencilUsageCreateOptions struct {
	StencilUsage common.ImageUsages

	common.HaveNext
}

func (o ImageStencilUsageCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageStencilUsageCreateInfo{})))
	}

	info := (*C.VkImageStencilUsageCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO
	info.pNext = next
	info.stencilUsage = C.VkImageUsageFlags(o.StencilUsage)

	return preallocatedPointer, nil
}

func (o ImageStencilUsageCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkImageStencilUsageCreateInfo)(cDataPointer)
	return info.pNext, nil
}

////

type ImageFormatListCreateOptions struct {
	ViewFormats []common.DataFormat

	common.HaveNext
}

func (o ImageFormatListCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageFormatListCreateInfo{})))
	}

	info := (*C.VkImageFormatListCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO
	info.pNext = next

	count := len(o.ViewFormats)
	info.viewFormatCount = C.uint32_t(count)
	info.pViewFormats = nil

	if count > 0 {
		info.pViewFormats = (*C.VkFormat)(allocator.Malloc(count * int(unsafe.Sizeof(C.VkFormat(0)))))
		viewFormatSlice := unsafe.Slice(info.pViewFormats, count)

		for i := 0; i < count; i++ {
			viewFormatSlice[i] = C.VkFormat(o.ViewFormats[i])
		}
	}

	return preallocatedPointer, nil
}

func (o ImageFormatListCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkImageFormatListCreateInfo)(cDataPointer)
	return info.pNext, nil
}
