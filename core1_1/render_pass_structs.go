package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
)

// InputAttachmentAspectReference specifies a subpass/input attachment pair and an aspect mask that
// can be read
type InputAttachmentAspectReference struct {
	// Subpass is an index into RenderPassCreateInfo.Subpasses
	Subpass int
	// InputAttachmentIndex is an index into the InputAttachments of the specified subpass
	InputAttachmentIndex int
	// AspectMask is a mask of which aspect(s) can be accessed within the specified subpass
	AspectMask core1_0.ImageAspectFlags
}

func (ref InputAttachmentAspectReference) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkInputAttachmentAspectReference{})))
	}

	val := (*C.VkInputAttachmentAspectReference)(preallocatedPointer)
	val.subpass = C.uint32_t(ref.Subpass)
	val.inputAttachmentIndex = C.uint32_t(ref.InputAttachmentIndex)
	val.aspectMask = C.VkImageAspectFlags(ref.AspectMask)

	return preallocatedPointer, nil
}

// RenderPassInputAttachmentAspectCreateInfo specifies, for a given subpass/input attachment
// pair, which aspect can be read
type RenderPassInputAttachmentAspectCreateInfo struct {
	// AspectReferences is a slice of InputAttachmentAspectReference structures containing
	// a mask describing which aspect(s) can be accessed for a given input attachment within a
	// given subpass
	AspectReferences []InputAttachmentAspectReference

	common.NextOptions
}

func (o RenderPassInputAttachmentAspectCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassInputAttachmentAspectCreateInfo{})))
	}

	createInfo := (*C.VkRenderPassInputAttachmentAspectCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO
	createInfo.pNext = next

	count := len(o.AspectReferences)
	if count < 1 {
		return nil, errors.New("options RenderPassInputAttachmentAspectCreateInfo must include at least 1 entry in AspectReferences")
	}

	createInfo.aspectReferenceCount = C.uint32_t(count)
	references, err := common.AllocSlice[C.VkInputAttachmentAspectReference, InputAttachmentAspectReference](allocator, o.AspectReferences)
	if err != nil {
		return nil, err
	}
	createInfo.pAspectReferences = references

	return preallocatedPointer, nil
}

////

// DeviceGroupRenderPassBeginInfo sets the initial Device mask and render areas for a RenderPass
// instance
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html
type DeviceGroupRenderPassBeginInfo struct {
	// DeviceMask is the deivce mask for the RenderPass instance
	DeviceMask uint32
	// DeviceRenderAreas is a slice of Rect2D structures defining the render area for each
	// PhysicalDevice
	DeviceRenderAreas []core1_0.Rect2D

	common.NextOptions
}

func (o DeviceGroupRenderPassBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupRenderPassBeginInfo{})))
	}

	info := (*C.VkDeviceGroupRenderPassBeginInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO
	info.pNext = next
	info.deviceMask = C.uint32_t(o.DeviceMask)

	count := len(o.DeviceRenderAreas)
	info.deviceRenderAreaCount = C.uint32_t(count)
	info.pDeviceRenderAreas = nil

	if count > 0 {
		areas := (*C.VkRect2D)(allocator.Malloc(count * C.sizeof_struct_VkRect2D))
		areaSlice := ([]C.VkRect2D)(unsafe.Slice(areas, count))

		for i := 0; i < count; i++ {
			areaSlice[i].offset.x = C.int32_t(o.DeviceRenderAreas[i].Offset.X)
			areaSlice[i].offset.y = C.int32_t(o.DeviceRenderAreas[i].Offset.Y)
			areaSlice[i].extent.width = C.uint32_t(o.DeviceRenderAreas[i].Extent.Width)
			areaSlice[i].extent.height = C.uint32_t(o.DeviceRenderAreas[i].Extent.Height)
		}

		info.pDeviceRenderAreas = areas
	}

	return preallocatedPointer, nil
}

////

// RenderPassMultiviewCreateInfo contains multiview information for all subpasses
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html
type RenderPassMultiviewCreateInfo struct {
	// ViewMasks is a slice of view masks, where each mask is a bitfield of view indices describing
	// which views rendering is broadcast to in each subpass, when multiview is enabled
	ViewMasks []uint32
	// ViewOffsets is a slice of view offsets, one for each subpass dependency. Each view offset
	// controls which view in the source subpass the views in the destination subpass depends on.
	ViewOffsets []int
	// CorrelationMasks is a slice of view masks indicating stes of views that may be
	// more efficient to render concurrently
	CorrelationMasks []uint32

	common.NextOptions
}

func (o RenderPassMultiviewCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassMultiviewCreateInfo{})))
	}

	info := (*C.VkRenderPassMultiviewCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO
	info.pNext = next

	count := len(o.ViewMasks)
	info.subpassCount = C.uint32_t(count)
	info.pViewMasks = nil
	if count > 0 {
		viewMasks := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		viewMaskSlice := ([]C.uint32_t)(unsafe.Slice(viewMasks, count))

		for i := 0; i < count; i++ {
			viewMaskSlice[i] = C.uint32_t(o.ViewMasks[i])
		}
		info.pViewMasks = viewMasks
	}

	count = len(o.ViewOffsets)
	info.dependencyCount = C.uint32_t(count)
	info.pViewOffsets = nil
	if count > 0 {
		viewOffsets := (*C.int32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.int32_t(0)))))
		viewOffsetSlice := ([]C.int32_t)(unsafe.Slice(viewOffsets, count))

		for i := 0; i < count; i++ {
			viewOffsetSlice[i] = C.int32_t(o.ViewOffsets[i])
		}
		info.pViewOffsets = viewOffsets
	}

	count = len(o.CorrelationMasks)
	info.correlationMaskCount = C.uint32_t(count)
	info.pCorrelationMasks = nil
	if count > 0 {
		correlationMasks := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		correlationMaskSlice := ([]C.uint32_t)(unsafe.Slice(correlationMasks, count))

		for i := 0; i < count; i++ {
			correlationMaskSlice[i] = C.uint32_t(o.CorrelationMasks[i])
		}
		info.pCorrelationMasks = correlationMasks
	}

	return preallocatedPointer, nil
}
