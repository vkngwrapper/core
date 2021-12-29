package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "strings"

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
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		checkBit := QueryResultFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := queryResultFlagsToString[checkBit]
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
