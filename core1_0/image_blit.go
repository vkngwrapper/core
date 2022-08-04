package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

// ImageBlit specifies an Image blit operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit.html
type ImageBlit struct {
	// SrcSubresource is the subresource to blit from
	SrcSubresource ImageSubresourceLayers
	// SrcOffsets is a slice of Offset3D structures specifying the bounds of the source region
	// within the source subresource
	SrcOffsets [2]Offset3D

	// DstSubresource is the subresource to blit to
	DstSubresource ImageSubresourceLayers
	// DstOffsets is a slice of Offset3D structures specifying the bounds of the destination region
	// within the destination subresource
	DstOffsets [2]Offset3D
}

func (b ImageBlit) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkImageBlit)
	}
	imageBlitInfo := (*C.VkImageBlit)(preallocatedPointer)
	imageBlitInfo.srcSubresource.aspectMask = C.VkImageAspectFlags(b.SrcSubresource.AspectMask)
	imageBlitInfo.srcSubresource.mipLevel = C.uint32_t(b.SrcSubresource.MipLevel)
	imageBlitInfo.srcSubresource.baseArrayLayer = C.uint32_t(b.SrcSubresource.BaseArrayLayer)
	imageBlitInfo.srcSubresource.layerCount = C.uint32_t(b.SrcSubresource.LayerCount)

	imageBlitInfo.dstSubresource.aspectMask = C.VkImageAspectFlags(b.DstSubresource.AspectMask)
	imageBlitInfo.dstSubresource.mipLevel = C.uint32_t(b.DstSubresource.MipLevel)
	imageBlitInfo.dstSubresource.baseArrayLayer = C.uint32_t(b.DstSubresource.BaseArrayLayer)
	imageBlitInfo.dstSubresource.layerCount = C.uint32_t(b.DstSubresource.LayerCount)

	imageBlitInfo.srcOffsets[0].x = C.int32_t(b.SrcOffsets[0].X)
	imageBlitInfo.srcOffsets[0].y = C.int32_t(b.SrcOffsets[0].Y)
	imageBlitInfo.srcOffsets[0].z = C.int32_t(b.SrcOffsets[0].Z)
	imageBlitInfo.srcOffsets[1].x = C.int32_t(b.SrcOffsets[1].X)
	imageBlitInfo.srcOffsets[1].y = C.int32_t(b.SrcOffsets[1].Y)
	imageBlitInfo.srcOffsets[1].z = C.int32_t(b.SrcOffsets[1].Z)

	imageBlitInfo.dstOffsets[0].x = C.int32_t(b.DstOffsets[0].X)
	imageBlitInfo.dstOffsets[0].y = C.int32_t(b.DstOffsets[0].Y)
	imageBlitInfo.dstOffsets[0].z = C.int32_t(b.DstOffsets[0].Z)
	imageBlitInfo.dstOffsets[1].x = C.int32_t(b.DstOffsets[1].X)
	imageBlitInfo.dstOffsets[1].y = C.int32_t(b.DstOffsets[1].Y)
	imageBlitInfo.dstOffsets[1].z = C.int32_t(b.DstOffsets[1].Z)

	return preallocatedPointer, nil
}
