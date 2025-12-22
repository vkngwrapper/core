package impl1_0

/*
#include <stdlib.h>
#include "../../common/vulkan.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyInstance(instance types.Instance, callbacks *driver.AllocationCallbacks) {
	if instance.Handle() == 0 {
		panic("instance was uninitialized")
	}
	v.Driver.VkDestroyInstance(instance.Handle(), callbacks.Handle())
}

func (v *Vulkan) EnumeratePhysicalDevices(instance types.Instance) ([]types.PhysicalDevice, common.VkResult, error) {
	if instance.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, fmt.Errorf("instance was uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*driver.Uint32)(allocator.Malloc(int(unsafe.Sizeof(driver.Uint32(0)))))

	res, err := v.Driver.VkEnumeratePhysicalDevices(instance.Handle(), count, nil)
	if err != nil {
		return nil, res, err
	}

	if *count == 0 {
		return nil, res, nil
	}

	allocatedHandles := allocator.Malloc(int(uintptr(*count) * unsafe.Sizeof([1]driver.VkPhysicalDevice{})))

	deviceHandles := ([]driver.VkPhysicalDevice)(unsafe.Slice((*driver.VkPhysicalDevice)(allocatedHandles), int(*count)))
	res, err = v.Driver.VkEnumeratePhysicalDevices(instance.Handle(), count, (*driver.VkPhysicalDevice)(allocatedHandles))
	if err != nil {
		return nil, res, err
	}

	goCount := uint32(*count)
	var devices []types.PhysicalDevice
	for ind := uint32(0); ind < goCount; ind++ {
		propertiesUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

		v.Driver.VkGetPhysicalDeviceProperties(deviceHandles[ind], (*driver.VkPhysicalDeviceProperties)(propertiesUnsafe))

		var properties core1_0.PhysicalDeviceProperties
		err = (&properties).PopulateFromCPointer(propertiesUnsafe)
		if err != nil {
			return nil, core1_0.VKErrorUnknown, err
		}

		deviceVersion := instance.APIVersion().Min(properties.APIVersion)
		physicalDevice := types.InternalPhysicalDevice(deviceHandles[ind], instance.APIVersion(), deviceVersion)

		devices = append(devices, physicalDevice)
	}

	return devices, res, nil
}
