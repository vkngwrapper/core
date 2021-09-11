package resources

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoparam"
	"strings"
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
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := FenceCreateFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := fenceCreateFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type FenceOptions struct {
	Flags FenceCreateFlags

	Next core.Options
}

func (o *FenceOptions) AllocForC(allocator *cgoparam.Allocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkFenceCreateInfo)(allocator.Malloc(C.sizeof_struct_VkFenceCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
	createInfo.flags = C.VkFenceCreateFlags(o.Flags)

	var err error
	var next unsafe.Pointer
	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}

	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
