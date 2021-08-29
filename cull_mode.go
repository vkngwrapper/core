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
	None         CullMode = C.VK_CULL_MODE_NONE
	Front        CullMode = C.VK_CULL_MODE_FRONT_BIT
	Back         CullMode = C.VK_CULL_MODE_BACK_BIT
	FrontAndBack CullMode = C.VK_CULL_MODE_FRONT_AND_BACK
)

var cullModeToString = map[CullMode]string{
	Front: "Front",
	Back:  "Back",
}

func (m CullMode) String() string {
	if m == None {
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
