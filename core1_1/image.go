package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	ImageAspectPlane0 common.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_0_BIT_KHR
	ImageAspectPlane1 common.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_1_BIT_KHR
	ImageAspectPlane2 common.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_2_BIT_KHR

	ImageCreateDisjoint common.ImageCreateFlags = C.VK_IMAGE_CREATE_DISJOINT_BIT_KHR
)

func init() {
	ImageAspectPlane0.Register("Plane 0")
	ImageAspectPlane1.Register("Plane 1")
	ImageAspectPlane2.Register("Plane 2")

	ImageCreateDisjoint.Register("Disjoint")
}
