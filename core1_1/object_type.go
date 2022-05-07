package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	ObjectTypeDescriptorTemplate     common.ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE
	ObjectTypeSamplerYcbcrConversion common.ObjectType = C.VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION
)

func init() {
	ObjectTypeDescriptorTemplate.Register("Descriptor Template")
	ObjectTypeSamplerYcbcrConversion.Register("Sampler Ycbcr Conversion")
}
