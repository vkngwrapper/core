package universal

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanCommandPool struct {
	driver driver.Driver
	handle driver.VkCommandPool
	device driver.VkDevice
}

func (p *VulkanCommandPool) Handle() driver.VkCommandPool {
	return p.handle
}

func (p *VulkanCommandPool) Device() driver.VkDevice {
	return p.device
}

func (p *VulkanCommandPool) Driver() driver.Driver {
	return p.driver
}

func (p *VulkanCommandPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.driver.VkDestroyCommandPool(p.device, p.handle, callbacks.Handle())
}

func (p *VulkanCommandPool) Reset(flags core.CommandPoolResetFlags) (common.VkResult, error) {
	return p.driver.VkResetCommandPool(p.device, p.handle, driver.VkCommandPoolResetFlags(flags))
}
