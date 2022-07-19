package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

// BufferCopy specifies a buffer copy operation via CommandBuffer.CmdCopyBuffer
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferCopy.html
type BufferCopy struct {
	// SrcOffset is the starting offset in bytes from the start of the source Buffer
	SrcOffset int
	// DstOffset is the starting offset in bytes from the start of the dest Buffer
	DstOffset int
	// Size is the number of bytes to copy
	Size int
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

// BufferImageCopy specifies a buffer image copy operation via CommandBuffer.CmdCopyBufferToImage
// or CommandBuffer.CmdCopyImageToBuffer
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy.html
type BufferImageCopy struct {
	// BufferOffset is the offset in bytes from the start of the Buffer
	BufferOffset int
	// BufferRowLength is the size in texels of the rows of the image stored in the Buffer.
	// 0 indicates that the ImageExtent controls this value
	BufferRowLength int
	// BufferImageHeight is the height in texels of the image stored in the Buffer
	// 0 indicates that the ImageExtent controls this value
	BufferImageHeight int

	// ImageSubresource is used to specify the specific image subresources of the Image
	ImageSubresource ImageSubresourceLayers
	// ImageOffset selects the initial x, y, and z offset in texels of the Image subregion
	ImageOffset Offset3D
	// ImageExtent is the size in texels of the Image subregion
	ImageExtent Extent3D
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
