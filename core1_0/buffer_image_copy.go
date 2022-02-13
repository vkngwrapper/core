package core1_0

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type BufferImageCopy struct {
	BufferOffset      int
	BufferRowLength   int
	BufferImageHeight int

	ImageSubresource common.ImageSubresourceLayers
	ImageOffset      common.Offset3D
	ImageExtent      common.Extent3D
}

func (c BufferImageCopy) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if c.BufferImageHeight < 0 {
		return nil, errors.New("provided BufferImageHeight of <0")
	}

	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkBufferImageCopy)
	}

	createInfo := (*C.VkBufferImageCopy)(preallocatedPointer)
	createInfo.bufferOffset = C.VkDeviceSize(c.BufferOffset)
	createInfo.bufferRowLength = C.uint32_t(c.BufferRowLength)
	createInfo.bufferImageHeight = C.uint32_t(c.BufferImageHeight)
	createInfo.imageSubresource.aspectMask = C.VkImageAspectFlags(c.ImageSubresource.AspectMask)
	createInfo.imageSubresource.mipLevel = C.uint32_t(c.ImageSubresource.MipLevel)
	createInfo.imageSubresource.baseArrayLayer = C.uint32_t(c.ImageSubresource.BaseArrayLayer)
	createInfo.imageSubresource.layerCount = C.uint32_t(c.ImageSubresource.LayerCount)
	createInfo.imageOffset.x = C.int32_t(c.ImageOffset.X)
	createInfo.imageOffset.y = C.int32_t(c.ImageOffset.Y)
	createInfo.imageOffset.z = C.int32_t(c.ImageOffset.Z)
	createInfo.imageExtent.width = C.uint32_t(c.ImageExtent.Width)
	createInfo.imageExtent.height = C.uint32_t(c.ImageExtent.Height)
	createInfo.imageExtent.depth = C.uint32_t(c.ImageExtent.Depth)

	return preallocatedPointer, nil
}
