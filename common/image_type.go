package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type ImageType int32

const (
	ImageType1D ImageType = C.VK_IMAGE_TYPE_1D
	ImageType2D ImageType = C.VK_IMAGE_TYPE_2D
	ImageType3D ImageType = C.VK_IMAGE_TYPE_3D
)

var imageTypeToString = map[ImageType]string{
	ImageType1D: "1D",
	ImageType2D: "2D",
	ImageType3D: "3D",
}

func (t ImageType) String() string {
	return imageTypeToString[t]
}
