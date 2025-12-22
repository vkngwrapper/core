package core

/*
#include <stdlib.h>
#include "common/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

//go:generate mockgen -source ./iface.go -destination mocks/loader_mocks.go -package mocks

// Loader is the root object of vkng - all usage begins by creating a Loader and then using
// the Loader to create a Vulkan instance with Loader.CreateInstance
type Loader interface {
	// Driver is the Vulkan wrapper driver used by this Loader
	Driver() driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Loader.
	APIVersion() common.APIVersion
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
