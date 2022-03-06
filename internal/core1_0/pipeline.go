package core1_0

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
	Driver         driver.Driver
	Device         driver.VkDevice
	PipelineHandle driver.VkPipeline

	MaximumAPIVersion common.APIVersion
}

func (p *VulkanPipeline) Handle() driver.VkPipeline {
	return p.PipelineHandle
}

func (p *VulkanPipeline) Destroy(callbacks *driver.AllocationCallbacks) {
	p.Driver.VkDestroyPipeline(p.Device, p.PipelineHandle, callbacks.Handle())
}
