package options

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
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
	return common.FlagsToString(f, commandPoolToString)
}

type CommandPoolOptions struct {
	GraphicsQueueFamily *int
	Flags               CommandPoolFlags

	core.HaveNext
}

func (o CommandPoolOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.GraphicsQueueFamily == nil {
		return nil, errors.New("attempted to create a command pool without setting GraphicsQueueFamilyIndex")
	}

	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkCommandPoolCreateInfo)
	}

	familyIndex := *o.GraphicsQueueFamily

	cmdPoolCreate := (*C.VkCommandPoolCreateInfo)(preallocatedPointer)
	cmdPoolCreate.sType = C.VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
	cmdPoolCreate.flags = C.VkCommandPoolCreateFlags(o.Flags)
	cmdPoolCreate.pNext = next

	cmdPoolCreate.queueFamilyIndex = C.uint32_t(familyIndex)

	return unsafe.Pointer(cmdPoolCreate), nil
}
