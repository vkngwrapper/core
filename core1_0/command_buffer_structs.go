package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
)

const (
	// CommandBufferResetReleaseResources specifies that most or all memory resources currently owned
	// by the CommandBuffer should be returned to the parent CommandPool
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferResetFlagBits.html
	CommandBufferResetReleaseResources CommandBufferResetFlags = C.VK_COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT

	// CommandBufferLevelPrimary specifies a primary CommandBuffer
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferLevel.html
	CommandBufferLevelPrimary CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_PRIMARY
	// CommandBufferLevelSecondary specifies a secondary CommandBuffer
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferLevel.html
	CommandBufferLevelSecondary CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_SECONDARY

	// IndexTypeUInt16 specifies that indices are 16-bit unsigned integer values
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html
	IndexTypeUInt16 IndexType = C.VK_INDEX_TYPE_UINT16
	// IndexTypeUInt32 specifies that indices are 32-bit unsigned integer values
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html
	IndexTypeUInt32 IndexType = C.VK_INDEX_TYPE_UINT32

	// StencilFaceFront specifies that only the front set of stencil state is updated
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilFaceFlagBits.html
	StencilFaceFront StencilFaceFlags = C.VK_STENCIL_FACE_FRONT_BIT
	// StencilFaceBack specifies that only the back set of stencil state is updated
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilFaceFlagBits.html
	StencilFaceBack StencilFaceFlags = C.VK_STENCIL_FACE_BACK_BIT
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

// CommandBufferAllocateInfo specifies the allocation parameters for the CommandBuffer object
type CommandBufferAllocateInfo struct {
	// Level specifies the CommandBuffer level
	Level CommandBufferLevel
	// CommandBufferCount is the number of CommandBuffer objects to allocate from the CommandPool
	CommandBufferCount int
	// CommandPool is the CommandPool from which the CommandBuffer objects are allocated
	CommandPool core.CommandPool

	common.NextOptions
}

func (o CommandBufferAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.CommandBufferCount == 0 {
		return nil, errors.New("attempted to create 0 command buffers")
	}
	if o.CommandPool.Handle() == 0 {
		return nil, errors.New("core1_0.CommandBufferAllocateInfo.CommandPool may not be left unset")
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
