package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const (
	SparseImageFormatSingleMipTail        SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT
	SparseImageFormatAlignedMipSize       SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT
	SparseImageFormatNonstandardBlockSize SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT

	AspectColor    ImageAspectFlags = C.VK_IMAGE_ASPECT_COLOR_BIT
	AspectDepth    ImageAspectFlags = C.VK_IMAGE_ASPECT_DEPTH_BIT
	AspectStencil  ImageAspectFlags = C.VK_IMAGE_ASPECT_STENCIL_BIT
	AspectMetadata ImageAspectFlags = C.VK_IMAGE_ASPECT_METADATA_BIT
)

func init() {
	AspectColor.Register("Color")
	AspectDepth.Register("Depth")
	AspectStencil.Register("Stencil")
	AspectMetadata.Register("Metadata")

	SparseImageFormatSingleMipTail.Register("Single Mip Tail")
	SparseImageFormatAlignedMipSize.Register("Aligned Mip Size")
	SparseImageFormatNonstandardBlockSize.Register("Nonstandard Block Size")
}

type SparseImageFormatProperties struct {
	AspectMask       ImageAspectFlags
	ImageGranularity Extent3D
	Flags            SparseImageFormatFlags
}

type SparseImageMemoryRequirements struct {
	FormatProperties     SparseImageFormatProperties
	ImageMipTailFirstLod int
	ImageMipTailSize     int
	ImageMipTailOffset   int
	ImageMipTailStride   int
}
