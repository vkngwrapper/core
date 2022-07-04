package core1_1

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type InputAttachmentAspectReference struct {
	Subpass              int
	InputAttachmentIndex int
	AspectMask           core1_0.ImageAspectFlags
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

type RenderPassInputAttachmentAspectOptions struct {
	AspectReferences []InputAttachmentAspectReference

	common.NextOptions
}

func (o RenderPassInputAttachmentAspectOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassInputAttachmentAspectCreateInfo{})))
	}

	createInfo := (*C.VkRenderPassInputAttachmentAspectCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO
	createInfo.pNext = next

	count := len(o.AspectReferences)
	if count < 1 {
		return nil, errors.New("options RenderPassInputAttachmentAspectOptions must include at least 1 entry in AspectReferences")
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

type DeviceGroupRenderPassBeginOptions struct {
	DeviceMask        uint32
	DeviceRenderAreas []core1_0.Rect2D

	common.NextOptions
}

func (o DeviceGroupRenderPassBeginOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

type RenderPassMultiviewOptions struct {
	SubpassViewMasks      []uint32
	DependencyViewOffsets []int
	CorrelationMasks      []uint32

	common.NextOptions
}

func (o RenderPassMultiviewOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassMultiviewCreateInfo{})))
	}

	info := (*C.VkRenderPassMultiviewCreateInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO
	info.pNext = next

	count := len(o.SubpassViewMasks)
	info.subpassCount = C.uint32_t(count)
	info.pViewMasks = nil
	if count > 0 {
		viewMasks := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		viewMaskSlice := ([]C.uint32_t)(unsafe.Slice(viewMasks, count))

		for i := 0; i < count; i++ {
			viewMaskSlice[i] = C.uint32_t(o.SubpassViewMasks[i])
		}
		info.pViewMasks = viewMasks
	}

	count = len(o.DependencyViewOffsets)
	info.dependencyCount = C.uint32_t(count)
	info.pViewOffsets = nil
	if count > 0 {
		viewOffsets := (*C.int32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.int32_t(0)))))
		viewOffsetSlice := ([]C.int32_t)(unsafe.Slice(viewOffsets, count))

		for i := 0; i < count; i++ {
			viewOffsetSlice[i] = C.int32_t(o.DependencyViewOffsets[i])
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
