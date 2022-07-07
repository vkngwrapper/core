package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type MemoryAllocateFlags int32

var memoryAllocateFlagsMapping = common.NewFlagStringMapping[MemoryAllocateFlags]()

func (f MemoryAllocateFlags) Register(str string) {
	memoryAllocateFlagsMapping.Register(f, str)
}

func (f MemoryAllocateFlags) String() string {
	return memoryAllocateFlagsMapping.FlagsToString(f)
}

////

const (
	MemoryAllocateDeviceMask MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT
)

func init() {
	MemoryAllocateDeviceMask.Register("Device Mask")
}

////

type DeviceGroupDeviceCreateInfo struct {
	PhysicalDevices []core1_0.PhysicalDevice

	common.NextOptions
}

func (o DeviceGroupDeviceCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupDeviceCreateInfoKHR{})))
	}

	if len(o.PhysicalDevices) < 1 {
		return nil, errors.New("must include at least one physical device in DeviceGroupDeviceCreateInfo")
	}

	createInfo := (*C.VkDeviceGroupDeviceCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR
	createInfo.pNext = next

	count := len(o.PhysicalDevices)
	createInfo.physicalDeviceCount = C.uint32_t(count)
	physicalDevicesPtr := (*C.VkPhysicalDevice)(allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkPhysicalDevice{}))))
	physicalDevicesSlice := ([]C.VkPhysicalDevice)(unsafe.Slice(physicalDevicesPtr, count))

	for i := 0; i < count; i++ {
		physicalDevicesSlice[i] = C.VkPhysicalDevice(unsafe.Pointer(o.PhysicalDevices[i].Handle()))
	}
	createInfo.pPhysicalDevices = physicalDevicesPtr
	return preallocatedPointer, nil
}

////

type MemoryAllocateFlagsInfo struct {
	Flags      MemoryAllocateFlags
	DeviceMask uint32

	common.NextOptions
}

func (o MemoryAllocateFlagsInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryAllocateFlagsInfo{})))
	}

	createInfo := (*C.VkMemoryAllocateFlagsInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO
	createInfo.pNext = next
	createInfo.flags = C.VkMemoryAllocateFlags(o.Flags)
	createInfo.deviceMask = C.uint32_t(o.DeviceMask)

	return preallocatedPointer, nil
}
