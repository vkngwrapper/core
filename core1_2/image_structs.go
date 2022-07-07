package core1_2

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

const (
	ImageLayoutDepthAttachmentOptimal   core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
	ImageLayoutDepthReadOnlyOptimal     core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL
	ImageLayoutStencilAttachmentOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL
	ImageLayoutStencilReadOnlyOptimal   core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL
)

func init() {
	ImageLayoutDepthAttachmentOptimal.Register("Depth Attachment Optimal")
	ImageLayoutDepthReadOnlyOptimal.Register("Depth Read-Only Optimal")
	ImageLayoutStencilAttachmentOptimal.Register("Stencil Attachment Optimal")
	ImageLayoutStencilReadOnlyOptimal.Register("Stencil Read-Only Optimal")
}

////

type ImageStencilUsageCreateInfo struct {
	StencilUsage core1_0.ImageUsageFlags

	common.NextOptions
}

func (o ImageStencilUsageCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageStencilUsageCreateInfo{})))
	}

	info := (*C.VkImageStencilUsageCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO
	info.pNext = next
	info.stencilUsage = C.VkImageUsageFlags(o.StencilUsage)

	return preallocatedPointer, nil
}

////

type ImageFormatListCreateInfo struct {
	ViewFormats []core1_0.Format

	common.NextOptions
}

func (o ImageFormatListCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
