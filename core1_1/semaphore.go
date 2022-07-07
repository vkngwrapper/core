package core1_1

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
)

type VulkanSemaphore struct {
	core1_0.Semaphore
}

func PromoteSemaphore(semaphore core1_0.Semaphore) Semaphore {
	if !semaphore.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return semaphore.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(semaphore.Handle()),
		driver.Core1_1,
		func() any {
			return &VulkanSemaphore{semaphore}
		}).(Semaphore)
}
