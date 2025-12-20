package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanSampler is an implementation of the Sampler interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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

func (s *VulkanSampler) DeviceHandle() driver.VkDevice {
	return s.Device
}

func (s *VulkanSampler) APIVersion() common.APIVersion {
	return s.MaximumAPIVersion
}

func (s *VulkanSampler) Destroy(callbacks *driver.AllocationCallbacks) {
	s.DeviceDriver.VkDestroySampler(s.Device, s.SamplerHandle, callbacks.Handle())
}
