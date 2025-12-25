//go:build !windows

package core

/*
#include <stdlib.h>
#include "common/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v3/loader"

// CreateSystemLoader generates a Loader from a vulkan-1.dll/so located on the local file system
func CreateSystemLoader() (loader.Loader, error) {
	return loader.CreateLoaderFromProcAddr(unsafe.Pointer(C.vkGetInstanceProcAddr))
}
