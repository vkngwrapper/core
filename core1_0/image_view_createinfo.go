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
	SwizzleIdentity common.ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_IDENTITY
	SwizzleZero     common.ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_ZERO
	SwizzleOne      common.ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_ONE
	SwizzleRed      common.ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_R
	SwizzleGreen    common.ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_G
	SwizzleBlue     common.ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_B
	SwizzleAlpha    common.ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_A

	ViewType1D        common.ImageViewType = C.VK_IMAGE_VIEW_TYPE_1D
	ViewType2D        common.ImageViewType = C.VK_IMAGE_VIEW_TYPE_2D
	ViewType3D        common.ImageViewType = C.VK_IMAGE_VIEW_TYPE_3D
	ViewTypeCube      common.ImageViewType = C.VK_IMAGE_VIEW_TYPE_CUBE
	ViewType1DArray   common.ImageViewType = C.VK_IMAGE_VIEW_TYPE_1D_ARRAY
	ViewType2DArray   common.ImageViewType = C.VK_IMAGE_VIEW_TYPE_2D_ARRAY
	ViewTypeCubeArray common.ImageViewType = C.VK_IMAGE_VIEW_TYPE_CUBE_ARRAY
)

func init() {
	SwizzleIdentity.Register("Identity")
	SwizzleZero.Register("Zero")
	SwizzleOne.Register("One")
	SwizzleRed.Register("Red")
	SwizzleGreen.Register("Green")
	SwizzleBlue.Register("Blue")
	SwizzleAlpha.Register("Alpha")

	ViewType1D.Register("1D")
	ViewType2D.Register("2D")
	ViewType3D.Register("3D")
	ViewTypeCube.Register("Cube")
	ViewType1DArray.Register("1D Array")
	ViewType2DArray.Register("2D Array")
	ViewTypeCubeArray.Register("Cube Array")
}

type ComponentMapping struct {
	R common.ComponentSwizzle
	G common.ComponentSwizzle
	B common.ComponentSwizzle
	A common.ComponentSwizzle
}

type ImageViewOptions struct {
	Image Image

	Flags            common.ImageViewCreateFlags
	ViewType         common.ImageViewType
	Format           common.DataFormat
	Components       ComponentMapping
	SubresourceRange common.ImageSubresourceRange

	common.HaveNext
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
