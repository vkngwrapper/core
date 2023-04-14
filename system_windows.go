//go:build windows

package core

import "C"
import (
	"syscall"
	"unsafe"
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
func CreateSystemLoader() (*VulkanLoader, error) {
	if getInstanceProcAddr == nil {
		err := loadProcAddr()
		if err != nil {
			return nil, err
		}
	}
	return CreateLoaderFromProcAddr(getInstanceProcAddr)
}
