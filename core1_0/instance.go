package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanInstance struct {
	instanceDriver driver.Driver
	instanceHandle driver.VkInstance
	maximumVersion common.APIVersion

	ActiveInstanceExtensions map[string]struct{}
}

func (i *VulkanInstance) Driver() driver.Driver {
	return i.instanceDriver
}

func (i *VulkanInstance) Handle() driver.VkInstance {
	return i.instanceHandle
}

func (i *VulkanInstance) APIVersion() common.APIVersion {
	return i.maximumVersion
}

func (i *VulkanInstance) Destroy(callbacks *driver.AllocationCallbacks) {
	i.instanceDriver.VkDestroyInstance(i.instanceHandle, callbacks.Handle())
	i.instanceDriver.ObjectStore().Delete(driver.VulkanHandle(i.instanceHandle))
	i.instanceDriver.Destroy()
}

func (i *VulkanInstance) IsInstanceExtensionActive(extensionName string) bool {
	_, active := i.ActiveInstanceExtensions[extensionName]
	return active
}

func (i *VulkanInstance) PhysicalDevices() ([]PhysicalDevice, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*driver.Uint32)(allocator.Malloc(int(unsafe.Sizeof(driver.Uint32(0)))))

	res, err := i.instanceDriver.VkEnumeratePhysicalDevices(i.instanceHandle, count, nil)
	if err != nil {
		return nil, res, err
	}

	if *count == 0 {
		return nil, res, nil
	}

	allocatedHandles := allocator.Malloc(int(uintptr(*count) * unsafe.Sizeof([1]driver.VkPhysicalDevice{})))

	deviceHandles := ([]driver.VkPhysicalDevice)(unsafe.Slice((*driver.VkPhysicalDevice)(allocatedHandles), int(*count)))
	res, err = i.instanceDriver.VkEnumeratePhysicalDevices(i.instanceHandle, count, (*driver.VkPhysicalDevice)(allocatedHandles))
	if err != nil {
		return nil, res, err
	}

	goCount := uint32(*count)
	var devices []PhysicalDevice
	for ind := uint32(0); ind < goCount; ind++ {
		propertiesUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

		i.instanceDriver.VkGetPhysicalDeviceProperties(deviceHandles[ind], (*driver.VkPhysicalDeviceProperties)(propertiesUnsafe))

		var properties PhysicalDeviceProperties
		err = (&properties).PopulateFromCPointer(propertiesUnsafe)
		if err != nil {
			return nil, VKErrorUnknown, err
		}

		deviceVersion := i.maximumVersion.Min(properties.APIVersion)
		physicalDevice := createPhysicalDeviceObject(i.instanceDriver, i.instanceHandle, deviceHandles[ind], i.maximumVersion, deviceVersion)

		devices = append(devices, physicalDevice)
	}

	return devices, res, nil
}
