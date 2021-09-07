package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type CommandBufferLevel int32

const (
	LevelPrimary   CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_PRIMARY
	LevelSecondary CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_SECONDARY
	LevelUnset     CommandBufferLevel = -1
)

var commandBufferLevelToString = map[CommandBufferLevel]string{
	LevelPrimary:   "Primary",
	LevelSecondary: "Secondary",
	LevelUnset:     "",
}

func (l CommandBufferLevel) String() string {
	return commandBufferLevelToString[l]
}
