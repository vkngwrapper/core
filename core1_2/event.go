package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanEvent is an implementation of the Event interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanEvent struct {
	core1_1.Event
}

// PromoteEvent accepts a Event object from any core version. If provided an event that supports
// at least core 1.2, it will return a core1_2.Event. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanEvent, even if it is provided a VulkanEvent from a higher
// core version. Two Vulkan 1.2 compatible Event objects with the same Event.Handle will
// return the same interface value when passed to this method.
func PromoteEvent(event core1_0.Event) Event {
	if event == nil {
		return nil
	}
	if !event.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedEvent := core1_1.PromoteEvent(event)
	return event.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(event.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanEvent{promotedEvent}
		}).(Event)
}
