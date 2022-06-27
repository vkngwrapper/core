package core1_0

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
	deviceDriver    driver.Driver
	device          driver.VkDevice
	semaphoreHandle driver.VkSemaphore

	maximumAPIVersion common.APIVersion
}

func (s *VulkanSemaphore) Handle() driver.VkSemaphore {
	return s.semaphoreHandle
}

func (s *VulkanSemaphore) DeviceHandle() driver.VkDevice {
	return s.device
}

func (s *VulkanSemaphore) Driver() driver.Driver {
	return s.deviceDriver
}

func (s *VulkanSemaphore) APIVersion() common.APIVersion {
	return s.maximumAPIVersion
}

func (s *VulkanSemaphore) Destroy(callbacks *driver.AllocationCallbacks) {
	s.deviceDriver.VkDestroySemaphore(s.device, s.semaphoreHandle, callbacks.Handle())
	s.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(s.semaphoreHandle))
}
