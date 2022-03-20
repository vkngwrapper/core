package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type VulkanFence struct {
	Driver      driver.Driver
	Device      driver.VkDevice
	FenceHandle driver.VkFence

	MaximumAPIVersion common.APIVersion

	Fence1_1 core1_1.Fence
}

func (f *VulkanFence) Handle() driver.VkFence {
	return f.FenceHandle
}

func (f *VulkanFence) Core1_1() core1_1.Fence {
	return f.Fence1_1
}

func (f *VulkanFence) Destroy(callbacks *driver.AllocationCallbacks) {
	f.Driver.VkDestroyFence(f.Device, f.FenceHandle, callbacks.Handle())
	f.Driver.ObjectStore().Delete(driver.VulkanHandle(f.FenceHandle), f)
}

func (f *VulkanFence) Wait(timeout time.Duration) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.FenceHandle

	return f.Driver.VkWaitForFences(f.Device, driver.Uint32(1), fencePtr, driver.VkBool32(C.VK_TRUE), driver.Uint64(common.TimeoutNanoseconds(timeout)))
}

func (f *VulkanFence) Reset() (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.FenceHandle

	return f.Driver.VkResetFences(f.Device, driver.Uint32(1), fencePtr)
}

func (f *VulkanFence) Status() (common.VkResult, error) {
	return f.Driver.VkGetFenceStatus(f.Device, f.FenceHandle)
}
