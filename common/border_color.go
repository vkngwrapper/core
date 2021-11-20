package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type BorderColor int32

const (
	BorderColorFloatTransparentBlack = C.VK_BORDER_COLOR_FLOAT_TRANSPARENT_BLACK
	BorderColorIntTransparentBlack   = C.VK_BORDER_COLOR_INT_TRANSPARENT_BLACK
	BorderColorFloatOpaqueBlack      = C.VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK
	BorderColorIntOpaqueBlack        = C.VK_BORDER_COLOR_INT_OPAQUE_BLACK
	BorderColorFloatOpaqueWhite      = C.VK_BORDER_COLOR_FLOAT_OPAQUE_WHITE
	BorderColorIntOpaqueWhite        = C.VK_BORDER_COLOR_INT_OPAQUE_WHITE
	BorderColorFloatCustomEXT        = C.VK_BORDER_COLOR_FLOAT_CUSTOM_EXT
	BorderColorIntCustomEXT          = C.VK_BORDER_COLOR_INT_CUSTOM_EXT
)

var borderColorToString = map[BorderColor]string{
	BorderColorFloatTransparentBlack: "Transparent Black - Float",
	BorderColorIntTransparentBlack:   "Transparent Black - Int",
	BorderColorFloatOpaqueBlack:      "Opaque Black - Float",
	BorderColorIntOpaqueBlack:        "Opaque Black - Int",
	BorderColorFloatOpaqueWhite:      "Opaque White - Float",
	BorderColorIntOpaqueWhite:        "Opaque White - Int",
	BorderColorFloatCustomEXT:        "Custom - Float (Extension)",
	BorderColorIntCustomEXT:          "Custom - Int (Extension)",
}

func (c BorderColor) String() string {
	return borderColorToString[c]
}
