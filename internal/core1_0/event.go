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
)

type VulkanEvent struct {
	EventHandle driver.VkEvent
	Device      driver.VkDevice
	Driver      driver.Driver

	MaximumAPIVersion common.APIVersion

	Event1_1 core1_1.Event
}

func (e *VulkanEvent) Handle() driver.VkEvent {
	return e.EventHandle
}

func (e *VulkanEvent) Core1_1() core1_1.Event {
	return e.Event1_1
}

func (e *VulkanEvent) Destroy(callbacks *driver.AllocationCallbacks) {
	e.Driver.VkDestroyEvent(e.Device, e.EventHandle, callbacks.Handle())
	e.Driver.ObjectStore().Delete(driver.VulkanHandle(e.EventHandle), e)
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
