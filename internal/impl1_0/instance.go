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
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *InstanceVulkanDriver) DestroyInstance(instance types.Instance, callbacks *loader.AllocationCallbacks) {
	if instance.Handle() == 0 {
		panic("instance was uninitialized")
	}
	v.LoaderObj.VkDestroyInstance(instance.Handle(), callbacks.Handle())
}

func (v *InstanceVulkanDriver) EnumeratePhysicalDevices(instance types.Instance) ([]types.PhysicalDevice, common.VkResult, error) {
	if instance.Handle() == 0 {
		return nil, core1_0.VKErrorUnknown, fmt.Errorf("instance was uninitialized")
	}

	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*loader.Uint32)(allocator.Malloc(int(unsafe.Sizeof(loader.Uint32(0)))))

	res, err := v.LoaderObj.VkEnumeratePhysicalDevices(instance.Handle(), count, nil)
	if err != nil {
		return nil, res, err
	}

	if *count == 0 {
		return nil, res, nil
	}

	allocatedHandles := allocator.Malloc(int(uintptr(*count) * unsafe.Sizeof([1]loader.VkPhysicalDevice{})))

	deviceHandles := ([]loader.VkPhysicalDevice)(unsafe.Slice((*loader.VkPhysicalDevice)(allocatedHandles), int(*count)))
	res, err = v.LoaderObj.VkEnumeratePhysicalDevices(instance.Handle(), count, (*loader.VkPhysicalDevice)(allocatedHandles))
	if err != nil {
		return nil, res, err
	}

	goCount := uint32(*count)
	var devices []types.PhysicalDevice
	for ind := uint32(0); ind < goCount; ind++ {
		propertiesUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

		v.LoaderObj.VkGetPhysicalDeviceProperties(deviceHandles[ind], (*loader.VkPhysicalDeviceProperties)(propertiesUnsafe))

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
