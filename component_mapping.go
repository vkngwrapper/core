package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type ComponentSwizzle int32

const (
	SwizzleIdentity ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_IDENTITY
	SwizzleZero     ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_ZERO
	SwizzleOne      ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_ONE
	SwizzleRed      ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_R
	SwizzleGreen    ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_G
	SwizzleBlue     ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_B
	SwizzleAlpha    ComponentSwizzle = C.VK_COMPONENT_SWIZZLE_A
)

var componentSwizzleToString = map[ComponentSwizzle]string{
	SwizzleIdentity: "Identity",
	SwizzleZero:     "Zero",
	SwizzleOne:      "One",
	SwizzleRed:      "Red",
	SwizzleGreen:    "Green",
	SwizzleBlue:     "Blue",
	SwizzleAlpha:    "Alpha",
}

func (s ComponentSwizzle) String() string {
	return componentSwizzleToString[s]
}

type ComponentMapping struct {
	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle
}
