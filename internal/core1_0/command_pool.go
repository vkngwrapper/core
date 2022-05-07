package internal1_0

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

type VulkanCommandPool struct {
	DeviceDriver      driver.Driver
	CommandPoolHandle driver.VkCommandPool
	DeviceHandle      driver.VkDevice
	MaximumAPIVersion common.APIVersion

	CommandPool1_1 core1_1.CommandPool
}

func (p *VulkanCommandPool) Handle() driver.VkCommandPool {
	return p.CommandPoolHandle
}

func (p *VulkanCommandPool) Device() driver.VkDevice {
	return p.DeviceHandle
}

func (p *VulkanCommandPool) Driver() driver.Driver {
	return p.DeviceDriver
}

func (p *VulkanCommandPool) APIVersion() common.APIVersion {
	return p.MaximumAPIVersion
}

func (p *VulkanCommandPool) Core1_1() core1_1.CommandPool {
	return p.CommandPool1_1
}

func (p *VulkanCommandPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.DeviceDriver.VkDestroyCommandPool(p.DeviceHandle, p.CommandPoolHandle, callbacks.Handle())
	p.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(p.CommandPoolHandle), p)
}

func (p *VulkanCommandPool) Reset(flags common.CommandPoolResetFlags) (common.VkResult, error) {
	return p.DeviceDriver.VkResetCommandPool(p.DeviceHandle, p.CommandPoolHandle, driver.VkCommandPoolResetFlags(flags))
}
