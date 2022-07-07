package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

type VulkanQueue struct {
	core1_1.Queue
}

func PromoteQueue(queue core1_0.Queue) Queue {
	if !queue.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedQueue := core1_1.PromoteQueue(queue)
	return queue.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(queue.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanQueue{promotedQueue}
		}).(Queue)
}
