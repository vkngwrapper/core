package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type MappedMemoryRange struct {
	Memory DeviceMemory
	Offset int
	Size   int

	core.HaveNext
}

func (r MappedMemoryRange) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
