package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

const (
	CommandBufferResetReleaseResources CommandBufferResetFlags = C.VK_COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT

	CommandBufferLevelPrimary   CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_PRIMARY
	CommandBufferLevelSecondary CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_SECONDARY

	IndexTypeUInt16 IndexType = C.VK_INDEX_TYPE_UINT16
	IndexTypeUInt32 IndexType = C.VK_INDEX_TYPE_UINT32

	StencilFaceFront StencilFaceFlags = C.VK_STENCIL_FACE_FRONT_BIT
	StencilFaceBack  StencilFaceFlags = C.VK_STENCIL_FACE_BACK_BIT
)

func init() {
	CommandBufferResetReleaseResources.Register("Reset Release Resources")

	CommandBufferLevelPrimary.Register("Primary")
	CommandBufferLevelSecondary.Register("Secondary")

	IndexTypeUInt16.Register("UInt16")
	IndexTypeUInt32.Register("UInt32")

	StencilFaceFront.Register("Stencil Front")
	StencilFaceBack.Register("Stencil Back")
}

type CommandBufferAllocateInfo struct {
	Level              CommandBufferLevel
	CommandBufferCount int
	CommandPool        CommandPool

	common.NextOptions
}

func (o CommandBufferAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.CommandBufferCount == 0 {
		return nil, errors.New("attempted to create 0 command buffers")
	}

	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkCommandBufferAllocateInfo{})))
	}

	createInfo := (*C.VkCommandBufferAllocateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	createInfo.pNext = next

	createInfo.level = C.VkCommandBufferLevel(o.Level)
	createInfo.commandBufferCount = C.uint32_t(o.CommandBufferCount)
	createInfo.commandPool = C.VkCommandPool(unsafe.Pointer(o.CommandPool.Handle()))

	return unsafe.Pointer(createInfo), nil
}
