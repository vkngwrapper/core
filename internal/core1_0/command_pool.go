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

type VulkanCommandPool struct {
	DeviceDriver      driver.Driver
	CommandPoolHandle driver.VkCommandPool
	Device            driver.VkDevice
	MaximumAPIVersion common.APIVersion
}

func (p *VulkanCommandPool) Handle() driver.VkCommandPool {
	return p.CommandPoolHandle
}

func (p *VulkanCommandPool) DeviceHandle() driver.VkDevice {
	return p.Device
}

func (p *VulkanCommandPool) Driver() driver.Driver {
	return p.DeviceDriver
}

func (p *VulkanCommandPool) APIVersion() common.APIVersion {
	return p.MaximumAPIVersion
}

func (p *VulkanCommandPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.DeviceDriver.VkDestroyCommandPool(p.Device, p.CommandPoolHandle, callbacks.Handle())
	p.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(p.CommandPoolHandle))
}

func (p *VulkanCommandPool) Reset(flags common.CommandPoolResetFlags) (common.VkResult, error) {
	return p.DeviceDriver.VkResetCommandPool(p.Device, p.CommandPoolHandle, driver.VkCommandPoolResetFlags(flags))
}
