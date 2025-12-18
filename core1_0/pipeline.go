package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanPipeline is an implementation of the Pipeline interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPipeline struct {
	deviceDriver   driver.Driver
	device         driver.VkDevice
	pipelineHandle driver.VkPipeline

	maximumAPIVersion common.APIVersion
}

func (p *VulkanPipeline) Handle() driver.VkPipeline {
	return p.pipelineHandle
}

func (p *VulkanPipeline) Driver() driver.Driver {
	return p.deviceDriver
}

func (p *VulkanPipeline) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p *VulkanPipeline) APIVersion() common.APIVersion {
	return p.maximumAPIVersion
}

func (p *VulkanPipeline) Destroy(callbacks *driver.AllocationCallbacks) {
	p.deviceDriver.VkDestroyPipeline(p.device, p.pipelineHandle, callbacks.Handle())
	p.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(p.pipelineHandle))
}
