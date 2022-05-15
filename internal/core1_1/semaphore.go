package internal1_1

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSemaphore struct {
	core1_0.Semaphore
}

func PromoteSemaphore(semaphore core1_0.Semaphore) core1_1.Semaphore {
	if !semaphore.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return semaphore.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(semaphore.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanSemaphore{semaphore}
		}).(core1_1.Semaphore)
}
