package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type SharingMode int32

const (
	Exclusive  SharingMode = C.VK_SHARING_MODE_EXCLUSIVE
	Concurrent SharingMode = C.VK_SHARING_MODE_CONCURRENT
)

var sharingModeToString = map[SharingMode]string {
	Exclusive:  "Exclusive",
	Concurrent: "Concurrent",
}

func (m SharingMode) String() string {
	return sharingModeToString[m]
}
