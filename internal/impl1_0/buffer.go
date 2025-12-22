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
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyBuffer(buffer types.Buffer, allocationCallbacks *driver.AllocationCallbacks) {
	if buffer.Handle() == 0 {
		panic("buffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	v.Driver.VkDestroyBuffer(buffer.DeviceHandle(), buffer.Handle(), allocationCallbacks.Handle())
}

func (v *Vulkan) GetBufferMemoryRequirements(buffer types.Buffer) *core1_0.MemoryRequirements {
	if buffer.Handle() == 0 {
		panic("buffer cannot be uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	requirementsUnsafe := allocator.Malloc(C.sizeof_struct_VkMemoryRequirements)

	v.Driver.VkGetBufferMemoryRequirements(buffer.DeviceHandle(), buffer.Handle(), (*driver.VkMemoryRequirements)(requirementsUnsafe))

	requirements := (*C.VkMemoryRequirements)(requirementsUnsafe)

	return &core1_0.MemoryRequirements{
		Size:           int(requirements.size),
		Alignment:      int(requirements.alignment),
		MemoryTypeBits: uint32(requirements.memoryTypeBits),
	}
}

func (v *Vulkan) BindBufferMemory(buffer types.Buffer, memory types.DeviceMemory, offset int) (common.VkResult, error) {
	if buffer.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("received uninitialized Buffer")
	}

	if memory.Handle() == 0 {
		return core1_0.VKErrorUnknown, errors.New("received uninitialized DeviceMemory")
	}

	return v.Driver.VkBindBufferMemory(buffer.DeviceHandle(), buffer.Handle(), memory.Handle(), driver.VkDeviceSize(offset))
}
