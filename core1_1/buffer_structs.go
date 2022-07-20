package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/core1_0"

const (
	// BufferCreateProtected specifies that the Buffer is a protected Buffer
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html
	BufferCreateProtected core1_0.BufferCreateFlags = C.VK_BUFFER_CREATE_PROTECTED_BIT
)

func init() {
	BufferCreateProtected.Register("Protected")
}
