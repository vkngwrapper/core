package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
)

type VulkanBuffer struct {
	deviceDriver driver.Driver
	device       driver.VkDevice
	bufferHandle driver.VkBuffer

	maximumAPIVersion common.APIVersion
}

func (b *VulkanBuffer) Handle() driver.VkBuffer {
	return b.bufferHandle
}

func (b *VulkanBuffer) DeviceHandle() driver.VkDevice {
	return b.device
}

func (b *VulkanBuffer) Driver() driver.Driver {
	return b.deviceDriver
}

func (b *VulkanBuffer) APIVersion() common.APIVersion {
	return b.maximumAPIVersion
}

func (b *VulkanBuffer) Destroy(allocationCallbacks *driver.AllocationCallbacks) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	b.deviceDriver.VkDestroyBuffer(b.device, b.bufferHandle, allocationCallbacks.Handle())
	b.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(b.bufferHandle))
}

func (b *VulkanBuffer) MemoryRequirements() *MemoryRequirements {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	requirementsUnsafe := allocator.Malloc(C.sizeof_struct_VkMemoryRequirements)

	b.deviceDriver.VkGetBufferMemoryRequirements(b.device, b.bufferHandle, (*driver.VkMemoryRequirements)(requirementsUnsafe))

	requirements := (*C.VkMemoryRequirements)(requirementsUnsafe)

	return &MemoryRequirements{
		Size:       int(requirements.size),
		Alignment:  int(requirements.alignment),
		MemoryType: uint32(requirements.memoryTypeBits),
	}
}

func (b *VulkanBuffer) BindBufferMemory(memory DeviceMemory, offset int) (common.VkResult, error) {
	if memory == nil {
		return VKErrorUnknown, errors.New("received nil DeviceMemory")
	}

	return b.deviceDriver.VkBindBufferMemory(b.device, b.bufferHandle, memory.Handle(), driver.VkDeviceSize(offset))
}
