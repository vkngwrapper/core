package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanDescriptorSetLayout is an implementation of the DescriptorSetLayout interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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

func (h *VulkanDescriptorSetLayout) DeviceHandle() driver.VkDevice {
	return h.Device
}

func (h *VulkanDescriptorSetLayout) APIVersion() common.APIVersion {
	return h.MaximumAPIVersion
}

func (h *VulkanDescriptorSetLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	h.DeviceDriver.VkDestroyDescriptorSetLayout(h.Device, h.DescriptorSetLayoutHandle, callbacks.Handle())
}
