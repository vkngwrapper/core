package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanEvent struct {
	core1_1.Event
}

func PromoteEvent(event core1_0.Event) core1_2.Event {
	if !event.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return event.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(event.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanEvent{core1_1.PromoteEvent(event)}
		}).(core1_2.Event)
}
