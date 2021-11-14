package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type IndexType int32

const (
	IndexUInt16   IndexType = C.VK_INDEX_TYPE_UINT16
	IndexUInt32   IndexType = C.VK_INDEX_TYPE_UINT32
	IndexNoneKHR  IndexType = C.VK_INDEX_TYPE_NONE_KHR
	IndexUInt8EXT IndexType = C.VK_INDEX_TYPE_UINT8_EXT
)

var indexTypeToString = map[IndexType]string{
	IndexUInt16:   "UInt16",
	IndexUInt32:   "UInt32",
	IndexNoneKHR:  "None (Khronos Extension)",
	IndexUInt8EXT: "UInt8 (Extension)",
}

func (t IndexType) String() string {
	return indexTypeToString[t]
}
