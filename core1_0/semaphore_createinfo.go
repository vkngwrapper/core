package core1_0

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

type SemaphoreCreateFlags int32

var semaphoreCreateFlagsMapping = common.NewFlagStringMapping[SemaphoreCreateFlags]()

func (f SemaphoreCreateFlags) Register(str string) {
	semaphoreCreateFlagsMapping.Register(f, str)
}

func (f SemaphoreCreateFlags) String() string {
	return semaphoreCreateFlagsMapping.FlagsToString(f)
}

////

type SemaphoreCreateOptions struct {
	Flags SemaphoreCreateFlags

	common.NextOptions
}

func (o SemaphoreCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkSemaphoreCreateInfo)
	}
	createInfo := (*C.VkSemaphoreCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO
	createInfo.flags = C.VkSemaphoreCreateFlags(o.Flags)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
