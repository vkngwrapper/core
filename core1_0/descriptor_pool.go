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
	deviceDriver         driver.Driver
	descriptorPoolHandle driver.VkDescriptorPool
	device               driver.VkDevice

	maximumAPIVersion common.APIVersion
}

func (p *VulkanDescriptorPool) Handle() driver.VkDescriptorPool {
	return p.descriptorPoolHandle
}

func (p *VulkanDescriptorPool) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p *VulkanDescriptorPool) Driver() driver.Driver {
	return p.deviceDriver
}

func (p *VulkanDescriptorPool) APIVersion() common.APIVersion {
	return p.maximumAPIVersion
}

func (p *VulkanDescriptorPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.deviceDriver.VkDestroyDescriptorPool(p.device, p.descriptorPoolHandle, callbacks.Handle())
	p.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(p.descriptorPoolHandle))
}

func (p *VulkanDescriptorPool) Reset(flags common.DescriptorPoolResetFlags) (common.VkResult, error) {
	return p.deviceDriver.VkResetDescriptorPool(p.device, p.descriptorPoolHandle, driver.VkDescriptorPoolResetFlags(flags))
}
