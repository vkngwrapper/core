package internal1_0

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSampler struct {
	Device        driver.VkDevice
	Driver        driver.Driver
	SamplerHandle driver.VkSampler

	MaximumAPIVersion common.APIVersion
}

func (s *VulkanSampler) Handle() driver.VkSampler {
	return s.SamplerHandle
}

func (s *VulkanSampler) Destroy(callbacks *driver.AllocationCallbacks) {
	s.Driver.VkDestroySampler(s.Device, s.SamplerHandle, callbacks.Handle())
	s.Driver.ObjectStore().Delete(driver.VulkanHandle(s.SamplerHandle), s)
}
