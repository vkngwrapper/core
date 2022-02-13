package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
)

type VulkanBuffer struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkBuffer
}

func (b *VulkanBuffer) Handle() driver.VkBuffer {
	return b.handle
}

func (b *VulkanBuffer) Destroy(allocationCallbacks *driver.AllocationCallbacks) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	b.driver.VkDestroyBuffer(b.device, b.handle, allocationCallbacks.Handle())
}

func (b *VulkanBuffer) MemoryRequirements() *core1_0.MemoryRequirements {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	requirementsUnsafe := allocator.Malloc(C.sizeof_struct_VkMemoryRequirements)

	b.driver.VkGetBufferMemoryRequirements(b.device, b.handle, (*driver.VkMemoryRequirements)(requirementsUnsafe))

	requirements := (*C.VkMemoryRequirements)(requirementsUnsafe)

	return &core1_0.MemoryRequirements{
		Size:       int(requirements.size),
		Alignment:  int(requirements.alignment),
		MemoryType: uint32(requirements.memoryTypeBits),
	}
}

func (b *VulkanBuffer) BindBufferMemory(memory iface.DeviceMemory, offset int) (common.VkResult, error) {
	if memory == nil {
		return common.VKErrorUnknown, errors.New("received nil DeviceMemory")
	}

	return b.driver.VkBindBufferMemory(b.device, b.handle, memory.Handle(), driver.VkDeviceSize(offset))
}
