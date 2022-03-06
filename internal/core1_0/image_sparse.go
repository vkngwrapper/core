package core1_0

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

func (i *VulkanImage) SparseMemoryRequirements() []core1_0.SparseImageMemoryRequirements {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	requirementsCount := (*C.uint32_t)(arena.Malloc(4))

	i.Driver.VkGetImageSparseMemoryRequirements(i.Device, i.ImageHandle, (*driver.Uint32)(requirementsCount), nil)

	if *requirementsCount == 0 {
		return nil
	}

	requirementsPtr := (*C.VkSparseImageMemoryRequirements)(arena.Malloc(int(*requirementsCount) * C.sizeof_struct_VkSparseImageMemoryRequirements))

	i.Driver.VkGetImageSparseMemoryRequirements(i.Device, i.ImageHandle, (*driver.Uint32)(unsafe.Pointer(requirementsCount)), (*driver.VkSparseImageMemoryRequirements)(unsafe.Pointer(requirementsPtr)))

	requirementsSlice := ([]C.VkSparseImageMemoryRequirements)(unsafe.Slice(requirementsPtr, int(*requirementsCount)))

	var outReqs []core1_0.SparseImageMemoryRequirements
	for j := 0; j < int(*requirementsCount); j++ {
		inReq := requirementsSlice[j]
		reqs := core1_0.SparseImageMemoryRequirements{
			FormatProperties: core1_0.SparseImageFormatProperties{
				AspectMask: common.ImageAspectFlags(inReq.formatProperties.aspectMask),
				ImageGranularity: common.Extent3D{
					Width:  int(inReq.formatProperties.imageGranularity.width),
					Height: int(inReq.formatProperties.imageGranularity.height),
					Depth:  int(inReq.formatProperties.imageGranularity.depth),
				},
				Flags: core1_0.SparseImageFormatFlags(inReq.formatProperties.flags),
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
