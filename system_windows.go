//go:build windows

package core

import "C"
import (
	"syscall"
	"unsafe"
)

// CreateSystemLoader generates a Loader from a vulkan-1.dll/so located on the local file system
//
// Allowing cgo to bring us the vkGetInstanceProcAddr method on windows, for whatever reason, causes heap corruption
// when the garbage collector runs. For whatever reason, manually loading it from dll does not have this issue
func CreateSystemLoader() (*VulkanLoader, error) {
	vulkan, err := syscall.LoadLibrary("vulkan-1.dll")
	if err != nil {
		return nil, err
	}

	getInstanceProcAddr, err := syscall.GetProcAddress(vulkan, "vkGetInstanceProcAddr")
	if err != nil {
		return nil, err
	}

	return CreateLoaderFromProcAddr(unsafe.Pointer(getInstanceProcAddr))
}
