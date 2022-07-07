package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanEvent struct {
	core1_0.Event
}

func PromoteEvent(event core1_0.Event) Event {
	if !event.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return event.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(event.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanEvent{event}
		}).(Event)
}
