package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

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

	QueryTypeOcclusion          QueryType = C.VK_QUERY_TYPE_OCCLUSION
	QueryTypePipelineStatistics QueryType = C.VK_QUERY_TYPE_PIPELINE_STATISTICS
	QueryTypeTimestamp          QueryType = C.VK_QUERY_TYPE_TIMESTAMP

	QueryResult64Bit            QueryResultFlags = C.VK_QUERY_RESULT_64_BIT
	QueryResultWait             QueryResultFlags = C.VK_QUERY_RESULT_WAIT_BIT
	QueryResultWithAvailability QueryResultFlags = C.VK_QUERY_RESULT_WITH_AVAILABILITY_BIT
	QueryResultPartial          QueryResultFlags = C.VK_QUERY_RESULT_PARTIAL_BIT
)

func init() {
	PipelineStatisticInputAssemblyVertices.Register("Input Assembly Vertices")
	PipelineStatisticInputAssemblyPrimitives.Register("Input Assembly Primitives")
	PipelineStatisticVertexShaderInvocations.Register("Vertex Shader Invocations")
	PipelineStatisticGeometryShaderInvocations.Register("Geometry Shader Invocations")
	PipelineStatisticGeometryShaderPrimitives.Register("GeometryShaderPrimitives")
	PipelineStatisticClippingInvocations.Register("Clipping Invocations")
	PipelineStatisticClippingPrimitives.Register("Clipping Primitives")
	PipelineStatisticFragmentShaderInvocations.Register("Fragment Shader Invocations")
	PipelineStatisticTessellationControlShaderPatches.Register("Tessellation Control Shader Patches")
	PipelineStatisticTessellationEvaluationShaderInvocations.Register("Tessellation Evaluation Shader Invocations")
	PipelineStatisticComputeShaderInvocations.Register("Compute Shader Invocations")

	QueryTypeOcclusion.Register("Occlusion")
	QueryTypePipelineStatistics.Register("Pipeline Statistics")
	QueryTypeTimestamp.Register("Timestamp")

	QueryResult64Bit.Register("64-Bit")
	QueryResultWait.Register("Wait")
	QueryResultWithAvailability.Register("With Availability")
	QueryResultPartial.Register("Partial")
}

type QueryPoolCreateOptions struct {
	QueryType          QueryType
	QueryCount         int
	PipelineStatistics PipelineStatistics

	common.NextOptions
}

func (o QueryPoolCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkQueryPoolCreateInfo)
	}
	createInfo := (*C.VkQueryPoolCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = 0
	createInfo.queryType = C.VkQueryType(o.QueryType)
	createInfo.queryCount = C.uint32_t(o.QueryCount)
	createInfo.pipelineStatistics = C.VkQueryPipelineStatisticFlags(o.PipelineStatistics)

	return preallocatedPointer, nil
}
