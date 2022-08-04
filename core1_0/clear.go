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

// ClearAttachment specifies a clear attachment
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearAttachment.html
type ClearAttachment struct {
	// AspectMask is a mask selecting the color, depth, and/or stencil aspects of the attachment
	// to be cleared
	AspectMask ImageAspectFlags
	// ColorAttachment is an index into the currently-bound color attachments
	ColorAttachment int
	// ClearValue is the color or depth/stencil value to clear the attachment to
	ClearValue ClearValue
}

func (c ClearAttachment) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkClearAttachment)
	}

	clearAttachment := (*C.VkClearAttachment)(preallocatedPointer)
	clearAttachment.aspectMask = C.VkImageAspectFlags(c.AspectMask)
	clearAttachment.colorAttachment = C.uint32_t(c.ColorAttachment)
	c.ClearValue.PopulateValueUnion(unsafe.Pointer(&clearAttachment.clearValue))

	return preallocatedPointer, nil
}

// ClearRect specifies a clear rectangle
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearRect.html
type ClearRect struct {
	// Rect is the two-dimensional region to be cleared
	Rect Rect2D
	// BaseArrayLayer is the first layer to be cleared
	BaseArrayLayer int
	// LayerCount is the number of layers to clear
	LayerCount int
}

func (r ClearRect) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkClearRect)
	}

	clearRect := (*C.VkClearRect)(preallocatedPointer)
	clearRect.baseArrayLayer = C.uint32_t(r.BaseArrayLayer)
	clearRect.layerCount = C.uint32_t(r.LayerCount)
	clearRect.rect.extent.width = C.uint32_t(r.Rect.Extent.Width)
	clearRect.rect.extent.height = C.uint32_t(r.Rect.Extent.Height)
	clearRect.rect.offset.x = C.int32_t(r.Rect.Offset.X)
	clearRect.rect.offset.y = C.int32_t(r.Rect.Offset.Y)

	return preallocatedPointer, nil
}
