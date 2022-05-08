package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	VkErrorInvalidExternalHandle common.VkResult = C.VK_ERROR_INVALID_EXTERNAL_HANDLE
	VkErrorOutOfPoolMemory       common.VkResult = C.VK_ERROR_OUT_OF_POOL_MEMORY
)

func init() {
	VkErrorInvalidExternalHandle.Register("invalid external handle")
	VkErrorOutOfPoolMemory.Register("out of pool memory")
}
