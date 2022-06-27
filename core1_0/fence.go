package core1_0

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
	deviceDriver driver.Driver
	device       driver.VkDevice
	fenceHandle  driver.VkFence

	maximumAPIVersion common.APIVersion
}

func (f *VulkanFence) Handle() driver.VkFence {
	return f.fenceHandle
}

func (f *VulkanFence) DeviceHandle() driver.VkDevice {
	return f.device
}

func (f *VulkanFence) Driver() driver.Driver {
	return f.deviceDriver
}

func (f *VulkanFence) APIVersion() common.APIVersion {
	return f.maximumAPIVersion
}

func (f *VulkanFence) Destroy(callbacks *driver.AllocationCallbacks) {
	f.deviceDriver.VkDestroyFence(f.device, f.fenceHandle, callbacks.Handle())
	f.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(f.fenceHandle))
}

func (f *VulkanFence) Wait(timeout time.Duration) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.fenceHandle

	return f.deviceDriver.VkWaitForFences(f.device, driver.Uint32(1), fencePtr, driver.VkBool32(C.VK_TRUE), driver.Uint64(common.TimeoutNanoseconds(timeout)))
}

func (f *VulkanFence) Reset() (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceUnsafePtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, 1))
	fenceSlice[0] = f.fenceHandle

	return f.deviceDriver.VkResetFences(f.device, driver.Uint32(1), fencePtr)
}

func (f *VulkanFence) Status() (common.VkResult, error) {
	return f.deviceDriver.VkGetFenceStatus(f.device, f.fenceHandle)
}
