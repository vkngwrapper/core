//go:build !windows

package bootstrap

/*
#include <stdlib.h>
#include "common/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/core1_0"
)

// CreateSystemLoader generates a Loader from a vulkan-1.dll/so located on the local file system
func CreateSystemDriver() (core1_0.GlobalDriver, error) {
	return CreateDriverFromProcAddr(unsafe.Pointer(C.vkGetInstanceProcAddr))
}
