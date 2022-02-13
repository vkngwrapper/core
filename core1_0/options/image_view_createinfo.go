package options

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ImageViewFlags int32

const (
	ImageViewCreateFragmentDensityMapDynamicEXT  ImageViewFlags = C.VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT
	ImageViewCreateFragmentDensityMapDeferredEXT ImageViewFlags = C.VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT
)

var imageViewFlagsToString = map[ImageViewFlags]string{
	ImageViewCreateFragmentDensityMapDynamicEXT:  "Create Fragment Density Map - Dynamic (Extension)",
	ImageViewCreateFragmentDensityMapDeferredEXT: "Create Fragment Density Map - Deferred (Extension)",
}

func (f ImageViewFlags) String() string {
	return common.FlagsToString(f, imageViewFlagsToString)
}

type ImageViewOptions struct {
	Image iface.Image

	Flags            ImageViewFlags
	ViewType         common.ImageViewType
	Format           common.DataFormat
	Components       common.ComponentMapping
	SubresourceRange common.ImageSubresourceRange

	core.HaveNext
}

func (o *ImageViewOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkImageViewCreateInfo{})))
	}
	createInfo := (*C.VkImageViewCreateInfo)(preallocatedPointer)

	createInfo.sType = C.VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
	createInfo.flags = C.VkImageViewCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))
	createInfo.viewType = C.VkImageViewType(o.ViewType)
	createInfo.format = C.VkFormat(o.Format)
	createInfo.components.r = C.VkComponentSwizzle(o.Components.R)
	createInfo.components.g = C.VkComponentSwizzle(o.Components.G)
	createInfo.components.b = C.VkComponentSwizzle(o.Components.B)
	createInfo.components.a = C.VkComponentSwizzle(o.Components.A)
	createInfo.subresourceRange.aspectMask = C.VkImageAspectFlags(o.SubresourceRange.AspectMask)
	createInfo.subresourceRange.baseMipLevel = C.uint32_t(o.SubresourceRange.BaseMipLevel)
	createInfo.subresourceRange.levelCount = C.uint32_t(o.SubresourceRange.LevelCount)
	createInfo.subresourceRange.baseArrayLayer = C.uint32_t(o.SubresourceRange.BaseArrayLayer)
	createInfo.subresourceRange.layerCount = C.uint32_t(o.SubresourceRange.LayerCount)

	return preallocatedPointer, nil
}
