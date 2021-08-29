package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type LogicOp int32

const (
	OpClear        LogicOp = C.VK_LOGIC_OP_CLEAR
	OpAnd          LogicOp = C.VK_LOGIC_OP_AND
	OpAndReverse   LogicOp = C.VK_LOGIC_OP_AND_REVERSE
	OpCopy         LogicOp = C.VK_LOGIC_OP_COPY
	OpAndInverted  LogicOp = C.VK_LOGIC_OP_AND_INVERTED
	OpNoop         LogicOp = C.VK_LOGIC_OP_NO_OP
	OpXor          LogicOp = C.VK_LOGIC_OP_XOR
	OpOr           LogicOp = C.VK_LOGIC_OP_OR
	OpNor          LogicOp = C.VK_LOGIC_OP_NOR
	OpEquivalent   LogicOp = C.VK_LOGIC_OP_EQUIVALENT
	OpInvert       LogicOp = C.VK_LOGIC_OP_INVERT
	OpOrReverse    LogicOp = C.VK_LOGIC_OP_OR_REVERSE
	OpCopyInverted LogicOp = C.VK_LOGIC_OP_COPY_INVERTED
	OpOrInverted   LogicOp = C.VK_LOGIC_OP_OR_INVERTED
	OpNand         LogicOp = C.VK_LOGIC_OP_NAND
	OpSet          LogicOp = C.VK_LOGIC_OP_SET
)

var logicOpToString = map[LogicOp]string{
	OpClear:        "Clear",
	OpAnd:          "And",
	OpAndReverse:   "And Reverse",
	OpCopy:         "Copy",
	OpAndInverted:  "And Inverted",
	OpNoop:         "No-Op",
	OpXor:          "Xor",
	OpOr:           "Or",
	OpNor:          "Nor",
	OpEquivalent:   "Equivalent",
	OpInvert:       "Invert",
	OpOrReverse:    "Or Reverse",
	OpCopyInverted: "Copy Inverted",
	OpOrInverted:   "Or Inverted",
	OpNand:         "Nand",
	OpSet:          "Set",
}

func (o LogicOp) String() string {
	return logicOpToString[o]
}
