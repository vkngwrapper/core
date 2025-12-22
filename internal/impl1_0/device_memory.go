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
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) MapMemory(deviceMemory types.DeviceMemory, offset int, size int, flags core1_0.MemoryMapFlags) (unsafe.Pointer, common.VkResult, error) {
	if deviceMemory.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, errors.New("deviceMemory was uninitialized")
	}

	var data unsafe.Pointer
	res, err := v.Driver.VkMapMemory(deviceMemory.DeviceHandle(), deviceMemory.Handle(), driver.VkDeviceSize(offset), driver.VkDeviceSize(size), driver.VkMemoryMapFlags(flags), &data)
	if err != nil {
		return nil, res, err
	}

	return data, res, nil
}

func (v *Vulkan) UnmapMemory(deviceMemory types.DeviceMemory) {
	if deviceMemory.Handle() == 0 {
		panic("deviceMemory was uninitialized")
	}
	v.Driver.VkUnmapMemory(deviceMemory.DeviceHandle(), deviceMemory.Handle())
}

func (v *Vulkan) FreeMemory(deviceMemory types.DeviceMemory, allocationCallbacks *driver.AllocationCallbacks) {
	if deviceMemory.Handle() == 0 {
		panic("deviceMemory was uninitialized")
	}

	v.Driver.VkFreeMemory(deviceMemory.DeviceHandle(), deviceMemory.Handle(), allocationCallbacks.Handle())
}

func (v *Vulkan) GetDeviceMemoryCommitment(deviceMemory types.DeviceMemory) int {
	if deviceMemory.Handle() == 0 {
		panic("deviceMemory was uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	committedMemoryPtr := (*driver.VkDeviceSize)(arena.Malloc(8))

	v.Driver.VkGetDeviceMemoryCommitment(deviceMemory.DeviceHandle(), deviceMemory.Handle(), committedMemoryPtr)

	return int(*committedMemoryPtr)
}
