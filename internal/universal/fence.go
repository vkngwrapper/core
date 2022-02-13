package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type VulkanFence struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkFence
}

func (f *VulkanFence) Handle() driver.VkFence {
	return f.handle
}

func (f *VulkanFence) Destroy(callbacks *driver.AllocationCallbacks) {
	f.driver.VkDestroyFence(f.device, f.handle, callbacks.Handle())
}

func (f *VulkanFence) Wait(timeout time.Duration) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.handle

	return f.driver.VkWaitForFences(f.device, driver.Uint32(1), fencePtr, driver.VkBool32(C.VK_TRUE), driver.Uint64(common.TimeoutNanoseconds(timeout)))
}

func (f *VulkanFence) Reset() (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.handle

	return f.driver.VkResetFences(f.device, driver.Uint32(1), fencePtr)
}

func (f *VulkanFence) Status() (common.VkResult, error) {
	return f.driver.VkGetFenceStatus(f.device, f.handle)
}
