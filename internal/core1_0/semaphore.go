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
	DeviceDriver    driver.Driver
	Device          driver.VkDevice
	SemaphoreHandle driver.VkSemaphore

	MaximumAPIVersion common.APIVersion
}

func (s *VulkanSemaphore) Handle() driver.VkSemaphore {
	return s.SemaphoreHandle
}

func (s *VulkanSemaphore) DeviceHandle() driver.VkDevice {
	return s.Device
}

func (s *VulkanSemaphore) Driver() driver.Driver {
	return s.DeviceDriver
}

func (s *VulkanSemaphore) APIVersion() common.APIVersion {
	return s.MaximumAPIVersion
}

func (s *VulkanSemaphore) Destroy(callbacks *driver.AllocationCallbacks) {
	s.DeviceDriver.VkDestroySemaphore(s.Device, s.SemaphoreHandle, callbacks.Handle())
	s.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(s.SemaphoreHandle))
}
