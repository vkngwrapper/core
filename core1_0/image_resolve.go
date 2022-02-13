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

type ImageResolve struct {
	SrcSubresource common.ImageSubresourceLayers
	SrcOffset      common.Offset3D
	DstSubresource common.ImageSubresourceLayers
	DstOffset      common.Offset3D
	Extent         common.Extent3D
}

func (r ImageResolve) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkImageResolve)
	}

	imageResolve := (*C.VkImageResolve)(preallocatedPointer)
	imageResolve.srcSubresource.aspectMask = C.VkImageAspectFlags(r.SrcSubresource.AspectMask)
	imageResolve.srcSubresource.mipLevel = C.uint32_t(r.SrcSubresource.MipLevel)
	imageResolve.srcSubresource.baseArrayLayer = C.uint32_t(r.SrcSubresource.BaseArrayLayer)
	imageResolve.srcSubresource.layerCount = C.uint32_t(r.SrcSubresource.LayerCount)

	imageResolve.srcOffset.x = C.int32_t(r.SrcOffset.X)
	imageResolve.srcOffset.y = C.int32_t(r.SrcOffset.Y)
	imageResolve.srcOffset.z = C.int32_t(r.SrcOffset.Z)

	imageResolve.dstSubresource.aspectMask = C.VkImageAspectFlags(r.DstSubresource.AspectMask)
	imageResolve.dstSubresource.mipLevel = C.uint32_t(r.DstSubresource.MipLevel)
	imageResolve.dstSubresource.baseArrayLayer = C.uint32_t(r.DstSubresource.BaseArrayLayer)
	imageResolve.dstSubresource.layerCount = C.uint32_t(r.DstSubresource.LayerCount)

	imageResolve.dstOffset.x = C.int32_t(r.DstOffset.X)
	imageResolve.dstOffset.y = C.int32_t(r.DstOffset.Y)
	imageResolve.dstOffset.z = C.int32_t(r.DstOffset.Z)

	imageResolve.extent.width = C.uint32_t(r.Extent.Width)
	imageResolve.extent.height = C.uint32_t(r.Extent.Height)
	imageResolve.extent.depth = C.uint32_t(r.Extent.Depth)

	return preallocatedPointer, nil
}
