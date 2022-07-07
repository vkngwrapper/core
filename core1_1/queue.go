package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanQueue struct {
	core1_0.Queue
}

func PromoteQueue(queue core1_0.Queue) Queue {
	if !queue.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return queue.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(queue.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanQueue{queue}
		}).(Queue)
}
