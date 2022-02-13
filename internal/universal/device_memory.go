package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0/options"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanDeviceMemory struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkDeviceMemory

	size int
}

func (m *VulkanDeviceMemory) Handle() driver.VkDeviceMemory {
	return m.handle
}

func (m *VulkanDeviceMemory) DeviceHandle() driver.VkDevice {
	return m.device
}

func (m *VulkanDeviceMemory) Driver() driver.Driver {
	return m.driver
}

func (m *VulkanDeviceMemory) MapMemory(offset int, size int, flags options.MemoryMapFlags) (unsafe.Pointer, common.VkResult, error) {
	var data unsafe.Pointer
	res, err := m.driver.VkMapMemory(m.device, m.handle, driver.VkDeviceSize(offset), driver.VkDeviceSize(size), driver.VkMemoryMapFlags(flags), &data)
	if err != nil {
		return nil, res, err
	}

	return data, res, nil
}

func (m *VulkanDeviceMemory) UnmapMemory() {
	m.driver.VkUnmapMemory(m.device, m.handle)
}

func freeDeviceMemory(memory iface.DeviceMemory, allocationCallbacks *driver.AllocationCallbacks) {
	memory.Driver().VkFreeMemory(memory.DeviceHandle(), memory.Handle(), allocationCallbacks.Handle())
}

func (m *VulkanDeviceMemory) Free(allocationCallbacks *driver.AllocationCallbacks) {
	freeDeviceMemory(m, allocationCallbacks)
}

func (m *VulkanDeviceMemory) Commitment() int {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	committedMemoryPtr := (*driver.VkDeviceSize)(arena.Malloc(8))

	m.driver.VkGetDeviceMemoryCommitment(m.device, m.handle, committedMemoryPtr)

	return int(*committedMemoryPtr)
}

func (m *VulkanDeviceMemory) Flush() (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	mappedRange := (*C.VkMappedMemoryRange)(arena.Malloc(C.sizeof_struct_VkMappedMemoryRange))
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = nil
	mappedRange.memory = C.VkDeviceMemory(unsafe.Pointer(m.handle))
	mappedRange.offset = 0
	mappedRange.size = C.VkDeviceSize(m.size)

	return m.driver.VkFlushMappedMemoryRanges(m.device, driver.Uint32(1), (*driver.VkMappedMemoryRange)(unsafe.Pointer(mappedRange)))
}

func (m *VulkanDeviceMemory) Invalidate() (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	mappedRange := (*C.VkMappedMemoryRange)(arena.Malloc(C.sizeof_struct_VkMappedMemoryRange))
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = nil
	mappedRange.memory = C.VkDeviceMemory(unsafe.Pointer(m.handle))
	mappedRange.offset = 0
	mappedRange.size = C.VkDeviceSize(m.size)

	return m.driver.VkInvalidateMappedMemoryRanges(m.device, driver.Uint32(1), (*driver.VkMappedMemoryRange)(unsafe.Pointer(mappedRange)))
}
