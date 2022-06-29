package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

const (
	CommandPoolResetReleaseResources CommandPoolResetFlags = C.VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT

	CommandPoolCreateTransient   CommandPoolCreateFlags = C.VK_COMMAND_POOL_CREATE_TRANSIENT_BIT
	CommandPoolCreateResetBuffer CommandPoolCreateFlags = C.VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT
)

func init() {
	CommandPoolResetReleaseResources.Register("Release Resources")

	CommandPoolCreateTransient.Register("Transient")
	CommandPoolCreateResetBuffer.Register("Reset Command Buffer")
}

type CommandPoolCreateOptions struct {
	GraphicsQueueFamily *int
	Flags               CommandPoolCreateFlags

	common.HaveNext
}

func (o CommandPoolCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

func (o CommandPoolCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkCommandPoolCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
