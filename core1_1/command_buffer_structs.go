package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DeviceGroupCommandBufferBeginInfo struct {
	DeviceMask uint32

	common.NextOptions
}

func (o DeviceGroupCommandBufferBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupCommandBufferBeginInfo{})))
	}

	createInfo := (*C.VkDeviceGroupCommandBufferBeginInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO
	createInfo.pNext = next
	createInfo.deviceMask = C.uint32_t(o.DeviceMask)

	return preallocatedPointer, nil
}
