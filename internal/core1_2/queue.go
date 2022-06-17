package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanQueue struct {
	core1_1.Queue
}

func PromoteQueue(queue core1_0.Queue) core1_2.Queue {
	if !queue.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return queue.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(queue.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanQueue{core1_1.PromoteQueue(queue)}
		}).(core1_2.Queue)
}
