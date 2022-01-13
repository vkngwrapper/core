package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

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
	return FlagsToString(f, dependencyFlagsToString)
}
