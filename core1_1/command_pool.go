package core1_1

import "github.com/CannibalVox/VKng/core/common"

type CommandPoolTrimFlags int32

var commandPoolTrimFlagsMapping = common.NewFlagStringMapping[CommandPoolTrimFlags]()

func (f CommandPoolTrimFlags) Register(str string) {
	commandPoolTrimFlagsMapping.Register(f, str)
}
func (f CommandPoolTrimFlags) String() string {
	return commandPoolTrimFlagsMapping.FlagsToString(f)
}
