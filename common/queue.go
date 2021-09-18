package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"strings"
)

type QueueFlags int32

const (
	Graphics      QueueFlags = C.VK_QUEUE_GRAPHICS_BIT
	Compute       QueueFlags = C.VK_QUEUE_COMPUTE_BIT
	Transfer      QueueFlags = C.VK_QUEUE_TRANSFER_BIT
	SparseBinding QueueFlags = C.VK_QUEUE_SPARSE_BINDING_BIT
	Protected     QueueFlags = C.VK_QUEUE_PROTECTED_BIT
)

var queueFlagsToString = map[QueueFlags]string{
	Graphics:      "Graphics",
	Compute:       "Compute",
	Transfer:      "Transfer",
	SparseBinding: "SparseBinding",
	Protected:     "Protected",
}

func (f QueueFlags) String() string {
	if f == 0 {
		return "None"
	}

	hasOne := false
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		shiftedBit := QueueFlags(1 << i)
		if f&shiftedBit != 0 {
			strVal, exists := queueFlagsToString[shiftedBit]
			if exists {
				if hasOne {
					sb.WriteString("|")
				}
				sb.WriteString(strVal)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type QueueFamily struct {
	Flags                       QueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity Extent3D
}
