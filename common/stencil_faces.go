package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

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
	return FlagsToString(f, stencilFacesToString)
}
