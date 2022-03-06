package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanDeviceMemory struct {
	DeviceDriver       driver.Driver
	Device             driver.VkDevice
	DeviceMemoryHandle driver.VkDeviceMemory

	MaximumAPIVersion common.APIVersion

	size int
}

func (m *VulkanDeviceMemory) Handle() driver.VkDeviceMemory {
	return m.DeviceMemoryHandle
}

func (m *VulkanDeviceMemory) DeviceHandle() driver.VkDevice {
	return m.Device
}

func (m *VulkanDeviceMemory) Driver() driver.Driver {
	return m.DeviceDriver
}

func (m *VulkanDeviceMemory) MapMemory(offset int, size int, flags core1_0.MemoryMapFlags) (unsafe.Pointer, common.VkResult, error) {
	var data unsafe.Pointer
	res, err := m.DeviceDriver.VkMapMemory(m.Device, m.DeviceMemoryHandle, driver.VkDeviceSize(offset), driver.VkDeviceSize(size), driver.VkMemoryMapFlags(flags), &data)
	if err != nil {
		return nil, res, err
	}

	return data, res, nil
}

func (m *VulkanDeviceMemory) UnmapMemory() {
	m.DeviceDriver.VkUnmapMemory(m.Device, m.DeviceMemoryHandle)
}

func freeDeviceMemory(memory core1_0.DeviceMemory, allocationCallbacks *driver.AllocationCallbacks) {
	memory.Driver().VkFreeMemory(memory.DeviceHandle(), memory.Handle(), allocationCallbacks.Handle())
}

func (m *VulkanDeviceMemory) Free(allocationCallbacks *driver.AllocationCallbacks) {
	freeDeviceMemory(m, allocationCallbacks)
}

func (m *VulkanDeviceMemory) Commitment() int {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	committedMemoryPtr := (*driver.VkDeviceSize)(arena.Malloc(8))

	m.DeviceDriver.VkGetDeviceMemoryCommitment(m.Device, m.DeviceMemoryHandle, committedMemoryPtr)

	return int(*committedMemoryPtr)
}

func (m *VulkanDeviceMemory) Flush() (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	mappedRange := (*C.VkMappedMemoryRange)(arena.Malloc(C.sizeof_struct_VkMappedMemoryRange))
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = nil
	mappedRange.memory = C.VkDeviceMemory(unsafe.Pointer(m.DeviceMemoryHandle))
	mappedRange.offset = 0
	mappedRange.size = C.VkDeviceSize(m.size)

	return m.DeviceDriver.VkFlushMappedMemoryRanges(m.Device, driver.Uint32(1), (*driver.VkMappedMemoryRange)(unsafe.Pointer(mappedRange)))
}

func (m *VulkanDeviceMemory) Invalidate() (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	mappedRange := (*C.VkMappedMemoryRange)(arena.Malloc(C.sizeof_struct_VkMappedMemoryRange))
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = nil
	mappedRange.memory = C.VkDeviceMemory(unsafe.Pointer(m.DeviceMemoryHandle))
	mappedRange.offset = 0
	mappedRange.size = C.VkDeviceSize(m.size)

	return m.DeviceDriver.VkInvalidateMappedMemoryRanges(m.Device, driver.Uint32(1), (*driver.VkMappedMemoryRange)(unsafe.Pointer(mappedRange)))
}
