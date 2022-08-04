package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/common"

const (
	// VkErrorInvalidExternalHandle indicates an external handle is not a valid handle
	// of the specified type
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VkErrorInvalidExternalHandle common.VkResult = C.VK_ERROR_INVALID_EXTERNAL_HANDLE
	// VkErrorOutOfPoolMemory indicates a pool memory allocation has failed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VkErrorOutOfPoolMemory common.VkResult = C.VK_ERROR_OUT_OF_POOL_MEMORY
)

func init() {
	VkErrorInvalidExternalHandle.Register("invalid external handle")
	VkErrorOutOfPoolMemory.Register("out of pool memory")
}
