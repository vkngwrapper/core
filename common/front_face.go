package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type FrontFace int32

const (
	FrontFaceCounterClockwise FrontFace = C.VK_FRONT_FACE_COUNTER_CLOCKWISE
	FrontFaceClockwise        FrontFace = C.VK_FRONT_FACE_CLOCKWISE
)

var frontFaceToString = map[FrontFace]string{
	FrontFaceCounterClockwise: "Counter-Clockwise",
	FrontFaceClockwise:        "Clockwise",
}

func (f FrontFace) String() string {
	return frontFaceToString[f]
}
