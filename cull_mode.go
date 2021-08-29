package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"strings"
)

type CullMode int32

const (
	CullNone         CullMode = C.VK_CULL_MODE_NONE
	CullFront        CullMode = C.VK_CULL_MODE_FRONT_BIT
	CullBack         CullMode = C.VK_CULL_MODE_BACK_BIT
	CullFrontAndBack CullMode = C.VK_CULL_MODE_FRONT_AND_BACK
)

var cullModeToString = map[CullMode]string{
	CullFront: "Front",
	CullBack:  "Back",
}

func (m CullMode) String() string {
	if m == CullNone {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := CullMode(1 << i)
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
