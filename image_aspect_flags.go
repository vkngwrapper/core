package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "strings"

type ImageAspectFlags int32

const (
	AspectColor        ImageAspectFlags = C.VK_IMAGE_ASPECT_COLOR_BIT
	AspectDepth        ImageAspectFlags = C.VK_IMAGE_ASPECT_DEPTH_BIT
	AspectStencil      ImageAspectFlags = C.VK_IMAGE_ASPECT_STENCIL_BIT
	AspectMetadata     ImageAspectFlags = C.VK_IMAGE_ASPECT_METADATA_BIT
	AspectPlane0       ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_0_BIT
	AspectPlane1       ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_1_BIT
	AspectPlane2       ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_2_BIT
	AspectMemoryPlane0 ImageAspectFlags = C.VK_IMAGE_ASPECT_MEMORY_PLANE_0_BIT_EXT
	AspectMemoryPlane1 ImageAspectFlags = C.VK_IMAGE_ASPECT_MEMORY_PLANE_1_BIT_EXT
	AspectMemoryPlane2 ImageAspectFlags = C.VK_IMAGE_ASPECT_MEMORY_PLANE_2_BIT_EXT
	AspectMemoryPlane3 ImageAspectFlags = C.VK_IMAGE_ASPECT_MEMORY_PLANE_3_BIT_EXT
)

var imageAspectToString = map[ImageAspectFlags]string{
	AspectColor:        "Color",
	AspectDepth:        "Depth",
	AspectStencil:      "Stencil",
	AspectMetadata:     "Metaddata",
	AspectPlane0:       "Plane 0",
	AspectPlane1:       "Plane 1",
	AspectPlane2:       "Plane 2",
	AspectMemoryPlane0: "Memory Plane 0",
	AspectMemoryPlane1: "Memory Plane 1",
	AspectMemoryPlane2: "Memory Plane 2",
	AspectMemoryPlane3: "Memory Plane 3",
}

func (f ImageAspectFlags) String() string {
	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := ImageAspectFlags(1 << i)
		if (f & checkBit) != 0 {
			if hasOne {
				sb.WriteString("|")
			}
			sb.WriteString(imageAspectToString[checkBit])
			hasOne = true
		}
	}

	return sb.String()
}
