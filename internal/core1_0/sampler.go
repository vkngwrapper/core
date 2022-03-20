package core1_0

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSampler struct {
	Device        driver.VkDevice
	Driver        driver.Driver
	SamplerHandle driver.VkSampler

	MaximumAPIVersion common.APIVersion

	Sampler1_1 core1_1.Sampler
}

func (s *VulkanSampler) Handle() driver.VkSampler {
	return s.SamplerHandle
}

func (s *VulkanSampler) Core1_1() core1_1.Sampler {
	return s.Sampler1_1
}

func (s *VulkanSampler) Destroy(callbacks *driver.AllocationCallbacks) {
	s.Driver.VkDestroySampler(s.Device, s.SamplerHandle, callbacks.Handle())
	s.Driver.ObjectStore().Delete(driver.VulkanHandle(s.SamplerHandle), s)
}
