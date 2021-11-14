package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type BufferUsages int

const (
	UsageTransferSrc                                BufferUsages = C.VK_BUFFER_USAGE_TRANSFER_SRC_BIT
	UsageTransferDst                                BufferUsages = C.VK_BUFFER_USAGE_TRANSFER_DST_BIT
	UsageUniformTexelBuffer                         BufferUsages = C.VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT
	UsageStorageTexelBuffer                         BufferUsages = C.VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT
	UsageUniformBuffer                              BufferUsages = C.VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT
	UsageStorageBuffer                              BufferUsages = C.VK_BUFFER_USAGE_STORAGE_BUFFER_BIT
	UsageIndexBuffer                                BufferUsages = C.VK_BUFFER_USAGE_INDEX_BUFFER_BIT
	UsageVertexBuffer                               BufferUsages = C.VK_BUFFER_USAGE_VERTEX_BUFFER_BIT
	UsageIndirectBuffer                             BufferUsages = C.VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT
	UsageShaderDeviceAddress                        BufferUsages = C.VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT
	UsageTransformFeedbackBufferEXT                 BufferUsages = C.VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT
	UsageTransformFeedbackCounterBufferEXT          BufferUsages = C.VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT
	UsageConditionalRenderingEXT                    BufferUsages = C.VK_BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT
	UsageAccelerationStructureBuildInputReadOnlyKHR BufferUsages = C.VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR
	UsageAccelerationStructureStorageKHR            BufferUsages = C.VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR
	UsageShaderBindingTableKHR                      BufferUsages = C.VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR
)

var bufferUsageToString = map[BufferUsages]string{
	UsageTransferSrc:                                "Transfer Source",
	UsageTransferDst:                                "Transfer Destination",
	UsageUniformTexelBuffer:                         "Uniform Texel Buffer",
	UsageStorageTexelBuffer:                         "Storage Texel Buffer",
	UsageUniformBuffer:                              "Uniform Buffer",
	UsageStorageBuffer:                              "Storage Buffer",
	UsageIndexBuffer:                                "Index Buffer",
	UsageVertexBuffer:                               "Vertex Buffer",
	UsageIndirectBuffer:                             "Indirect Buffer",
	UsageShaderDeviceAddress:                        "Shader Device Address",
	UsageTransformFeedbackBufferEXT:                 "Transform Feedback Buffer (Extension)",
	UsageTransformFeedbackCounterBufferEXT:          "Transform Feedback Counter Buffer (Extension)",
	UsageConditionalRenderingEXT:                    "Conditional Rendering (Extension)",
	UsageAccelerationStructureBuildInputReadOnlyKHR: "Acceleration Structure Build Input- Read Only (Khronos Extension)",
	UsageAccelerationStructureStorageKHR:            "Acceleration Structure Storage (Khronos Extension)",
	UsageShaderBindingTableKHR:                      "Shader Binding Table (Khronos Extension)",
}

func (u BufferUsages) String() string {
	if u == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := BufferUsages(1 << i)
		if (u & checkBit) != 0 {
			str, hasStr := bufferUsageToString[checkBit]
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
