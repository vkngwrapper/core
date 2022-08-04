package core1_0

/*
#include <stdlib.h>
#include "../common/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
)

// VulkanFramebuffer is an implementation of the Framebuffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
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
