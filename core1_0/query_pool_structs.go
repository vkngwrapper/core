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
	QueryTypeOcclusion          QueryType = C.VK_QUERY_TYPE_OCCLUSION
	QueryTypePipelineStatistics QueryType = C.VK_QUERY_TYPE_PIPELINE_STATISTICS
	QueryTypeTimestamp          QueryType = C.VK_QUERY_TYPE_TIMESTAMP

	QueryResult64Bit            QueryResultFlags = C.VK_QUERY_RESULT_64_BIT
	QueryResultWait             QueryResultFlags = C.VK_QUERY_RESULT_WAIT_BIT
	QueryResultWithAvailability QueryResultFlags = C.VK_QUERY_RESULT_WITH_AVAILABILITY_BIT
	QueryResultPartial          QueryResultFlags = C.VK_QUERY_RESULT_PARTIAL_BIT
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

type QueryPoolCreateInfo struct {
	Flags QueryPoolCreateFlags

	QueryType          QueryType
	QueryCount         int
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
