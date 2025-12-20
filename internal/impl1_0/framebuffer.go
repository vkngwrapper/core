package impl1_0

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/driver"
)

// VulkanFramebuffer is an implementation of the Framebuffer interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanFramebuffer struct {
	DeviceDriver      driver.Driver
	Device            driver.VkDevice
	FramebufferHandle driver.VkFramebuffer

	MaximumAPIVersion common.APIVersion
}

func (b *VulkanFramebuffer) Handle() driver.VkFramebuffer {
	return b.FramebufferHandle
}

func (b *VulkanFramebuffer) DeviceHandle() driver.VkDevice {
	return b.Device
}

func (b *VulkanFramebuffer) Driver() driver.Driver {
	return b.DeviceDriver
}

func (b *VulkanFramebuffer) APIVersion() common.APIVersion {
	return b.MaximumAPIVersion
}

func (b *VulkanFramebuffer) Destroy(callbacks *driver.AllocationCallbacks) {
	b.DeviceDriver.VkDestroyFramebuffer(b.Device, b.FramebufferHandle, callbacks.Handle())
}
