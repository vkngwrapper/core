package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type CommandBufferLevel int32

const (
	Primary   CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_PRIMARY
	Secondary CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_SECONDARY
	Unset     CommandBufferLevel = -1
)

var commandBufferLevelToString = map[CommandBufferLevel]string{
	Primary:   "Primary",
	Secondary: "Secondary",
	Unset:     "",
}

func (l CommandBufferLevel) String() string {
	return commandBufferLevelToString[l]
}
