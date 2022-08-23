package driver

import "C"

/*
#include "func_ptrs.h"
#include "func_ptrs_def.h"

PFN_vkVoidFunction instance_proc_addr(DriverFuncPtrs *funcPtrs, VkInstance instance, const char *procName) {
	PFN_vkGetInstanceProcAddr procAddr = funcPtrs->vkGetInstanceProcAddr;
	return procAddr(instance, procName);
}

PFN_vkVoidFunction device_proc_addr(DriverFuncPtrs *funcPtrs, VkDevice device, const char *procName) {
	PFN_vkGetDeviceProcAddr procAddr = funcPtrs->vkGetDeviceProcAddr;
	return procAddr(device, procName);
}
*/
import "C"
import (
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

type vulkanDriver struct {
	instance VkInstance
	device   VkDevice
	funcPtrs *C.DriverFuncPtrs

	version common.APIVersion

	objStore *VulkanObjectStore
}

func createVulkanDriver(funcPtrs *C.DriverFuncPtrs, objStore *VulkanObjectStore, instance VkInstance, device VkDevice) (*vulkanDriver, error) {
	version := common.Vulkan1_0
	driver := &vulkanDriver{
		funcPtrs: funcPtrs,
		instance: instance,
		device:   device,

		objStore: objStore,
	}

	if funcPtrs.vkEnumerateInstanceVersion != nil {
		var versionBits Uint32
		_, err := driver.VkEnumerateInstanceVersion(&versionBits)
		if err != nil {
			return nil, err
		}

		version = common.APIVersion(versionBits)
	}

	driver.version = version
	return driver, nil
}

func CreateDriverFromProcAddr(procAddr unsafe.Pointer) (*vulkanDriver, error) {
	baseFuncPtr := (C.PFN_vkGetInstanceProcAddr)(procAddr)
	funcPtrs := (*C.DriverFuncPtrs)(C.malloc(C.sizeof_struct_DriverFuncPtrs))
	C.driverFuncPtrs_populate(baseFuncPtr, funcPtrs)

	return createVulkanDriver(funcPtrs, NewObjectStore(), VkInstance(NullHandle), VkDevice(NullHandle))
}

func (l *vulkanDriver) ObjectStore() *VulkanObjectStore {
	return l.objStore
}

func (l *vulkanDriver) Destroy() {
	C.free(unsafe.Pointer(l.funcPtrs))
}

func (l *vulkanDriver) CreateInstanceDriver(instance VkInstance) (Driver, error) {
	instanceFuncPtrs := (*C.DriverFuncPtrs)(C.malloc(C.sizeof_struct_DriverFuncPtrs))
	C.instanceFuncPtrs_populate((C.VkInstance)(unsafe.Pointer(instance)), l.funcPtrs, instanceFuncPtrs)

	return createVulkanDriver(instanceFuncPtrs, l.objStore, instance, VkDevice(NullHandle))
}

func (l *vulkanDriver) CreateDeviceDriver(device VkDevice) (Driver, error) {
	if l.instance == VkInstance(NullHandle) {
		return nil, errors.New("attempted to call instance driver function on a basic driver")
	}

	deviceFuncPtrs := (*C.DriverFuncPtrs)(C.malloc(C.sizeof_struct_DriverFuncPtrs))
	C.deviceFuncPtrs_populate((C.VkDevice)(unsafe.Pointer(device)), l.funcPtrs, deviceFuncPtrs)

	return createVulkanDriver(deviceFuncPtrs, l.objStore, l.instance, device)
}

func (l *vulkanDriver) LoadProcAddr(name *Char) unsafe.Pointer {
	if l.device != VkDevice(NullHandle) {
		return unsafe.Pointer(C.device_proc_addr(l.funcPtrs, C.VkDevice(unsafe.Pointer(l.device)), (*C.char)(name)))
	} else {
		return unsafe.Pointer(C.instance_proc_addr(l.funcPtrs, C.VkInstance(unsafe.Pointer(l.instance)), (*C.char)(name)))
	}
}

func (l *vulkanDriver) Version() common.APIVersion {
	return l.version
}
