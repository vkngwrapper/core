package resource

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

func CreateInstance(allocator cgoalloc.Allocator, options *InstanceOptions) (*Instance, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	var instanceHandle C.VkInstance

	res := core.Result(C.vkCreateInstance((*C.VkInstanceCreateInfo)(createInfo), nil, &instanceHandle))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Instance{
		handle: instanceHandle,
	}, res, nil
}

type InstanceHandle C.VkInstance
type Instance struct {
	handle C.VkInstance
}

func (i *Instance) Handle() C.VkInstance {
	return i.handle
}

func (i *Instance) Destroy() {
	C.vkDestroyInstance(i.handle, nil)
}

func (i *Instance) PhysicalDevices(allocator cgoalloc.Allocator) ([]*PhysicalDevice, core.Result, error) {
	count := (*C.uint32_t)(allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	defer allocator.Free(unsafe.Pointer(count))

	res := core.Result(C.vkEnumeratePhysicalDevices(i.handle, count, nil))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	if *count == 0 {
		return nil, res, nil
	}

	allocatedHandles := allocator.Malloc(int(uintptr(*count) * unsafe.Sizeof([1]C.VkPhysicalDevice{})))
	defer allocator.Free(allocatedHandles)

	deviceHandles := ([]C.VkPhysicalDevice)(unsafe.Slice((*C.VkPhysicalDevice)(allocatedHandles), int(*count)))
	res = core.Result(C.vkEnumeratePhysicalDevices(i.handle, count, (*C.VkPhysicalDevice)(allocatedHandles)))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	goCount := uint32(*count)
	var devices []*PhysicalDevice
	for i := uint32(0); i < goCount; i++ {
		devices = append(devices, &PhysicalDevice{handle: deviceHandles[i]})
	}

	return devices, res, nil
}
