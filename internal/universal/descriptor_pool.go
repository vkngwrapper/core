package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorPool struct {
	driver driver.Driver
	handle driver.VkDescriptorPool
	device driver.VkDevice
}

func (p *VulkanDescriptorPool) Handle() driver.VkDescriptorPool {
	return p.handle
}

func (p *VulkanDescriptorPool) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p *VulkanDescriptorPool) Driver() driver.Driver {
	return p.driver
}

func (p *VulkanDescriptorPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.driver.VkDestroyDescriptorPool(p.device, p.handle, callbacks.Handle())
}

func (p *VulkanDescriptorPool) Reset(flags core1_0.DescriptorPoolResetFlags) (common.VkResult, error) {
	return p.driver.VkResetDescriptorPool(p.device, p.handle, driver.VkDescriptorPoolResetFlags(flags))
}
