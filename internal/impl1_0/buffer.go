package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanBuffer is an implementation of the Buffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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
}

func (b *VulkanBuffer) MemoryRequirements() *core1_0.MemoryRequirements {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	requirementsUnsafe := allocator.Malloc(C.sizeof_struct_VkMemoryRequirements)

	b.DeviceDriver.VkGetBufferMemoryRequirements(b.Device, b.BufferHandle, (*driver.VkMemoryRequirements)(requirementsUnsafe))

	requirements := (*C.VkMemoryRequirements)(requirementsUnsafe)

	return &core1_0.MemoryRequirements{
		Size:           int(requirements.size),
		Alignment:      int(requirements.alignment),
		MemoryTypeBits: uint32(requirements.memoryTypeBits),
	}
}

func (b *VulkanBuffer) BindBufferMemory(memory core1_0.DeviceMemory, offset int) (common.VkResult, error) {
	if memory == nil {
		return core1_0.VKErrorUnknown, errors.New("received nil DeviceMemory")
	}

	return b.DeviceDriver.VkBindBufferMemory(b.Device, b.BufferHandle, memory.Handle(), driver.VkDeviceSize(offset))
}
