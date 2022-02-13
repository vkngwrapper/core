package core1_0

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type BufferCopy struct {
	SrcOffset int
	DstOffset int
	Size      int
}

func (c BufferCopy) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkBufferCopy)
	}

	copyRegion := (*C.VkBufferCopy)(preallocatedPointer)
	copyRegion.srcOffset = C.VkDeviceSize(c.SrcOffset)
	copyRegion.dstOffset = C.VkDeviceSize(c.DstOffset)
	copyRegion.size = C.VkDeviceSize(c.Size)

	return preallocatedPointer, nil
}
