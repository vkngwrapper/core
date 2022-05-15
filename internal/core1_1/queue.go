package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanQueue struct {
	core1_0.Queue
}

func PromoteQueue(queue core1_0.Queue) core1_1.Queue {
	if !queue.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return queue.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(queue.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanQueue{queue}
		}).(core1_1.Queue)
}
