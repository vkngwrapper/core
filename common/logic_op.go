package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type LogicOp int32

const (
	LogicOpClear        LogicOp = C.VK_LOGIC_OP_CLEAR
	LogicOpAnd          LogicOp = C.VK_LOGIC_OP_AND
	LogicOpAndReverse   LogicOp = C.VK_LOGIC_OP_AND_REVERSE
	LogicOpCopy         LogicOp = C.VK_LOGIC_OP_COPY
	LogicOpAndInverted  LogicOp = C.VK_LOGIC_OP_AND_INVERTED
	LogicOpNoop         LogicOp = C.VK_LOGIC_OP_NO_OP
	LogicOpXor          LogicOp = C.VK_LOGIC_OP_XOR
	LogicOpOr           LogicOp = C.VK_LOGIC_OP_OR
	LogicOpNor          LogicOp = C.VK_LOGIC_OP_NOR
	LogicOpEquivalent   LogicOp = C.VK_LOGIC_OP_EQUIVALENT
	LogicOpInvert       LogicOp = C.VK_LOGIC_OP_INVERT
	LogicOpOrReverse    LogicOp = C.VK_LOGIC_OP_OR_REVERSE
	LogicOpCopyInverted LogicOp = C.VK_LOGIC_OP_COPY_INVERTED
	LogicOpOrInverted   LogicOp = C.VK_LOGIC_OP_OR_INVERTED
	LogicOpNand         LogicOp = C.VK_LOGIC_OP_NAND
	LogicOpSet          LogicOp = C.VK_LOGIC_OP_SET
)

var logicOpToString = map[LogicOp]string{
	LogicOpClear:        "Clear",
	LogicOpAnd:          "And",
	LogicOpAndReverse:   "And Reverse",
	LogicOpCopy:         "Copy",
	LogicOpAndInverted:  "And Inverted",
	LogicOpNoop:         "No-Op",
	LogicOpXor:          "Xor",
	LogicOpOr:           "Or",
	LogicOpNor:          "Nor",
	LogicOpEquivalent:   "Equivalent",
	LogicOpInvert:       "Invert",
	LogicOpOrReverse:    "Or Reverse",
	LogicOpCopyInverted: "Copy Inverted",
	LogicOpOrInverted:   "Or Inverted",
	LogicOpNand:         "Nand",
	LogicOpSet:          "Set",
}

func (o LogicOp) String() string {
	return logicOpToString[o]
}
