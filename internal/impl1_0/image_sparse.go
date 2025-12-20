package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

func (i *VulkanImage) SparseMemoryRequirements() []core1_0.SparseImageMemoryRequirements {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	requirementsCount := (*C.uint32_t)(arena.Malloc(4))

	i.DeviceDriver.VkGetImageSparseMemoryRequirements(i.Device, i.ImageHandle, (*driver.Uint32)(requirementsCount), nil)

	if *requirementsCount == 0 {
		return nil
	}

	requirementsPtr := (*C.VkSparseImageMemoryRequirements)(arena.Malloc(int(*requirementsCount) * C.sizeof_struct_VkSparseImageMemoryRequirements))

	i.DeviceDriver.VkGetImageSparseMemoryRequirements(i.Device, i.ImageHandle, (*driver.Uint32)(unsafe.Pointer(requirementsCount)), (*driver.VkSparseImageMemoryRequirements)(unsafe.Pointer(requirementsPtr)))

	requirementsSlice := ([]C.VkSparseImageMemoryRequirements)(unsafe.Slice(requirementsPtr, int(*requirementsCount)))

	var outReqs []core1_0.SparseImageMemoryRequirements
	for j := 0; j < int(*requirementsCount); j++ {
		inReq := requirementsSlice[j]
		reqs := core1_0.SparseImageMemoryRequirements{
			FormatProperties: core1_0.SparseImageFormatProperties{
				AspectMask: core1_0.ImageAspectFlags(inReq.formatProperties.aspectMask),
				ImageGranularity: core1_0.Extent3D{
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
