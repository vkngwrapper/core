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
	ResetReleaseResources common.CommandBufferResetFlags = C.VK_COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT

	LevelPrimary   common.CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_PRIMARY
	LevelSecondary common.CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_SECONDARY

	IndexUInt16 common.IndexType = C.VK_INDEX_TYPE_UINT16
	IndexUInt32 common.IndexType = C.VK_INDEX_TYPE_UINT32

	StencilFaceFront common.StencilFaces = C.VK_STENCIL_FACE_FRONT_BIT
	StencilFaceBack  common.StencilFaces = C.VK_STENCIL_FACE_BACK_BIT
)

func init() {
	ResetReleaseResources.Register("Reset Release Resources")

	LevelPrimary.Register("Primary")
	LevelSecondary.Register("Secondary")

	IndexUInt16.Register("UInt16")
	IndexUInt32.Register("UInt32")

	StencilFaceFront.Register("Stencil Front")
	StencilFaceBack.Register("Stencil Back")
}

type CommandBufferOptions struct {
	Level       common.CommandBufferLevel
	BufferCount int
	CommandPool CommandPool

	common.HaveNext
}

func (o CommandBufferOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.BufferCount == 0 {
		return nil, errors.New("attempted to create 0 command buffers")
	}

	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkCommandBufferAllocateInfo{})))
	}

	createInfo := (*C.VkCommandBufferAllocateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	createInfo.pNext = next

	createInfo.level = C.VkCommandBufferLevel(o.Level)
	createInfo.commandBufferCount = C.uint32_t(o.BufferCount)
	createInfo.commandPool = C.VkCommandPool(unsafe.Pointer(o.CommandPool.Handle()))

	return unsafe.Pointer(createInfo), nil
}
