package impl1_0

import (
	"fmt"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyEvent(event types.Event, callbacks *driver.AllocationCallbacks) {
	if event.Handle() == 0 {
		panic("event was uninitialized")
	}

	v.Driver.VkDestroyEvent(event.DeviceHandle(), event.Handle(), callbacks.Handle())
}

func (v *Vulkan) Set(event types.Event) (common.VkResult, error) {
	if event.Handle() == 0 {
		return core1_0.VKErrorUnknown, fmt.Errorf("event was uninitialized")
	}

	return v.Driver.VkSetEvent(event.DeviceHandle(), event.Handle())
}

func (v *Vulkan) ResetEvent(event types.Event) (common.VkResult, error) {
	if event.Handle() == 0 {
		return core1_0.VKErrorUnknown, fmt.Errorf("event was uninitialized")
	}

	return v.Driver.VkResetEvent(event.DeviceHandle(), event.Handle())
}

func (v *Vulkan) GetEventStatus(event types.Event) (common.VkResult, error) {
	if event.Handle() == 0 {
		return core1_0.VKErrorUnknown, fmt.Errorf("event was uninitialized")
	}

	return v.Driver.VkGetEventStatus(event.DeviceHandle(), event.Handle())
}
