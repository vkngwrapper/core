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

type VulkanPipelineLayout struct {
	Driver               driver.Driver
	Device               driver.VkDevice
	PipelineLayoutHandle driver.VkPipelineLayout

	MaximumAPIVersion common.APIVersion
}

func (l *VulkanPipelineLayout) Handle() driver.VkPipelineLayout {
	return l.PipelineLayoutHandle
}

func (l *VulkanPipelineLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	l.Driver.VkDestroyPipelineLayout(l.Device, l.PipelineLayoutHandle, callbacks.Handle())
	l.Driver.ObjectStore().Delete(driver.VulkanHandle(l.PipelineLayoutHandle), l)
}
