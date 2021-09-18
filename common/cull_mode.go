package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"strings"
)

type CullModes int32

const (
	CullNone         CullModes = C.VK_CULL_MODE_NONE
	CullFront        CullModes = C.VK_CULL_MODE_FRONT_BIT
	CullBack         CullModes = C.VK_CULL_MODE_BACK_BIT
	CullFrontAndBack CullModes = C.VK_CULL_MODE_FRONT_AND_BACK
)

var cullModeToString = map[CullModes]string{
	CullFront: "Front",
	CullBack:  "Back",
}

func (m CullModes) String() string {
	if m == CullNone {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := CullModes(1 << i)
		if (m & checkBit) != 0 {
			str, hasStr := cullModeToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteString("|")
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}
