package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type QueryControlFlags int32

const (
	QueryPrecise QueryControlFlags = C.VK_QUERY_CONTROL_PRECISE_BIT
)

var queryControlFlagsToString = map[QueryControlFlags]string{
	QueryPrecise: "Precise",
}

func (f QueryControlFlags) String() string {
	return FlagsToString(f, queryControlFlagsToString)
}
