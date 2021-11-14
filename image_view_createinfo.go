package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"strings"
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
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		checkBit := ImageViewFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := imageViewFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type ImageViewOptions struct {
	Image Image

	Flags            ImageViewFlags
	ViewType         common.ImageViewType
	Format           common.DataFormat
	Components       common.ComponentMapping
	SubresourceRange common.ImageSubresourceRange

	common.HaveNext
}

func (o *ImageViewOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkImageViewCreateInfo)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkImageViewCreateInfo{}))))

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

	return unsafe.Pointer(createInfo), nil
}
