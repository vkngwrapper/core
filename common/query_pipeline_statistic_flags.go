package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

type QueryPipelineStatisticFlags int32

const (
	StatisticInputAssemblyVertices                   QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT
	StatisticInputAssemblyPrimitives                 QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT
	StatisticVertexShaderInvocations                 QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT
	StatisticGeometryShaderInvocations               QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT
	StatisticGeometryShaderPrimitives                QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT
	StatisticClippingInvocations                     QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT
	StatisticClippingPrimitives                      QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT
	StatisticFragmentShaderInvocations               QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_FRAGMENT_SHADER_INVOCATIONS_BIT
	StatisticTessellationControlShaderPatches        QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT
	StatisticTessellationEvaluationShaderInvocations QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT
	StatisticComputeShaderInvocations                QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_COMPUTE_SHADER_INVOCATIONS_BIT
)

var queryPipelineStatisticFlagsToString = map[QueryPipelineStatisticFlags]string{
	StatisticInputAssemblyVertices:                   "Input Assembly Vertices",
	StatisticInputAssemblyPrimitives:                 "Input Assembly Primitives",
	StatisticVertexShaderInvocations:                 "Vertex Shader Invocations",
	StatisticGeometryShaderInvocations:               "Geometry Shader Invocations",
	StatisticGeometryShaderPrimitives:                "Geometry Shader Primitives",
	StatisticClippingInvocations:                     "Clipping Invocations",
	StatisticClippingPrimitives:                      "Clipping Primitives",
	StatisticFragmentShaderInvocations:               "Fragment Shader Invocations",
	StatisticTessellationControlShaderPatches:        "Tessellation Control Shader Patches",
	StatisticTessellationEvaluationShaderInvocations: "Tessellation Evaluation Shader Invocations",
	StatisticComputeShaderInvocations:                "Compute Shader Invocations",
}

func (f QueryPipelineStatisticFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := QueryPipelineStatisticFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := queryPipelineStatisticFlagsToString[checkBit]
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
