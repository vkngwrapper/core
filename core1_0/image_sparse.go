package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
)

type SparseImageFormatFlags int32

const (
	SparseImageFormatSingleMipTail        SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT
	SparseImageFormatAlignedMipSize       SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT
	SparseImageFormatNonstandardBlockSize SparseImageFormatFlags = C.VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT
)

var sparseImageFormatFlagsToString = map[SparseImageFormatFlags]string{
	SparseImageFormatSingleMipTail:        "Single Mip Tail",
	SparseImageFormatAlignedMipSize:       "Aligned Mip Size",
	SparseImageFormatNonstandardBlockSize: "Nonstandard Block Size",
}

func (f SparseImageFormatFlags) String() string {
	return common.FlagsToString(f, sparseImageFormatFlagsToString)
}

type SparseImageFormatProperties struct {
	AspectMask       common.ImageAspectFlags
	ImageGranularity common.Extent3D
	Flags            SparseImageFormatFlags
}

type SparseImageMemoryRequirements struct {
	FormatProperties     SparseImageFormatProperties
	ImageMipTailFirstLod int
	ImageMipTailSize     int
	ImageMipTailOffset   int
	ImageMipTailStride   int
}
