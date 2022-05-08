package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

const (
	BufferCreateProtected common.BufferCreateFlags = C.VK_BUFFER_CREATE_PROTECTED_BIT
)

func init() {
	BufferCreateProtected.Register("Protected")
}
