package core1_1

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanQueue is an implementation of the Queue interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanQueue struct {
	core1_0.Queue
}

// PromoteQueue accepts a Queue object from any core version. If provided a queue that supports
// at least core 1.1, it will return a core1_1.Queue. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanQueue, even if it is provided a VulkanQueue from a higher
// core version. Two Vulkan 1.1 compatible Queue objects with the same Queue.Handle will
// return the same interface value when passed to this method.
func PromoteQueue(queue core1_0.Queue) Queue {
	if queue == nil {
		return nil
	}
	if !queue.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	promoted, alreadyPromoted := queue.(Queue)
	if alreadyPromoted {
		return promoted
	}

	return queue.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(queue.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanQueue{queue}
		}).(Queue)
}
