package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"unsafe"
)

//go:generate mockgen -source ./iface.go -destination mocks/loader_mocks.go -package mocks

type Loader interface {
	Driver() driver.Driver
	APIVersion() common.APIVersion

	AvailableExtensions() (map[string]*core1_0.ExtensionProperties, common.VkResult, error)
	AvailableExtensionsForLayer(layerName string) (map[string]*core1_0.ExtensionProperties, common.VkResult, error)
	AvailableLayers() (map[string]*core1_0.LayerProperties, common.VkResult, error)

	CreateInstance(allocationCallbacks *driver.AllocationCallbacks, options core1_0.InstanceCreateInfo) (core1_0.Instance, common.VkResult, error)
}

func CreateStaticLinkedLoader() (*VulkanLoader, error) {
	return CreateLoaderFromProcAddr(unsafe.Pointer(C.vkGetInstanceProcAddr))
}

func CreateLoaderFromProcAddr(addr unsafe.Pointer) (*VulkanLoader, error) {
	driver, err := driver.CreateDriverFromProcAddr(addr)
	if err != nil {
		return nil, err
	}

	return CreateLoaderFromDriver(driver)
}

func CreateLoaderFromDriver(driver driver.Driver) (*VulkanLoader, error) {
	return NewLoader(driver), nil
}
