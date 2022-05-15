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

type VulkanFramebuffer struct {
	DeviceDriver      driver.Driver
	Device            driver.VkDevice
	FramebufferHandle driver.VkFramebuffer

	MaximumAPIVersion common.APIVersion
}

func (b *VulkanFramebuffer) Handle() driver.VkFramebuffer {
	return b.FramebufferHandle
}

func (b *VulkanFramebuffer) Driver() driver.Driver {
	return b.DeviceDriver
}

func (b *VulkanFramebuffer) APIVersion() common.APIVersion {
	return b.MaximumAPIVersion
}

func (b *VulkanFramebuffer) Destroy(callbacks *driver.AllocationCallbacks) {
	b.DeviceDriver.VkDestroyFramebuffer(b.Device, b.FramebufferHandle, callbacks.Handle())
	b.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(b.FramebufferHandle))
}
