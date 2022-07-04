package core1_0

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

type MemoryMapFlags int32

func (f MemoryMapFlags) String() string {
	return "None"
}

type MemoryAllocateOptions struct {
	AllocationSize  int
	MemoryTypeIndex int

	common.NextOptions
}

func (o MemoryAllocateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

type MappedMemoryRangeOptions struct {
	Memory DeviceMemory
	Offset int
	Size   int

	common.NextOptions
}

func (r MappedMemoryRangeOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
