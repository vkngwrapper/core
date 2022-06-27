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

type VulkanFramebuffer struct {
	deviceDriver      driver.Driver
	device            driver.VkDevice
	framebufferHandle driver.VkFramebuffer

	maximumAPIVersion common.APIVersion
}

func (b *VulkanFramebuffer) Handle() driver.VkFramebuffer {
	return b.framebufferHandle
}

func (b *VulkanFramebuffer) DeviceHandle() driver.VkDevice {
	return b.device
}

func (b *VulkanFramebuffer) Driver() driver.Driver {
	return b.deviceDriver
}

func (b *VulkanFramebuffer) APIVersion() common.APIVersion {
	return b.maximumAPIVersion
}

func (b *VulkanFramebuffer) Destroy(callbacks *driver.AllocationCallbacks) {
	b.deviceDriver.VkDestroyFramebuffer(b.device, b.framebufferHandle, callbacks.Handle())
	b.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(b.framebufferHandle))
}
