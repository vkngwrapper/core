package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanDescriptorPool is an implementation of the DescriptorPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDescriptorPool struct {
	DeviceDriver         driver.Driver
	DescriptorPoolHandle driver.VkDescriptorPool
	Device               driver.VkDevice

	MaximumAPIVersion common.APIVersion
}

func (p *VulkanDescriptorPool) Handle() driver.VkDescriptorPool {
	return p.DescriptorPoolHandle
}

func (p *VulkanDescriptorPool) DeviceHandle() driver.VkDevice {
	return p.Device
}

func (p *VulkanDescriptorPool) Driver() driver.Driver {
	return p.DeviceDriver
}

func (p *VulkanDescriptorPool) APIVersion() common.APIVersion {
	return p.MaximumAPIVersion
}

func (p *VulkanDescriptorPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.DeviceDriver.VkDestroyDescriptorPool(p.Device, p.DescriptorPoolHandle, callbacks.Handle())
}

func (p *VulkanDescriptorPool) Reset(flags core1_0.DescriptorPoolResetFlags) (common.VkResult, error) {
	return p.DeviceDriver.VkResetDescriptorPool(p.Device, p.DescriptorPoolHandle, driver.VkDescriptorPoolResetFlags(flags))
}
