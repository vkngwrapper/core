package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSemaphore struct {
	core1_1.Semaphore

	DeviceDriver    driver.Driver
	DeviceHandle    driver.VkDevice
	SemaphoreHandle driver.VkSemaphore
}

func PromoteSemaphore(semaphore core1_0.Semaphore) core1_2.Semaphore {
	if !semaphore.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return semaphore.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(semaphore.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanSemaphore{
				Semaphore: core1_1.PromoteSemaphore(semaphore),

				DeviceDriver:    semaphore.Driver(),
				DeviceHandle:    semaphore.DeviceHandle(),
				SemaphoreHandle: semaphore.Handle(),
			}
		}).(core1_2.Semaphore)
}

func (s *VulkanSemaphore) CounterValue() (uint64, common.VkResult, error) {

	var value driver.Uint64
	res, err := s.DeviceDriver.VkGetSemaphoreCounterValue(
		s.DeviceHandle,
		s.SemaphoreHandle,
		&value,
	)
	if err != nil {
		return 0, res, err
	}

	return uint64(value), res, nil
}
