package core1_0

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorSetLayout struct {
	Driver                    driver.Driver
	Device                    driver.VkDevice
	DescriptorSetLayoutHandle driver.VkDescriptorSetLayout

	MaximumAPIVersion common.APIVersion
}

func (h *VulkanDescriptorSetLayout) Handle() driver.VkDescriptorSetLayout {
	return h.DescriptorSetLayoutHandle
}

func (h *VulkanDescriptorSetLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	h.Driver.VkDestroyDescriptorSetLayout(h.Device, h.DescriptorSetLayoutHandle, callbacks.Handle())
}
