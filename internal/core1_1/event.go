package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanEvent struct {
	core1_0.Event
}

func PromoteEvent(event core1_0.Event) core1_1.Event {
	if !event.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return event.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(event.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanEvent{event}
		}).(core1_1.Event)
}
