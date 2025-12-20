package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

////

// VulkanDeviceMemory is an implementation of the DeviceMemory interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDeviceMemory struct {
	DeviceDriver       driver.Driver
	Device             driver.VkDevice
	DeviceMemoryHandle driver.VkDeviceMemory

	MaximumAPIVersion common.APIVersion

	Size int
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

func (m *VulkanDeviceMemory) APIVersion() common.APIVersion {
	return m.MaximumAPIVersion
}

func (m *VulkanDeviceMemory) Map(offset int, size int, flags core1_0.MemoryMapFlags) (unsafe.Pointer, common.VkResult, error) {
	var data unsafe.Pointer
	res, err := m.DeviceDriver.VkMapMemory(m.Device, m.DeviceMemoryHandle, driver.VkDeviceSize(offset), driver.VkDeviceSize(size), driver.VkMemoryMapFlags(flags), &data)
	if err != nil {
		return nil, res, err
	}

	return data, res, nil
}

func (m *VulkanDeviceMemory) Unmap() {
	m.DeviceDriver.VkUnmapMemory(m.Device, m.DeviceMemoryHandle)
}

func (m *VulkanDeviceMemory) Free(allocationCallbacks *driver.AllocationCallbacks) {
	m.Driver().VkFreeMemory(m.Device, m.DeviceMemoryHandle, allocationCallbacks.Handle())
}

func (m *VulkanDeviceMemory) Commitment() int {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	committedMemoryPtr := (*driver.VkDeviceSize)(arena.Malloc(8))

	m.DeviceDriver.VkGetDeviceMemoryCommitment(m.Device, m.DeviceMemoryHandle, committedMemoryPtr)

	return int(*committedMemoryPtr)
}

func (m *VulkanDeviceMemory) FlushAll() (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	mappedRange := (*C.VkMappedMemoryRange)(arena.Malloc(C.sizeof_struct_VkMappedMemoryRange))
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = nil
	mappedRange.memory = C.VkDeviceMemory(unsafe.Pointer(m.DeviceMemoryHandle))
	mappedRange.offset = 0
	mappedRange.size = C.VkDeviceSize(m.Size)

	return m.DeviceDriver.VkFlushMappedMemoryRanges(m.Device, driver.Uint32(1), (*driver.VkMappedMemoryRange)(unsafe.Pointer(mappedRange)))
}

func (m *VulkanDeviceMemory) InvalidateAll() (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	mappedRange := (*C.VkMappedMemoryRange)(arena.Malloc(C.sizeof_struct_VkMappedMemoryRange))
	mappedRange.sType = C.VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	mappedRange.pNext = nil
	mappedRange.memory = C.VkDeviceMemory(unsafe.Pointer(m.DeviceMemoryHandle))
	mappedRange.offset = 0
	mappedRange.size = C.VkDeviceSize(m.Size)

	return m.DeviceDriver.VkInvalidateMappedMemoryRanges(m.Device, driver.Uint32(1), (*driver.VkMappedMemoryRange)(unsafe.Pointer(mappedRange)))
}
