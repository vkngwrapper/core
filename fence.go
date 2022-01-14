package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type vulkanFence struct {
	driver Driver
	device VkDevice
	handle VkFence
}

func (f *vulkanFence) Handle() VkFence {
	return f.handle
}

func (f *vulkanFence) Destroy(callbacks *AllocationCallbacks) {
	f.driver.VkDestroyFence(f.device, f.handle, callbacks.Handle())
}

func (f *vulkanFence) Wait(timeout time.Duration) (VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*VkFence)(fenceUnsafePtr)

	fenceSlice := ([]VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.handle

	return f.driver.VkWaitForFences(f.device, Uint32(1), fencePtr, VkBool32(C.VK_TRUE), Uint64(common.TimeoutNanoseconds(timeout)))
}

func (f *vulkanFence) Reset() (VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*VkFence)(fenceUnsafePtr)
	fenceSlice := ([]VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.handle

	return f.driver.VkResetFences(f.device, Uint32(1), fencePtr)
}

func (f *vulkanFence) Status() (VkResult, error) {
	return f.driver.VkGetFenceStatus(f.device, f.handle)
}
