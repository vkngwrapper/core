package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/core1_0"

const (
	ObjectTypeDescriptorUpdateTemplate core1_0.ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE
	ObjectTypeSamplerYcbcrConversion   core1_0.ObjectType = C.VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION
)

func init() {
	ObjectTypeDescriptorUpdateTemplate.Register("Descriptor Template")
	ObjectTypeSamplerYcbcrConversion.Register("Sampler Ycbcr Conversion")
}
