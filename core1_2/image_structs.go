package core1_2

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

const (
	// ImageLayoutDepthAttachmentOptimal specifies a layout for the depth aspect of a depth/stencil
	// format Image allowing read and write access as a depth attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutDepthAttachmentOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL
	// ImageLayoutDepthReadOnlyOptimal specifies a layout for the depth aspect of a depth/stencil
	// format Image allowing read-only access as a depth attachment or in shaders as a sampled Image,
	// combined Image/Sampler, or input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutDepthReadOnlyOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL
	// ImageLayoutStencilAttachmentOptimal specifies a layout for the stencil aspect of a
	// depth/stencil format Image allowing read and write access as a stencil attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutStencilAttachmentOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL
	// ImageLayoutStencilReadOnlyOptimal specifies a layout for the stencil aspect of a depth/stencil
	// format Image allowing read-only access as a stencil attachment or in shaders as a sampled
	// Image, combined Image/Sampler, or input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutStencilReadOnlyOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL
)

func init() {
	ImageLayoutDepthAttachmentOptimal.Register("Depth Attachment Optimal")
	ImageLayoutDepthReadOnlyOptimal.Register("Depth Read-Only Optimal")
	ImageLayoutStencilAttachmentOptimal.Register("Stencil Attachment Optimal")
	ImageLayoutStencilReadOnlyOptimal.Register("Stencil Read-Only Optimal")
}

////

// ImageStencilUsageCreateInfo specifies separate usage flags for the stencil aspect of a
// depth-stencil Image
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html
type ImageStencilUsageCreateInfo struct {
	// StencilUsage describes the intended usage of the stencil aspect of the Image
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

// ImageFormatListCreateInfo specifies that an Image can be used with a particular set of formats
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html
type ImageFormatListCreateInfo struct {
	// ViewFormats is a slice of core1_0.Format values specifying all formats which can be used
	// when creating views of this Image
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
