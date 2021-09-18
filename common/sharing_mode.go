package common

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type SharingMode int32

const (
	SharingExclusive  SharingMode = C.VK_SHARING_MODE_EXCLUSIVE
	SharingConcurrent SharingMode = C.VK_SHARING_MODE_CONCURRENT
)

var sharingModeToString = map[SharingMode]string{
	SharingExclusive:  "Exclusive",
	SharingConcurrent: "Concurrent",
}

func (m SharingMode) String() string {
	return sharingModeToString[m]
}
