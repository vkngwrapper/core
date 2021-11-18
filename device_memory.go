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

func (m *vulkanDeviceMemory) UnmapMemory() error {
	return m.driver.VkUnmapMemory(m.device, m.handle)
}
