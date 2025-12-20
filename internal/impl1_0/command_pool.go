package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanCommandPool is an implementation of the CommandPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanCommandPool struct {
	DeviceDriver      driver.Driver
	CommandPoolHandle driver.VkCommandPool
	Device            driver.VkDevice
	MaximumAPIVersion common.APIVersion
}

func (p *VulkanCommandPool) Handle() driver.VkCommandPool {
	return p.CommandPoolHandle
}

func (p *VulkanCommandPool) DeviceHandle() driver.VkDevice {
	return p.Device
}

func (p *VulkanCommandPool) Driver() driver.Driver {
	return p.DeviceDriver
}

func (p *VulkanCommandPool) APIVersion() common.APIVersion {
	return p.MaximumAPIVersion
}

func (p *VulkanCommandPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.DeviceDriver.VkDestroyCommandPool(p.Device, p.CommandPoolHandle, callbacks.Handle())
}

func (p *VulkanCommandPool) Reset(flags core1_0.CommandPoolResetFlags) (common.VkResult, error) {
	return p.DeviceDriver.VkResetCommandPool(p.Device, p.CommandPoolHandle, driver.VkCommandPoolResetFlags(flags))
}
