package internal1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

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
	p.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(p.PipelineHandle))
}
