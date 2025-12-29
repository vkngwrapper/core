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
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

// BindBufferMemoryInfo specifies how to bind a Buffer to DeviceMemory
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html
type BindBufferMemoryInfo struct {
	// Buffer is the Buffer to be attached to memory
	Buffer core1_0.Buffer
	// Memory describes the DeviceMemory object to attach
	Memory core1_0.DeviceMemory
	// MemoryOffset is the start offset of the region of memory which is to be bound to the Buffer
	MemoryOffset int

	common.NextOptions
}

func (o BindBufferMemoryInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.Buffer.Initialized() {
		return nil, errors.Errorf("core1_1.BindBufferMemoryInfo.Buffer cannot be left unset")
	}
	if !o.Memory.Initialized() {
		return nil, errors.Errorf("core1_1.BindBufferMemoryInfo.Memory cannot be left unset")
	}
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

// BindImageMemoryInfo specifies how to bind an Image to DeviceMemory
type BindImageMemoryInfo struct {
	// Image is the image to be attached to DeviceMemory
	Image core1_0.Image
	// Memory describes the DeviceMemory to attach
	Memory core1_0.DeviceMemory
	// MemoryOffset is the start offset of the region of DeviceMemory to be bound to the Image
	MemoryOffset uint64

	common.NextOptions
}

func (o BindImageMemoryInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.Image.Initialized() {
		return nil, errors.Errorf("core1_1.BindImageMemoryInfo.Image cannot be left unset")
	}
	if !o.Memory.Initialized() {
		return nil, errors.Errorf("core1_1.BindImageMemoryInfo.Memory cannot be left unset")
	}
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

// BindBufferMemoryDeviceGroupInfo specifies Device within a group to bind to
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryDeviceGroupInfo.html
type BindBufferMemoryDeviceGroupInfo struct {
	// DeviceIndices is a slice of Device indices
	DeviceIndices []int

	common.NextOptions
}

func (o BindBufferMemoryDeviceGroupInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

// BindImageMemoryDeviceGroupInfo specifies Device within a group to bind to
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html
type BindImageMemoryDeviceGroupInfo struct {
	// DeviceIndices is a slice of Device indices
	DeviceIndices []int
	// SplitInstanceBindRegions is a slice of Rect2D structures describing which regions of
	// the Image are attached to each instance of DeviceMemory
	SplitInstanceBindRegions []core1_0.Rect2D

	common.NextOptions
}

func (o BindImageMemoryDeviceGroupInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

// BindImagePlaneMemoryInfo specifies how to bind an Image plane to DeviceMemory
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html
type BindImagePlaneMemoryInfo struct {
	// PlaneAspect specifies the aspect of the disjoint Image plane to bind
	PlaneAspect core1_0.ImageAspectFlags

	common.NextOptions
}

func (o BindImagePlaneMemoryInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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

// DeviceGroupBindSparseInfo indicates which instances are bound
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupBindSparseInfo.html
type DeviceGroupBindSparseInfo struct {
	// ResourceDeviceIndex is a Device index indicating which instance of the resource is bound
	ResourceDeviceIndex int
	// MemoryDeviceIndex is a Device index indicating which instance of the memory the resource instance is bound to
	MemoryDeviceIndex int

	common.NextOptions
}

func (o DeviceGroupBindSparseInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
