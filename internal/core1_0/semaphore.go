package internal1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSemaphore struct {
	Driver          driver.Driver
	Device          driver.VkDevice
	SemaphoreHandle driver.VkSemaphore

	MaximumAPIVersion common.APIVersion
}

func (s *VulkanSemaphore) Handle() driver.VkSemaphore {
	return s.SemaphoreHandle
}

func (s *VulkanSemaphore) Destroy(callbacks *driver.AllocationCallbacks) {
	s.Driver.VkDestroySemaphore(s.Device, s.SemaphoreHandle, callbacks.Handle())
	s.Driver.ObjectStore().Delete(driver.VulkanHandle(s.SemaphoreHandle), s)
}
