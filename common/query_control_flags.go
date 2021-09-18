package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type QueryControlFlags int32

const (
	QueryPrecise QueryControlFlags = C.VK_QUERY_CONTROL_PRECISE_BIT
)

var queryControlFlagsToString = map[QueryControlFlags]string{
	QueryPrecise: "Precise",
}

func (f QueryControlFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := QueryControlFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := queryControlFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}
