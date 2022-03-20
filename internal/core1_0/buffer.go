package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
)

type VulkanBuffer struct {
	Driver       driver.Driver
	Device       driver.VkDevice
	BufferHandle driver.VkBuffer

	Buffer1_1 core1_1.Buffer

	MaximumAPIVersion common.APIVersion
}

func (b *VulkanBuffer) Handle() driver.VkBuffer {
	return b.BufferHandle
}

func (b *VulkanBuffer) Core1_1() core1_1.Buffer {
	return b.Buffer1_1
}

func (b *VulkanBuffer) Destroy(allocationCallbacks *driver.AllocationCallbacks) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	b.Driver.VkDestroyBuffer(b.Device, b.BufferHandle, allocationCallbacks.Handle())
	b.Driver.ObjectStore().Delete(driver.VulkanHandle(b.BufferHandle), b)
}

func (b *VulkanBuffer) MemoryRequirements() *core1_0.MemoryRequirements {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	requirementsUnsafe := allocator.Malloc(C.sizeof_struct_VkMemoryRequirements)

	b.Driver.VkGetBufferMemoryRequirements(b.Device, b.BufferHandle, (*driver.VkMemoryRequirements)(requirementsUnsafe))

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

	return b.Driver.VkBindBufferMemory(b.Device, b.BufferHandle, memory.Handle(), driver.VkDeviceSize(offset))
}
