package core1_0

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
)

type VulkanSampler struct {
	device        driver.VkDevice
	deviceDriver  driver.Driver
	samplerHandle driver.VkSampler

	maximumAPIVersion common.APIVersion
}

func (s *VulkanSampler) Handle() driver.VkSampler {
	return s.samplerHandle
}

func (s *VulkanSampler) Driver() driver.Driver {
	return s.deviceDriver
}

func (s *VulkanSampler) DeviceHandle() driver.VkDevice {
	return s.device
}

func (s *VulkanSampler) APIVersion() common.APIVersion {
	return s.maximumAPIVersion
}

func (s *VulkanSampler) Destroy(callbacks *driver.AllocationCallbacks) {
	s.deviceDriver.VkDestroySampler(s.device, s.samplerHandle, callbacks.Handle())
	s.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(s.samplerHandle))
}
