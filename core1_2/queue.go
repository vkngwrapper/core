package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
)

// VulkanQueue is an implementation of the Queue interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanQueue struct {
	core1_1.Queue
}

// PromoteQueue accepts a Queue object from any core version. If provided a queue that supports
// at least core 1.2, it will return a core1_2.Queue. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanQueue, even if it is provided a VulkanQueue from a higher
// core version. Two Vulkan 1.2 compatible Queue objects with the same Queue.Handle will
// return the same interface value when passed to this method.
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
