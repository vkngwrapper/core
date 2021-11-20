package core

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
	SrcAccessMask  common.AccessFlags
	DestAccessMask common.AccessFlags

	common.HaveNext
}

func (o *MemoryBarrierOptions) populate(createInfo *C.VkMemoryBarrier, next unsafe.Pointer) error {
	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_BARRIER
	createInfo.pNext = next
	createInfo.srcAccessMask = C.VkAccessFlags(o.SrcAccessMask)
	createInfo.dstAccessMask = C.VkAccessFlags(o.DestAccessMask)

	return nil
}

func (o *MemoryBarrierOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkMemoryBarrier)(allocator.Malloc(C.sizeof_struct_VkMemoryBarrier))
	err := o.populate(createInfo, next)
	return unsafe.Pointer(createInfo), err
}

type BufferMemoryBarrierOptions struct {
	SrcAccessMask  common.AccessFlags
	DestAccessMask common.AccessFlags

	SrcQueueFamilyIndex  int
	DestQueueFamilyIndex int

	Buffer Buffer

	Offset uint64
	Size   uint64

	common.HaveNext
}

func (o *BufferMemoryBarrierOptions) populate(createInfo *C.VkBufferMemoryBarrier, next unsafe.Pointer) error {
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER
	createInfo.pNext = next
	createInfo.srcAccessMask = C.VkAccessFlags(o.SrcAccessMask)
	createInfo.dstAccessMask = C.VkAccessFlags(o.DestAccessMask)
	createInfo.srcQueueFamilyIndex = C.uint32_t(o.SrcQueueFamilyIndex)
	createInfo.dstQueueFamilyIndex = C.uint32_t(o.DestQueueFamilyIndex)
	createInfo.buffer = C.VkBuffer(o.Buffer.Handle())
	createInfo.offset = C.VkDeviceSize(o.Offset)
	createInfo.size = C.VkDeviceSize(o.Size)

	return nil
}

func (o *BufferMemoryBarrierOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkBufferMemoryBarrier)(allocator.Malloc(C.sizeof_struct_VkBufferMemoryBarrier))
	err := o.populate(createInfo, next)
	return unsafe.Pointer(createInfo), err
}

type ImageMemoryBarrierOptions struct {
	SrcAccessMask  common.AccessFlags
	DestAccessMask common.AccessFlags

	OldLayout common.ImageLayout
	NewLayout common.ImageLayout

	SrcQueueFamilyIndex  int
	DestQueueFamilyIndex int

	Image            Image
	SubresourceRange common.ImageSubresourceRange

	common.HaveNext
}

func (o *ImageMemoryBarrierOptions) populate(createInfo *C.VkImageMemoryBarrier, next unsafe.Pointer) error {
	createInfo.sType = C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER
	createInfo.pNext = next
	createInfo.srcAccessMask = C.VkAccessFlags(o.SrcAccessMask)
	createInfo.dstAccessMask = C.VkAccessFlags(o.DestAccessMask)
	createInfo.oldLayout = C.VkImageLayout(o.OldLayout)
	createInfo.newLayout = C.VkImageLayout(o.NewLayout)
	createInfo.srcQueueFamilyIndex = C.uint32_t(o.SrcQueueFamilyIndex)
	createInfo.dstQueueFamilyIndex = C.uint32_t(o.DestQueueFamilyIndex)
	createInfo.image = C.VkImage(o.Image.Handle())
	createInfo.subresourceRange.aspectMask = C.VkImageAspectFlags(o.SubresourceRange.AspectMask)
	createInfo.subresourceRange.baseMipLevel = C.uint32_t(o.SubresourceRange.BaseMipLevel)
	createInfo.subresourceRange.levelCount = C.uint32_t(o.SubresourceRange.LevelCount)
	createInfo.subresourceRange.baseArrayLayer = C.uint32_t(o.SubresourceRange.BaseArrayLayer)
	createInfo.subresourceRange.layerCount = C.uint32_t(o.SubresourceRange.LayerCount)

	return nil
}

func (o *ImageMemoryBarrierOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkImageMemoryBarrier)(allocator.Malloc(C.sizeof_struct_VkImageMemoryBarrier))
	err := o.populate(createInfo, next)
	return unsafe.Pointer(createInfo), err
}
