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
	"github.com/vkngwrapper/core/v3/common"
)

// MemoryMapFlags reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapFlags.html
type MemoryMapFlags int32

func (f MemoryMapFlags) String() string {
	return "None"
}

// MemoryAllocateInfo contains parameters of a memory allocation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html
type MemoryAllocateInfo struct {
	// AllocationSize is the size of the allocation in bytes
	AllocationSize int
	// MemoryTypeIndex is an index identifying a memory type from the MemoryTypes slice of
	// PhysicalDeviceMemoryProperties
	MemoryTypeIndex int

	common.NextOptions
}

func (o MemoryAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkMemoryAllocateInfo)
	}
	createInfo := (*C.VkMemoryAllocateInfo)(preallocatedPointer)

	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
	createInfo.allocationSize = C.VkDeviceSize(o.AllocationSize)
	createInfo.memoryTypeIndex = C.uint32_t(o.MemoryTypeIndex)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}

// MappedMemoryRange specifies a mapped memory range
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMappedMemoryRange.html
type MappedMemoryRange struct {
	// Memory is the DeviceMemory object to which this range belongs
	Memory DeviceMemory
	// Offset is the zero-based byte offset from the beginning of the DeviceMemory objects
	Offset int
	// Size is either the size of the range or -1 to affect the range from offset to the end
	// of the current mapping of the allocation
	Size int

	common.NextOptions
}

func (r MappedMemoryRange) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !r.Memory.Initialized() {
		return nil, errors.New("MappedMemoryRange.Memory cannot be left unset")
	}
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkMappedMemoryRange)
	}

	mappedRange := (*C.VkMappedMemoryRange)(preallocatedPointer)
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = next
	mappedRange.memory = C.VkDeviceMemory(unsafe.Pointer(r.Memory.Handle()))
	mappedRange.offset = C.VkDeviceSize(r.Offset)
	mappedRange.size = C.VkDeviceSize(r.Size)

	return preallocatedPointer, nil
}
