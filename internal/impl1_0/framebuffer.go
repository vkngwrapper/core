package impl1_0

import (
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyFramebuffer(framebuffer types.Framebuffer, callbacks *driver.AllocationCallbacks) {
	v.Driver.VkDestroyFramebuffer(framebuffer.DeviceHandle(), framebuffer.Handle(), callbacks.Handle())
}
