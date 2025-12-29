package impl1_0

import (
	"fmt"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

func (v *DeviceVulkanDriver) DestroyEvent(event core1_0.Event, callbacks *loader.AllocationCallbacks) {
	if !event.Initialized() {
		panic("event was uninitialized")
	}

	v.LoaderObj.VkDestroyEvent(event.DeviceHandle(), event.Handle(), callbacks.Handle())
}

func (v *DeviceVulkanDriver) SetEvent(event core1_0.Event) (common.VkResult, error) {
	if !event.Initialized() {
		return core1_0.VKErrorUnknown, fmt.Errorf("event was uninitialized")
	}

	return v.LoaderObj.VkSetEvent(event.DeviceHandle(), event.Handle())
}

func (v *DeviceVulkanDriver) ResetEvent(event core1_0.Event) (common.VkResult, error) {
	if !event.Initialized() {
		return core1_0.VKErrorUnknown, fmt.Errorf("event was uninitialized")
	}

	return v.LoaderObj.VkResetEvent(event.DeviceHandle(), event.Handle())
}

func (v *DeviceVulkanDriver) GetEventStatus(event core1_0.Event) (common.VkResult, error) {
	if !event.Initialized() {
		return core1_0.VKErrorUnknown, fmt.Errorf("event was uninitialized")
	}

	return v.LoaderObj.VkGetEventStatus(event.DeviceHandle(), event.Handle())
}
