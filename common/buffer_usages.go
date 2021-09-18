package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type BufferUsages int

const (
	UsageTransferSrc                             BufferUsages = C.VK_BUFFER_USAGE_TRANSFER_SRC_BIT
	UsageTransferDst                             BufferUsages = C.VK_BUFFER_USAGE_TRANSFER_DST_BIT
	UsageUniformTexelBuffer                      BufferUsages = C.VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT
	UsageStorageTexelBuffer                      BufferUsages = C.VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT
	UsageUniformBuffer                           BufferUsages = C.VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT
	UsageStorageBuffer                           BufferUsages = C.VK_BUFFER_USAGE_STORAGE_BUFFER_BIT
	UsageIndexBuffer                             BufferUsages = C.VK_BUFFER_USAGE_INDEX_BUFFER_BIT
	UsageVertexBuffer                            BufferUsages = C.VK_BUFFER_USAGE_VERTEX_BUFFER_BIT
	UsageIndirectBuffer                          BufferUsages = C.VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT
	UsageShaderDeviceAddress                     BufferUsages = C.VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT
	UsageTransformFeedbackBuffer                 BufferUsages = C.VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT
	UsageTransformFeedbackCounterBuffer          BufferUsages = C.VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT
	UsageConditionalRendering                    BufferUsages = C.VK_BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT
	UsageAccelerationStructureBuildInputReadOnly BufferUsages = C.VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR
	UsageAccelerationStructureStorage            BufferUsages = C.VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR
	UsageShaderBindingTable                      BufferUsages = C.VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR
)

var bufferUsageToString = map[BufferUsages]string{
	UsageTransferSrc:                             "Transfer Source",
	UsageTransferDst:                             "Transfer Destination",
	UsageUniformTexelBuffer:                      "Uniform Texel Buffer",
	UsageStorageTexelBuffer:                      "Storage Texel Buffer",
	UsageUniformBuffer:                           "Uniform Buffer",
	UsageStorageBuffer:                           "Storage Buffer",
	UsageIndexBuffer:                             "Index Buffer",
	UsageVertexBuffer:                            "Vertex Buffer",
	UsageIndirectBuffer:                          "Indirect Buffer",
	UsageShaderDeviceAddress:                     "Shader Device Address",
	UsageTransformFeedbackBuffer:                 "Transform Feedback Buffer",
	UsageTransformFeedbackCounterBuffer:          "Transform Feedback Counter Buffer",
	UsageConditionalRendering:                    "Conditional Rendering",
	UsageAccelerationStructureBuildInputReadOnly: "Acceleration Structure Build Input- Read Only",
	UsageAccelerationStructureStorage:            "Acceleration Structure Storage",
	UsageShaderBindingTable:                      "Shader Binding Table",
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
