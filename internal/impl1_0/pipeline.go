package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanPipeline is an implementation of the Pipeline interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPipeline struct {
	DeviceDriver   driver.Driver
	Device         driver.VkDevice
	PipelineHandle driver.VkPipeline

	MaximumAPIVersion common.APIVersion
}

func (p *VulkanPipeline) Handle() driver.VkPipeline {
	return p.PipelineHandle
}

func (p *VulkanPipeline) Driver() driver.Driver {
	return p.DeviceDriver
}

func (p *VulkanPipeline) DeviceHandle() driver.VkDevice {
	return p.Device
}

func (p *VulkanPipeline) APIVersion() common.APIVersion {
	return p.MaximumAPIVersion
}

func (p *VulkanPipeline) Destroy(callbacks *driver.AllocationCallbacks) {
	p.DeviceDriver.VkDestroyPipeline(p.Device, p.PipelineHandle, callbacks.Handle())
}
