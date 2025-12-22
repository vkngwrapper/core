package impl1_0

import (
	"github.com/vkngwrapper/core/v3/driver"
	"github.com/vkngwrapper/core/v3/types"
)

func (v *Vulkan) DestroyDescriptorSetLayout(layout types.DescriptorSetLayout, callbacks *driver.AllocationCallbacks) {
	if layout.Handle() == 0 {
		panic("layout was uninitialiazed")
	}
	v.Driver.VkDestroyDescriptorSetLayout(layout.DeviceHandle(), layout.Handle(), callbacks.Handle())
}
