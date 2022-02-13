package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
)

type CommandBufferResetFlags int32

const (
	ResetReleaseResources CommandBufferResetFlags = C.VK_COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT
)

var resetFlagsToString = map[CommandBufferResetFlags]string{
	ResetReleaseResources: "Reset Release Resources",
}

func (f CommandBufferResetFlags) String() string {
	return common.FlagsToString(f, resetFlagsToString)
}
