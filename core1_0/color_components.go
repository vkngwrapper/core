package core1_0

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type ColorComponentFlags int32

const (
	ColorComponentRed   ColorComponentFlags = C.VK_COLOR_COMPONENT_R_BIT
	ColorComponentGreen ColorComponentFlags = C.VK_COLOR_COMPONENT_G_BIT
	ColorComponentBlue  ColorComponentFlags = C.VK_COLOR_COMPONENT_B_BIT
	ColorComponentAlpha ColorComponentFlags = C.VK_COLOR_COMPONENT_A_BIT
)

func (c ColorComponentFlags) String() string {
	if c == 0 {
		return "None"
	}

	var sb strings.Builder
	if (c & ColorComponentRed) != 0 {
		sb.WriteRune('R')
	}
	if (c & ColorComponentGreen) != 0 {
		sb.WriteRune('G')
	}
	if (c & ColorComponentBlue) != 0 {
		sb.WriteRune('B')
	}
	if (c & ColorComponentAlpha) != 0 {
		sb.WriteRune('A')
	}

	return sb.String()
}
