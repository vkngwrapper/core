package internal1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanEvent struct {
	EventHandle  driver.VkEvent
	Device       driver.VkDevice
	DeviceDriver driver.Driver

	MaximumAPIVersion common.APIVersion
}

func (e *VulkanEvent) Handle() driver.VkEvent {
	return e.EventHandle
}

func (e *VulkanEvent) Driver() driver.Driver {
	return e.DeviceDriver
}

func (e *VulkanEvent) APIVersion() common.APIVersion {
	return e.MaximumAPIVersion
}

func (e *VulkanEvent) Destroy(callbacks *driver.AllocationCallbacks) {
	e.DeviceDriver.VkDestroyEvent(e.Device, e.EventHandle, callbacks.Handle())
	e.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(e.EventHandle))
}

func (e *VulkanEvent) Set() (common.VkResult, error) {
	return e.DeviceDriver.VkSetEvent(e.Device, e.EventHandle)
}

func (e *VulkanEvent) Reset() (common.VkResult, error) {
	return e.DeviceDriver.VkResetEvent(e.Device, e.EventHandle)
}

func (e *VulkanEvent) Status() (common.VkResult, error) {
	return e.DeviceDriver.VkGetEventStatus(e.Device, e.EventHandle)
}
