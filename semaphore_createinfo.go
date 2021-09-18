package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type SemaphoreOptions struct {
	common.HaveNext
}

func (o *SemaphoreOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkSemaphoreCreateInfo)(allocator.Malloc(C.sizeof_struct_VkSemaphoreCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
