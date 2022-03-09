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

type VulkanDescriptorPool struct {
	DeviceDriver         driver.Driver
	DescriptorPoolHandle driver.VkDescriptorPool
	Device               driver.VkDevice

	MaximumAPIVersion common.APIVersion
}

func (p *VulkanDescriptorPool) Handle() driver.VkDescriptorPool {
	return p.DescriptorPoolHandle
}

func (p *VulkanDescriptorPool) DeviceHandle() driver.VkDevice {
	return p.Device
}

func (p *VulkanDescriptorPool) Driver() driver.Driver {
	return p.DeviceDriver
}

func (p *VulkanDescriptorPool) APIVersion() common.APIVersion {
	return p.MaximumAPIVersion
}

func (p *VulkanDescriptorPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.DeviceDriver.VkDestroyDescriptorPool(p.Device, p.DescriptorPoolHandle, callbacks.Handle())
}

func (p *VulkanDescriptorPool) Reset(flags common.DescriptorPoolResetFlags) (common.VkResult, error) {
	return p.DeviceDriver.VkResetDescriptorPool(p.Device, p.DescriptorPoolHandle, driver.VkDescriptorPoolResetFlags(flags))
}
