package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const (
	ComponentSwizzleIdentity ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_IDENTITY
	ComponentSwizzleZero     ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_ZERO
	ComponentSwizzleOne      ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_ONE
	ComponentSwizzleRed      ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_R
	ComponentSwizzleGreen    ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_G
	ComponentSwizzleBlue     ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_B
	ComponentSwizzleAlpha    ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_A

	ImageViewType1D        ImageViewType = C.VK_IMAGE_VIEW_TYPE_1D
	ImageViewType2D        ImageViewType = C.VK_IMAGE_VIEW_TYPE_2D
	ImageViewType3D        ImageViewType = C.VK_IMAGE_VIEW_TYPE_3D
	ImageViewTypeCube      ImageViewType = C.VK_IMAGE_VIEW_TYPE_CUBE
	ImageViewType1DArray   ImageViewType = C.VK_IMAGE_VIEW_TYPE_1D_ARRAY
	ImageViewType2DArray   ImageViewType = C.VK_IMAGE_VIEW_TYPE_2D_ARRAY
	ImageViewTypeCubeArray ImageViewType = C.VK_IMAGE_VIEW_TYPE_CUBE_ARRAY
)

func init() {
	ComponentSwizzleIdentity.Register("Identity")
	ComponentSwizzleZero.Register("Zero")
	ComponentSwizzleOne.Register("One")
	ComponentSwizzleRed.Register("Red")
	ComponentSwizzleGreen.Register("Green")
	ComponentSwizzleBlue.Register("Blue")
	ComponentSwizzleAlpha.Register("Alpha")

	ImageViewType1D.Register("1D")
	ImageViewType2D.Register("2D")
	ImageViewType3D.Register("3D")
	ImageViewTypeCube.Register("Cube")
	ImageViewType1DArray.Register("1D Array")
	ImageViewType2DArray.Register("2D Array")
	ImageViewTypeCubeArray.Register("Cube Array")
}

type ComponentMapping struct {
	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle
}

type ImageViewCreateInfo struct {
	Image Image

	Flags            ImageViewCreateFlags
	ViewType         ImageViewType
	Format           Format
	Components       ComponentMapping
	SubresourceRange ImageSubresourceRange

	common.NextOptions
}

func (o ImageViewCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
