package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type StencilFaces int32

const (
	StencilFaceFront StencilFaces = C.VK_STENCIL_FACE_FRONT_BIT
	StencilFaceBack  StencilFaces = C.VK_STENCIL_FACE_BACK_BIT
)

var stencilFacesToString = map[StencilFaces]string{
	StencilFaceFront: "Stencil Front",
	StencilFaceBack:  "Stencil Back",
}

func (f StencilFaces) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := StencilFaces(1 << i)
		if (f & checkBit) != 0 {
			if hasOne {
				sb.WriteString("|")
			}
			sb.WriteString(stencilFacesToString[checkBit])
			hasOne = true
		}
	}

	return sb.String()
}
