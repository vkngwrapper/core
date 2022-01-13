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

type MemoryMapFlags int32

func (f MemoryMapFlags) String() string {
	return "None"
}

type DeviceMemoryOptions struct {
	AllocationSize  int
	MemoryTypeIndex int

	common.HaveNext
}

func (o *DeviceMemoryOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkMemoryAllocateInfo)(allocator.Malloc(C.sizeof_struct_VkMemoryAllocateInfo))

	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
	createInfo.allocationSize = C.VkDeviceSize(o.AllocationSize)
	createInfo.memoryTypeIndex = C.uint32_t(o.MemoryTypeIndex)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}

type vulkanDeviceMemory struct {
	driver Driver
	device VkDevice
	handle VkDeviceMemory

	size int
}

func (m *vulkanDeviceMemory) Handle() VkDeviceMemory {
	return m.handle
}

func (m *vulkanDeviceMemory) MapMemory(offset int, size int, flags MemoryMapFlags) (unsafe.Pointer, VkResult, error) {
	var data unsafe.Pointer
	res, err := m.driver.VkMapMemory(m.device, m.handle, VkDeviceSize(offset), VkDeviceSize(size), VkMemoryMapFlags(flags), &data)
	if err != nil {
		return nil, res, err
	}

	return data, res, nil
}

func (m *vulkanDeviceMemory) UnmapMemory() {
	m.driver.VkUnmapMemory(m.device, m.handle)
}

func (m *vulkanDeviceMemory) Free(allocationCallbacks *AllocationCallbacks) {
	m.driver.VkFreeMemory(m.device, m.handle, allocationCallbacks.Handle())
}

func (m *vulkanDeviceMemory) Commitment() int {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	committedMemoryPtr := (*VkDeviceSize)(arena.Malloc(8))

	m.driver.VkGetDeviceMemoryCommitment(m.device, m.handle, committedMemoryPtr)

	return int(*committedMemoryPtr)
}

func (m *vulkanDeviceMemory) Flush() (VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	mappedRange := (*C.VkMappedMemoryRange)(arena.Malloc(C.sizeof_struct_VkMappedMemoryRange))
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = nil
	mappedRange.memory = C.VkDeviceMemory(m.handle)
	mappedRange.offset = 0
	mappedRange.size = C.VkDeviceSize(m.size)

	return m.driver.VkFlushMappedMemoryRanges(m.device, Uint32(1), (*VkMappedMemoryRange)(mappedRange))
}

func (m *vulkanDeviceMemory) Invalidate() (VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	mappedRange := (*C.VkMappedMemoryRange)(arena.Malloc(C.sizeof_struct_VkMappedMemoryRange))
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = nil
	mappedRange.memory = C.VkDeviceMemory(m.handle)
	mappedRange.offset = 0
	mappedRange.size = C.VkDeviceSize(m.size)

	return m.driver.VkInvalidateMappedMemoryRanges(m.device, Uint32(1), (*VkMappedMemoryRange)(mappedRange))
}
