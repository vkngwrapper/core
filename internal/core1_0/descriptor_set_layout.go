package core1_0

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorSetLayout struct {
	Driver                    driver.Driver
	Device                    driver.VkDevice
	DescriptorSetLayoutHandle driver.VkDescriptorSetLayout

	MaximumAPIVersion common.APIVersion

	DescriptorSetLayout1_1 core1_1.DescriptorSetLayout
}

func (h *VulkanDescriptorSetLayout) Handle() driver.VkDescriptorSetLayout {
	return h.DescriptorSetLayoutHandle
}

func (h *VulkanDescriptorSetLayout) Core1_1() core1_1.DescriptorSetLayout {
	return h.DescriptorSetLayout1_1
}

func (h *VulkanDescriptorSetLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	h.Driver.VkDestroyDescriptorSetLayout(h.Device, h.DescriptorSetLayoutHandle, callbacks.Handle())
}
