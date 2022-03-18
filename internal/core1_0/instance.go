package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	internal1_1 "github.com/CannibalVox/VKng/core/internal/core1_1"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanInstance struct {
	InstanceDriver driver.Driver
	InstanceHandle driver.VkInstance
	MaximumVersion common.APIVersion

	Instance1_1 core1_1.Instance
}

func (i *VulkanInstance) Driver() driver.Driver {
	return i.InstanceDriver
}

func (i *VulkanInstance) Handle() driver.VkInstance {
	return i.InstanceHandle
}

func (i *VulkanInstance) APIVersion() common.APIVersion {
	return i.MaximumVersion
}

func (i *VulkanInstance) Core1_1() core1_1.Instance {
	return i.Instance1_1
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

		version := i.MaximumVersion.Min(properties.APIVersion)
		physicalDevice := &VulkanPhysicalDevice{
			InstanceDriver:       i.InstanceDriver,
			PhysicalDeviceHandle: deviceHandles[ind],
			MaximumVersion:       version,
		}

		if version.IsAtLeast(common.Vulkan1_1) {
			physicalDevice.PhysicalDevice1_1 = &internal1_1.VulkanPhysicalDevice{
				InstanceDriver:       i.InstanceDriver,
				PhysicalDeviceHandle: deviceHandles[ind],
			}
		}

		devices = append(devices, physicalDevice)
	}

	return devices, res, nil
}
