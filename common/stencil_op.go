package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type StencilOp int32

const (
	StencilKeep              StencilOp = C.VK_STENCIL_OP_KEEP
	StencilZero              StencilOp = C.VK_STENCIL_OP_ZERO
	StencilReplace           StencilOp = C.VK_STENCIL_OP_REPLACE
	StencilIncrementAndClamp StencilOp = C.VK_STENCIL_OP_INCREMENT_AND_CLAMP
	StencilDecrementAndClamp StencilOp = C.VK_STENCIL_OP_DECREMENT_AND_CLAMP
	StencilInvert            StencilOp = C.VK_STENCIL_OP_INVERT
	StencilIncrementAndWrap  StencilOp = C.VK_STENCIL_OP_INCREMENT_AND_WRAP
	StencilDecrementAndWrap  StencilOp = C.VK_STENCIL_OP_DECREMENT_AND_WRAP
)

var stencilOpToString = map[StencilOp]string{
	StencilKeep:              "Keep",
	StencilZero:              "Zero",
	StencilReplace:           "Replace",
	StencilIncrementAndClamp: "Increment and Clamp",
	StencilDecrementAndClamp: "Decrement and Clamp",
	StencilInvert:            "Invert",
	StencilIncrementAndWrap:  "Increment and Wrap",
	StencilDecrementAndWrap:  "Decrement and Wrap",
}

func (o StencilOp) String() string {
	return stencilOpToString[o]
}
