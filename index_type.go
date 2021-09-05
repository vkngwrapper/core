package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type IndexType int32

const (
	IndexUInt16 IndexType = C.VK_INDEX_TYPE_UINT16
	IndexUInt32 IndexType = C.VK_INDEX_TYPE_UINT32
	IndexNone   IndexType = C.VK_INDEX_TYPE_NONE_KHR
	IndexUInt8  IndexType = C.VK_INDEX_TYPE_UINT8_EXT
)

var indexTypeToString = map[IndexType]string{
	IndexUInt16: "UInt16",
	IndexUInt32: "UInt32",
	IndexNone:   "None",
	IndexUInt8:  "UInt8",
}

func (t IndexType) String() string {
	return indexTypeToString[t]
}
