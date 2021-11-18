package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type ImageTiling int32

const (
	ImageTilingOptimal              ImageTiling = C.VK_IMAGE_TILING_OPTIMAL
	ImageTilingLinear               ImageTiling = C.VK_IMAGE_TILING_LINEAR
	ImageTilingDRMFormatModifierEXT ImageTiling = C.VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT
)

var imageTilingToString = map[ImageTiling]string{
	ImageTilingOptimal:              "Optimal",
	ImageTilingLinear:               "Linear",
	ImageTilingDRMFormatModifierEXT: "DRM Format Modifier (Extension)",
}

func (t ImageTiling) String() string {
	return imageTilingToString[t]
}
