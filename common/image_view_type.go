package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type ImageViewType int32

const (
	View1D        ImageViewType = C.VK_IMAGE_VIEW_TYPE_1D
	View2D        ImageViewType = C.VK_IMAGE_VIEW_TYPE_2D
	View3D        ImageViewType = C.VK_IMAGE_VIEW_TYPE_3D
	ViewCube      ImageViewType = C.VK_IMAGE_VIEW_TYPE_CUBE
	View1DArray   ImageViewType = C.VK_IMAGE_VIEW_TYPE_1D_ARRAY
	View2DArray   ImageViewType = C.VK_IMAGE_VIEW_TYPE_2D_ARRAY
	ViewCubeArray ImageViewType = C.VK_IMAGE_VIEW_TYPE_CUBE_ARRAY
)

var imageViewTypeToString = map[ImageViewType]string{
	View1D:        "1D",
	View2D:        "2D",
	View3D:        "3D",
	ViewCube:      "Cube",
	View1DArray:   "1D Array",
	View2DArray:   "2D Array",
	ViewCubeArray: "Cube Array",
}

func (t ImageViewType) String() string {
	return imageViewTypeToString[t]
}
