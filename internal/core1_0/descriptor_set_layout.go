package internal1_0

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorSetLayout struct {
	DeviceDriver              driver.Driver
	Device                    driver.VkDevice
	DescriptorSetLayoutHandle driver.VkDescriptorSetLayout

	MaximumAPIVersion common.APIVersion
}

func (h *VulkanDescriptorSetLayout) Handle() driver.VkDescriptorSetLayout {
	return h.DescriptorSetLayoutHandle
}

func (h *VulkanDescriptorSetLayout) Driver() driver.Driver {
	return h.DeviceDriver
}

func (h *VulkanDescriptorSetLayout) APIVersion() common.APIVersion {
	return h.MaximumAPIVersion
}

func (h *VulkanDescriptorSetLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	h.DeviceDriver.VkDestroyDescriptorSetLayout(h.Device, h.DescriptorSetLayoutHandle, callbacks.Handle())
	h.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(h.DescriptorSetLayoutHandle))
}
