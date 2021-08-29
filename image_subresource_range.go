package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type ImageSubresourceRange struct {
	AspectMask     ImageAspectFlags
	BaseMipLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}
