package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type APIVersion uint32

const (
	Vulkan1_0 APIVersion = C.VK_API_VERSION_1_0
	Vulkan1_1 APIVersion = C.VK_API_VERSION_1_1
	Vulkan1_2 APIVersion = C.VK_API_VERSION_1_2
)

var apiVersionToString = map[APIVersion]string{
	Vulkan1_0: "Vulkan 1.0",
	Vulkan1_1: "Vulkan 1.1",
	Vulkan1_2: "Vulkan 1.2",
}

func (v APIVersion) String() string {
	return apiVersionToString[v]
}
