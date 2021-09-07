package commands

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/resource"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type BufferCopy struct {
	SrcOffset int
	DstOffset int
	Size      int
}

func (c *CommandBuffer) CmdCopyBuffer(allocator cgoalloc.Allocator, srcBuffer *resource.Buffer, dstBuffer *resource.Buffer, copyRegions []BufferCopy) {
	copyRegionCount := len(copyRegions)
	copyRegionUnsafe := allocator.Malloc(copyRegionCount * C.sizeof_struct_VkBufferCopy)
	defer allocator.Free(copyRegionUnsafe)

	copyRegionPtr := (*C.VkBufferCopy)(copyRegionUnsafe)
	copyRegionSlice := ([]C.VkBufferCopy)(unsafe.Slice(copyRegionPtr, copyRegionCount))

	for i := 0; i < copyRegionCount; i++ {
		copyRegionSlice[i].srcOffset = C.VkDeviceSize(copyRegions[i].SrcOffset)
		copyRegionSlice[i].dstOffset = C.VkDeviceSize(copyRegions[i].DstOffset)
		copyRegionSlice[i].size = C.VkDeviceSize(copyRegions[i].Size)
	}

	srcBufferHandle := C.VkBuffer(unsafe.Pointer(srcBuffer.Handle()))
	dstBufferHandle := C.VkBuffer(unsafe.Pointer(dstBuffer.Handle()))

	C.vkCmdCopyBuffer(c.handle, srcBufferHandle, dstBufferHandle, C.uint32_t(copyRegionCount), copyRegionPtr)
}
