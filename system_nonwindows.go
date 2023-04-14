//go:build !windows

package core

/*
#include <stdlib.h>
#include "common/vulkan.h"
*/
import "C"
import (
	"unsafe"
)

// CreateSystemLoader generates a Loader from a vulkan-1.dll/so located on the local file system
func CreateSystemLoader() (*VulkanLoader, error) {
	return CreateLoaderFromProcAddr(unsafe.Pointer(C.vkGetInstanceProcAddr))
}
