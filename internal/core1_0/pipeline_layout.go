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

func (l *VulkanPipelineLayout) APIVersion() common.APIVersion {
	return l.MaximumAPIVersion
}

func (l *VulkanPipelineLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	l.DeviceDriver.VkDestroyPipelineLayout(l.Device, l.PipelineLayoutHandle, callbacks.Handle())
	l.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(l.PipelineLayoutHandle))
}
