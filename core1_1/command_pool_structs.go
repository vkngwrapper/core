package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
)

type CommandPoolTrimFlags int32

var commandPoolTrimFlagsMapping = common.NewFlagStringMapping[CommandPoolTrimFlags]()

func (f CommandPoolTrimFlags) Register(str string) {
	commandPoolTrimFlagsMapping.Register(f, str)
}
func (f CommandPoolTrimFlags) String() string {
	return commandPoolTrimFlagsMapping.FlagsToString(f)
}

////

const (
	CommandPoolCreateProtected core1_0.CommandPoolCreateFlags = C.VK_COMMAND_POOL_CREATE_PROTECTED_BIT
)

func init() {
	CommandPoolCreateProtected.Register("Protected")
}
