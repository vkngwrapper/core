package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

const (
	// MemoryHeapMultiInstance specifies that ina  logical Device representing more than one
	// PhysicalDevice, there is a per-PhysicalDevice instance of the heap memory
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeapFlagBits.html
	MemoryHeapMultiInstance core1_0.MemoryHeapFlags = C.VK_MEMORY_HEAP_MULTI_INSTANCE_BIT

	// MemoryPropertyProtected specifies that the memory type only allows Device access to the
	// memory, and allows protected Queue operations to access the DeviceMemory
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPropertyFlagBits.html
	MemoryPropertyProtected core1_0.MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_PROTECTED_BIT
)

func init() {
	MemoryHeapMultiInstance.Register("Multi-Instance")

	MemoryPropertyProtected.Register("Protected")
}

////

// BufferMemoryRequirementsInfo2 has no documentation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryRequirementsInfo2.html
type BufferMemoryRequirementsInfo2 struct {
	// Buffer is the Buffer to query
	Buffer core1_0.Buffer

	common.NextOptions
}

func (o BufferMemoryRequirementsInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Buffer == nil {
		return nil, errors.New("core1_1.BufferMemoryRequirementsInfo2.Buffer cannot be nil")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBufferMemoryRequirementsInfo2{})))
	}

	options := (*C.VkBufferMemoryRequirementsInfo2)(preallocatedPointer)
	options.sType = C.VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
	options.pNext = next
	options.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))

	return preallocatedPointer, nil
}

////

// ImageMemoryRequirementsInfo2 has no documentation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2.html
type ImageMemoryRequirementsInfo2 struct {
	// Image is the Image to query
	Image core1_0.Image

	common.NextOptions
}

func (o ImageMemoryRequirementsInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Image == nil {
		return nil, errors.New("core1_1.ImageMemoryRequirementsInfo2.Image cannot be nil")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageMemoryRequirementsInfo2{})))
	}

	options := (*C.VkImageMemoryRequirementsInfo2)(preallocatedPointer)
	options.sType = C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
	options.pNext = next
	options.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))

	return preallocatedPointer, nil
}

////

// MemoryRequirements2 specifies memory requirements
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html
type MemoryRequirements2 struct {
	// MemoryRequirements describes the memory requirements of the resource
	MemoryRequirements core1_0.MemoryRequirements

	common.NextOutData
}

func (o *MemoryRequirements2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryRequirements2{})))
	}

	outData := (*C.VkMemoryRequirements2)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *MemoryRequirements2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkMemoryRequirements2)(cDataPointer)
	o.MemoryRequirements.Size = int(outData.memoryRequirements.size)
	o.MemoryRequirements.Alignment = int(outData.memoryRequirements.alignment)
	o.MemoryRequirements.MemoryTypeBits = uint32(outData.memoryRequirements.memoryTypeBits)

	return outData.pNext, nil
}

////

// ImageSparseMemoryRequirementsInfo2 has no documentation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSparseMemoryRequirementsInfo2.html
type ImageSparseMemoryRequirementsInfo2 struct {
	// Image is the Image to query
	Image core1_0.Image

	common.NextOptions
}

func (o ImageSparseMemoryRequirementsInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Image == nil {
		return nil, errors.New("core1_1.ImageSparseMemoryRequirementsInfo2.Image cannot be nil")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageSparseMemoryRequirementsInfo2{})))
	}

	options := (*C.VkImageSparseMemoryRequirementsInfo2)(preallocatedPointer)
	options.sType = C.VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
	options.pNext = next
	options.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))

	return preallocatedPointer, nil
}

////

// SparseImageMemoryRequirements2 has no documentation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements2.html
type SparseImageMemoryRequirements2 struct {
	// MemoryRequirements describes the memory requirements of the sparse image
	MemoryRequirements core1_0.SparseImageMemoryRequirements

	common.NextOutData
}

func (o *SparseImageMemoryRequirements2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSparseImageMemoryRequirements2{})))
	}

	outData := (*C.VkSparseImageMemoryRequirements2)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *SparseImageMemoryRequirements2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
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
