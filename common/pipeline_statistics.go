package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type PipelineStatistics int32

const (
	PipelineStatisticInputAssemblyVertices                   PipelineStatistics = C.VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT
	PipelineStatisticInputAssemblyPrimitives                 PipelineStatistics = C.VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT
	PipelineStatisticVertexShaderInvocations                 PipelineStatistics = C.VK_QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT
	PipelineStatisticGeometryShaderInvocations               PipelineStatistics = C.VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT
	PipelineStatisticGeometryShaderPrimitives                PipelineStatistics = C.VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT
	PipelineStatisticClippingInvocations                     PipelineStatistics = C.VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT
	PipelineStatisticClippingPrimitives                      PipelineStatistics = C.VK_QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT
	PipelineStatisticFragmentShaderInvocations               PipelineStatistics = C.VK_QUERY_PIPELINE_STATISTIC_FRAGMENT_SHADER_INVOCATIONS_BIT
	PipelineStatisticTessellationControlShaderPatches        PipelineStatistics = C.VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT
	PipelineStatisticTessellationEvaluationShaderInvocations PipelineStatistics = C.VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT
	PipelineStatisticComputeShaderInvocations                PipelineStatistics = C.VK_QUERY_PIPELINE_STATISTIC_COMPUTE_SHADER_INVOCATIONS_BIT
)

var pipelineStatisticsToString = map[PipelineStatistics]string{
	PipelineStatisticInputAssemblyVertices:                   "Input Assembly Vertices",
	PipelineStatisticInputAssemblyPrimitives:                 "Input Assembly Primitives",
	PipelineStatisticVertexShaderInvocations:                 "Vertex Shader Invocations",
	PipelineStatisticGeometryShaderInvocations:               "Geometry Shader Invocations",
	PipelineStatisticGeometryShaderPrimitives:                "GeometryShaderPrimitives",
	PipelineStatisticClippingInvocations:                     "Clipping Invocations",
	PipelineStatisticClippingPrimitives:                      "Clipping Primitives",
	PipelineStatisticFragmentShaderInvocations:               "Fragment Shader Invocations",
	PipelineStatisticTessellationControlShaderPatches:        "Tessellation Control Shader Patches",
	PipelineStatisticTessellationEvaluationShaderInvocations: "Tessellation Evaluation Shader Invocations",
	PipelineStatisticComputeShaderInvocations:                "Compute Shader Invocations",
}

func (f PipelineStatistics) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		checkBit := PipelineStatistics(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := pipelineStatisticsToString[checkBit]
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
