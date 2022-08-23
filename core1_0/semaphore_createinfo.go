package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// SemaphoreCreateFlags is reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateFlags.html
type SemaphoreCreateFlags int32

var semaphoreCreateFlagsMapping = common.NewFlagStringMapping[SemaphoreCreateFlags]()

func (f SemaphoreCreateFlags) Register(str string) {
	semaphoreCreateFlagsMapping.Register(f, str)
}

func (f SemaphoreCreateFlags) String() string {
	return semaphoreCreateFlagsMapping.FlagsToString(f)
}

////

// SemaphoreCreateInfo specifies parameters of a newly-created Semaphore
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html
type SemaphoreCreateInfo struct {
	// Flags is reserved future use
	Flags SemaphoreCreateFlags

	common.NextOptions
}

func (o SemaphoreCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkSemaphoreCreateInfo)
	}
	createInfo := (*C.VkSemaphoreCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO
	createInfo.flags = C.VkSemaphoreCreateFlags(o.Flags)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
