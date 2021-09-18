package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

func CreateInstance(load Driver, options *InstanceOptions) (Instance, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	var instanceHandle VkInstance

	res, err := load.VkCreateInstance((*VkInstanceCreateInfo)(createInfo), nil, &instanceHandle)
	if err != nil {
		return nil, res, err
	}

	instanceDriver, err := load.CreateInstanceDriver((VkInstance)(instanceHandle))
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	return &vulkanInstance{
		driver: instanceDriver,
		handle: instanceHandle,
	}, res, nil
}

type vulkanInstance struct {
	driver Driver
	handle VkInstance
}

func (i *vulkanInstance) Driver() Driver {
	return i.driver
}

func (i *vulkanInstance) Handle() VkInstance {
	return i.handle
}

func (i *vulkanInstance) Destroy() error {
	err := i.driver.VkDestroyInstance(i.handle, nil)
	if err != nil {
		return err
	}

	i.driver.Destroy()
	return nil
}

func (i *vulkanInstance) PhysicalDevices() ([]PhysicalDevice, VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*Uint32)(allocator.Malloc(int(unsafe.Sizeof(Uint32(0)))))

	res, err := i.driver.VkEnumeratePhysicalDevices(i.handle, count, nil)
	if err != nil {
		return nil, res, err
	}

	if *count == 0 {
		return nil, res, nil
	}

	allocatedHandles := allocator.Malloc(int(uintptr(*count) * unsafe.Sizeof([1]VkPhysicalDevice{})))

	deviceHandles := ([]VkPhysicalDevice)(unsafe.Slice((*VkPhysicalDevice)(allocatedHandles), int(*count)))
	res, err = i.driver.VkEnumeratePhysicalDevices(i.handle, count, (*VkPhysicalDevice)(allocatedHandles))
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
