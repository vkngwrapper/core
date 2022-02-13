package core1_0

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type QueryResultFlags int32

const (
	QueryResult64Bit            QueryResultFlags = C.VK_QUERY_RESULT_64_BIT
	QueryResultWait             QueryResultFlags = C.VK_QUERY_RESULT_WAIT_BIT
	QueryResultWithAvailability QueryResultFlags = C.VK_QUERY_RESULT_WITH_AVAILABILITY_BIT
	QueryResultPartial          QueryResultFlags = C.VK_QUERY_RESULT_PARTIAL_BIT
)

var queryResultFlagsToString = map[QueryResultFlags]string{
	QueryResult64Bit:            "64-Bit",
	QueryResultWait:             "Wait",
	QueryResultWithAvailability: "With Availability",
	QueryResultPartial:          "Partial",
}

func (f QueryResultFlags) String() string {
	return common.FlagsToString(f, queryResultFlagsToString)
}
