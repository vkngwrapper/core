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

	common.HaveNext
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

func (o BindBufferMemoryOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkBindBufferMemoryInfo)(cDataPointer)
	return createInfo.pNext, nil
}

////

type BindImageMemoryOptions struct {
	Image        core1_0.Image
	Memory       core1_0.DeviceMemory
	MemoryOffset uint64

	common.HaveNext
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

func (o BindImageMemoryOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkBindImageMemoryInfo)(cDataPointer)
	return createInfo.pNext, nil
}
