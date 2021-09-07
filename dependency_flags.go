package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "strings"

type DependencyFlags int32

const (
	DependencyByRegion    DependencyFlags = C.VK_DEPENDENCY_BY_REGION_BIT
	DependencyDeviceGroup DependencyFlags = C.VK_DEPENDENCY_DEVICE_GROUP_BIT
	DependencyViewLocal   DependencyFlags = C.VK_DEPENDENCY_VIEW_LOCAL_BIT
)

var dependencyFlagsToString = map[DependencyFlags]string{
	DependencyByRegion:    "By Region",
	DependencyDeviceGroup: "Device Group",
	DependencyViewLocal:   "View Local",
}

func (f DependencyFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		checkBit := DependencyFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := dependencyFlagsToString[checkBit]
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
