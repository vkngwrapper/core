package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const (
	// SparseImageFormatSingleMipTail specifies that the Image uses a single mip tail region for all
	// array layers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatFlagBits.html
	SparseImageFormatSingleMipTail SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT
	// SparseImageFormatAlignedMipSize specifies that the first mip level whose dimensions are not
	// integer multiples of the corresponding dimensions of the sparse Image block begins the
	// mip tail region
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatFlagBits.html
	SparseImageFormatAlignedMipSize SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT
	// SparseImageFormatNonstandardBlockSize specifies that the Image uses non-standard sparse Image
	// block dimensions
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatFlagBits.html
	SparseImageFormatNonstandardBlockSize SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT

	// ImageAspectColor specifies the color aspect
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html
	ImageAspectColor ImageAspectFlags = C.VK_IMAGE_ASPECT_COLOR_BIT
	// ImageAspectDepth speciifies the depth aspect
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html
	ImageAspectDepth ImageAspectFlags = C.VK_IMAGE_ASPECT_DEPTH_BIT
	// ImageAspectStencil specifies the stencil aspect
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html
	ImageAspectStencil ImageAspectFlags = C.VK_IMAGE_ASPECT_STENCIL_BIT
	// ImageAspectMetadata specifies the metadata aspect, used for sparse resource operations
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html
	ImageAspectMetadata ImageAspectFlags = C.VK_IMAGE_ASPECT_METADATA_BIT
)

func init() {
	ImageAspectColor.Register("Color")
	ImageAspectDepth.Register("Depth")
	ImageAspectStencil.Register("Stencil")
	ImageAspectMetadata.Register("Metadata")

	SparseImageFormatSingleMipTail.Register("Single Mip Tail")
	SparseImageFormatAlignedMipSize.Register("Aligned Mip Size")
	SparseImageFormatNonstandardBlockSize.Register("Nonstandard Block Size")
}

// SparseImageFormatProperties specifies sparse Image format properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties.html
type SparseImageFormatProperties struct {
	// AspectMask specifies which aspects of the Image the properties apply to
	AspectMask ImageAspectFlags
	// ImageGranularity is the width, height, and depth of the sparse Image block in texels
	// or compressed texel blocks
	ImageGranularity Extent3D
	// Flags specifies additional information about the sparse resource
	Flags SparseImageFormatFlags
}

// SparseImageMemoryRequirements specifies sparse image memory requirements
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements.html
type SparseImageMemoryRequirements struct {
	// FormatProperties specifies properties of the Image format
	FormatProperties SparseImageFormatProperties
	// ImageMipTailFirstLod is the first mip level at which Image subresources are included in the mip tail region
	ImageMipTailFirstLod int
	// ImageMipTailSize is the memory size in (in bytes) of the mip tail region
	ImageMipTailSize int
	// ImageMipTailOffset is the opaque memory offset used to bind the mip tail region(s)
	ImageMipTailOffset int
	// ImageMipTailStride is the offset stride between each array-layer's mip tail
	ImageMipTailStride int
}
