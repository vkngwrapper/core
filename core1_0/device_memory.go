package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
	"unsafe"
)

type MemoryType struct {
	PropertyFlags MemoryPropertyFlags
	HeapIndex     int
}

type MemoryHeap struct {
	Size  int
	Flags MemoryHeapFlags
}

////

type VulkanDeviceMemory struct {
	deviceDriver       driver.Driver
	device             driver.VkDevice
	deviceMemoryHandle driver.VkDeviceMemory

	maximumAPIVersion common.APIVersion

	size int
}

func (m *VulkanDeviceMemory) Handle() driver.VkDeviceMemory {
	return m.deviceMemoryHandle
}

func (m *VulkanDeviceMemory) DeviceHandle() driver.VkDevice {
	return m.device
}

func (m *VulkanDeviceMemory) Driver() driver.Driver {
	return m.deviceDriver
}

func (m *VulkanDeviceMemory) APIVersion() common.APIVersion {
	return m.maximumAPIVersion
}

func (m *VulkanDeviceMemory) Map(offset int, size int, flags MemoryMapFlags) (unsafe.Pointer, common.VkResult, error) {
	var data unsafe.Pointer
	res, err := m.deviceDriver.VkMapMemory(m.device, m.deviceMemoryHandle, driver.VkDeviceSize(offset), driver.VkDeviceSize(size), driver.VkMemoryMapFlags(flags), &data)
	if err != nil {
		return nil, res, err
	}

	return data, res, nil
}

func (m *VulkanDeviceMemory) Unmap() {
	m.deviceDriver.VkUnmapMemory(m.device, m.deviceMemoryHandle)
}

func (m *VulkanDeviceMemory) Free(allocationCallbacks *driver.AllocationCallbacks) {
	m.Driver().VkFreeMemory(m.device, m.deviceMemoryHandle, allocationCallbacks.Handle())
	m.Driver().ObjectStore().Delete(driver.VulkanHandle(m.deviceMemoryHandle))
}

func (m *VulkanDeviceMemory) Commitment() int {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	committedMemoryPtr := (*driver.VkDeviceSize)(arena.Malloc(8))

	m.deviceDriver.VkGetDeviceMemoryCommitment(m.device, m.deviceMemoryHandle, committedMemoryPtr)

	return int(*committedMemoryPtr)
}

func (m *VulkanDeviceMemory) FlushAll() (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	mappedRange := (*C.VkMappedMemoryRange)(arena.Malloc(C.sizeof_struct_VkMappedMemoryRange))
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = nil
	mappedRange.memory = C.VkDeviceMemory(unsafe.Pointer(m.deviceMemoryHandle))
	mappedRange.offset = 0
	mappedRange.size = C.VkDeviceSize(m.size)

	return m.deviceDriver.VkFlushMappedMemoryRanges(m.device, driver.Uint32(1), (*driver.VkMappedMemoryRange)(unsafe.Pointer(mappedRange)))
}

func (m *VulkanDeviceMemory) InvalidateAll() (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	mappedRange := (*C.VkMappedMemoryRange)(arena.Malloc(C.sizeof_struct_VkMappedMemoryRange))
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = nil
	mappedRange.memory = C.VkDeviceMemory(unsafe.Pointer(m.deviceMemoryHandle))
	mappedRange.offset = 0
	mappedRange.size = C.VkDeviceSize(m.size)

	return m.deviceDriver.VkInvalidateMappedMemoryRanges(m.device, driver.Uint32(1), (*driver.VkMappedMemoryRange)(unsafe.Pointer(mappedRange)))
}
