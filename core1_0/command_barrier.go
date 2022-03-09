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

type MemoryBarrierOptions struct {
	SrcAccessMask common.AccessFlags
	DstAccessMask common.AccessFlags

	common.HaveNext
}

func (o MemoryBarrierOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkMemoryBarrier)
	}
	createInfo := (*C.VkMemoryBarrier)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_BARRIER
	createInfo.pNext = next
	createInfo.srcAccessMask = C.VkAccessFlags(o.SrcAccessMask)
	createInfo.dstAccessMask = C.VkAccessFlags(o.DstAccessMask)

	return preallocatedPointer, nil
}

type BufferMemoryBarrierOptions struct {
	SrcAccessMask common.AccessFlags
	DstAccessMask common.AccessFlags

	SrcQueueFamilyIndex int
	DstQueueFamilyIndex int

	Buffer Buffer

	Offset int
	Size   int

	common.HaveNext
}

func (o BufferMemoryBarrierOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkBufferMemoryBarrier)
	}
	createInfo := (*C.VkBufferMemoryBarrier)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER
	createInfo.pNext = next
	createInfo.srcAccessMask = C.VkAccessFlags(o.SrcAccessMask)
	createInfo.dstAccessMask = C.VkAccessFlags(o.DstAccessMask)
	createInfo.srcQueueFamilyIndex = C.uint32_t(o.SrcQueueFamilyIndex)
	createInfo.dstQueueFamilyIndex = C.uint32_t(o.DstQueueFamilyIndex)
	createInfo.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))
	createInfo.offset = C.VkDeviceSize(o.Offset)
	createInfo.size = C.VkDeviceSize(o.Size)

	return preallocatedPointer, nil
}

type ImageMemoryBarrierOptions struct {
	SrcAccessMask common.AccessFlags
	DstAccessMask common.AccessFlags

	OldLayout common.ImageLayout
	NewLayout common.ImageLayout

	SrcQueueFamilyIndex int
	DstQueueFamilyIndex int

	Image            Image
	SubresourceRange common.ImageSubresourceRange

	common.HaveNext
}

func (o ImageMemoryBarrierOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkImageMemoryBarrier)
	}
	createInfo := (*C.VkImageMemoryBarrier)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER
	createInfo.pNext = next
	createInfo.srcAccessMask = C.VkAccessFlags(o.SrcAccessMask)
	createInfo.dstAccessMask = C.VkAccessFlags(o.DstAccessMask)
	createInfo.oldLayout = C.VkImageLayout(o.OldLayout)
	createInfo.newLayout = C.VkImageLayout(o.NewLayout)
	createInfo.srcQueueFamilyIndex = C.uint32_t(o.SrcQueueFamilyIndex)
	createInfo.dstQueueFamilyIndex = C.uint32_t(o.DstQueueFamilyIndex)
	createInfo.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))
	createInfo.subresourceRange.aspectMask = C.VkImageAspectFlags(o.SubresourceRange.AspectMask)
	createInfo.subresourceRange.baseMipLevel = C.uint32_t(o.SubresourceRange.BaseMipLevel)
	createInfo.subresourceRange.levelCount = C.uint32_t(o.SubresourceRange.LevelCount)
	createInfo.subresourceRange.baseArrayLayer = C.uint32_t(o.SubresourceRange.BaseArrayLayer)
	createInfo.subresourceRange.layerCount = C.uint32_t(o.SubresourceRange.LayerCount)

	return preallocatedPointer, nil
}
