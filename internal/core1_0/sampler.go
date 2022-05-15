package internal1_0

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSampler struct {
	Device        driver.VkDevice
	DeviceDriver  driver.Driver
	SamplerHandle driver.VkSampler

	MaximumAPIVersion common.APIVersion
}

func (s *VulkanSampler) Handle() driver.VkSampler {
	return s.SamplerHandle
}

func (s *VulkanSampler) Driver() driver.Driver {
	return s.DeviceDriver
}

func (s *VulkanSampler) APIVersion() common.APIVersion {
	return s.MaximumAPIVersion
}

func (s *VulkanSampler) Destroy(callbacks *driver.AllocationCallbacks) {
	s.DeviceDriver.VkDestroySampler(s.Device, s.SamplerHandle, callbacks.Handle())
	s.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(s.SamplerHandle))
}
