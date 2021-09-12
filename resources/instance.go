package resources

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

func CreateInstance(load loader.Loader, options *InstanceOptions) (Instance, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, options)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var instanceHandle loader.VkInstance

	res, err := load.VkCreateInstance((*loader.VkInstanceCreateInfo)(createInfo), nil, &instanceHandle)
	if err != nil {
		return nil, res, err
	}

	instanceLoader, err := load.CreateInstanceLoader((loader.VkInstance)(instanceHandle))
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	return &vulkanInstance{
		loader: instanceLoader,
		handle: instanceHandle,
	}, res, nil
}

type vulkanInstance struct {
	loader loader.Loader
	handle loader.VkInstance
}

func (i *vulkanInstance) Loader() loader.Loader {
	return i.loader
}

func (i *vulkanInstance) Handle() loader.VkInstance {
	return i.handle
}

func (i *vulkanInstance) Destroy() error {
	err := i.loader.VkDestroyInstance(i.handle, nil)
	if err != nil {
		return err
	}

	i.loader.Destroy()
	return nil
}

func (i *vulkanInstance) PhysicalDevices() ([]PhysicalDevice, loader.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	count := (*loader.Uint32)(allocator.Malloc(int(unsafe.Sizeof(loader.Uint32(0)))))

	res, err := i.loader.VkEnumeratePhysicalDevices(i.handle, count, nil)
	if err != nil {
		return nil, res, err
	}

	if *count == 0 {
		return nil, res, nil
	}

	allocatedHandles := allocator.Malloc(int(uintptr(*count) * unsafe.Sizeof([1]loader.VkPhysicalDevice{})))

	deviceHandles := ([]loader.VkPhysicalDevice)(unsafe.Slice((*loader.VkPhysicalDevice)(allocatedHandles), int(*count)))
	res, err = i.loader.VkEnumeratePhysicalDevices(i.handle, count, (*loader.VkPhysicalDevice)(allocatedHandles))
	if err != nil {
		return nil, res, err
	}

	goCount := uint32(*count)
	var devices []PhysicalDevice
	for ind := uint32(0); ind < goCount; ind++ {
		devices = append(devices, &vulkanPhysicalDevice{loader: i.loader, handle: deviceHandles[ind]})
	}

	return devices, res, nil
}
