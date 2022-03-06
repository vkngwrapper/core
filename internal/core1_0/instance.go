package core1_0

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
	"unsafe"
)

type VulkanInstance struct {
	InstanceDriver driver.Driver
	InstanceHandle driver.VkInstance
	MaximumVersion common.APIVersion
}

func (i *VulkanInstance) Driver() driver.Driver {
	return i.InstanceDriver
}

func (i *VulkanInstance) Handle() driver.VkInstance {
	return i.InstanceHandle
}

func (i *VulkanInstance) Destroy(callbacks *driver.AllocationCallbacks) {
	i.InstanceDriver.VkDestroyInstance(i.InstanceHandle, callbacks.Handle())
	i.InstanceDriver.Destroy()
}

func (i *VulkanInstance) PhysicalDevices() ([]core1_0.PhysicalDevice, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*driver.Uint32)(allocator.Malloc(int(unsafe.Sizeof(driver.Uint32(0)))))

	res, err := i.InstanceDriver.VkEnumeratePhysicalDevices(i.InstanceHandle, count, nil)
	if err != nil {
		return nil, res, err
	}

	if *count == 0 {
		return nil, res, nil
	}

	allocatedHandles := allocator.Malloc(int(uintptr(*count) * unsafe.Sizeof([1]driver.VkPhysicalDevice{})))

	deviceHandles := ([]driver.VkPhysicalDevice)(unsafe.Slice((*driver.VkPhysicalDevice)(allocatedHandles), int(*count)))
	res, err = i.InstanceDriver.VkEnumeratePhysicalDevices(i.InstanceHandle, count, (*driver.VkPhysicalDevice)(allocatedHandles))
	if err != nil {
		return nil, res, err
	}

	goCount := uint32(*count)
	var devices []core1_0.PhysicalDevice
	for ind := uint32(0); ind < goCount; ind++ {
		propertiesUnsafe := allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

		i.InstanceDriver.VkGetPhysicalDeviceProperties(deviceHandles[ind], (*driver.VkPhysicalDeviceProperties)(propertiesUnsafe))

		properties := createPhysicalDeviceProperties((*C.VkPhysicalDeviceProperties)(propertiesUnsafe))

		devices = append(devices, &VulkanPhysicalDevice{
			InstanceDriver:       i.InstanceDriver,
			PhysicalDeviceHandle: deviceHandles[ind],
			DeviceProperties:     properties,
			MaximumVersion:       i.MaximumVersion.Min(properties.APIVersion),
		})
	}

	return devices, res, nil
}
