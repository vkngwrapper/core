package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const (
	MemoryHeapMultiInstance core1_0.MemoryHeapFlags = C.VK_MEMORY_HEAP_MULTI_INSTANCE_BIT

	MemoryPropertyProtected core1_0.MemoryProperties = C.VK_MEMORY_PROPERTY_PROTECTED_BIT
)

func init() {
	MemoryHeapMultiInstance.Register("Multi-Instance")

	MemoryPropertyProtected.Register("Protected")
}

////

type BufferMemoryRequirementsOptions struct {
	Buffer core1_0.Buffer

	common.HaveNext
}

func (o BufferMemoryRequirementsOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBufferMemoryRequirementsInfo2{})))
	}

	options := (*C.VkBufferMemoryRequirementsInfo2)(preallocatedPointer)
	options.sType = C.VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
	options.pNext = next
	options.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))

	return preallocatedPointer, nil
}

func (o BufferMemoryRequirementsOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	options := (*C.VkBufferMemoryRequirementsInfo2)(cDataPointer)
	return options.pNext, nil
}

////

type ImageMemoryRequirementsOptions struct {
	Image core1_0.Image

	common.HaveNext
}

func (o ImageMemoryRequirementsOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageMemoryRequirementsInfo2{})))
	}

	options := (*C.VkImageMemoryRequirementsInfo2)(preallocatedPointer)
	options.sType = C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
	options.pNext = next
	options.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))

	return preallocatedPointer, nil
}

func (o ImageMemoryRequirementsOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	options := (*C.VkImageMemoryRequirementsInfo2)(cDataPointer)
	return options.pNext, nil
}

////

type MemoryRequirementsOutData struct {
	MemoryRequirements core1_0.MemoryRequirements
	common.HaveNext
}

func (o *MemoryRequirementsOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryRequirements2{})))
	}

	outData := (*C.VkMemoryRequirements2)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *MemoryRequirementsOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkMemoryRequirements2)(cDataPointer)
	o.MemoryRequirements.Size = int(outData.memoryRequirements.size)
	o.MemoryRequirements.Alignment = int(outData.memoryRequirements.alignment)
	o.MemoryRequirements.MemoryType = uint32(outData.memoryRequirements.memoryTypeBits)

	return outData.pNext, nil
}

////

type ImageSparseMemoryRequirementsOptions struct {
	Image core1_0.Image

	common.HaveNext
}

func (o ImageSparseMemoryRequirementsOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageSparseMemoryRequirementsInfo2{})))
	}

	options := (*C.VkImageSparseMemoryRequirementsInfo2)(preallocatedPointer)
	options.sType = C.VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
	options.pNext = next
	options.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))

	return preallocatedPointer, nil
}

func (o ImageSparseMemoryRequirementsOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	options := (*C.VkImageSparseMemoryRequirementsInfo2)(cDataPointer)
	return options.pNext, nil
}

////

type SparseImageMemoryRequirementsOutData struct {
	MemoryRequirements core1_0.SparseImageMemoryRequirements

	common.HaveNext
}

func (o *SparseImageMemoryRequirementsOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSparseImageMemoryRequirements2{})))
	}

	outData := (*C.VkSparseImageMemoryRequirements2)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *SparseImageMemoryRequirementsOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkSparseImageMemoryRequirements2)(cDataPointer)
	o.MemoryRequirements.FormatProperties.Flags = core1_0.SparseImageFormatFlags(outData.memoryRequirements.formatProperties.flags)
	o.MemoryRequirements.FormatProperties.ImageGranularity = core1_0.Extent3D{
		Width:  int(outData.memoryRequirements.formatProperties.imageGranularity.width),
		Height: int(outData.memoryRequirements.formatProperties.imageGranularity.height),
		Depth:  int(outData.memoryRequirements.formatProperties.imageGranularity.depth),
	}
	o.MemoryRequirements.FormatProperties.AspectMask = core1_0.ImageAspectFlags(outData.memoryRequirements.formatProperties.aspectMask)
	o.MemoryRequirements.ImageMipTailSize = int(outData.memoryRequirements.imageMipTailSize)
	o.MemoryRequirements.ImageMipTailStride = int(outData.memoryRequirements.imageMipTailStride)
	o.MemoryRequirements.ImageMipTailOffset = int(outData.memoryRequirements.imageMipTailOffset)
	o.MemoryRequirements.ImageMipTailFirstLod = int(outData.memoryRequirements.imageMipTailFirstLod)

	return outData.pNext, nil
}
