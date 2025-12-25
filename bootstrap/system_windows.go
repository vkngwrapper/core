//go:build windows

package bootstrap

import "C"
import (
	"syscall"
	"unsafe"

	"github.com/vkngwrapper/core/v3/core1_0"
)

var vulkanDLL syscall.Handle
var getInstanceProcAddr unsafe.Pointer

func loadProcAddr() error {
	var err error
	vulkanDLL, err = syscall.LoadLibrary("vulkan-1.dll")
	if err != nil {
		return err
	}

	getInstanceProcAddrHandle, err := syscall.GetProcAddress(vulkanDLL, "vkGetInstanceProcAddr")
	if err != nil {
		return err
	}
	getInstanceProcAddr = unsafe.Pointer(getInstanceProcAddrHandle)

	return nil
}

// CreateSystemLoader generates a Loader from a vulkan-1.dll/so located on the local file system
//
// Allowing cgo to bring us the vkGetInstanceProcAddr method on windows, for whatever reason, causes heap corruption
// when the garbage collector runs. For whatever reason, manually loading it from dll does not have this issue
func CreateSystemDriver() (core1_0.GlobalDriver, error) {
	if getInstanceProcAddr == nil {
		err := loadProcAddr()
		if err != nil {
			return nil, err
		}
	}
	return CreateDriverFromProcAddr(getInstanceProcAddr)
}
