package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

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

func (m *vulkanDeviceMemory) Free() error {
	return m.driver.VkFreeMemory(m.device, m.handle, nil)
}

func (m *vulkanDeviceMemory) MapMemory(offset int, size int) (unsafe.Pointer, VkResult, error) {
	var data unsafe.Pointer
	res, err := m.driver.VkMapMemory(m.device, m.handle, VkDeviceSize(offset), VkDeviceSize(size), 0, &data)
	if err != nil {
		return nil, res, err
	}

	return data, res, nil
}

func (m *vulkanDeviceMemory) UnmapMemory() error {
	return m.driver.VkUnmapMemory(m.device, m.handle)
}

func (m *vulkanDeviceMemory) WriteData(offset int, data interface{}) (VkResult, error) {
	bufferSize := binary.Size(data)

	memoryPtr, res, err := m.MapMemory(offset, bufferSize)
	if err != nil {
		return res, err
	}
	defer m.UnmapMemory()

	dataBuffer := unsafe.Slice((*byte)(memoryPtr), bufferSize)

	buf := &bytes.Buffer{}
	err = binary.Write(buf, common.ByteOrder, data)
	if err != nil {
		return VKErrorUnknown, err
	}

	copy(dataBuffer, buf.Bytes())
	return res, nil
}
