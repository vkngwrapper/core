package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
)

const (
	// ComponentSwizzleIdentity specifies that the component is set to the identity swizzle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html
	ComponentSwizzleIdentity ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_IDENTITY
	// ComponentSwizzleZero specifies that the component is set to zero
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html
	ComponentSwizzleZero ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_ZERO
	// ComponentSwizzleOne specifies that hte component is set to either 1 or 1.0, depending
	// on the type of the ImageView format
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html
	ComponentSwizzleOne ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_ONE
	// ComponentSwizzleRed specifies that the component is set to the value of the R component
	// of the Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html
	ComponentSwizzleRed ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_R
	// ComponentSwizzleGreen specifies that the component is set to the value of the G component
	// of the Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html
	ComponentSwizzleGreen ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_G
	// ComponentSwizzleBlue specifies that the component is set to the value of the B component
	// of the Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html
	ComponentSwizzleBlue ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_B
	// ComponentSwizzleAlpha specifies that the component is set to the value of the A component
	// of the Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentSwizzle.html
	ComponentSwizzleAlpha ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_A

	// ImageViewType1D specifies a 1-dimensional ImageView
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html
	ImageViewType1D ImageViewType = C.VK_IMAGE_VIEW_TYPE_1D
	// ImageViewType2D specifies a 2-dimensional ImageView
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html
	ImageViewType2D ImageViewType = C.VK_IMAGE_VIEW_TYPE_2D
	// ImageViewType3D specifies a 3-dimensional ImageView
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html
	ImageViewType3D ImageViewType = C.VK_IMAGE_VIEW_TYPE_3D
	// ImageViewTypeCube specifies a cube map
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html
	ImageViewTypeCube ImageViewType = C.VK_IMAGE_VIEW_TYPE_CUBE
	// ImageViewType1DArray specifies an ImageView that is an array of 1-dimensional images
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html
	ImageViewType1DArray ImageViewType = C.VK_IMAGE_VIEW_TYPE_1D_ARRAY
	// ImageViewType2DArray specifies an ImageView that is an array of 2-dimensional images
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html
	ImageViewType2DArray ImageViewType = C.VK_IMAGE_VIEW_TYPE_2D_ARRAY
	// ImageViewTypeCubeArray specifies an ImageView that is an array of cube maps
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html
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

// ComponentMapping specifies a color component mapping
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html
type ComponentMapping struct {
	// R specifies the component value placed in the R component of the output vector
	R ComponentSwizzle
	// G specifies the component value placed in the G component of the output vector
	G ComponentSwizzle
	// B specifies the component value placed in the B component of the output vector
	B ComponentSwizzle
	// A specifies the component value placed in the A component of the output vector
	A ComponentSwizzle
}

// ImageViewCreateInfo specifies parameters of a newly-created ImageView
type ImageViewCreateInfo struct {
	// Image is an Image on which the view will be created
	Image core.Image

	// Flags describes additional parameters of the ImageView
	Flags ImageViewCreateFlags
	// ViewType specifies the type of the ImageView
	ViewType ImageViewType
	// Format describes the format and type used to interpret texel blocks in the Image
	Format Format
	// Components specifies a remapping of color components
	Components ComponentMapping
	// SubresourceRange selects the set of mipmap levels and array layers to be accessible
	// to the view
	SubresourceRange ImageSubresourceRange

	common.NextOptions
}

func (o ImageViewCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Image.Handle() == 0 {
		return nil, errors.New("core1_0.ImageViewCreateInfo.Image cannot be left unset")
	}
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
