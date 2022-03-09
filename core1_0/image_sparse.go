package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
)

const (
	SparseImageFormatSingleMipTail        common.SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT
	SparseImageFormatAlignedMipSize       common.SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT
	SparseImageFormatNonstandardBlockSize common.SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT

	AspectColor    common.ImageAspectFlags = C.VK_IMAGE_ASPECT_COLOR_BIT
	AspectDepth    common.ImageAspectFlags = C.VK_IMAGE_ASPECT_DEPTH_BIT
	AspectStencil  common.ImageAspectFlags = C.VK_IMAGE_ASPECT_STENCIL_BIT
	AspectMetadata common.ImageAspectFlags = C.VK_IMAGE_ASPECT_METADATA_BIT
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
	AspectMask       common.ImageAspectFlags
	ImageGranularity common.Extent3D
	Flags            common.SparseImageFormatFlags
}

type SparseImageMemoryRequirements struct {
	FormatProperties     SparseImageFormatProperties
	ImageMipTailFirstLod int
	ImageMipTailSize     int
	ImageMipTailOffset   int
	ImageMipTailStride   int
}
