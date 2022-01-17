package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	driver3 "github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type vulkanInstance struct {
	driver driver3.Driver
	handle driver3.VkInstance
}

func (i *vulkanInstance) Driver() driver3.Driver {
	return i.driver
}

func (i *vulkanInstance) Handle() driver3.VkInstance {
	return i.handle
}

func (i *vulkanInstance) Destroy(callbacks *AllocationCallbacks) {
	i.driver.VkDestroyInstance(i.handle, callbacks.Handle())
	i.driver.Destroy()
}

func (i *vulkanInstance) PhysicalDevices() ([]PhysicalDevice, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*driver3.Uint32)(allocator.Malloc(int(unsafe.Sizeof(driver3.Uint32(0)))))

	res, err := i.driver.VkEnumeratePhysicalDevices(i.handle, count, nil)
	if err != nil {
		return nil, res, err
	}

	if *count == 0 {
		return nil, res, nil
	}

	allocatedHandles := allocator.Malloc(int(uintptr(*count) * unsafe.Sizeof([1]driver3.VkPhysicalDevice{})))

	deviceHandles := ([]driver3.VkPhysicalDevice)(unsafe.Slice((*driver3.VkPhysicalDevice)(allocatedHandles), int(*count)))
	res, err = i.driver.VkEnumeratePhysicalDevices(i.handle, count, (*driver3.VkPhysicalDevice)(allocatedHandles))
	if err != nil {
		return nil, res, err
	}

	goCount := uint32(*count)
	var devices []PhysicalDevice
	for ind := uint32(0); ind < goCount; ind++ {
		devices = append(devices, &vulkanPhysicalDevice{driver: i.driver, handle: deviceHandles[ind]})
	}

	return devices, res, nil
}
