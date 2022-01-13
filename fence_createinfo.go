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

type FenceCreateFlags int32

const (
	FenceSignaled FenceCreateFlags = C.VK_FENCE_CREATE_SIGNALED_BIT
)

var fenceCreateFlagsToString = map[FenceCreateFlags]string{
	FenceSignaled: "Signaled",
}

func (f FenceCreateFlags) String() string {
	return common.FlagsToString(f, fenceCreateFlagsToString)
}

type FenceOptions struct {
	Flags FenceCreateFlags

	common.HaveNext
}

func (o *FenceOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkFenceCreateInfo)(allocator.Malloc(C.sizeof_struct_VkFenceCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
	createInfo.flags = C.VkFenceCreateFlags(o.Flags)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
