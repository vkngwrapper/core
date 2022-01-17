package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
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

func (i *vulkanImage) SparseMemoryRequirements() []SparseImageMemoryRequirements {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	requirementsCount := (*C.uint32_t)(arena.Malloc(4))

	i.driver.VkGetImageSparseMemoryRequirements(i.device, i.handle, (*driver.Uint32)(requirementsCount), nil)

	if *requirementsCount == 0 {
		return nil
	}

	requirementsPtr := (*C.VkSparseImageMemoryRequirements)(arena.Malloc(int(*requirementsCount) * C.sizeof_struct_VkSparseImageMemoryRequirements))

	i.driver.VkGetImageSparseMemoryRequirements(i.device, i.handle, (*driver.Uint32)(unsafe.Pointer(requirementsCount)), (*driver.VkSparseImageMemoryRequirements)(unsafe.Pointer(requirementsPtr)))

	requirementsSlice := ([]C.VkSparseImageMemoryRequirements)(unsafe.Slice(requirementsPtr, int(*requirementsCount)))

	var outReqs []SparseImageMemoryRequirements
	for j := 0; j < int(*requirementsCount); j++ {
		inReq := requirementsSlice[j]
		reqs := SparseImageMemoryRequirements{
			FormatProperties: SparseImageFormatProperties{
				AspectMask: common.ImageAspectFlags(inReq.formatProperties.aspectMask),
				ImageGranularity: common.Extent3D{
					Width:  int(inReq.formatProperties.imageGranularity.width),
					Height: int(inReq.formatProperties.imageGranularity.height),
					Depth:  int(inReq.formatProperties.imageGranularity.depth),
				},
				Flags: SparseImageFormatFlags(inReq.formatProperties.flags),
			},
			ImageMipTailFirstLod: int(inReq.imageMipTailFirstLod),
			ImageMipTailOffset:   int(inReq.imageMipTailOffset),
			ImageMipTailSize:     int(inReq.imageMipTailSize),
			ImageMipTailStride:   int(inReq.imageMipTailStride),
		}

		outReqs = append(outReqs, reqs)
	}

	return outReqs
}
