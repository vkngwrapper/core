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
	Driver            driver.Driver
	Device            driver.VkDevice
	FramebufferHandle driver.VkFramebuffer

	MaximumAPIVersion common.APIVersion
}

func (b *VulkanFramebuffer) Handle() driver.VkFramebuffer {
	return b.FramebufferHandle
}

func (b *VulkanFramebuffer) Destroy(callbacks *driver.AllocationCallbacks) {
	b.Driver.VkDestroyFramebuffer(b.Device, b.FramebufferHandle, callbacks.Handle())
}
