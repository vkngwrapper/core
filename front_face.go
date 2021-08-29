package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type FrontFace int32

const (
	CounterClockwise FrontFace = C.VK_FRONT_FACE_COUNTER_CLOCKWISE
	Clockwise        FrontFace = C.VK_FRONT_FACE_CLOCKWISE
)

var frontFaceToString = map[FrontFace]string{
	CounterClockwise: "Counter-Clockwise",
	Clockwise:        "Clockwise",
}

func (f FrontFace) String() string {
	return frontFaceToString[f]
}
