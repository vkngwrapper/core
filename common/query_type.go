package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type QueryType int32

const (
	QueryTypeOcclusion                                 QueryType = C.VK_QUERY_TYPE_OCCLUSION
	QueryTypePipelineStatistics                        QueryType = C.VK_QUERY_TYPE_PIPELINE_STATISTICS
	QueryTypeTimestamp                                 QueryType = C.VK_QUERY_TYPE_TIMESTAMP
	QueryTypeTransformFeedbackStreamEXT                QueryType = C.VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT
	QueryTypePerformanceQueryKHR                       QueryType = C.VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR
	QueryTypeAccelerationStructureCompactedSizeKHR     QueryType = C.VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR
	QueryTypeAccelerationStructureSerializationSizeKHR QueryType = C.VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR
	QueryTypeAccelerationStructureCompactedSizeNV      QueryType = C.VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV
	QueryTypePerformanceQueryIntel                     QueryType = C.VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL
)

var queryTypeToString = map[QueryType]string{
	QueryTypeOcclusion:                                 "Occlusion",
	QueryTypePipelineStatistics:                        "Pipeline Statistics",
	QueryTypeTimestamp:                                 "Timestamp",
	QueryTypeTransformFeedbackStreamEXT:                "Transform Feedback Stream (Extension)",
	QueryTypePerformanceQueryKHR:                       "Performance Query (Khronos extension)",
	QueryTypeAccelerationStructureCompactedSizeKHR:     "Acceleration Structure Compacted Size (Khronos Extension)",
	QueryTypeAccelerationStructureSerializationSizeKHR: "Acceleration Structure Serialization Size (Khronos Extension)",
	QueryTypeAccelerationStructureCompactedSizeNV:      "Acceleration Structure Compacted Size (Nvidia Extension)",
	QueryTypePerformanceQueryIntel:                     "Performance Query (Intel Extension)",
}

func (f QueryType) String() string {
	return queryTypeToString[f]
}
