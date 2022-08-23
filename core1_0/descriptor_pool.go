package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/driver"
)

// VulkanDescriptorPool is an implementation of the DescriptorPool interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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

func (p *VulkanDescriptorPool) Reset(flags DescriptorPoolResetFlags) (common.VkResult, error) {
	return p.deviceDriver.VkResetDescriptorPool(p.device, p.descriptorPoolHandle, driver.VkDescriptorPoolResetFlags(flags))
}
