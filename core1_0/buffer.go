package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanBuffer is an implementation of the Buffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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
		Size:           int(requirements.size),
		Alignment:      int(requirements.alignment),
		MemoryTypeBits: uint32(requirements.memoryTypeBits),
	}
}

func (b *VulkanBuffer) BindBufferMemory(memory DeviceMemory, offset int) (common.VkResult, error) {
	if memory == nil {
		return VKErrorUnknown, errors.New("received nil DeviceMemory")
	}

	return b.deviceDriver.VkBindBufferMemory(b.device, b.bufferHandle, memory.Handle(), driver.VkDeviceSize(offset))
}
