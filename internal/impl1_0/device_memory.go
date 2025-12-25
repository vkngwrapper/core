package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) MapMemory(deviceMemory core.DeviceMemory, offset int, size int, flags core1_0.MemoryMapFlags) (unsafe.Pointer, common.VkResult, error) {
	if deviceMemory.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, errors.New("deviceMemory was uninitialized")
	}

	var data unsafe.Pointer
	res, err := v.LoaderObj.VkMapMemory(deviceMemory.DeviceHandle(), deviceMemory.Handle(), loader.VkDeviceSize(offset), loader.VkDeviceSize(size), loader.VkMemoryMapFlags(flags), &data)
	if err != nil {
		return nil, res, err
	}

	return data, res, nil
}

func (v *DeviceVulkanDriver) UnmapMemory(deviceMemory core.DeviceMemory) {
	if deviceMemory.Handle() == 0 {
		panic("deviceMemory was uninitialized")
	}
	v.LoaderObj.VkUnmapMemory(deviceMemory.DeviceHandle(), deviceMemory.Handle())
}

func (v *DeviceVulkanDriver) FreeMemory(deviceMemory core.DeviceMemory, allocationCallbacks *loader.AllocationCallbacks) {
	if deviceMemory.Handle() == 0 {
		panic("deviceMemory was uninitialized")
	}

	v.LoaderObj.VkFreeMemory(deviceMemory.DeviceHandle(), deviceMemory.Handle(), allocationCallbacks.Handle())
}

func (v *DeviceVulkanDriver) GetDeviceMemoryCommitment(deviceMemory core.DeviceMemory) int {
	if deviceMemory.Handle() == 0 {
		panic("deviceMemory was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	committedMemoryPtr := (*loader.VkDeviceSize)(arena.Malloc(8))

	v.LoaderObj.VkGetDeviceMemoryCommitment(deviceMemory.DeviceHandle(), deviceMemory.Handle(), committedMemoryPtr)

	return int(*committedMemoryPtr)
}
