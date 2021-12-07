package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type ImageViewType int32

const (
	ViewType1D        ImageViewType = C.VK_IMAGE_VIEW_TYPE_1D
	ViewType2D        ImageViewType = C.VK_IMAGE_VIEW_TYPE_2D
	ViewType3D        ImageViewType = C.VK_IMAGE_VIEW_TYPE_3D
	ViewTypeCube      ImageViewType = C.VK_IMAGE_VIEW_TYPE_CUBE
	ViewType1DArray   ImageViewType = C.VK_IMAGE_VIEW_TYPE_1D_ARRAY
	ViewType2DArray   ImageViewType = C.VK_IMAGE_VIEW_TYPE_2D_ARRAY
	ViewTypeCubeArray ImageViewType = C.VK_IMAGE_VIEW_TYPE_CUBE_ARRAY
)

var imageViewTypeToString = map[ImageViewType]string{
	ViewType1D:        "1D",
	ViewType2D:        "2D",
	ViewType3D:        "3D",
	ViewTypeCube:      "Cube",
	ViewType1DArray:   "1D Array",
	ViewType2DArray:   "2D Array",
	ViewTypeCubeArray: "Cube Array",
}

func (t ImageViewType) String() string {
	return imageViewTypeToString[t]
}
