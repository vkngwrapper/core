package core1_1

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanSemaphore is an implementation of the Semaphore interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanSemaphore struct {
	core1_0.Semaphore
}

// PromoteSemaphore accepts a Semaphore object from any core version. If provided a semaphore that supports
// at least core 1.1, it will return a core1_1.Semaphore. Otherwise, it will return nil. This method
// will always return a core1_1.VulkanSemaphore, even if it is provided a VulkanSemaphore from a higher
// core version. Two Vulkan 1.1 compatible Semaphore objects with the same Semaphore.Handle will
// return the same interface value when passed to this method.
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
