package internal1_0

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
	DeviceDriver driver.Driver
	Device       driver.VkDevice
	FenceHandle  driver.VkFence

	MaximumAPIVersion common.APIVersion
}

func (f *VulkanFence) Handle() driver.VkFence {
	return f.FenceHandle
}

func (f *VulkanFence) DeviceHandle() driver.VkDevice {
	return f.Device
}

func (f *VulkanFence) Driver() driver.Driver {
	return f.DeviceDriver
}

func (f *VulkanFence) APIVersion() common.APIVersion {
	return f.MaximumAPIVersion
}

func (f *VulkanFence) Destroy(callbacks *driver.AllocationCallbacks) {
	f.DeviceDriver.VkDestroyFence(f.Device, f.FenceHandle, callbacks.Handle())
	f.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(f.FenceHandle))
}

func (f *VulkanFence) Wait(timeout time.Duration) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.FenceHandle

	return f.DeviceDriver.VkWaitForFences(f.Device, driver.Uint32(1), fencePtr, driver.VkBool32(C.VK_TRUE), driver.Uint64(common.TimeoutNanoseconds(timeout)))
}

func (f *VulkanFence) Reset() (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.FenceHandle

	return f.DeviceDriver.VkResetFences(f.Device, driver.Uint32(1), fencePtr)
}

func (f *VulkanFence) Status() (common.VkResult, error) {
	return f.DeviceDriver.VkGetFenceStatus(f.Device, f.FenceHandle)
}
