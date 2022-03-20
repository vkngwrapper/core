package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanPipeline struct {
	Driver         driver.Driver
	Device         driver.VkDevice
	PipelineHandle driver.VkPipeline

	MaximumAPIVersion common.APIVersion

	Pipeline1_1 core1_1.Pipeline
}

func (p *VulkanPipeline) Handle() driver.VkPipeline {
	return p.PipelineHandle
}

func (p *VulkanPipeline) Core1_1() core1_1.Pipeline {
	return p.Pipeline1_1
}

func (p *VulkanPipeline) Destroy(callbacks *driver.AllocationCallbacks) {
	p.Driver.VkDestroyPipeline(p.Device, p.PipelineHandle, callbacks.Handle())
	p.Driver.ObjectStore().Delete(driver.VulkanHandle(p.PipelineHandle), p)
}
