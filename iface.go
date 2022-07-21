package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	"unsafe"
)

//go:generate mockgen -source ./iface.go -destination mocks/loader_mocks.go -package mocks

// Loader is the root object of vkng - all usage begins by creating a Loader and then using
// the Loader to create a Vulkan instance with Loader.CreateInstance
type Loader interface {
	// Driver is the Vulkan wrapper driver used by this Loader
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Loader.
	APIVersion() common.APIVersion

	// AvailableExtensions returns all of the instance extensions available on this Loader,
	// in the form of a map of extension name to ExtensionProperties
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceExtensionProperties.html
	AvailableExtensions() (map[string]*core1_0.ExtensionProperties, common.VkResult, error)
	// AvailableExtensionsForLayer returns all of the layer extensions available on this Loader
	// for the requested layer, in the form of a map of extension name to ExtensionProperties
	//
	// layerName - a string naming the layer to retrieve extensions from
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceExtensionProperties.html
	AvailableExtensionsForLayer(layerName string) (map[string]*core1_0.ExtensionProperties, common.VkResult, error)
	// AvailableLayers returns all of the layers available on this Loader, in the form of a
	// map of layer name to LayerProperties
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceLayerProperties.html
	AvailableLayers() (map[string]*core1_0.LayerProperties, common.VkResult, error)

	// CreateInstance creates a new Vulkan Instance
	//
	// allocationCallbacks - controls host memory allocation
	//
	// options - Controls creation of the Instance
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html
	CreateInstance(allocationCallbacks *driver.AllocationCallbacks, options core1_0.InstanceCreateInfo) (core1_0.Instance, common.VkResult, error)
}

// CreateSystemLoader generates a Loader from a vulkan-1.dll/so located on the local file system
func CreateSystemLoader() (*VulkanLoader, error) {
	return CreateLoaderFromProcAddr(unsafe.Pointer(C.vkGetInstanceProcAddr))
}

// CreateLoaderFromProcAddr generates a Loader from a ProcAddr provided by another library,
// such as the address provided by sdl2's SDL_Vulkan_GetVkInstanceProcAddr
func CreateLoaderFromProcAddr(addr unsafe.Pointer) (*VulkanLoader, error) {
	driver, err := driver.CreateDriverFromProcAddr(addr)
	if err != nil {
		return nil, err
	}

	return CreateLoaderFromDriver(driver)
}

// CreateLoaderFromDriver generates a Loader from a driver.Driver object- this is usually
// used in tests to build a Loader from mock drivers
func CreateLoaderFromDriver(driver driver.Driver) (*VulkanLoader, error) {
	return NewLoader(driver), nil
}
