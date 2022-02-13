package universal

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
	handle driver.VkEvent
	device driver.VkDevice
	driver driver.Driver
}

func (e *VulkanEvent) Handle() driver.VkEvent {
	return e.handle
}

func (e *VulkanEvent) Destroy(callbacks *driver.AllocationCallbacks) {
	e.driver.VkDestroyEvent(e.device, e.handle, callbacks.Handle())
}

func (e *VulkanEvent) Set() (common.VkResult, error) {
	return e.driver.VkSetEvent(e.device, e.handle)
}

func (e *VulkanEvent) Reset() (common.VkResult, error) {
	return e.driver.VkResetEvent(e.device, e.handle)
}

func (e *VulkanEvent) Status() (common.VkResult, error) {
	return e.driver.VkGetEventStatus(e.device, e.handle)
}
