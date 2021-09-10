package loader

import "C"

/*
#include "func_ptrs.h"
#include "func_ptrs_def.h"

PFN_vkVoidFunction instance_proc_addr(LoaderFuncPtrs *funcPtrs, VkInstance instance, const char *procName) {
	PFN_vkGetInstanceProcAddr procAddr = funcPtrs->vkGetInstanceProcAddr;
	return procAddr(instance, procName);
}

PFN_vkVoidFunction device_proc_addr(LoaderFuncPtrs *funcPtrs, VkDevice device, const char *procName) {
	PFN_vkGetDeviceProcAddr procAddr = funcPtrs->vkGetDeviceProcAddr;
	return procAddr(device, procName);
}
*/
import "C"
import (
	"github.com/cockroachdb/errors"
	"unsafe"
)

type Loader struct {
	instance VkInstance
	device   VkDevice
	funcPtrs *C.LoaderFuncPtrs
}

func CreateStaticLinkedLoader() (*Loader, error) {
	return CreateLoaderFromProcAddr(unsafe.Pointer(C.vkGetInstanceProcAddr))
}

func CreateLoaderFromProcAddr(procAddr unsafe.Pointer) (*Loader, error) {
	baseFuncPtr := (C.PFN_vkGetInstanceProcAddr)(procAddr)
	funcPtrs := (*C.LoaderFuncPtrs)(C.malloc(C.sizeof_struct_LoaderFuncPtrs))
	C.loaderFuncPtrs_populate(baseFuncPtr, funcPtrs)

	return &Loader{funcPtrs: funcPtrs}, nil
}

func (l *Loader) Destroy() {
	C.free(unsafe.Pointer(l.funcPtrs))
}

func (l *Loader) CreateInstanceLoader(instance VkInstance) (*Loader, error) {
	instanceFuncPtrs := (*C.LoaderFuncPtrs)(C.malloc(C.sizeof_struct_LoaderFuncPtrs))
	C.instanceFuncPtrs_populate((C.VkInstance)(instance), l.funcPtrs, instanceFuncPtrs)

	return &Loader{instance: instance, funcPtrs: instanceFuncPtrs}, nil
}

func (l *Loader) CreateDeviceLoader(device VkDevice) (*Loader, error) {
	if l.instance == nil {
		return nil, errors.New("attempted to call instance loader function on a basic loader")
	}

	deviceFuncPtrs := (*C.LoaderFuncPtrs)(C.malloc(C.sizeof_struct_LoaderFuncPtrs))
	C.deviceFuncPtrs_populate((C.VkDevice)(device), l.funcPtrs, deviceFuncPtrs)

	return &Loader{instance: l.instance, device: device, funcPtrs: deviceFuncPtrs}, nil
}

func (l *Loader) LoadProcAddr(name *Char) unsafe.Pointer {
	if l.device != nil {
		return unsafe.Pointer(C.device_proc_addr(l.funcPtrs, l.device, (*C.char)(name)))
	} else {
		return unsafe.Pointer(C.instance_proc_addr(l.funcPtrs, l.instance, (*C.char)(name)))
	}
}
