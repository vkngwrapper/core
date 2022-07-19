package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

// ImageCopy specifies an Image copy operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy.html
type ImageCopy struct {
	// SrcSubresource specifies the Image subresources of the Image objects used for the
	// source Image data
	SrcSubresource ImageSubresourceLayers
	// SrcOffset selects the initial x, y, and z offsets in texels of the sub-regions of the
	// source Image data
	SrcOffset Offset3D
	// DstSubresource specifies the Image subresource of the Image objects used for the
	// destination Image data
	DstSubresource ImageSubresourceLayers
	// DstOffset selects the initial x, y, and z offsets in texels of the sub-regions of the
	// destination Image data
	DstOffset Offset3D
	// Extent is the size in texels of the Image to copy in width, height, and depth
	Extent Extent3D
}

func (c ImageCopy) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkImageCopy)
	}

	copyRegion := (*C.VkImageCopy)(preallocatedPointer)

	copyRegion.srcSubresource.aspectMask = C.VkImageAspectFlags(c.SrcSubresource.AspectMask)
	copyRegion.srcSubresource.mipLevel = C.uint32_t(c.SrcSubresource.MipLevel)
	copyRegion.srcSubresource.baseArrayLayer = C.uint32_t(c.SrcSubresource.BaseArrayLayer)
	copyRegion.srcSubresource.layerCount = C.uint32_t(c.SrcSubresource.LayerCount)

	copyRegion.dstSubresource.aspectMask = C.VkImageAspectFlags(c.DstSubresource.AspectMask)
	copyRegion.dstSubresource.mipLevel = C.uint32_t(c.DstSubresource.MipLevel)
	copyRegion.dstSubresource.baseArrayLayer = C.uint32_t(c.DstSubresource.BaseArrayLayer)
	copyRegion.dstSubresource.layerCount = C.uint32_t(c.DstSubresource.LayerCount)

	copyRegion.srcOffset.x = C.int32_t(c.SrcOffset.X)
	copyRegion.srcOffset.y = C.int32_t(c.SrcOffset.Y)
	copyRegion.srcOffset.z = C.int32_t(c.SrcOffset.Z)

	copyRegion.dstOffset.x = C.int32_t(c.DstOffset.X)
	copyRegion.dstOffset.y = C.int32_t(c.DstOffset.Y)
	copyRegion.dstOffset.z = C.int32_t(c.DstOffset.Z)

	copyRegion.extent.width = C.uint32_t(c.Extent.Width)
	copyRegion.extent.height = C.uint32_t(c.Extent.Height)
	copyRegion.extent.depth = C.uint32_t(c.Extent.Depth)

	return preallocatedPointer, nil
}
