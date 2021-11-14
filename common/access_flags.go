package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type AccessFlags int32

const (
	AccessNoneKHR                              AccessFlags = C.VK_ACCESS_NONE_KHR
	AccessIndirectCommandRead                  AccessFlags = C.VK_ACCESS_INDIRECT_COMMAND_READ_BIT
	AccessIndexRead                            AccessFlags = C.VK_ACCESS_INDEX_READ_BIT
	AccessVertexAttributeRead                  AccessFlags = C.VK_ACCESS_VERTEX_ATTRIBUTE_READ_BIT
	AccessUniformRead                          AccessFlags = C.VK_ACCESS_UNIFORM_READ_BIT
	AccessInputAttachmentRead                  AccessFlags = C.VK_ACCESS_INPUT_ATTACHMENT_READ_BIT
	AccessShaderRead                           AccessFlags = C.VK_ACCESS_SHADER_READ_BIT
	AccessShaderWrite                          AccessFlags = C.VK_ACCESS_SHADER_WRITE_BIT
	AccessColorAttachmentRead                  AccessFlags = C.VK_ACCESS_COLOR_ATTACHMENT_READ_BIT
	AccessColorAttachmentWrite                 AccessFlags = C.VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT
	AccessDepthStencilAttachmentRead           AccessFlags = C.VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT
	AccessDepthStencilAttachmentWrite          AccessFlags = C.VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT
	AccessTransferRead                         AccessFlags = C.VK_ACCESS_TRANSFER_READ_BIT
	AccessTransferWrite                        AccessFlags = C.VK_ACCESS_TRANSFER_WRITE_BIT
	AccessHostRead                             AccessFlags = C.VK_ACCESS_HOST_READ_BIT
	AccessHostWrite                            AccessFlags = C.VK_ACCESS_HOST_WRITE_BIT
	AccessMemoryRead                           AccessFlags = C.VK_ACCESS_MEMORY_READ_BIT
	AccessMemoryWrite                          AccessFlags = C.VK_ACCESS_MEMORY_WRITE_BIT
	AccessTransformFeedbackWriteEXT            AccessFlags = C.VK_ACCESS_TRANSFORM_FEEDBACK_WRITE_BIT_EXT
	AccessTransformFeedbackCounterReadEXT      AccessFlags = C.VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT
	AccessTransformFeedbackCounterWriteEXT     AccessFlags = C.VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT
	AccessConditionalRenderingReadEXT          AccessFlags = C.VK_ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT
	AccessColorAttachmentReadNonCoherentEXT    AccessFlags = C.VK_ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT
	AccessAccelerationStructureReadKHR         AccessFlags = C.VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR
	AccessAccelerationStructureWriteKHR        AccessFlags = C.VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR
	AccessFragmentDensityMapReadEXT            AccessFlags = C.VK_ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT
	AccessFragmentShadingRateAttachmentReadKHR AccessFlags = C.VK_ACCESS_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR
	AccessPreProcessReadNV                     AccessFlags = C.VK_ACCESS_COMMAND_PREPROCESS_READ_BIT_NV
	AccessPreProcessWriteNV                    AccessFlags = C.VK_ACCESS_COMMAND_PREPROCESS_WRITE_BIT_NV
)

var accessFlagsToString = map[AccessFlags]string{
	AccessIndirectCommandRead:                  "Indirect Command Read",
	AccessIndexRead:                            "Index Read",
	AccessVertexAttributeRead:                  "Vertex Attribute Read",
	AccessUniformRead:                          "Uniform Read",
	AccessInputAttachmentRead:                  "Input Attachment Read",
	AccessShaderRead:                           "Shader Read",
	AccessShaderWrite:                          "Shader Write",
	AccessColorAttachmentRead:                  "Color Attachment Read",
	AccessColorAttachmentWrite:                 "Color Attachment Write",
	AccessDepthStencilAttachmentRead:           "Depth/Stencil Attachment Read",
	AccessDepthStencilAttachmentWrite:          "Depth/Stencil Attachment Write",
	AccessTransferRead:                         "Transfer Read",
	AccessTransferWrite:                        "Transfer Write",
	AccessHostRead:                             "Host Read",
	AccessHostWrite:                            "Host Write",
	AccessMemoryRead:                           "Memory Read",
	AccessMemoryWrite:                          "Memory Write",
	AccessTransformFeedbackWriteEXT:            "Transform Feedback Write (Extension)",
	AccessTransformFeedbackCounterReadEXT:      "Transform Feedback Counter Read (Extension)",
	AccessTransformFeedbackCounterWriteEXT:     "Transform Feedback Counter Write (Extension)",
	AccessConditionalRenderingReadEXT:          "Conditional Rendering Read (Extension)",
	AccessColorAttachmentReadNonCoherentEXT:    "Color Attachment Read Non-Coherent (Extension)",
	AccessAccelerationStructureReadKHR:         "Acceleration Structure Read (Khronos Extension)",
	AccessAccelerationStructureWriteKHR:        "Acceleration Structure Write (Khronos Extension)",
	AccessFragmentDensityMapReadEXT:            "Fragment Density Map Read (Extension)",
	AccessFragmentShadingRateAttachmentReadKHR: "Fragment Shading Rate Attachment Read (Khronos Extension)",
	AccessPreProcessReadNV:                     "Pre-Process Read (Nvidia Extension)",
	AccessPreProcessWriteNV:                    "Pre-Process Write (Nvidia Extension)",
}

func (f AccessFlags) String() string {
	if f == 0 {
		return "None (Khronos Extension)"
	}

	var hasOne bool
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		checkBit := AccessFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := accessFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}
