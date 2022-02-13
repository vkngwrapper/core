package universal

import (
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorSetLayout struct {
	driver driver.Driver
	device driver.VkDevice
	handle driver.VkDescriptorSetLayout
}

func (h *VulkanDescriptorSetLayout) Handle() driver.VkDescriptorSetLayout {
	return h.handle
}

func (h *VulkanDescriptorSetLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	h.driver.VkDestroyDescriptorSetLayout(h.device, h.handle, callbacks.Handle())
}
