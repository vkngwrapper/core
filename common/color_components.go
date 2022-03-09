package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type ColorComponents int32

const (
	ComponentRed   ColorComponents = C.VK_COLOR_COMPONENT_R_BIT
	ComponentGreen ColorComponents = C.VK_COLOR_COMPONENT_G_BIT
	ComponentBlue  ColorComponents = C.VK_COLOR_COMPONENT_B_BIT
	ComponentAlpha ColorComponents = C.VK_COLOR_COMPONENT_A_BIT
)

func (c ColorComponents) String() string {
	if c == 0 {
		return "None"
	}

	var sb strings.Builder
	if (c & ComponentRed) != 0 {
		sb.WriteRune('R')
	}
	if (c & ComponentGreen) != 0 {
		sb.WriteRune('G')
	}
	if (c & ComponentBlue) != 0 {
		sb.WriteRune('B')
	}
	if (c & ComponentAlpha) != 0 {
		sb.WriteRune('A')
	}

	return sb.String()
}
