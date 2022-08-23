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

const (
	// FenceCreateSignaled specifies that the Fence object is created in the signaled state
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceCreateFlagBits.html
	FenceCreateSignaled FenceCreateFlags = C.VK_FENCE_CREATE_SIGNALED_BIT
)

func init() {
	FenceCreateSignaled.Register("Signaled")
}

// FenceCreateInfo specifies parameters of a newly-created Fence
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceCreateInfo.html
type FenceCreateInfo struct {
	// Flags specifies the initial state and behavior of the Fence
	Flags FenceCreateFlags

	common.NextOptions
}

func (o FenceCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkFenceCreateInfo)
	}
	createInfo := (*C.VkFenceCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
	createInfo.flags = C.VkFenceCreateFlags(o.Flags)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
