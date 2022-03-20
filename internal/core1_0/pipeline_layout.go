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

type VulkanPipelineLayout struct {
	Driver               driver.Driver
	Device               driver.VkDevice
	PipelineLayoutHandle driver.VkPipelineLayout

	MaximumAPIVersion common.APIVersion

	PipelineLayout1_1 core1_1.PipelineLayout
}

func (l *VulkanPipelineLayout) Handle() driver.VkPipelineLayout {
	return l.PipelineLayoutHandle
}

func (l *VulkanPipelineLayout) Core1_1() core1_1.PipelineLayout {
	return l.PipelineLayout1_1
}

func (l *VulkanPipelineLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	l.Driver.VkDestroyPipelineLayout(l.Device, l.PipelineLayoutHandle, callbacks.Handle())
	l.Driver.ObjectStore().Delete(driver.VulkanHandle(l.PipelineLayoutHandle), l)
}
