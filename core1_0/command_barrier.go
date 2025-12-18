package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
)

// MemoryBarrier specifies a global memory barrier
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier.html
type MemoryBarrier struct {
	// SrcAccessMask specifies a source access mask
	SrcAccessMask AccessFlags
	// DstAccessMask specifies a destination access mask
	DstAccessMask AccessFlags

	common.NextOptions
}

func (o MemoryBarrier) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

// BufferMemoryBarrier specifies a buffer memory barrier
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier.html
type BufferMemoryBarrier struct {
	// SrcAccessMask specifies a source access mask
	SrcAccessMask AccessFlags
	// DstAccessMask specifies a destination access mask
	DstAccessMask AccessFlags

	// SrcQueueFamilyIndex is the source queue family for a queue family ownership transfer
	SrcQueueFamilyIndex int
	// DstQueueFamilyIndex is the source queue family for a queue family ownership transfer
	DstQueueFamilyIndex int

	// Buffer is the buffer whose backing memory is affected by the barrier
	Buffer Buffer

	// Offset is an offset in bytes into the backing memory for Buffer
	Offset int
	// Size is a size in bytes of the affected area of backing memory for Buffer
	Size int

	common.NextOptions
}

func (o BufferMemoryBarrier) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Buffer == nil {
		return nil, errors.New("core1_0.BufferMemoryBarrier.Buffer cannot be nil")
	}
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

// ImageMemoryBarrier specifies the parameters of an image memory barrier
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier.html
type ImageMemoryBarrier struct {
	// SrcAccessMask specifies a source access mask
	SrcAccessMask AccessFlags
	// DstAccessMask specifies a destination access mask
	DstAccessMask AccessFlags

	// OldLayout is the old layout in an image layout transition
	OldLayout ImageLayout
	// NewLayout is the new layout in an image layout transition
	NewLayout ImageLayout

	// SrcQueueFamilyIndex is the source queue family for a queue family ownership transfer
	SrcQueueFamilyIndex int
	// DstQueueFamilyIndex is the destination queue family for a queue family ownership transfer
	DstQueueFamilyIndex int

	// Image is the Image object affected by this barrier
	Image Image
	// SubresourceRange describes the image subresource range within Image that is affected by this barrier
	SubresourceRange ImageSubresourceRange

	common.NextOptions
}

func (o ImageMemoryBarrier) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Image == nil {
		return nil, errors.New("core1_0.ImageMemoryBarrier.Image cannot be nil")
	}
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
