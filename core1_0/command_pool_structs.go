package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

const (
	// CommandPoolResetReleaseResources specifies that resetting a CommandPool recycles all of the
	// resources from the CommandPool back to the system
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolResetFlagBits.html
	CommandPoolResetReleaseResources CommandPoolResetFlags = C.VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT

	// CommandPoolCreateTransient specifies that CommandBuffer objects allocated from the pool
	// will be short-lived, meaning that they will be reset within a relatively short timeframe.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateFlagBits.html
	CommandPoolCreateTransient CommandPoolCreateFlags = C.VK_COMMAND_POOL_CREATE_TRANSIENT_BIT
	// CommandPoolCreateResetBuffer allows any CommandBuffer allocated from a pool to be individually
	// reset to the initial state.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateFlagBits.html
	CommandPoolCreateResetBuffer CommandPoolCreateFlags = C.VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT
)

func init() {
	CommandPoolResetReleaseResources.Register("Release Resources")

	CommandPoolCreateTransient.Register("Transient")
	CommandPoolCreateResetBuffer.Register("Reset Command Buffer")
}

// CommandPoolCreateInfo specifies parameters of a newly-created CommandPool
type CommandPoolCreateInfo struct {
	// QueueFamilyIndex designates a queue family. All CommandBuffer objects allocated from this
	// CommandPool must be submitted on queues from the same queue family
	QueueFamilyIndex int
	// Flags indicates usage behavior for the pool and CommandBuffer objects allocated from it
	Flags CommandPoolCreateFlags

	common.NextOptions
}

func (o CommandPoolCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkCommandPoolCreateInfo)
	}

	cmdPoolCreate := (*C.VkCommandPoolCreateInfo)(preallocatedPointer)
	cmdPoolCreate.sType = C.VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
	cmdPoolCreate.flags = C.VkCommandPoolCreateFlags(o.Flags)
	cmdPoolCreate.pNext = next

	cmdPoolCreate.queueFamilyIndex = C.uint32_t(o.QueueFamilyIndex)

	return unsafe.Pointer(cmdPoolCreate), nil
}
