package commands

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"github.com/palantir/stacktrace"
	"strings"
	"unsafe"
)

type CommandPoolFlags int32

const (
	CommandPoolTransient   CommandPoolFlags = C.VK_COMMAND_POOL_CREATE_TRANSIENT_BIT
	CommandPoolResetBuffer CommandPoolFlags = C.VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT
)

var commandPoolToString = map[CommandPoolFlags]string{
	CommandPoolTransient:   "Transient",
	CommandPoolResetBuffer: "Reset Command Buffer",
}

func (f CommandPoolFlags) String() string {
	if f == 0 {
		return "None"
	}
	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := CommandPoolFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := commandPoolToString[checkBit]
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

type CommandPoolOptions struct {
	GraphicsQueueFamily *int
	Flags               CommandPoolFlags

	Next core.Options
}

func (o *CommandPoolOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	if o.GraphicsQueueFamily == nil {
		return nil, stacktrace.NewError("attempted to create a command pool without setting GraphicsQueueFamilyIndex")
	}

	familyIndex := *o.GraphicsQueueFamily

	cmdPoolCreate := (*C.VkCommandPoolCreateInfo)(allocator.Malloc(C.sizeof_struct_VkCommandPoolCreateInfo))
	cmdPoolCreate.sType = C.VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
	cmdPoolCreate.flags = C.VkCommandPoolCreateFlags(o.Flags)
	cmdPoolCreate.queueFamilyIndex = C.uint32_t(familyIndex)

	var next unsafe.Pointer
	var err error

	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}
	if err != nil {
		return nil, err
	}

	cmdPoolCreate.pNext = next

	return unsafe.Pointer(cmdPoolCreate), nil
}
