package core1_0

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
)

type VulkanDescriptorSetLayout struct {
	deviceDriver              driver.Driver
	device                    driver.VkDevice
	descriptorSetLayoutHandle driver.VkDescriptorSetLayout

	maximumAPIVersion common.APIVersion
}

func (h *VulkanDescriptorSetLayout) Handle() driver.VkDescriptorSetLayout {
	return h.descriptorSetLayoutHandle
}

func (h *VulkanDescriptorSetLayout) Driver() driver.Driver {
	return h.deviceDriver
}

func (h *VulkanDescriptorSetLayout) DeviceHandle() driver.VkDevice {
	return h.device
}

func (h *VulkanDescriptorSetLayout) APIVersion() common.APIVersion {
	return h.maximumAPIVersion
}

func (h *VulkanDescriptorSetLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	h.deviceDriver.VkDestroyDescriptorSetLayout(h.device, h.descriptorSetLayoutHandle, callbacks.Handle())
	h.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(h.descriptorSetLayoutHandle))
}
