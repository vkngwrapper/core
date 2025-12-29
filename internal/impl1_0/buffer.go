package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) DestroyBuffer(buffer core.Buffer, allocationCallbacks *loader.AllocationCallbacks) {
	if !buffer.Initialized() {
		panic("buffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	v.LoaderObj.VkDestroyBuffer(buffer.DeviceHandle(), buffer.Handle(), allocationCallbacks.Handle())
}

func (v *DeviceVulkanDriver) GetBufferMemoryRequirements(buffer core.Buffer) *core1_0.MemoryRequirements {
	if !buffer.Initialized() {
		panic("buffer cannot be uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	requirementsUnsafe := allocator.Malloc(C.sizeof_struct_VkMemoryRequirements)

	v.LoaderObj.VkGetBufferMemoryRequirements(buffer.DeviceHandle(), buffer.Handle(), (*loader.VkMemoryRequirements)(requirementsUnsafe))

	requirements := (*C.VkMemoryRequirements)(requirementsUnsafe)

	return &core1_0.MemoryRequirements{
		Size:           int(requirements.size),
		Alignment:      int(requirements.alignment),
		MemoryTypeBits: uint32(requirements.memoryTypeBits),
	}
}

func (v *DeviceVulkanDriver) BindBufferMemory(buffer core.Buffer, memory core.DeviceMemory, offset int) (common.VkResult, error) {
	if !buffer.Initialized() {
		return core1_0.VKErrorUnknown, errors.New("received uninitialized Buffer")
	}

	if !memory.Initialized() {
		return core1_0.VKErrorUnknown, errors.New("received uninitialized DeviceMemory")
	}

	return v.LoaderObj.VkBindBufferMemory(buffer.DeviceHandle(), buffer.Handle(), memory.Handle(), loader.VkDeviceSize(offset))
}
