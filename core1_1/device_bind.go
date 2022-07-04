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

type BindBufferMemoryOptions struct {
	Buffer       core1_0.Buffer
	Memory       core1_0.DeviceMemory
	MemoryOffset int

	common.NextOptions
}

func (o BindBufferMemoryOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindBufferMemoryInfo{})))
	}

	createInfo := (*C.VkBindBufferMemoryInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
	createInfo.pNext = next
	createInfo.buffer = (C.VkBuffer)(unsafe.Pointer(o.Buffer.Handle()))
	createInfo.memory = (C.VkDeviceMemory)(unsafe.Pointer(o.Memory.Handle()))
	createInfo.memoryOffset = C.VkDeviceSize(o.MemoryOffset)

	return preallocatedPointer, nil
}

////

type BindImageMemoryOptions struct {
	Image        core1_0.Image
	Memory       core1_0.DeviceMemory
	MemoryOffset uint64

	common.NextOptions
}

func (o BindImageMemoryOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindImageMemoryInfo{})))
	}

	createInfo := (*C.VkBindImageMemoryInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
	createInfo.pNext = next
	createInfo.image = (C.VkImage)(unsafe.Pointer(o.Image.Handle()))
	createInfo.memory = (C.VkDeviceMemory)(unsafe.Pointer(o.Memory.Handle()))
	createInfo.memoryOffset = C.VkDeviceSize(o.MemoryOffset)

	return preallocatedPointer, nil
}

////

type BindBufferMemoryDeviceGroupOptions struct {
	DeviceIndices []int

	common.NextOptions
}

func (o BindBufferMemoryDeviceGroupOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindBufferMemoryDeviceGroupInfo{})))
	}

	info := (*C.VkBindBufferMemoryDeviceGroupInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO
	info.pNext = next

	count := len(o.DeviceIndices)
	info.deviceIndexCount = C.uint32_t(count)
	info.pDeviceIndices = nil

	if count > 0 {
		indices := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		indexSlice := ([]C.uint32_t)(unsafe.Slice(indices, count))

		for i := 0; i < count; i++ {
			indexSlice[i] = C.uint32_t(o.DeviceIndices[i])
		}

		info.pDeviceIndices = indices
	}

	return preallocatedPointer, nil
}

////

type BindImageMemoryDeviceGroupOptions struct {
	DeviceIndices            []int
	SplitInstanceBindRegions []core1_0.Rect2D

	common.NextOptions
}

func (o BindImageMemoryDeviceGroupOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindImageMemoryDeviceGroupInfo{})))
	}

	info := (*C.VkBindImageMemoryDeviceGroupInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO
	info.pNext = next

	count := len(o.DeviceIndices)
	info.deviceIndexCount = C.uint32_t(count)
	info.pDeviceIndices = nil
	if count > 0 {
		indices := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		indexSlice := ([]C.uint32_t)(unsafe.Slice(indices, count))

		for i := 0; i < count; i++ {
			indexSlice[i] = C.uint32_t(o.DeviceIndices[i])
		}

		info.pDeviceIndices = indices
	}

	count = len(o.SplitInstanceBindRegions)
	info.splitInstanceBindRegionCount = C.uint32_t(count)
	info.pSplitInstanceBindRegions = nil
	if count > 0 {
		regions := (*C.VkRect2D)(allocator.Malloc(count * C.sizeof_struct_VkRect2D))
		regionSlice := ([]C.VkRect2D)(unsafe.Slice(regions, count))

		for i := 0; i < count; i++ {
			regionSlice[i].offset.x = C.int32_t(o.SplitInstanceBindRegions[i].Offset.X)
			regionSlice[i].offset.y = C.int32_t(o.SplitInstanceBindRegions[i].Offset.Y)
			regionSlice[i].extent.width = C.uint32_t(o.SplitInstanceBindRegions[i].Extent.Width)
			regionSlice[i].extent.height = C.uint32_t(o.SplitInstanceBindRegions[i].Extent.Height)
		}

		info.pSplitInstanceBindRegions = regions
	}

	return preallocatedPointer, nil
}

////

type BindImagePlaneMemoryOptions struct {
	PlaneAspect core1_0.ImageAspectFlags

	common.NextOptions
}

func (o BindImagePlaneMemoryOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindImagePlaneMemoryInfo{})))
	}

	info := (*C.VkBindImagePlaneMemoryInfo)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO
	info.pNext = next
	info.planeAspect = C.VkImageAspectFlagBits(o.PlaneAspect)

	return preallocatedPointer, nil
}

////

type DeviceGroupBindSparseOptions struct {
	ResourceDeviceIndex int
	MemoryDeviceIndex   int

	common.NextOptions
}

func (o DeviceGroupBindSparseOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupBindSparseInfo{})))
	}

	createInfo := (*C.VkDeviceGroupBindSparseInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO
	createInfo.pNext = next
	createInfo.resourceDeviceIndex = C.uint32_t(o.ResourceDeviceIndex)
	createInfo.memoryDeviceIndex = C.uint32_t(o.MemoryDeviceIndex)

	return preallocatedPointer, nil
}
