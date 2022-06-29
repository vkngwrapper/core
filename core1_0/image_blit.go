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

type ImageBlit struct {
	SourceSubresource ImageSubresourceLayers
	SourceOffsets     [2]Offset3D

	DestinationSubresource ImageSubresourceLayers
	DestinationOffsets     [2]Offset3D
}

func (b ImageBlit) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkImageBlit)
	}
	imageBlitInfo := (*C.VkImageBlit)(preallocatedPointer)
	imageBlitInfo.srcSubresource.aspectMask = C.VkImageAspectFlags(b.SourceSubresource.AspectMask)
	imageBlitInfo.srcSubresource.mipLevel = C.uint32_t(b.SourceSubresource.MipLevel)
	imageBlitInfo.srcSubresource.baseArrayLayer = C.uint32_t(b.SourceSubresource.BaseArrayLayer)
	imageBlitInfo.srcSubresource.layerCount = C.uint32_t(b.SourceSubresource.LayerCount)

	imageBlitInfo.dstSubresource.aspectMask = C.VkImageAspectFlags(b.DestinationSubresource.AspectMask)
	imageBlitInfo.dstSubresource.mipLevel = C.uint32_t(b.DestinationSubresource.MipLevel)
	imageBlitInfo.dstSubresource.baseArrayLayer = C.uint32_t(b.DestinationSubresource.BaseArrayLayer)
	imageBlitInfo.dstSubresource.layerCount = C.uint32_t(b.DestinationSubresource.LayerCount)

	imageBlitInfo.srcOffsets[0].x = C.int32_t(b.SourceOffsets[0].X)
	imageBlitInfo.srcOffsets[0].y = C.int32_t(b.SourceOffsets[0].Y)
	imageBlitInfo.srcOffsets[0].z = C.int32_t(b.SourceOffsets[0].Z)
	imageBlitInfo.srcOffsets[1].x = C.int32_t(b.SourceOffsets[1].X)
	imageBlitInfo.srcOffsets[1].y = C.int32_t(b.SourceOffsets[1].Y)
	imageBlitInfo.srcOffsets[1].z = C.int32_t(b.SourceOffsets[1].Z)

	imageBlitInfo.dstOffsets[0].x = C.int32_t(b.DestinationOffsets[0].X)
	imageBlitInfo.dstOffsets[0].y = C.int32_t(b.DestinationOffsets[0].Y)
	imageBlitInfo.dstOffsets[0].z = C.int32_t(b.DestinationOffsets[0].Z)
	imageBlitInfo.dstOffsets[1].x = C.int32_t(b.DestinationOffsets[1].X)
	imageBlitInfo.dstOffsets[1].y = C.int32_t(b.DestinationOffsets[1].Y)
	imageBlitInfo.dstOffsets[1].z = C.int32_t(b.DestinationOffsets[1].Z)

	return preallocatedPointer, nil
}
