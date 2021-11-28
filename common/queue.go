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
	QueueGraphics      QueueFlags = C.VK_QUEUE_GRAPHICS_BIT
	QueueCompute       QueueFlags = C.VK_QUEUE_COMPUTE_BIT
	QueueTransfer      QueueFlags = C.VK_QUEUE_TRANSFER_BIT
	QueueSparseBinding QueueFlags = C.VK_QUEUE_SPARSE_BINDING_BIT
	QueueProtected     QueueFlags = C.VK_QUEUE_PROTECTED_BIT
)

var queueFlagsToString = map[QueueFlags]string{
	QueueGraphics:      "Graphics",
	QueueCompute:       "Compute",
	QueueTransfer:      "Transfer",
	QueueSparseBinding: "Sparse Binding",
	QueueProtected:     "Protected",
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
