package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
)

// VulkanPipelineLayout is an implementation of the PipelineLayout interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanPipelineLayout struct {
	deviceDriver         driver.Driver
	device               driver.VkDevice
	pipelineLayoutHandle driver.VkPipelineLayout

	maximumAPIVersion common.APIVersion
}

func (l *VulkanPipelineLayout) Handle() driver.VkPipelineLayout {
	return l.pipelineLayoutHandle
}

func (l *VulkanPipelineLayout) Driver() driver.Driver {
	return l.deviceDriver
}

func (l *VulkanPipelineLayout) DeviceHandle() driver.VkDevice {
	return l.device
}

func (l *VulkanPipelineLayout) APIVersion() common.APIVersion {
	return l.maximumAPIVersion
}

func (l *VulkanPipelineLayout) Destroy(callbacks *driver.AllocationCallbacks) {
	l.deviceDriver.VkDestroyPipelineLayout(l.device, l.pipelineLayoutHandle, callbacks.Handle())
	l.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(l.pipelineLayoutHandle))
}
