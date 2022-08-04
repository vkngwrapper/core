package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/core1_0"

const (
	// ObjectTypeDescriptorUpdateTemplate specifies a DescriptorUpdateTemplate handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeDescriptorUpdateTemplate core1_0.ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE
	// ObjectTypeSamplerYcbcrConversion specifies a SamplerYcbcrConversion handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeSamplerYcbcrConversion core1_0.ObjectType = C.VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION
)

func init() {
	ObjectTypeDescriptorUpdateTemplate.Register("Descriptor Template")
	ObjectTypeSamplerYcbcrConversion.Register("Sampler Ycbcr Conversion")
}
