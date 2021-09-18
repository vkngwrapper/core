package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
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

func (c *vulkanCommandBuffer) CmdCopyBuffer(srcBuffer Buffer, dstBuffer Buffer, copyRegions []BufferCopy) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionCount := len(copyRegions)
	copyRegionUnsafe := allocator.Malloc(copyRegionCount * C.sizeof_struct_VkBufferCopy)
	copyRegionPtr := (*C.VkBufferCopy)(copyRegionUnsafe)
	copyRegionSlice := ([]C.VkBufferCopy)(unsafe.Slice(copyRegionPtr, copyRegionCount))

	for i := 0; i < copyRegionCount; i++ {
		copyRegionSlice[i].srcOffset = C.VkDeviceSize(copyRegions[i].SrcOffset)
		copyRegionSlice[i].dstOffset = C.VkDeviceSize(copyRegions[i].DstOffset)
		copyRegionSlice[i].size = C.VkDeviceSize(copyRegions[i].Size)
	}

	return c.driver.VkCmdCopyBuffer(c.handle, srcBuffer.Handle(), dstBuffer.Handle(), Uint32(copyRegionCount), (*VkBufferCopy)(copyRegionUnsafe))
}
