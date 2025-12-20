package impl1_2

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/internal/impl1_1"
)

// VulkanSemaphore is an implementation of the Semaphore interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanSemaphore struct {
	impl1_1.VulkanSemaphore
}

func (s *VulkanSemaphore) CounterValue() (uint64, common.VkResult, error) {
	var value driver.Uint64
	res, err := s.Driver().VkGetSemaphoreCounterValue(
		s.DeviceHandle(),
		s.Handle(),
		&value,
	)
	if err != nil {
		return 0, res, err
	}

	return uint64(value), res, nil
}
