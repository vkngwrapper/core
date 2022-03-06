package core1_0

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
	EventHandle driver.VkEvent
	Device      driver.VkDevice
	Driver      driver.Driver

	MaximumAPIVersion common.APIVersion
}

func (e *VulkanEvent) Handle() driver.VkEvent {
	return e.EventHandle
}

func (e *VulkanEvent) Destroy(callbacks *driver.AllocationCallbacks) {
	e.Driver.VkDestroyEvent(e.Device, e.EventHandle, callbacks.Handle())
}

func (e *VulkanEvent) Set() (common.VkResult, error) {
	return e.Driver.VkSetEvent(e.Device, e.EventHandle)
}

func (e *VulkanEvent) Reset() (common.VkResult, error) {
	return e.Driver.VkResetEvent(e.Device, e.EventHandle)
}

func (e *VulkanEvent) Status() (common.VkResult, error) {
	return e.Driver.VkGetEventStatus(e.Device, e.EventHandle)
}
