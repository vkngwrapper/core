package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanEvent is an implementation of the Event interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanEvent struct {
	EventHandle  driver.VkEvent
	Device       driver.VkDevice
	DeviceDriver driver.Driver

	MaximumAPIVersion common.APIVersion
}

func (e *VulkanEvent) Handle() driver.VkEvent {
	return e.EventHandle
}

func (e *VulkanEvent) DeviceHandle() driver.VkDevice {
	return e.Device
}

func (e *VulkanEvent) Driver() driver.Driver {
	return e.DeviceDriver
}

func (e *VulkanEvent) APIVersion() common.APIVersion {
	return e.MaximumAPIVersion
}

func (e *VulkanEvent) Destroy(callbacks *driver.AllocationCallbacks) {
	e.DeviceDriver.VkDestroyEvent(e.Device, e.EventHandle, callbacks.Handle())
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
