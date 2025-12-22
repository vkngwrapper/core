package core1_1

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/types"
)

// MemoryAllocateFlags specifies flags for a DeviceMemory allocation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBits.html
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
	// MemoryAllocateDeviceMask specifies that memory will be allocated for the devices
	// in MemoryAllocateFlagsInfo.DeviceMask
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBitsKHR.html
	MemoryAllocateDeviceMask MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT
)

func init() {
	MemoryAllocateDeviceMask.Register("Device Mask")
}

////

// DeviceGroupDeviceCreateInfo creates a logical Device from multiple PhysicalDevice objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupDeviceCreateInfo.html
type DeviceGroupDeviceCreateInfo struct {
	// PhysicalDevices is a slice of PhysicalDevice objects belonging to the same Device group
	PhysicalDevices []types.PhysicalDevice

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
		if o.PhysicalDevices[i].Handle() == 0 {
			return nil, errors.Errorf("core1_1.DeviceGroupDeviceCreateInfo.PhysicalDevices cannot contain unset elements "+
				"elements, but elements %d is unset", i)
		}
		physicalDevicesSlice[i] = C.VkPhysicalDevice(unsafe.Pointer(o.PhysicalDevices[i].Handle()))
	}
	createInfo.pPhysicalDevices = physicalDevicesPtr
	return preallocatedPointer, nil
}

////

// MemoryAllocateFlagsInfo controls how many instances of memory will be allocated
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagsInfoKHR.html
type MemoryAllocateFlagsInfo struct {
	// Flags controls the allocation
	Flags MemoryAllocateFlags
	// DeviceMask is a mask of PhysicalDevice objects in the logical Device, indicating that
	// memory must be allocated on each Device in the mask, if MemoryAllocateDeviceMask is set
	// in flags
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
