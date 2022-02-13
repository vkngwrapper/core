package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type CommandPoolResetFlags int32

const (
	CommandPoolResetReleaseResources CommandPoolResetFlags = C.VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT
)

var commandPoolResetFlagsToString = map[CommandPoolResetFlags]string{
	CommandPoolResetReleaseResources: "Release Resources",
}

func (f CommandPoolResetFlags) String() string {
	return common.FlagsToString(f, commandPoolResetFlagsToString)
}
