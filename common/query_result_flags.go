package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

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
	return FlagsToString(f, queryResultFlagsToString)
}
