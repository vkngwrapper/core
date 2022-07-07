package core1_0

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/driver"
	"unsafe"
)

func (i *VulkanImage) SparseMemoryRequirements() []SparseImageMemoryRequirements {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	requirementsCount := (*C.uint32_t)(arena.Malloc(4))

	i.deviceDriver.VkGetImageSparseMemoryRequirements(i.device, i.imageHandle, (*driver.Uint32)(requirementsCount), nil)

	if *requirementsCount == 0 {
		return nil
	}

	requirementsPtr := (*C.VkSparseImageMemoryRequirements)(arena.Malloc(int(*requirementsCount) * C.sizeof_struct_VkSparseImageMemoryRequirements))

	i.deviceDriver.VkGetImageSparseMemoryRequirements(i.device, i.imageHandle, (*driver.Uint32)(unsafe.Pointer(requirementsCount)), (*driver.VkSparseImageMemoryRequirements)(unsafe.Pointer(requirementsPtr)))

	requirementsSlice := ([]C.VkSparseImageMemoryRequirements)(unsafe.Slice(requirementsPtr, int(*requirementsCount)))

	var outReqs []SparseImageMemoryRequirements
	for j := 0; j < int(*requirementsCount); j++ {
		inReq := requirementsSlice[j]
		reqs := SparseImageMemoryRequirements{
			FormatProperties: SparseImageFormatProperties{
				AspectMask: ImageAspectFlags(inReq.formatProperties.aspectMask),
				ImageGranularity: Extent3D{
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
