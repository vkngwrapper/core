package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanSemaphore is an implementation of the Semaphore interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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
}
