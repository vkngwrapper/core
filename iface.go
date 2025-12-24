package core

/*
#include <stdlib.h>
#include "common/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate mockgen -source ./iface.go -destination mocks/loader_mocks.go -package mocks

// Loader is the root object of vkng - all usage begins by creating a Loader and then using
// the Loader to create a Vulkan instance with loader.CreateInstance
type Loader interface {
	// Driver is the Vulkan wrapper loader used by this Loader
	Driver() loader.Loader
	// APIVersion is the maximum Vulkan API version supported by this Loader.
	APIVersion() common.APIVersion
}

//// CreateLoaderFromProcAddr generates a Loader from a ProcAddr provided by another library,
//// such as the address provided by sdl2's SDL_Vulkan_GetVkInstanceProcAddr
//func CreateLoaderFromProcAddr(addr unsafe.Pointer) (*VulkanLoader, error) {
//	driver, err := loader.CreateLoaderFromProcAddr(addr)
//	if err != nil {
//		return nil, err
//	}
//
//	return CreateLoaderFromDriver(driver)
//}
//
//// CreateLoaderFromDriver generates a Loader from a loader.Loader object- this is usually
//// used in tests to build a Loader from mock drivers
//func CreateLoaderFromDriver(driver loader.Loader) (*VulkanLoader, error) {
//	return NewLoader(driver), nil
//}
