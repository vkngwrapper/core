package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

// QueryPoolCreateFlags is reserved for future use
type QueryPoolCreateFlags uint32

var queryPoolCreateFlags = common.NewFlagStringMapping[QueryPoolCreateFlags]()

func (f QueryPoolCreateFlags) Register(str string) {
	queryPoolCreateFlags.Register(f, str)
}

func (f QueryPoolCreateFlags) String() string {
	return queryPoolCreateFlags.FlagsToString(f)
}

////

const (
	// QueryTypeOcclusion specifies an occlusion query
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html
	QueryTypeOcclusion QueryType = C.VK_QUERY_TYPE_OCCLUSION
	// QueryTypePipelineStatistics specifies a pipeline statistics query
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html
	QueryTypePipelineStatistics QueryType = C.VK_QUERY_TYPE_PIPELINE_STATISTICS
	// QueryTypeTimestamp specifies a timestamp query
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html
	QueryTypeTimestamp QueryType = C.VK_QUERY_TYPE_TIMESTAMP

	// QueryResult64Bit specifies the results will be written as an array of 64-bit unsigned
	// integer values (instead of 32-bit unsigned integer values)
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlagBits.html
	QueryResult64Bit QueryResultFlags = C.VK_QUERY_RESULT_64_BIT
	// QueryResultWait specifies that Vulkan will wait for each query's status to become available
	// before retrieving its results
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlagBits.html
	QueryResultWait QueryResultFlags = C.VK_QUERY_RESULT_WAIT_BIT
	// QueryResultWithAvailability specifies that the availability status accompanies the results
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlagBits.html
	QueryResultWithAvailability QueryResultFlags = C.VK_QUERY_RESULT_WITH_AVAILABILITY_BIT
	// QueryResultPartial specifies that returning partial results is acceptable
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlagBits.html
	QueryResultPartial QueryResultFlags = C.VK_QUERY_RESULT_PARTIAL_BIT
)

func init() {
	QueryTypeOcclusion.Register("Occlusion")
	QueryTypePipelineStatistics.Register("Pipeline Statistics")
	QueryTypeTimestamp.Register("Timestamp")

	QueryResult64Bit.Register("64-Bit")
	QueryResultWait.Register("Wait")
	QueryResultWithAvailability.Register("With Availability")
	QueryResultPartial.Register("Partial")
}

// QueryPoolCreateInfo specifies parameters of a newly-created QueryPool
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html
type QueryPoolCreateInfo struct {
	// Flags is reserved for future use
	Flags QueryPoolCreateFlags

	// QueryType specifies the type of queries managed by the QueryPool
	QueryType QueryType
	// QueryCount is the number of queries managed by the QueryPool
	QueryCount int
	// PipelineStatistics specifies which counters will be returned in queries on the
	// new QueryPool
	PipelineStatistics QueryPipelineStatisticFlags

	common.NextOptions
}

func (o QueryPoolCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkQueryPoolCreateInfo)
	}
	createInfo := (*C.VkQueryPoolCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkQueryPoolCreateFlags(o.Flags)
	createInfo.queryType = C.VkQueryType(o.QueryType)
	createInfo.queryCount = C.uint32_t(o.QueryCount)
	createInfo.pipelineStatistics = C.VkQueryPipelineStatisticFlags(o.PipelineStatistics)

	return preallocatedPointer, nil
}
