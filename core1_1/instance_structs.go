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
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

// PhysicalDeviceGroupProperties specifies PhysicalDevice group properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGroupProperties.html
type PhysicalDeviceGroupProperties struct {
	// PhysicalDevices is a slice of PhysicalDevice objects that represent all PhysicalDevice
	// objects in the group
	PhysicalDevices []core.PhysicalDevice
	// SubsetAllocation specifies whether logical Device objects created from the group support
	// allocating DeviceMemory on a subset of Device objects, via MemoryAllocateFlagsInfo
	SubsetAllocation bool

	common.NextOutData
}

func (o *PhysicalDeviceGroupProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceGroupProperties{})))
	}

	createInfo := (*C.VkPhysicalDeviceGroupProperties)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
	createInfo.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceGroupProperties) PopulateOutData(cPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo := (*C.VkPhysicalDeviceGroupProperties)(cPointer)
	o.SubsetAllocation = createInfo.subsetAllocation != C.VkBool32(0)

	instanceDriver, ok := common.OfType[loader.Loader](helpers)
	if !ok {
		return nil, errors.New("outdata population requires an instance loader passed to populate helpers")
	}
	instanceVersion, ok := common.OfType[common.APIVersion](helpers)
	if !ok {
		return nil, errors.New("outdata population requires an instance version passed to populate helpers")
	}

	count := int(createInfo.physicalDeviceCount)
	o.PhysicalDevices = make([]core.PhysicalDevice, count)

	propertiesUnsafe := arena.Malloc(C.sizeof_struct_VkPhysicalDeviceProperties)

	for i := 0; i < count; i++ {
		handle := loader.VkPhysicalDevice(unsafe.Pointer(createInfo.physicalDevices[i]))
		instanceDriver.VkGetPhysicalDeviceProperties(handle, (*loader.VkPhysicalDeviceProperties)(propertiesUnsafe))

		var properties core1_0.PhysicalDeviceProperties
		err = (&properties).PopulateFromCPointer(propertiesUnsafe)
		if err != nil {
			return nil, err
		}

		deviceVersion := instanceVersion.Min(properties.APIVersion)

		o.PhysicalDevices[i] = core.InternalPhysicalDevice(handle, instanceVersion, deviceVersion)
	}

	return createInfo.pNext, nil
}
