package internal1_0

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
	"github.com/cockroachdb/errors"
)

type VulkanBuffer struct {
	DeviceDriver driver.Driver
	Device       driver.VkDevice
	BufferHandle driver.VkBuffer

	MaximumAPIVersion common.APIVersion
}

func (b *VulkanBuffer) Handle() driver.VkBuffer {
	return b.BufferHandle
}

func (b *VulkanBuffer) DeviceHandle() driver.VkDevice {
	return b.Device
}

func (b *VulkanBuffer) Driver() driver.Driver {
	return b.DeviceDriver
}

func (b *VulkanBuffer) APIVersion() common.APIVersion {
	return b.MaximumAPIVersion
}

func (b *VulkanBuffer) Destroy(allocationCallbacks *driver.AllocationCallbacks) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	b.DeviceDriver.VkDestroyBuffer(b.Device, b.BufferHandle, allocationCallbacks.Handle())
	b.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(b.BufferHandle))
}

func (b *VulkanBuffer) MemoryRequirements() *core1_0.MemoryRequirements {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	requirementsUnsafe := allocator.Malloc(C.sizeof_struct_VkMemoryRequirements)

	b.DeviceDriver.VkGetBufferMemoryRequirements(b.Device, b.BufferHandle, (*driver.VkMemoryRequirements)(requirementsUnsafe))

	requirements := (*C.VkMemoryRequirements)(requirementsUnsafe)

	return &core1_0.MemoryRequirements{
		Size:       int(requirements.size),
		Alignment:  int(requirements.alignment),
		MemoryType: uint32(requirements.memoryTypeBits),
	}
}

func (b *VulkanBuffer) BindBufferMemory(memory core1_0.DeviceMemory, offset int) (common.VkResult, error) {
	if memory == nil {
		return core1_0.VKErrorUnknown, errors.New("received nil DeviceMemory")
	}

	return b.DeviceDriver.VkBindBufferMemory(b.Device, b.BufferHandle, memory.Handle(), driver.VkDeviceSize(offset))
}
