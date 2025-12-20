package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanPipelineLayout is an implementation of the PipelineLayout interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPipelineLayout struct {
	DeviceDriver         driver.Driver
	Device               driver.VkDevice
	PipelineLayoutHandle driver.VkPipelineLayout

	MaximumAPIVersion common.APIVersion
}

func (l *VulkanPipelineLayout) Handle() driver.VkPipelineLayout {
	return l.PipelineLayoutHandle
}

func (l *VulkanPipelineLayout) Driver() driver.Driver {
	return l.DeviceDriver
}

func (l *VulkanPipelineLayout) DeviceHandle() driver.VkDevice {
	return l.Device
}

func (l *VulkanPipelineLayout) APIVersion() common.APIVersion {
	return l.MaximumAPIVersion
}

func (l *VulkanPipelineLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	l.DeviceDriver.VkDestroyPipelineLayout(l.Device, l.PipelineLayoutHandle, callbacks.Handle())
}
