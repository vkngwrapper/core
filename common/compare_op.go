package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type CompareOp int32

const (
	CompareNever          = C.VK_COMPARE_OP_NEVER
	CompareLess           = C.VK_COMPARE_OP_LESS
	CompareEqual          = C.VK_COMPARE_OP_EQUAL
	CompareLessOrEqual    = C.VK_COMPARE_OP_LESS_OR_EQUAL
	CompareGreater        = C.VK_COMPARE_OP_GREATER
	CompareNotEqual       = C.VK_COMPARE_OP_NOT_EQUAL
	CompareGreaterOrEqual = C.VK_COMPARE_OP_GREATER_OR_EQUAL
	CompareAlways         = C.VK_COMPARE_OP_ALWAYS
)

var compareOpToString = map[CompareOp]string{
	CompareNever:          "Never",
	CompareLess:           "Less Than",
	CompareEqual:          "Equal",
	CompareLessOrEqual:    "Less Than Or Equal",
	CompareGreater:        "Greater Than",
	CompareNotEqual:       "Not Equal",
	CompareGreaterOrEqual: "Greater Than Or Equal",
	CompareAlways:         "Always",
}

func (o CompareOp) String() string {
	return compareOpToString[o]
}
