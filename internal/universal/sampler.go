package universal

import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanSampler struct {
	device driver.VkDevice
	driver driver.Driver
	handle driver.VkSampler
}

func (s *VulkanSampler) Handle() driver.VkSampler {
	return s.handle
}

func (s *VulkanSampler) Destroy(callbacks *driver.AllocationCallbacks) {
	s.driver.VkDestroySampler(s.device, s.handle, callbacks.Handle())
}
