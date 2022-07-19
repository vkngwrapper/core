package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import "strings"

// ColorComponentFlags controls which components are written to the framebuffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlagBits.html
type ColorComponentFlags int32

const (
	// ColorComponentRed specifies that the R value is written to the color attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlagBits.html
	ColorComponentRed ColorComponentFlags = C.VK_COLOR_COMPONENT_R_BIT
	// ColorComponentGreen specifies that the G value is written to the color attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlagBits.html
	ColorComponentGreen ColorComponentFlags = C.VK_COLOR_COMPONENT_G_BIT
	// ColorComponentBlue specifies that the B value is written to the color attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlagBits.html
	ColorComponentBlue ColorComponentFlags = C.VK_COLOR_COMPONENT_B_BIT
	// ColorComponentAlpha specifies that the A value is written to the color attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlagBits.html
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
