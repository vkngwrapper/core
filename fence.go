package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	driver3 "github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type vulkanFence struct {
	driver driver3.Driver
	device driver3.VkDevice
	handle driver3.VkFence
}

func (f *vulkanFence) Handle() driver3.VkFence {
	return f.handle
}

func (f *vulkanFence) Destroy(callbacks *AllocationCallbacks) {
	f.driver.VkDestroyFence(f.device, f.handle, callbacks.Handle())
}

func (f *vulkanFence) Wait(timeout time.Duration) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver3.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]driver3.VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.handle

	return f.driver.VkWaitForFences(f.device, driver3.Uint32(1), fencePtr, driver3.VkBool32(C.VK_TRUE), driver3.Uint64(common.TimeoutNanoseconds(timeout)))
}

func (f *vulkanFence) Reset() (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver3.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]driver3.VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.handle

	return f.driver.VkResetFences(f.device, driver3.Uint32(1), fencePtr)
}

func (f *vulkanFence) Status() (common.VkResult, error) {
	return f.driver.VkGetFenceStatus(f.device, f.handle)
}
