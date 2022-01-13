package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

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
	return FlagsToString(m, cullModeToString)
}
