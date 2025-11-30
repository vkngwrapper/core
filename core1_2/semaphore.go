package core1_2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanSemaphore is an implementation of the Semaphore interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanSemaphore struct {
	core1_1.Semaphore

	DeviceDriver    driver.Driver
	Device          driver.VkDevice
	SemaphoreHandle driver.VkSemaphore
}

// PromoteSemaphore accepts a Semaphore object from any core version. If provided a semaphore that supports
// at least core 1.2, it will return a core1_2.Semaphore. Otherwise, it will return nil. This method
// will always return a core1_2.VulkanSemaphore, even if it is provided a VulkanSemaphore from a higher
// core version. Two Vulkan 1.2 compatible Semaphore objects with the same Semaphore.Handle will
// return the same interface value when passed to this method.
func PromoteSemaphore(semaphore core1_0.Semaphore) Semaphore {
	if semaphore == nil {
		return nil
	}
	if !semaphore.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promoted, alreadyPromoted := semaphore.(Semaphore)
	if alreadyPromoted {
		return promoted
	}

	promotedSemaphore := core1_1.PromoteSemaphore(semaphore)
	return semaphore.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(semaphore.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanSemaphore{
				Semaphore: promotedSemaphore,

				DeviceDriver:    semaphore.Driver(),
				Device:          semaphore.DeviceHandle(),
				SemaphoreHandle: semaphore.Handle(),
			}
		}).(Semaphore)
}

func (s *VulkanSemaphore) CounterValue() (uint64, common.VkResult, error) {
	var value driver.Uint64
	res, err := s.DeviceDriver.VkGetSemaphoreCounterValue(
		s.Device,
		s.SemaphoreHandle,
		&value,
	)
	if err != nil {
		return 0, res, err
	}

	return uint64(value), res, nil
}
