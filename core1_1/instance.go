package core1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type DeviceGroupOutData struct {
	PhysicalDevices  []core1_0.PhysicalDevice
	SubsetAllocation bool

	common.HaveNext
}

func (o *DeviceGroupOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceGroupProperties{})))
	}

	createInfo := (*C.VkPhysicalDeviceGroupProperties)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
	createInfo.pNext = next

	return preallocatedPointer, nil
}

// We had a circular dependency problem here- objects.Create* methods all must interact with
// every core/core* version, which creates zany circular dependencies between the versions.
// In order to keep the dep graph acyclical, dependency flow must be very particular:
// core/core* may only include lower versions of core/core*. core/internal/core* may only
// include HIGHER versions of core/internal/core* but can include any version of core/core*.
//
// This is all no problem when objects.Create* is only included from core and core/internal/core*,
// but it poses a serious problem right here, in core/core*. I'm breaking the circular dependency
// by using a go:linkname and may god have mercy on my soul.

//go:linkname createPhysicalDevice github.com/CannibalVox/VKng/core.CreatePhysicalDevice
func createPhysicalDevice(coreDriver driver.Driver, instance driver.VkInstance, handle driver.VkPhysicalDevice, version common.APIVersion) core1_0.PhysicalDevice

func (o *DeviceGroupOutData) PopulateOutData(cPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo := (*C.VkPhysicalDeviceGroupProperties)(cPointer)
	o.SubsetAllocation = createInfo.subsetAllocation != C.VkBool32(0)

	instanceHandle, ok := common.OfType[driver.VkInstance](helpers)
	if !ok {
		return nil, errors.New("outdata population requires an instance handle passed to populate helpers")
	}
	instanceDriver, ok := common.OfType[driver.Driver](helpers)
	if !ok {
		return nil, errors.New("outdata population requires an instance driver passed to populate helpers")
	}
	instanceVersion, ok := common.OfType[common.APIVersion](helpers)
	if !ok {
		return nil, errors.New("outdata population requires an instance version passed to populate helpers")
	}

	count := int(createInfo.physicalDeviceCount)
	o.PhysicalDevices = make([]core1_0.PhysicalDevice, count)

	propertiesUnsafe := arena.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

	for i := 0; i < count; i++ {
		handle := driver.VkPhysicalDevice(unsafe.Pointer(createInfo.physicalDevices[i]))
		instanceDriver.VkGetPhysicalDeviceProperties(handle, (*driver.VkPhysicalDeviceProperties)(propertiesUnsafe))

		var properties core1_0.PhysicalDeviceProperties
		err = (&properties).PopulateFromCPointer(propertiesUnsafe)
		if err != nil {
			return nil, err
		}

		version := instanceVersion.Min(properties.APIVersion)

		o.PhysicalDevices[i] = createPhysicalDevice(instanceDriver, instanceHandle, handle, version)
	}

	return createInfo.pNext, nil
}
