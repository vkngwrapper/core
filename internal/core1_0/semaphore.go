package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSemaphore struct {
	Driver          driver.Driver
	Device          driver.VkDevice
	SemaphoreHandle driver.VkSemaphore

	MaximumAPIVersion common.APIVersion

	Semaphore1_1 core1_1.Semaphore
}

func (s *VulkanSemaphore) Handle() driver.VkSemaphore {
	return s.SemaphoreHandle
}

func (s *VulkanSemaphore) Core1_1() core1_1.Semaphore {
	return s.Semaphore1_1
}

func (s *VulkanSemaphore) Destroy(callbacks *driver.AllocationCallbacks) {
	s.Driver.VkDestroySemaphore(s.Device, s.SemaphoreHandle, callbacks.Handle())
}
